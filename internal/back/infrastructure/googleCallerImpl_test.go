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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vitorsalgado/mocha/v3"
	"github.com/vitorsalgado/mocha/v3/expect"
	"github.com/vitorsalgado/mocha/v3/reply"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const accessToken = "ya29.a0AWY7CklOxN71GhDLKDpGKyHeaPDbvKdpZjL7btStvCIVj3AoKWzz8ttt60diwyhusAgQrMGhPNGa95a3XrlVQSIVkMGaEn3C7sma0iqaPVeQgL2cD7esKc40Viq19huXBzh-9Oreg7lvymX_tw3ZL6XMnTpHDTYqaCgYKASoSARASFQG1tDrpqFJkEkSiMaydX5a51q1_ew0167"

func TestGoogleAccessTokenCaller_Get(t *testing.T) {
	m := mocha.New(t).CloseOnCleanup(t)
	m.Start()

	response := successResponse{
		Aud:           aud,
		Sub:           sub,
		Exp:           expStr,
		Email:         email,
		EmailVerified: emailVerified,
	}

	scoped := m.AddMocks(
		mocha.
			Get(expect.URLPath("/tokeninfo")).
			Header("Content-Type", expect.ToEqual("application/json; charset=utf-8")).
			Header("Accept", expect.ToEqual("application/json; charset=utf-8")).
			Query("access_token", expect.ToEqual(accessToken)).
			Reply(reply.OK().BodyJSON(response)))

	caller := GoogleAccessTokenCaller{
		accessTokenBaseUrl: m.URL(),
	}

	actual, err := caller.Get(context.Background(), accessToken)
	if err != nil {
		assert.Failf(t, "Fail to call /tokeninfo", "%v", err)
	}

	expected := &domain.AccessToken{
		Aud:         aud,
		Sub:         sub,
		Exp:         time.Unix(exp, 0),
		Email:       email,
		OpaqueToken: accessToken,
		Provenance:  domain.Api,
	}

	assert.True(t, scoped.Called(), "google endpoint not called")
	assert.Equal(t, expected, actual)
}
