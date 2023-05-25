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

var pathRoleMapping = map[*endPointDef]domain.Role{
	&endPointDef{
		regex:  regexp.MustCompile(`^/api/v1/quiz`),
		method: "GET",
	}: domain.Student,
	&endPointDef{
		regex:  regexp.MustCompile(`^/api/v1/quiz/[^/]+`),
		method: "GET",
	}: domain.Student,
	&endPointDef{
		regex:  regexp.MustCompile(`^/api/v1/user`),
		method: "GET",
	}: domain.Admin,
	&endPointDef{
		regex:  regexp.MustCompile(`^/api/v1/user/[^/]+`),
		method: "DELETE",
	}: domain.Admin,
	&endPointDef{
		regex:  regexp.MustCompile(`^/api/v1/user/[^/]+/activate`),
		method: "POST",
	}: domain.Admin,
	&endPointDef{
		regex:  regexp.MustCompile(`^/api/v1/session`),
		method: "GET",
	}: domain.Student,
}

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

	public.POST("/register", c.register)

	private.GET("/quiz", c.quizList)
	private.GET("/quiz/:sha1", c.quizBySha1)

	private.GET("/user", c.userList)
	private.DELETE("/user/:id", c.deactivateUser)
	private.POST("/user/:id/activate", c.activateUser)

	private.GET("/session", c.sessionList)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on http://0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
