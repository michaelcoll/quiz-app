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

// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// AssignUserToClass provides a mock function with given fields: ctx, userId, classId
func (_m *MockUserRepository) AssignUserToClass(ctx context.Context, userId string, classId uuid.UUID) error {
	ret := _m.Called(ctx, userId, classId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) error); ok {
		r0 = rf(ctx, userId, classId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_AssignUserToClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignUserToClass'
type MockUserRepository_AssignUserToClass_Call struct {
	*mock.Call
}

// AssignUserToClass is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - classId uuid.UUID
func (_e *MockUserRepository_Expecter) AssignUserToClass(ctx interface{}, userId interface{}, classId interface{}) *MockUserRepository_AssignUserToClass_Call {
	return &MockUserRepository_AssignUserToClass_Call{Call: _e.mock.On("AssignUserToClass", ctx, userId, classId)}
}

func (_c *MockUserRepository_AssignUserToClass_Call) Run(run func(ctx context.Context, userId string, classId uuid.UUID)) *MockUserRepository_AssignUserToClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *MockUserRepository_AssignUserToClass_Call) Return(_a0 error) *MockUserRepository_AssignUserToClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_AssignUserToClass_Call) RunAndReturn(run func(context.Context, string, uuid.UUID) error) *MockUserRepository_AssignUserToClass_Call {
	_c.Call.Return(run)
	return _c
}

// CreateOrReplaceUser provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) CreateOrReplaceUser(ctx context.Context, user *User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_CreateOrReplaceUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrReplaceUser'
type MockUserRepository_CreateOrReplaceUser_Call struct {
	*mock.Call
}

// CreateOrReplaceUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *User
func (_e *MockUserRepository_Expecter) CreateOrReplaceUser(ctx interface{}, user interface{}) *MockUserRepository_CreateOrReplaceUser_Call {
	return &MockUserRepository_CreateOrReplaceUser_Call{Call: _e.mock.On("CreateOrReplaceUser", ctx, user)}
}

func (_c *MockUserRepository_CreateOrReplaceUser_Call) Run(run func(ctx context.Context, user *User)) *MockUserRepository_CreateOrReplaceUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*User))
	})
	return _c
}

func (_c *MockUserRepository_CreateOrReplaceUser_Call) Return(_a0 error) *MockUserRepository_CreateOrReplaceUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_CreateOrReplaceUser_Call) RunAndReturn(run func(context.Context, *User) error) *MockUserRepository_CreateOrReplaceUser_Call {
	_c.Call.Return(run)
	return _c
}

// FindActiveUserById provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) FindActiveUserById(ctx context.Context, id string) (*User, error) {
	ret := _m.Called(ctx, id)

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_FindActiveUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindActiveUserById'
type MockUserRepository_FindActiveUserById_Call struct {
	*mock.Call
}

// FindActiveUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockUserRepository_Expecter) FindActiveUserById(ctx interface{}, id interface{}) *MockUserRepository_FindActiveUserById_Call {
	return &MockUserRepository_FindActiveUserById_Call{Call: _e.mock.On("FindActiveUserById", ctx, id)}
}

func (_c *MockUserRepository_FindActiveUserById_Call) Run(run func(ctx context.Context, id string)) *MockUserRepository_FindActiveUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_FindActiveUserById_Call) Return(_a0 *User, _a1 error) *MockUserRepository_FindActiveUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_FindActiveUserById_Call) RunAndReturn(run func(context.Context, string) (*User, error)) *MockUserRepository_FindActiveUserById_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllUser provides a mock function with given fields: ctx
func (_m *MockUserRepository) FindAllUser(ctx context.Context) ([]*User, error) {
	ret := _m.Called(ctx)

	var r0 []*User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_FindAllUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllUser'
type MockUserRepository_FindAllUser_Call struct {
	*mock.Call
}

// FindAllUser is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUserRepository_Expecter) FindAllUser(ctx interface{}) *MockUserRepository_FindAllUser_Call {
	return &MockUserRepository_FindAllUser_Call{Call: _e.mock.On("FindAllUser", ctx)}
}

func (_c *MockUserRepository_FindAllUser_Call) Run(run func(ctx context.Context)) *MockUserRepository_FindAllUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUserRepository_FindAllUser_Call) Return(_a0 []*User, _a1 error) *MockUserRepository_FindAllUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_FindAllUser_Call) RunAndReturn(run func(context.Context) ([]*User, error)) *MockUserRepository_FindAllUser_Call {
	_c.Call.Return(run)
	return _c
}

