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
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBearerToken(t *testing.T) {
	ctx := &gin.Context{}
	ctx.Request = &http.Request{Header: http.Header{}}

	t.Run("returns error when no Authorization header", func(t *testing.T) {
		_, err := getBearerToken(ctx)
		assert.Error(t, err)
		assert.Equal(t, "no Authorization header", err.Error())
	})

	t.Run("returns error when Authorization header is not a Bearer type", func(t *testing.T) {
		ctx.Request.Header = http.Header{"Authorization": []string{"Basic token"}}
		_, err := getBearerToken(ctx)
		assert.Error(t, err)
		assert.Equal(t, "authorization header is not a Bearer type", err.Error())
	})

	t.Run("returns error when token has not a valid format", func(t *testing.T) {
		ctx.Request.Header = http.Header{"Authorization": []string{"Bearer token"}}
		_, err := getBearerToken(ctx)
		assert.Error(t, err)
		assert.Equal(t, "token has not a valid format, token=token", err.Error())
	})

	t.Run("returns token when Authorization header is a Bearer type and token has a valid format", func(t *testing.T) {
		ctx.Request.Header = http.Header{"Authorization": []string{"Bearer gho_token"}}
		token, err := getBearerToken(ctx)
		assert.NoError(t, err)
		assert.Equal(t, "gho_token", token)
	})
}

func TestGetApiKey(t *testing.T) {
	ctx := &gin.Context{}
	ctx.Request = &http.Request{Header: http.Header{}}

	t.Run("returns error when no Authorization header", func(t *testing.T) {
		_, err := getApiKey(ctx)
		assert.Error(t, err)
		assert.Equal(t, "no X-Api-Key header", err.Error())
	})

	t.Run("returns Apikey when Authorization header is a Apikey", func(t *testing.T) {
		ctx.Request.Header = http.Header{"X-Api-Key": []string{"key"}}
		key, err := getApiKey(ctx)
		assert.NoError(t, err)
		assert.Equal(t, "key", key)
	})
}
