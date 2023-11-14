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
	"time"

	"github.com/gin-gonic/gin"
)

func (c *ApiController) sync(ctx *gin.Context) {

	if time.Since(c.lastSyncUpdate) > 10*time.Second {
		c.lastSyncUpdate = time.Now()
		err := c.quizService.Sync(ctx.Request.Context())
		if err != nil {
			handleError(ctx, err)
			return
		}
	} else {
		handleHttpError(ctx, http.StatusTooManyRequests, "too many sync requests")
		return
	}
}
