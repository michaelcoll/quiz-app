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

import mock "github.com/stretchr/testify/mock"

// MockMaintenanceRepository is an autogenerated mock type for the MaintenanceRepository type
type MockMaintenanceRepository struct {
	mock.Mock
}

type MockMaintenanceRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMaintenanceRepository) EXPECT() *MockMaintenanceRepository_Expecter {
	return &MockMaintenanceRepository_Expecter{mock: &_m.Mock}
}

// Dump provides a mock function with given fields:
func (_m *MockMaintenanceRepository) Dump() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMaintenanceRepository_Dump_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dump'
type MockMaintenanceRepository_Dump_Call struct {
	*mock.Call
}

// Dump is a helper method to define mock.On call
func (_e *MockMaintenanceRepository_Expecter) Dump() *MockMaintenanceRepository_Dump_Call {
	return &MockMaintenanceRepository_Dump_Call{Call: _e.mock.On("Dump")}
}

func (_c *MockMaintenanceRepository_Dump_Call) Run(run func()) *MockMaintenanceRepository_Dump_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMaintenanceRepository_Dump_Call) Return(_a0 string, _a1 error) *MockMaintenanceRepository_Dump_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMaintenanceRepository_Dump_Call) RunAndReturn(run func() (string, error)) *MockMaintenanceRepository_Dump_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockMaintenanceRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMaintenanceRepository creates a new instance of MockMaintenanceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMaintenanceRepository(t mockConstructorTestingTNewMockMaintenanceRepository) *MockMaintenanceRepository {
	mock := &MockMaintenanceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}