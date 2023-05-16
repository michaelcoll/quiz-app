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
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var rangeRxp = regexp.MustCompile(`(?P<Unit>.*)=(?P<Start>[0-9]+)-(?P<End>[0-9]*)`)

func (c *ApiController) quizList(ctx *gin.Context) {

	start, end, err := extractRangeHeader(ctx.GetHeader("Range"))
	if err != nil {
		handleError(ctx, err)
		return
	}

	quizzes, total, err := c.quizService.FindAllActive(ctx.Request.Context(), end-start, start)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", "quiz", start, int(start)+len(quizzes), total))
	ctx.JSON(http.StatusOK, fromDomains(quizzes))
}

func extractRangeHeader(rangeHeader string) (uint16, uint16, error) {
	r := rangeRxp.FindStringSubmatch(rangeHeader)
	st := http.StatusRequestedRangeNotSatisfiable

	if len(r) < 4 {
		return 0, 0, Errorf(st, "Range is not valid, supported format : quiz=0-25")
	}

	if r[1] != "quiz" {
		return 0, 0, Errorf(st, "Unit in range is not valid, supported unit : quiz")
	}

	start, errStart := strconv.Atoi(r[2])
	end, errEnd := strconv.Atoi(r[3])

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

func (c *ApiController) quizBySha1(ctx *gin.Context) {
	sha1 := ctx.Param("sha1")

	quiz, err := c.quizService.FindFullBySha1(ctx, sha1)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fromDomain(quiz))
}
