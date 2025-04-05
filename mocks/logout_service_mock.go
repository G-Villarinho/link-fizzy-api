// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// LogoutServiceMock is an autogenerated mock type for the LogoutService type
type LogoutServiceMock struct {
	mock.Mock
}

type LogoutServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *LogoutServiceMock) EXPECT() *LogoutServiceMock_Expecter {
	return &LogoutServiceMock_Expecter{mock: &_m.Mock}
}

// CreateLogout provides a mock function with given fields: ctx, userID, token
func (_m *LogoutServiceMock) CreateLogout(ctx context.Context, userID *string, token string) error {
	ret := _m.Called(ctx, userID, token)

	if len(ret) == 0 {
		panic("no return value specified for CreateLogout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *string, string) error); ok {
		r0 = rf(ctx, userID, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogoutServiceMock_CreateLogout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLogout'
type LogoutServiceMock_CreateLogout_Call struct {
	*mock.Call
}

// CreateLogout is a helper method to define mock.On call
//   - ctx context.Context
//   - userID *string
//   - token string
func (_e *LogoutServiceMock_Expecter) CreateLogout(ctx interface{}, userID interface{}, token interface{}) *LogoutServiceMock_CreateLogout_Call {
	return &LogoutServiceMock_CreateLogout_Call{Call: _e.mock.On("CreateLogout", ctx, userID, token)}
}

func (_c *LogoutServiceMock_CreateLogout_Call) Run(run func(ctx context.Context, userID *string, token string)) *LogoutServiceMock_CreateLogout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*string), args[2].(string))
	})
	return _c
}

func (_c *LogoutServiceMock_CreateLogout_Call) Return(_a0 error) *LogoutServiceMock_CreateLogout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LogoutServiceMock_CreateLogout_Call) RunAndReturn(run func(context.Context, *string, string) error) *LogoutServiceMock_CreateLogout_Call {
	_c.Call.Return(run)
	return _c
}

// IsLogoutRegistered provides a mock function with given fields: ctx, token
func (_m *LogoutServiceMock) IsLogoutRegistered(ctx context.Context, token string) (bool, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for IsLogoutRegistered")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogoutServiceMock_IsLogoutRegistered_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsLogoutRegistered'
type LogoutServiceMock_IsLogoutRegistered_Call struct {
	*mock.Call
}

// IsLogoutRegistered is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *LogoutServiceMock_Expecter) IsLogoutRegistered(ctx interface{}, token interface{}) *LogoutServiceMock_IsLogoutRegistered_Call {
	return &LogoutServiceMock_IsLogoutRegistered_Call{Call: _e.mock.On("IsLogoutRegistered", ctx, token)}
}

func (_c *LogoutServiceMock_IsLogoutRegistered_Call) Run(run func(ctx context.Context, token string)) *LogoutServiceMock_IsLogoutRegistered_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LogoutServiceMock_IsLogoutRegistered_Call) Return(_a0 bool, _a1 error) *LogoutServiceMock_IsLogoutRegistered_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LogoutServiceMock_IsLogoutRegistered_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *LogoutServiceMock_IsLogoutRegistered_Call {
	_c.Call.Return(run)
	return _c
}

// NewLogoutServiceMock creates a new instance of LogoutServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogoutServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogoutServiceMock {
	mock := &LogoutServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
