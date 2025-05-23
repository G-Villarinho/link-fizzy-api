// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/g-villarinho/link-fizz-api/models"
	mock "github.com/stretchr/testify/mock"
)

// LinkRepositoryMock is an autogenerated mock type for the LinkRepository type
type LinkRepositoryMock struct {
	mock.Mock
}

type LinkRepositoryMock_Expecter struct {
	mock *mock.Mock
}

func (_m *LinkRepositoryMock) EXPECT() *LinkRepositoryMock_Expecter {
	return &LinkRepositoryMock_Expecter{mock: &_m.Mock}
}

// CreateLink provides a mock function with given fields: ctx, link
func (_m *LinkRepositoryMock) CreateLink(ctx context.Context, link models.Link) error {
	ret := _m.Called(ctx, link)

	if len(ret) == 0 {
		panic("no return value specified for CreateLink")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Link) error); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkRepositoryMock_CreateLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLink'
type LinkRepositoryMock_CreateLink_Call struct {
	*mock.Call
}

// CreateLink is a helper method to define mock.On call
//   - ctx context.Context
//   - link models.Link
func (_e *LinkRepositoryMock_Expecter) CreateLink(ctx interface{}, link interface{}) *LinkRepositoryMock_CreateLink_Call {
	return &LinkRepositoryMock_CreateLink_Call{Call: _e.mock.On("CreateLink", ctx, link)}
}

func (_c *LinkRepositoryMock_CreateLink_Call) Run(run func(ctx context.Context, link models.Link)) *LinkRepositoryMock_CreateLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Link))
	})
	return _c
}

func (_c *LinkRepositoryMock_CreateLink_Call) Return(_a0 error) *LinkRepositoryMock_CreateLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LinkRepositoryMock_CreateLink_Call) RunAndReturn(run func(context.Context, models.Link) error) *LinkRepositoryMock_CreateLink_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllShortCodesByUserID provides a mock function with given fields: ctx, userID
func (_m *LinkRepositoryMock) GetAllShortCodesByUserID(ctx context.Context, userID string) ([]string, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllShortCodesByUserID")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkRepositoryMock_GetAllShortCodesByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllShortCodesByUserID'
type LinkRepositoryMock_GetAllShortCodesByUserID_Call struct {
	*mock.Call
}

// GetAllShortCodesByUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *LinkRepositoryMock_Expecter) GetAllShortCodesByUserID(ctx interface{}, userID interface{}) *LinkRepositoryMock_GetAllShortCodesByUserID_Call {
	return &LinkRepositoryMock_GetAllShortCodesByUserID_Call{Call: _e.mock.On("GetAllShortCodesByUserID", ctx, userID)}
}

func (_c *LinkRepositoryMock_GetAllShortCodesByUserID_Call) Run(run func(ctx context.Context, userID string)) *LinkRepositoryMock_GetAllShortCodesByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LinkRepositoryMock_GetAllShortCodesByUserID_Call) Return(_a0 []string, _a1 error) *LinkRepositoryMock_GetAllShortCodesByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LinkRepositoryMock_GetAllShortCodesByUserID_Call) RunAndReturn(run func(context.Context, string) ([]string, error)) *LinkRepositoryMock_GetAllShortCodesByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinkByID provides a mock function with given fields: ctx, ID
func (_m *LinkRepositoryMock) GetLinkByID(ctx context.Context, ID string) (*models.Link, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetLinkByID")
	}

	var r0 *models.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Link, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Link); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Link)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkRepositoryMock_GetLinkByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinkByID'
type LinkRepositoryMock_GetLinkByID_Call struct {
	*mock.Call
}

// GetLinkByID is a helper method to define mock.On call
//   - ctx context.Context
//   - ID string
func (_e *LinkRepositoryMock_Expecter) GetLinkByID(ctx interface{}, ID interface{}) *LinkRepositoryMock_GetLinkByID_Call {
	return &LinkRepositoryMock_GetLinkByID_Call{Call: _e.mock.On("GetLinkByID", ctx, ID)}
}

