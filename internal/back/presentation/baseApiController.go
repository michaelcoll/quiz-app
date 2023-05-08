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
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"

	"github.com/schoolbyhiit/quizz-app/internal/back/domain/service"
)

const apiPort = ":8080"

type ApiController struct {
	quizzService *service.QuizzService
}

func NewApiController(s *service.QuizzService) ApiController {
	return ApiController{quizzService: s}
}

func (c *ApiController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	serveStatic(router)
	addCommonMiddlewares(router)

	public := router.Group("/api/v1")
	private := router.Group("/api/v1")

	addJWTMiddlewares(private)

	//private.GET("/daemon", c.daemonList)
	//private.GET("/daemon/:id", c.daemonById)
	//private.GET("/daemon/:id/media", c.mediaList)

	public.GET("/quizz", c.quizzList)

	//mediaGroup.GET("/daemon/:id/media/:hash", c.contentByHash)
	//mediaGroup.GET("/daemon/:id/thumbnail/:hash", c.thumbnailByHash)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on http://0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
