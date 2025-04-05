// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/g-villarinho/link-fizz-api/models"
	mock "github.com/stretchr/testify/mock"
)

// AuthServiceMock is an autogenerated mock type for the AuthService type
type AuthServiceMock struct {
	mock.Mock
}

type AuthServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthServiceMock) EXPECT() *AuthServiceMock_Expecter {
	return &AuthServiceMock_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: ctx, email, password, ipAdress, userAgent
func (_m *AuthServiceMock) Login(ctx context.Context, email string, password string, ipAdress string, userAgent string) (*models.LoginResponse, error) {
	ret := _m.Called(ctx, email, password, ipAdress, userAgent)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *models.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (*models.LoginResponse, error)); ok {
		return rf(ctx, email, password, ipAdress, userAgent)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *models.LoginResponse); ok {
		r0 = rf(ctx, email, password, ipAdress, userAgent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, email, password, ipAdress, userAgent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthServiceMock_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AuthServiceMock_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
//   - password string
//   - ipAdress string
//   - userAgent string
func (_e *AuthServiceMock_Expecter) Login(ctx interface{}, email interface{}, password interface{}, ipAdress interface{}, userAgent interface{}) *AuthServiceMock_Login_Call {
	return &AuthServiceMock_Login_Call{Call: _e.mock.On("Login", ctx, email, password, ipAdress, userAgent)}
}

func (_c *AuthServiceMock_Login_Call) Run(run func(ctx context.Context, email string, password string, ipAdress string, userAgent string)) *AuthServiceMock_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *AuthServiceMock_Login_Call) Return(_a0 *models.LoginResponse, _a1 error) *AuthServiceMock_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthServiceMock_Login_Call) RunAndReturn(run func(context.Context, string, string, string, string) (*models.LoginResponse, error)) *AuthServiceMock_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Logout provides a mock function with given fields: ctx, userID, sessionID, token
func (_m *AuthServiceMock) Logout(ctx context.Context, userID string, sessionID string, token string) error {
	ret := _m.Called(ctx, userID, sessionID, token)

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, userID, sessionID, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthServiceMock_Logout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logout'
type AuthServiceMock_Logout_Call struct {
	*mock.Call
}

// Logout is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - sessionID string
//   - token string
func (_e *AuthServiceMock_Expecter) Logout(ctx interface{}, userID interface{}, sessionID interface{}, token interface{}) *AuthServiceMock_Logout_Call {
	return &AuthServiceMock_Logout_Call{Call: _e.mock.On("Logout", ctx, userID, sessionID, token)}
}

func (_c *AuthServiceMock_Logout_Call) Run(run func(ctx context.Context, userID string, sessionID string, token string)) *AuthServiceMock_Logout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *AuthServiceMock_Logout_Call) Return(_a0 error) *AuthServiceMock_Logout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthServiceMock_Logout_Call) RunAndReturn(run func(context.Context, string, string, string) error) *AuthServiceMock_Logout_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: ctx, name, email, password, ipAdress, userAgent
func (_m *AuthServiceMock) Register(ctx context.Context, name string, email string, password string, ipAdress string, userAgent string) (*models.LoginResponse, error) {
	ret := _m.Called(ctx, name, email, password, ipAdress, userAgent)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *models.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) (*models.LoginResponse, error)); ok {
		return rf(ctx, name, email, password, ipAdress, userAgent)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) *models.LoginResponse); ok {
		r0 = rf(ctx, name, email, password, ipAdress, userAgent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string) error); ok {
		r1 = rf(ctx, name, email, password, ipAdress, userAgent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthServiceMock_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type AuthServiceMock_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - email string
//   - password string
//   - ipAdress string
//   - userAgent string
func (_e *AuthServiceMock_Expecter) Register(ctx interface{}, name interface{}, email interface{}, password interface{}, ipAdress interface{}, userAgent interface{}) *AuthServiceMock_Register_Call {
	return &AuthServiceMock_Register_Call{Call: _e.mock.On("Register", ctx, name, email, password, ipAdress, userAgent)}
}

func (_c *AuthServiceMock_Register_Call) Run(run func(ctx context.Context, name string, email string, password string, ipAdress string, userAgent string)) *AuthServiceMock_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *AuthServiceMock_Register_Call) Return(_a0 *models.LoginResponse, _a1 error) *AuthServiceMock_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthServiceMock_Register_Call) RunAndReturn(run func(context.Context, string, string, string, string, string) (*models.LoginResponse, error)) *AuthServiceMock_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthServiceMock creates a new instance of AuthServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthServiceMock {
	mock := &AuthServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
