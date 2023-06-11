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

	"github.com/gin-gonic/gin"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

func (c *ApiController) login(ctx *gin.Context) {
	token, exists := ctx.Get("token")
	if !exists {
		handleHttpError(ctx, http.StatusUnauthorized, "no token found in headers")
		return
	}

	user, err := c.authService.Login(ctx, token.(string))
	if err != nil {
		handleError(ctx, err)
		return
	}

	dto := User{}
	ctx.JSON(http.StatusCreated, dto.fromDomain(user))
}

func getRoleFromContext(ctx *gin.Context) (domain.Role, bool) {
	if r, found := ctx.Get(roleCtxKey); found {
		return r.(domain.Role), true
	}

	return 0, false
}

func isAdmin(ctx *gin.Context) bool {
	if role, found := getRoleFromContext(ctx); found {
		return role == domain.Admin
	}

	return false
}

func getUserIdFromContext(ctx *gin.Context) (string, bool) {
	if r, found := ctx.Get(userIdCtxKey); found {
		return r.(string), true
	}

	return "", false
}
