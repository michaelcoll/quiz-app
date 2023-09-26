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
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type AuthService struct {
	authRepository AuthRepository
	userRepository UserRepository
	c              AccessTokenCaller
}

func NewAuthService(authRepository AuthRepository, userRepository UserRepository, c AccessTokenCaller) AuthService {
	adminUsername := viper.GetString("default-admin-username")

	if len(adminUsername) > 0 {
		fmt.Printf("%s Default admin username set to %s\n", color.HiYellowString("i"), color.BlueString(adminUsername))
	}

	return AuthService{authRepository: authRepository, userRepository: userRepository, c: c}
}

func (s *AuthService) Login(ctx context.Context, accessToken string) (*User, error) {
	token, err := s.validateAndGetToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	user := &User{
		Id:      token.Sub,
		Login:   token.Login,
		Name:    token.Name,
		Picture: token.Picture,
		Active:  true,
	}

	adminUsername := viper.GetString("default-admin-username")
	if token.Login == adminUsername {
		user.Role = Admin
	} else {
		user.Role = Student
	}

	if dbUser, err := s.userRepository.FindUserById(ctx, token.Sub); dbUser != nil && err == nil {
		dbUser.Name = user.Name
		dbUser.Login = user.Login
		dbUser.Picture = user.Picture

		err = s.userRepository.UpdateUserInfo(ctx, dbUser)
		if err != nil {
			return nil, err
		}

		return dbUser, nil
	}

	err = s.userRepository.CreateOrReplaceUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) validateAndGetToken(ctx context.Context, tokenStr string) (*AccessToken, error) {

	// Try to get the token from the cache
	token, err := s.authRepository.FindTokenByTokenStr(tokenStr)
	if err != nil {
		return nil, err
	}

	// If it's not in cache, then get it from API
	if token == nil {
		token, err = s.c.Get(ctx, tokenStr)
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

func (s *AuthService) ValidateTokenAndGetUser(ctx context.Context, accessToken string) (*User, error) {
	token, err := s.validateAndGetToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindActiveUserById(ctx, token.Sub)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Errorf(UnAuthorized, "unknown user '%s'", token.Sub)
	}

	if token.Provenance == Api {
		err := s.authRepository.CacheToken(token)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (role Role) CanAccess(other Role) bool {

	if role == other {
		return true
	}

	if role == Admin && (other == Teacher || other == Student) {
		return true
	}

	if role == Teacher && other == Student {
		return true
	}

	return false
}
