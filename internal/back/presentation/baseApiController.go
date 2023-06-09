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
	"regexp"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const apiPort = ":8080"

type ApiController struct {
	quizService *domain.QuizService
	authService *domain.AuthService
}

func NewApiController(quizService *domain.QuizService, authService *domain.AuthService) ApiController {
	return ApiController{quizService: quizService, authService: authService}
}

var pathRoleMapping = map[*endPointDef]domain.Role{}

func (c *ApiController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(injectTokenIfPresent)

	serveStatic(router)
	addCommonMiddlewares(router)

	public := router.Group("/api/v1")
	private := router.Group("/api/v1")

	private.Use(validateAuthHeaderAndGetUser(c.authService))
	private.Use(enforceRoles)

	addPostEndpoint(public, "/login", domain.NoRole, c.login)

	addGetEndpoint(private, "/quiz", domain.Student, c.quizList)
	addGetEndpoint(private, "/quiz/:sha1", domain.Student, c.quizBySha1)

	addGetEndpoint(private, "/user", domain.Admin, c.userList)
	addDeleteEndpoint(private, "/user/:id", domain.Admin, c.deactivateUser)
	addPostEndpoint(private, "/user/:id/activate", domain.Admin, c.activateUser)

	addGetEndpoint(private, "/session", domain.Student, c.sessionList)
	addPostEndpoint(private, "/session", domain.Student, c.startSession)
	addPostEndpoint(private, "/session/:sessionId/answer", domain.Student, c.addSessionAnswer)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on http://0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}

func addGetEndpoint(routerGroup *gin.RouterGroup, path string, role domain.Role, handler gin.HandlerFunc) {
	appendEndpointDef(routerGroup, path, "GET", role)
	routerGroup.GET(path, handler)
}

func addPostEndpoint(routerGroup *gin.RouterGroup, path string, role domain.Role, handler gin.HandlerFunc) {
	appendEndpointDef(routerGroup, path, "POST", role)
	routerGroup.POST(path, handler)
}

func addDeleteEndpoint(routerGroup *gin.RouterGroup, path string, role domain.Role, handler gin.HandlerFunc) {
	appendEndpointDef(routerGroup, path, "DELETE", role)
	routerGroup.DELETE(path, handler)
}

func appendEndpointDef(routerGroup *gin.RouterGroup, path string, method string, role domain.Role) {
	if role != domain.NoRole {
		pathRoleMapping[&endPointDef{
			regex:  toRegExPath(routerGroup, path),
			method: method,
		}] = role
	}
}

func toRegExPath(routerGroup *gin.RouterGroup, path string) *regexp.Regexp {
	r := regexp.MustCompile(`:[0-9a-zA-Z]+`)
	replacedPath := r.ReplaceAllString(path, "[^/]+")

	return regexp.MustCompile(fmt.Sprintf("^%s%s$", routerGroup.BasePath(), replacedPath))
}
