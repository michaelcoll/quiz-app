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

	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

func TestUserDBRepository_FindUserById(t *testing.T) {
	connection := getDBConnection(t, true)
	defer connection.Close()

	r := NewUserRepository(NewConnectionWrapperForTest("data", connection))

	user, err := r.FindActiveUserById(context.Background(), "42")
	if err != nil {
		assert.Failf(t, "Fail to get user", "%v", err)
	}

	assert.Nil(t, user)

	err = r.CreateOrReplaceUser(context.Background(), &domain.User{
		Id:      subStr,
		Login:   login,
		Name:    name,
		Picture: picture,
		Role:    domain.Admin,
	})
	if err != nil {
		assert.Failf(t, "Fail to create user", "%v", err)
	}

	user, err = r.FindActiveUserById(context.Background(), subStr)
	if err != nil {
		assert.Failf(t, "Fail to get user", "%v", err)
	}

	assert.Equal(t, subStr, user.Id)
	assert.Equal(t, login, user.Login)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, picture, user.Picture)
	assert.True(t, user.Active)

	user, err = r.FindActiveUserById(context.Background(), subStr)
	if err != nil {
		assert.Failf(t, "Fail to get user", "%v", err)
	}

	assert.Equal(t, subStr, user.Id)
	assert.Equal(t, login, user.Login)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, picture, user.Picture)
	assert.True(t, user.Active)
	assert.Equal(t, user.Role, domain.Admin)
}
