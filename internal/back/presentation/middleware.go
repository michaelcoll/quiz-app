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
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const (
	userCtxKey   = "user"
	userIdCtxKey = "userId"
	roleCtxKey   = "role"
)

func addCommonMiddlewares(group *gin.Engine) {
	// CORS middleware
	group.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4040"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control", "Range"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Gzip middleware
	group.Use(gzip.Gzip(gzip.DefaultCompression))

	// Recovery middleware
	group.Use(gin.Recovery())
}

func validateAuthHeaderAndGetUser(s *domain.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getBearerToken(ctx)
		if err != nil {
			handleError(ctx, err)
			return
		}

		user, err := s.ValidateTokenAndGetUser(ctx.Request.Context(), token)
		if err != nil {
			handleError(ctx, err)
			return
		}

		ctx.Set(userCtxKey, user)
		ctx.Set(userIdCtxKey, user.Id)
		ctx.Set(roleCtxKey, user.Role)
	}
}

func injectTokenIfPresent(ctx *gin.Context) {
	if token, err := getBearerToken(ctx); err == nil {
		ctx.Set("token", token)
	}
}

func getBearerToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return "", Errorf(http.StatusUnauthorized, "no Authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return "", Errorf(http.StatusUnauthorized, "authorization header is not a Bearer type")
	}

	return token, nil
}

func enforceRoles(ctx *gin.Context) {

	role := findRoleMatchingEndpointDef(ctx)
	if role == 0 {
		handleHttpError(ctx, http.StatusForbidden, "forbidden access (path undefined)")
		return
	}

	if userRole, found := getRoleFromContext(ctx); found {
		if !userRole.CanAccess(role) {
			handleHttpError(ctx,
				http.StatusForbidden,
				fmt.Sprintf("forbidden access (userRole %s, required role %s)", toRoleDto(userRole), toRoleDto(role)))
			return
		}

		return
	}

	handleHttpError(ctx, http.StatusForbidden, "forbidden access (no role in context)")
}

func findRoleMatchingEndpointDef(ctx *gin.Context) domain.Role {
	for def, role := range pathRoleMapping {
		if def.match(ctx.Request) {
			return role
		}
	}

	return 0
}
