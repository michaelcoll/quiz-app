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
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const accessTokenBaseUrl = "https://oauth2.googleapis.com"

type GoogleAccessTokenCaller struct {
	domain.AccessTokenCaller

	accessTokenBaseUrl string
}

type successResponse struct {
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Exp           string `json:"exp"`
	ExpiresIn     string `json:"expires_in"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func NewGoogleAccessTokenCaller() *GoogleAccessTokenCaller {
	return &GoogleAccessTokenCaller{
		accessTokenBaseUrl: accessTokenBaseUrl,
	}
}

func (c *GoogleAccessTokenCaller) Get(ctx context.Context, token string) (*domain.AccessToken, error) {
	url := fmt.Sprintf("%s/tokeninfo?access_token=%s", c.accessTokenBaseUrl, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return c.toAccessToken(&res, token)
}

func (c *GoogleAccessTokenCaller) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return domain.Errorf(domain.UnAuthorized, errRes.Error)
		}

		return domain.Errorf(domain.UnexpectedError, "unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return domain.Errorf(domain.UnexpectedError, "can't decode access token response (%v)", err)
	}

	return nil
}

func (c *GoogleAccessTokenCaller) toAccessToken(res *successResponse, token string) (*domain.AccessToken, error) {

	expUnix, err := strconv.ParseInt(res.Exp, 10, 64)
	if err != nil {
		return nil, domain.Errorf(domain.UnexpectedError, "can't parse token exp %s (%v)", res.Exp, err)
	}

	expIn, err := strconv.ParseInt(res.ExpiresIn, 10, 16)
	if err != nil {
		return nil, domain.Errorf(domain.UnexpectedError, "can't parse token expires_in %s (%v)", res.Exp, err)
	}

	return &domain.AccessToken{
		Aud:         res.Aud,
		Sub:         res.Sub,
		Exp:         time.Unix(expUnix, 0),
		ExpiresIn:   int(expIn),
		Email:       res.Email,
		Provenance:  domain.Api,
		OpaqueToken: token,
	}, nil
}
