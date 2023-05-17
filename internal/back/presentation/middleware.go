/*
 * Copyright (c) 2022-2023 Michaël COLL.
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
	"context"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

const (
	auth0IssuerUrl = "https://dev-f6i8nn7snfrfpiop.eu.auth0.com/"
	auth0Audience  = "https://gallery-api/"
)

func addCommonMiddlewares(group *gin.Engine) {
	// CORS middleware
	group.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4040"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control", "Range"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Gzip middleware
	group.Use(gzip.Gzip(gzip.DefaultCompression))

	// Recovery middleware
	group.Use(gin.Recovery())
}

func addJWTMiddlewares(group *gin.RouterGroup) {
	issuerURL, _ := url.Parse(auth0IssuerUrl)

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{auth0Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &EmailCustomClaims{}
			},
		),
	)

	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken,
		jwtmiddleware.WithTokenExtractor(
			jwtmiddleware.MultiTokenExtractor(
				jwtmiddleware.AuthHeaderTokenExtractor,
				jwtmiddleware.ParameterTokenExtractor("access-token"),
			)),
		jwtmiddleware.WithValidateOnOptions(false))
	group.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
}

// EmailCustomClaims contains custom data we want from the token.
type EmailCustomClaims struct {
	Email string `json:"email"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c EmailCustomClaims) Validate(_ context.Context) error {
	return nil
}

//func extractEmailFromToken(ctx *gin.Context) (string, error) {
//	if claims, ok := ctx.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims); ok {
//		return claims.CustomClaims.(*EmailCustomClaims).Email, nil
//	} else {
//		return "", Errorf(http.StatusUnauthorized, "email claim not found")
//	}
//}
