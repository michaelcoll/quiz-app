/*
 * Copyright (c) 2023-2025 Michaël COLL.
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

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const accessTokenBaseUrl = "https://api.github.com"

type GithubAccessTokenCaller struct {
	domain.AccessTokenCaller

	accessTokenBaseUrl string
}

type successResponse struct {
	Sub     int    `json:"id"`
	Name    string `json:"name"`
	Login   string `json:"login"`
	Picture string `json:"avatar_url"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewGithubAccessTokenCaller() *GithubAccessTokenCaller {
	return &GithubAccessTokenCaller{
		accessTokenBaseUrl: accessTokenBaseUrl,
	}
}

func (c *GithubAccessTokenCaller) Get(ctx context.Context, token string) (*domain.AccessToken, error) {
	url := fmt.Sprintf("%s/user", c.accessTokenBaseUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := successResponse{}
	if err := c.sendRequest(req, token, &res); err != nil {
		return nil, err
	}

	return c.toAccessToken(&res, token)
}

func (c *GithubAccessTokenCaller) sendRequest(req *http.Request, token string, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	authHeader := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", authHeader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return domain.Errorf(domain.UnAuthorized, "%s", errRes.Message)
		}

		return domain.Errorf(domain.UnexpectedError, "unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return domain.Errorf(domain.UnexpectedError, "can't decode access token response (%v)", err)
	}

	return nil
}

func (c *GithubAccessTokenCaller) toAccessToken(res *successResponse, token string) (*domain.AccessToken, error) {

	return &domain.AccessToken{
		Sub:         strconv.Itoa(res.Sub),
		Login:       res.Login,
		Name:        res.Name,
		Picture:     res.Picture,
		Provenance:  domain.Api,
		OpaqueToken: token,
	}, nil
}
