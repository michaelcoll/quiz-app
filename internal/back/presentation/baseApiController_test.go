/*
 * Copyright (c) 2023 Michaël COLL.
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
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

func Test_toRegExPath(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	public := router.Group("/api/v1")

	r := toRegExPath(public, "/quiz/:sha1")
	assert.True(t, r.MatchString("/api/v1/quiz/1234564687"))
	assert.False(t, r.MatchString("/api/v1/quiz/1234564687/wrong-path"))

}

func Test_addGetEndpoint(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	public := router.Group("/api/v1")

	u := &url.URL{
		Path: "/api/v1/quiz/7126ef9c-b16d-442b-b773-37a793ddb89b/action",
	}
	r := &http.Request{
		Method: "GET",
		URL:    u,
	}
	ctx := &gin.Context{
		Request: r,
	}

	addGetEndpoint(public, "/quiz/:sha1/action", domain.Student, testHandlerFunc)

	role := findRoleMatchingEndpointDef(ctx)

	assert.Equal(t, domain.Student, role)
}

func Test_addPostEndpoint(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	public := router.Group("/api/v1")

	u := &url.URL{
		Path: "/api/v1/register",
	}
	r := &http.Request{
		Method: "POST",
		URL:    u,
	}
	ctx := &gin.Context{
		Request: r,
	}

	addPostEndpoint(public, "/register", domain.Student, testHandlerFunc)

	role := findRoleMatchingEndpointDef(ctx)

	assert.Equal(t, domain.Student, role)
}

func Test_addDeleteEndpoint(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	public := router.Group("/api/v1")

	u := &url.URL{
		Path: "/api/v1/user/1321656151",
	}
	r := &http.Request{
		Method: "DELETE",
		URL:    u,
	}
	ctx := &gin.Context{
		Request: r,
	}

	addDeleteEndpoint(public, "/user/:id", domain.Admin, testHandlerFunc)

	role := findRoleMatchingEndpointDef(ctx)

	assert.Equal(t, domain.Admin, role)
}

func testHandlerFunc(_ *gin.Context) {
}
