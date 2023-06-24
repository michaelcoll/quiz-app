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
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"google.golang.org/api/idtoken"
)

type AuthService struct {
	authRepository AuthRepository
	userRepository UserRepository
}

func NewAuthService(authRepository AuthRepository, userRepository UserRepository) AuthService {
	adminEmail := viper.GetString("default-admin-email")

	if len(adminEmail) > 0 {
		fmt.Printf("%s Default admin email set to %s\n", color.HiYellowString("i"), color.BlueString(adminEmail))
	}

	return AuthService{authRepository: authRepository, userRepository: userRepository}
}

func (s *AuthService) Login(ctx context.Context, idToken string) (*User, error) {
	token, err := s.validateToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	emailDomain := strings.Split(token.Email, "@")[1]
	restrictedDomainName := viper.GetString("restrict-email-domain")
	if len(restrictedDomainName) > 0 && emailDomain != restrictedDomainName {
		return nil, Errorf(InvalidArgument, "user is not in a valid domain (%s not in domain %s)", token.Email, restrictedDomainName)
	}

	user := &User{
		Id:        token.Sub,
		Email:     token.Email,
		Firstname: token.FirstName,
		Lastname:  token.LastName,
		Active:    true,
	}

	adminEmail := viper.GetString("default-admin-email")
	if token.Email == adminEmail {
		user.Role = Admin
	} else {
		user.Role = Student
	}

	err = s.userRepository.CreateOrReplaceUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) parseToken(ctx context.Context, tokenStr string) (*IdToken, error) {
	aud := viper.GetString("auth0-audience")
	payload, err := idtoken.Validate(ctx, tokenStr, aud)
	if err != nil {
		return nil, err
	}

	iat := (int64)(payload.Claims["iat"].(float64))
	exp := (int64)(payload.Claims["exp"].(float64))
	return &IdToken{
		Aud:         payload.Claims["aud"].(string),
		Sub:         payload.Claims["sub"].(string),
		Exp:         time.Unix(exp, 0),
		ExpiresIn:   (int)(exp - iat),
		Email:       payload.Claims["email"].(string),
		FirstName:   payload.Claims["given_name"].(string),
		LastName:    payload.Claims["family_name"].(string),
		Provenance:  Parse,
		JwtStrToken: tokenStr,
	}, nil
}

func (s *AuthService) validateToken(ctx context.Context, tokenStr string) (*IdToken, error) {

	// parse token
	token, err := s.parseToken(ctx, tokenStr)
	if err != nil {
		return nil, Errorf(UnAuthorized, "error while parsing token : %v", err)
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

	user, err := s.userRepository.FindUserById(ctx, token.Sub)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Errorf(UnAuthorized, "unknown user '%s'", token.Sub)
	}

	if token.Provenance == Parse {
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
