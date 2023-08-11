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
	"database/sql"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

type AuthDBRepository struct {
	domain.AuthRepository

	q  *sqlc.Queries
	tc *cache.Cache
}

func NewAuthRepository(c *sql.DB) *AuthDBRepository {
	tokenCache := cache.New(1*time.Hour, 1*time.Second)
	return &AuthDBRepository{q: sqlc.New(c), tc: tokenCache}
}

func (r *AuthDBRepository) CacheToken(token *domain.AccessToken) error {

	r.tc.Set(token.OpaqueToken, token, 1*time.Hour)

	return nil
}

func (r *AuthDBRepository) FindTokenByTokenStr(tokenStr string) (*domain.AccessToken, error) {

	if t, found := r.tc.Get(tokenStr); found {
		token := t.(*domain.AccessToken)
		token.Provenance = domain.Cache

		return token, nil
	}

	return nil, nil
}
