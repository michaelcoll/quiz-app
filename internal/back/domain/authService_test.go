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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_FindUserById(t *testing.T) {
	mockAuthRepository := NewMockAuthRepository(t)
	service := NewAuthService(mockAuthRepository)

	mockAuthRepository.On("FindUserById", context.Background(), sub).Return(nil, nil)

	_, err := service.FindUserById(context.Background(), sub)
	if err != nil {
		if code, found := GetCodeFromError(err); found {
			if code != NotFound {
				assert.Failf(t, "Error while getting user", "%v", err)
			}
		}
	}
}

func TestRole_canAccess(t *testing.T) {

	assert.True(t, Admin.CanAccess(Teacher))
	assert.True(t, Admin.CanAccess(Student))
	assert.True(t, Teacher.CanAccess(Student))
	assert.False(t, Teacher.CanAccess(Admin))
	assert.False(t, Student.CanAccess(Admin))

}
