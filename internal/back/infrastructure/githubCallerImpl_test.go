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

package infrastructure

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorsalgado/mocha/v3"
	"github.com/vitorsalgado/mocha/v3/expect"
	"github.com/vitorsalgado/mocha/v3/reply"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const accessToken = "gho_1oVpR3GWd0ZnaHQ15JG63HDFPyFdWN2hbEyK"

func TestGoogleAccessTokenCaller_Get(t *testing.T) {
	m := mocha.New(t).CloseOnCleanup(t)
	m.Start()

	response := successResponse{
		Sub:     sub,
		Name:    name,
		Login:   login,
		Picture: picture,
	}

	scoped := m.AddMocks(
		mocha.
			Get(expect.URLPath("/user")).
			Header("Content-Type", expect.ToEqual("application/json; charset=utf-8")).
			Header("Accept", expect.ToEqual("application/json; charset=utf-8")).
			Header("Authorization", expect.ToEqual(fmt.Sprintf("Bearer %s", accessToken))).
			Reply(reply.OK().BodyJSON(response)))

	caller := GithubAccessTokenCaller{
		accessTokenBaseUrl: m.URL(),
	}

	actual, err := caller.Get(context.Background(), accessToken)
	if err != nil {
		assert.Failf(t, "Fail to call /user", "%v", err)
	}

	expected := &domain.AccessToken{
		Sub:         subStr,
		Login:       login,
		Name:        name,
		Picture:     picture,
		OpaqueToken: accessToken,
		Provenance:  domain.Api,
	}

	assert.True(t, scoped.Called(), "google endpoint not called")
	assert.Equal(t, expected, actual)
}
