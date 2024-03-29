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
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const apiPort = ":8080"

var rangeRxp = regexp.MustCompile(`(?P<Unit>.*)=(?P<Start>[0-9]+)-(?P<End>[0-9]*)`)

type ApiController struct {
	lastSyncUpdate time.Time

	authService        *domain.AuthService
	classService       *domain.ClassService
	quizService        *domain.QuizService
	userService        *domain.UserService
	healthService      *domain.HealthService
	maintenanceService *domain.MaintenanceService
}

func NewApiController(
	authService *domain.AuthService,
	classService *domain.ClassService,
	quizService *domain.QuizService,
	userService *domain.UserService,
	healthService *domain.HealthService,
	maintenanceService *domain.MaintenanceService) ApiController {
	return ApiController{lastSyncUpdate: time.Now(), authService: authService, classService: classService,
		quizService: quizService, userService: userService, healthService: healthService,
		maintenanceService: maintenanceService}
}

var pathRoleMapping = map[*endPointDef]domain.Role{}

func (c *ApiController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/started", "/health/ready", "/health/live"),
		gin.Recovery(),
	)
	router.Use(injectTokenIfPresent)

	addCommonMiddlewares(router)

	public := router.Group("/api/v1")
	private := router.Group("/api/v1")
	health := router.Group("/health")
	maintenance := router.Group("/maintenance")

	private.Use(validateAuthHeaderAndGetUser(c.authService))
	private.Use(enforceRoles)

	maintenance.Use(validateAuthHeaderAndGetApiKey)
	maintenance.Use(enforceApiKey)

	addPostEndpoint(public, "/login", domain.NoRole, c.login)
	addPostEndpoint(public, "/sync", domain.NoRole, c.sync)

	addGetEndpoint(health, "/started", domain.NoRole, c.started)
	addGetEndpoint(health, "/ready", domain.NoRole, c.ready)
	addGetEndpoint(health, "/live", domain.NoRole, c.live)

	addGetEndpoint(maintenance, "/database/dump", domain.Machine, c.dbDump)

	addGetEndpoint(private, "/quiz", domain.Student, c.quizList)
	addGetEndpoint(private, "/quiz/:sha1", domain.Student, c.quizBySha1)
	addPostEndpoint(private, "/quiz/:sha1/class/:uuid", domain.Teacher, c.createQuizClassVisibility)
	addDeleteEndpoint(private, "/quiz/:sha1/class/:uuid", domain.Teacher, c.deleteQuizClassVisibility)

	addGetEndpoint(private, "/user", domain.Teacher, c.userList)
	addGetEndpoint(private, "/user/me", domain.Student, c.me)
	addDeleteEndpoint(private, "/user/:id", domain.Admin, c.deactivateUser)
	addPutEndpoint(private, "/user/:id/activate", domain.Admin, c.activateUser)
	addPutEndpoint(private, "/user/:id/role/:roleName", domain.Admin, c.updateUserRole)
	addPutEndpoint(private, "/user/:id/class/:uuid", domain.Teacher, c.assignUserToClass)

	addGetEndpoint(private, "/session", domain.Student, c.sessionList)
	addPostEndpoint(private, "/session", domain.Student, c.startSession)
	addPostEndpoint(private, "/session/:uuid/answer", domain.Student, c.addSessionAnswer)

	addGetEndpoint(private, "/class", domain.Teacher, c.classList)
	addPostEndpoint(private, "/class", domain.Admin, c.classCreate)
	addPutEndpoint(private, "/class/:uuid", domain.Admin, c.classUpdate)
	addDeleteEndpoint(private, "/class/:uuid", domain.Admin, c.classDelete)

	addGetEndpoint(private, "/quiz-session", domain.Student, c.quizSessionList)
	addGetEndpoint(private, "/quiz-session/:uuid", domain.Student, c.quizSessionByUuid)

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

func addPutEndpoint(routerGroup *gin.RouterGroup, path string, role domain.Role, handler gin.HandlerFunc) {
	appendEndpointDef(routerGroup, path, "PUT", role)
	routerGroup.PUT(path, handler)
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

func extractRangeHeader(rangeHeader string, unit string) (uint16, uint16, error) {
	r := rangeRxp.FindStringSubmatch(rangeHeader)
	st := http.StatusRequestedRangeNotSatisfiable

	if len(r) < 4 {
		return 0, 0, Errorf(st, "Range is not valid, supported format : %s=0-25", unit)
	}

	if r[1] != unit {
		return 0, 0, Errorf(st, "Unit in range is not valid, supported unit : %s", unit)
	}

	start, errStart := strconv.ParseUint(r[2], 10, 16)
	end, errEnd := strconv.ParseUint(r[3], 10, 16)

	if len(r[3]) == 0 {
		end = 0
	}

	if errStart != nil {
		return 0, 0, Errorf(st, "Start range is not valid")
	}

	if len(r[3]) != 0 && errEnd != nil {
		return 0, 0, Errorf(st, "End range is not valid")
	}

	if end != 0 && start >= end {
		return 0, 0, Errorf(st, "Range is not valid, start > end")
	}

	return uint16(start), uint16(end), nil
}
