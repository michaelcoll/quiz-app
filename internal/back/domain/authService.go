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

package domain

import (
	"context"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type AuthService struct {
	r AuthRepository
	c AccessTokenCaller
}

func NewAuthService(r AuthRepository, c AccessTokenCaller) AuthService {
	return AuthService{r: r, c: c}
}

func (s *AuthService) Register(ctx context.Context, user *User, accessToken string) error {
	token, err := s.validateToken(ctx, accessToken)
	if err != nil {
		return err
	}

	if token.Email != user.Email {
		return Errorf(UnAuthorized, "token email and user email mismatch (%s != %s)", token.Email, user.Email)
	}

	if token.Email != user.Email {
		return Errorf(UnAuthorized, "token email and user email mismatch (%s != %s)", token.Email, user.Email)
	}

	if token.Sub != user.Id {
		return Errorf(UnAuthorized, "token sub and user id mismatch (%s != %s)", token.Sub, user.Id)
	}

	emailDomain := strings.Split(user.Email, "@")[1]
	restrictedDomainName := viper.GetString("restrict-email-domain")
	if len(restrictedDomainName) > 0 && emailDomain != restrictedDomainName {
		return Errorf(UnAuthorized, "user is not in a valid domain (%s not in domain %s)", user.Email, restrictedDomainName)
	}

	err = s.r.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	err = s.r.AddRoleToUser(ctx, user.Id, Student)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) validateToken(ctx context.Context, tokenStr string) (token *AccessToken, err error) {

	// Try to get the token from the cache
	token, err = s.r.FindTokenByTokenStr(ctx, tokenStr)
	if err != nil {
		return nil, err
	}

	// If it's not in cache then get it from API
	if token == nil {
		token, err = s.c.Get(ctx, tokenStr)
		if err != nil {
			return nil, err
		}
	}

	aud := viper.GetString("auth0-audience")
	if len(aud) > 0 && token.Aud != aud {
		return nil, Errorf(UnAuthorized, "token is using a different audience than the one specified in config (%s != %s)", token.Aud, aud)
	}

	if token.Exp.Before(time.Now()) {
		return nil, Errorf(UnAuthorized, "token is expired (%s)", token.Exp.Format(time.RFC3339))
	}

	return token, nil
}

func (s *AuthService) ValidateTokenAndGetUser(ctx context.Context, accessToken string) (*User, error) {
	token, err := s.validateToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.r.FindUserById(ctx, token.Sub)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Errorf(UnAuthorized, "unknown user '%s'", token.Sub)
	}

	if token.Provenance == Api {
		err := s.r.CreateToken(ctx, token)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *AuthService) FindUserById(ctx context.Context, id string) (*User, error) {
	user, err := s.r.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Errorf(NotFound, "user with id '%s' not found", id)
	}

	return user, nil
}
