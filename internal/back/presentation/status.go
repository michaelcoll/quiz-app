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
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

var statusMapping = map[domain.ErrorCode]int{
	domain.NotFound:        http.StatusNotFound,
	domain.InvalidArgument: http.StatusBadRequest,
	domain.UnAuthorized:    http.StatusUnauthorized,
	domain.UnexpectedError: http.StatusInternalServerError,
}

func handleError(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	if code, match := domain.GetCodeFromError(err); match {
		if st, match := statusMapping[code]; match {
			status = st
		}
	}

	if st, match := GetCodeFromError(err); match {
		status = st
	}

	handleHttpError(ctx, status, err.Error())
}

func handleHttpError(ctx *gin.Context, st int, message string) {
	ctx.AbortWithStatusJSON(st, gin.H{"message": message})
}
