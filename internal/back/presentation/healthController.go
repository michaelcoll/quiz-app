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

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var ready = false

func (c *ApiController) started(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func (c *ApiController) ready(ctx *gin.Context) {
	if ready {
		ctx.String(http.StatusOK, "OK")
	} else {
		ctx.String(http.StatusServiceUnavailable, "Not ready yet !")
	}
}

func (c *ApiController) SetReady() {
	ready = true
	fmt.Printf("%s Ready to serve\n", color.GreenString("✓"))
}

func (c *ApiController) live(ctx *gin.Context) {
	if c.healthService.PingDB(ctx) {
		ctx.String(http.StatusOK, "OK")
	} else {
		ctx.String(http.StatusServiceUnavailable, "DB Connection down !")
	}
}
