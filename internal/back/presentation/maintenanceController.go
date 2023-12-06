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
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (c *ApiController) dbDump(ctx *gin.Context) {
	fullDbLocation, err := c.maintenanceService.Dump()
	if err != nil {
		handleError(ctx, err)
		return
	}

	fi, err := os.Stat(fullDbLocation)
	if err != nil {
		handleHttpError(ctx, http.StatusInternalServerError, "database file not found")
		return
	}

	// Set headers for download
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename=data.db")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", fmt.Sprintf("%d", fi.Size()))

	// Send the file as a response
	ctx.File(fullDbLocation)

}
