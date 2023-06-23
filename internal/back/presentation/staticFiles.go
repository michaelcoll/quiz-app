/*
 * Copyright (c) 2022-2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package presentation

import (
	"io/fs"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	cachecontrol "github.com/joeig/gin-cachecontrol"

	"github.com/michaelcoll/quiz-app/internal/web"

	"github.com/gin-gonic/gin"
)

func serveStatic(router *gin.Engine) {
	static, _ := fs.Sub(web.Static, ".output/public")
	staticIndexFS(http.FS(static), router)

	staticNuxtResources, _ := fs.Sub(web.Static, ".output/public/_nuxt")
	staticQuizPage, _ := fs.Sub(web.Static, ".output/public/quiz")

	// Pages
	router.StaticFS("/quiz", http.FS(staticQuizPage))

	cachedStatic := router.Group("/")

	// Add middlewares
	cachedStatic.Use(cachecontrol.New(&cachecontrol.Config{
		Public:    true,
		MaxAge:    cachecontrol.Duration(30 * 24 * time.Hour),
		Immutable: true,
	}))
	cachedStatic.Use(gzip.Gzip(gzip.DefaultCompression))

	cachedStatic.StaticFS("/_nuxt", http.FS(staticNuxtResources))
}

func staticIndexFS(fs http.FileSystem, router *gin.Engine) {
	relativePath := "/"
	handler := createStaticHandler(relativePath, fs)

	// Register GET and HEAD handlers
	router.GET(relativePath, handler)
	router.HEAD(relativePath, handler)
}

func createStaticHandler(relativePath string, fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.StripPrefix(relativePath, http.FileServer(fs))

	return func(c *gin.Context) {
		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open("/")
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		_ = f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