func (_c *LinkRepositoryMock_GetLinkByID_Call) Run(run func(ctx context.Context, ID string)) *LinkRepositoryMock_GetLinkByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LinkRepositoryMock_GetLinkByID_Call) Return(_a0 *models.Link, _a1 error) *LinkRepositoryMock_GetLinkByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LinkRepositoryMock_GetLinkByID_Call) RunAndReturn(run func(context.Context, string) (*models.Link, error)) *LinkRepositoryMock_GetLinkByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinkByShortCode provides a mock function with given fields: ctx, shortCode
func (_m *LinkRepositoryMock) GetLinkByShortCode(ctx context.Context, shortCode string) (*models.Link, error) {
	ret := _m.Called(ctx, shortCode)

	if len(ret) == 0 {
		panic("no return value specified for GetLinkByShortCode")
	}

	var r0 *models.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Link, error)); ok {
		return rf(ctx, shortCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Link); ok {
		r0 = rf(ctx, shortCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Link)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkRepositoryMock_GetLinkByShortCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinkByShortCode'
type LinkRepositoryMock_GetLinkByShortCode_Call struct {
	*mock.Call
}

// GetLinkByShortCode is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
func (_e *LinkRepositoryMock_Expecter) GetLinkByShortCode(ctx interface{}, shortCode interface{}) *LinkRepositoryMock_GetLinkByShortCode_Call {
	return &LinkRepositoryMock_GetLinkByShortCode_Call{Call: _e.mock.On("GetLinkByShortCode", ctx, shortCode)}
}

func (_c *LinkRepositoryMock_GetLinkByShortCode_Call) Run(run func(ctx context.Context, shortCode string)) *LinkRepositoryMock_GetLinkByShortCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LinkRepositoryMock_GetLinkByShortCode_Call) Return(_a0 *models.Link, _a1 error) *LinkRepositoryMock_GetLinkByShortCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LinkRepositoryMock_GetLinkByShortCode_Call) RunAndReturn(run func(context.Context, string) (*models.Link, error)) *LinkRepositoryMock_GetLinkByShortCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinksByUserID provides a mock function with given fields: ctx, userID
func (_m *LinkRepositoryMock) GetLinksByUserID(ctx context.Context, userID string) ([]models.Link, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetLinksByUserID")
	}

	var r0 []models.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Link, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Link); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Link)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkRepositoryMock_GetLinksByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinksByUserID'
type LinkRepositoryMock_GetLinksByUserID_Call struct {
	*mock.Call
}

// GetLinksByUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *LinkRepositoryMock_Expecter) GetLinksByUserID(ctx interface{}, userID interface{}) *LinkRepositoryMock_GetLinksByUserID_Call {
	return &LinkRepositoryMock_GetLinksByUserID_Call{Call: _e.mock.On("GetLinksByUserID", ctx, userID)}
}

func (_c *LinkRepositoryMock_GetLinksByUserID_Call) Run(run func(ctx context.Context, userID string)) *LinkRepositoryMock_GetLinksByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LinkRepositoryMock_GetLinksByUserID_Call) Return(_a0 []models.Link, _a1 error) *LinkRepositoryMock_GetLinksByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LinkRepositoryMock_GetLinksByUserID_Call) RunAndReturn(run func(context.Context, string) ([]models.Link, error)) *LinkRepositoryMock_GetLinksByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetOriginalURLByShortCode provides a mock function with given fields: ctx, shortCode
func (_m *LinkRepositoryMock) GetOriginalURLByShortCode(ctx context.Context, shortCode string) (string, error) {
	ret := _m.Called(ctx, shortCode)

	if len(ret) == 0 {
		panic("no return value specified for GetOriginalURLByShortCode")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, shortCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortCode)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkRepositoryMock_GetOriginalURLByShortCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOriginalURLByShortCode'
type LinkRepositoryMock_GetOriginalURLByShortCode_Call struct {
	*mock.Call
}

// GetOriginalURLByShortCode is a helper method to define mock.On call
//   - ctx context.Context
//   - shortCode string
func (_e *LinkRepositoryMock_Expecter) GetOriginalURLByShortCode(ctx interface{}, shortCode interface{}) *LinkRepositoryMock_GetOriginalURLByShortCode_Call {
	return &LinkRepositoryMock_GetOriginalURLByShortCode_Call{Call: _e.mock.On("GetOriginalURLByShortCode", ctx, shortCode)}
}

func (_c *LinkRepositoryMock_GetOriginalURLByShortCode_Call) Run(run func(ctx context.Context, shortCode string)) *LinkRepositoryMock_GetOriginalURLByShortCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LinkRepositoryMock_GetOriginalURLByShortCode_Call) Return(_a0 string, _a1 error) *LinkRepositoryMock_GetOriginalURLByShortCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LinkRepositoryMock_GetOriginalURLByShortCode_Call) RunAndReturn(run func(context.Context, string) (string, error)) *LinkRepositoryMock_GetOriginalURLByShortCode_Call {
	_c.Call.Return(run)
	return _c
}

// NewLinkRepositoryMock creates a new instance of LinkRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLinkRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *LinkRepositoryMock {
	mock := &LinkRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
