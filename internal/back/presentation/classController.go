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

	"github.com/gin-gonic/gin"
)

func (c *ApiController) classList(ctx *gin.Context) {

	start, end, err := extractRangeHeader(ctx.GetHeader("Range"), "class")
	if err != nil {
		handleError(ctx, err)
		return
	}

	classes, total, err := c.classService.FindAllClasses(ctx.Request.Context(), end-start, start)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", "class", start, int(start)+len(classes), total))
	ctx.JSON(http.StatusOK, toClassDtos(classes))
}