// FindUserById provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) FindUserById(ctx context.Context, id string) (*User, error) {
	ret := _m.Called(ctx, id)

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_FindUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserById'
type MockUserRepository_FindUserById_Call struct {
	*mock.Call
}

// FindUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockUserRepository_Expecter) FindUserById(ctx interface{}, id interface{}) *MockUserRepository_FindUserById_Call {
	return &MockUserRepository_FindUserById_Call{Call: _e.mock.On("FindUserById", ctx, id)}
}

func (_c *MockUserRepository_FindUserById_Call) Run(run func(ctx context.Context, id string)) *MockUserRepository_FindUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_FindUserById_Call) Return(_a0 *User, _a1 error) *MockUserRepository_FindUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_FindUserById_Call) RunAndReturn(run func(context.Context, string) (*User, error)) *MockUserRepository_FindUserById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserActive provides a mock function with given fields: ctx, id, active
func (_m *MockUserRepository) UpdateUserActive(ctx context.Context, id string, active bool) error {
	ret := _m.Called(ctx, id, active)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, id, active)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_UpdateUserActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserActive'
type MockUserRepository_UpdateUserActive_Call struct {
	*mock.Call
}

// UpdateUserActive is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - active bool
func (_e *MockUserRepository_Expecter) UpdateUserActive(ctx interface{}, id interface{}, active interface{}) *MockUserRepository_UpdateUserActive_Call {
	return &MockUserRepository_UpdateUserActive_Call{Call: _e.mock.On("UpdateUserActive", ctx, id, active)}
}

func (_c *MockUserRepository_UpdateUserActive_Call) Run(run func(ctx context.Context, id string, active bool)) *MockUserRepository_UpdateUserActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUserActive_Call) Return(_a0 error) *MockUserRepository_UpdateUserActive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_UpdateUserActive_Call) RunAndReturn(run func(context.Context, string, bool) error) *MockUserRepository_UpdateUserActive_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserInfo provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) UpdateUserInfo(ctx context.Context, user *User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_UpdateUserInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserInfo'
type MockUserRepository_UpdateUserInfo_Call struct {
	*mock.Call
}

// UpdateUserInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - user *User
func (_e *MockUserRepository_Expecter) UpdateUserInfo(ctx interface{}, user interface{}) *MockUserRepository_UpdateUserInfo_Call {
	return &MockUserRepository_UpdateUserInfo_Call{Call: _e.mock.On("UpdateUserInfo", ctx, user)}
}

func (_c *MockUserRepository_UpdateUserInfo_Call) Run(run func(ctx context.Context, user *User)) *MockUserRepository_UpdateUserInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*User))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUserInfo_Call) Return(_a0 error) *MockUserRepository_UpdateUserInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_UpdateUserInfo_Call) RunAndReturn(run func(context.Context, *User) error) *MockUserRepository_UpdateUserInfo_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserRole provides a mock function with given fields: ctx, userId, role
func (_m *MockUserRepository) UpdateUserRole(ctx context.Context, userId string, role Role) error {
	ret := _m.Called(ctx, userId, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, Role) error); ok {
		r0 = rf(ctx, userId, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_UpdateUserRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserRole'
type MockUserRepository_UpdateUserRole_Call struct {
	*mock.Call
}

// UpdateUserRole is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - role Role
func (_e *MockUserRepository_Expecter) UpdateUserRole(ctx interface{}, userId interface{}, role interface{}) *MockUserRepository_UpdateUserRole_Call {
	return &MockUserRepository_UpdateUserRole_Call{Call: _e.mock.On("UpdateUserRole", ctx, userId, role)}
}

func (_c *MockUserRepository_UpdateUserRole_Call) Run(run func(ctx context.Context, userId string, role Role)) *MockUserRepository_UpdateUserRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(Role))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUserRole_Call) Return(_a0 error) *MockUserRepository_UpdateUserRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_UpdateUserRole_Call) RunAndReturn(run func(context.Context, string, Role) error) *MockUserRepository_UpdateUserRole_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUserRepository(t mockConstructorTestingTNewMockUserRepository) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
