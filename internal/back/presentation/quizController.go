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
	"github.com/google/uuid"
)

func (c *ApiController) quizList(ctx *gin.Context) {

	start, end, err := extractRangeHeader(ctx.GetHeader("Range"), "quiz")
	if err != nil {
		handleError(ctx, err)
		return
	}

	userId := ""
	if isStudent(ctx) {
		if id, found := getUserIdFromContext(ctx); found {
			userId = id
		} else {
			handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
			return
		}
	}

	quizzes, total, err := c.quizService.FindAllActive(ctx.Request.Context(), userId, end-start, start)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", "quiz", start, int(start)+len(quizzes), total))
	ctx.JSON(http.StatusOK, toQuizDtos(quizzes))
}

func (c *ApiController) quizBySha1(ctx *gin.Context) {
	sha1 := ctx.Param("sha1")

	userId := ""
	if isStudent(ctx) {
		if id, found := getUserIdFromContext(ctx); found {
			userId = id
		} else {
			handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
			return
		}
	}

	quiz, err := c.quizService.FindFullBySha1(ctx, sha1, userId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	dto := Quiz{}
	ctx.JSON(http.StatusOK, dto.fromDomain(quiz))
}

func (c *ApiController) sessionList(ctx *gin.Context) {

	start, end, err := extractRangeHeader(ctx.GetHeader("Range"), "session")
	if err != nil {
		handleError(ctx, err)
		return
	}

	quizActive := true
	if query, present := ctx.GetQuery("quizActive"); present && query == "false" {
		quizActive = false
	}

	userId := ""
	if isAdmin(ctx) {
		if query, present := ctx.GetQuery("userId"); present {
			userId = query
		}
	} else {
		if id, found := getUserIdFromContext(ctx); found {
			userId = id
		} else {
			handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
			return
		}
	}

	sessions, total, err := c.quizService.FindAllSessions(ctx.Request.Context(), quizActive, userId, end-start, start)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", "session", start, int(start)+len(sessions), total))
	ctx.JSON(http.StatusOK, toSessionDtos(sessions))
}

func (c *ApiController) startSession(ctx *gin.Context) {
	quizSha1, present := ctx.GetQuery("quizSha1")
	if !present {
		handleHttpError(ctx, http.StatusBadRequest, "quizSha1 is required")
		return
	}

	userId, present := getUserIdFromContext(ctx)
	if !present {
		handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
		return
	}

	sessionId, err := c.quizService.StartSession(ctx, userId, quizSha1)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, Session{Id: sessionId})
}

func (c *ApiController) addSessionAnswer(ctx *gin.Context) {
	sessionIdStr := ctx.Param("uuid")
	sessionId, err := uuid.Parse(sessionIdStr)
	if err != nil {
		handleHttpError(ctx, http.StatusBadRequest, "invalid sessionId")
		return
	}

	var r SessionAnswerRequestBody
	if err := ctx.BindJSON(&r); err != nil {
		handleError(ctx, err)
		return
	}

	userId, present := getUserIdFromContext(ctx)
	if !present {
		handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
		return
	}

	err = c.quizService.AddSessionAnswer(ctx, sessionId, userId, r.QuestionSha1, r.AnswerSha1, r.Checked)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "answer saved"})
}

func (c *ApiController) quizSessionList(ctx *gin.Context) {

	unit := "quiz-session"
	start, end, err := extractRangeHeader(ctx.GetHeader("Range"), unit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	userId := ""
	if id, found := getUserIdFromContext(ctx); found {
		userId = id
	} else {
		handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
		return
	}

	studentUserId := ""
	if isStudent(ctx) {
		studentUserId = userId
	}

	classId, _ := ctx.GetQuery("classId")

	sessions, total, err := c.quizService.FindAllQuizSessions(ctx.Request.Context(), studentUserId, classId, end-start, start)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", unit, start, int(start)+len(sessions), total))
	ctx.JSON(http.StatusOK, toQuizSessionDtos(sessions, userId))
}

func (c *ApiController) quizSessionByUuid(ctx *gin.Context) {
	sessionIdStr := ctx.Param("uuid")
	sessionId, err := uuid.Parse(sessionIdStr)
	if err != nil {
		handleHttpError(ctx, http.StatusBadRequest, "invalid sessionId")
		return
	}

	userId := ""
	if id, found := getUserIdFromContext(ctx); found {
		userId = id
	} else {
		handleHttpError(ctx, http.StatusUnauthorized, "userId not present in context")
		return
	}

	sessionDetail, err := c.quizService.FindQuizSessionByUuid(ctx.Request.Context(), sessionId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if isStudent(ctx) && userId != sessionDetail.UserId {
		handleHttpError(ctx, http.StatusNotFound, "session not found")
		return
	}

	ctx.JSON(http.StatusOK, toQuizSessionDetail(sessionDetail))
}
