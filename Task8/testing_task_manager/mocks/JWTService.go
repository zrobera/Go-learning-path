// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// JWTService is an autogenerated mock type for the JWTService type
type JWTService struct {
	mock.Mock
}

type JWTService_Expecter struct {
	mock *mock.Mock
}

func (_m *JWTService) EXPECT() *JWTService_Expecter {
	return &JWTService_Expecter{mock: &_m.Mock}
}

// GenerateToken provides a mock function with given fields: username, role
func (_m *JWTService) GenerateToken(username string, role string) (string, error) {
	ret := _m.Called(username, role)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(username, role)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JWTService_GenerateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateToken'
type JWTService_GenerateToken_Call struct {
	*mock.Call
}

// GenerateToken is a helper method to define mock.On call
//   - username string
//   - role string
func (_e *JWTService_Expecter) GenerateToken(username interface{}, role interface{}) *JWTService_GenerateToken_Call {
	return &JWTService_GenerateToken_Call{Call: _e.mock.On("GenerateToken", username, role)}
}

func (_c *JWTService_GenerateToken_Call) Run(run func(username string, role string)) *JWTService_GenerateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *JWTService_GenerateToken_Call) Return(_a0 string, _a1 error) *JWTService_GenerateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *JWTService_GenerateToken_Call) RunAndReturn(run func(string, string) (string, error)) *JWTService_GenerateToken_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateToken provides a mock function with given fields: token
func (_m *JWTService) ValidateToken(token string) (map[string]interface{}, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateToken")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (map[string]interface{}, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JWTService_ValidateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateToken'
type JWTService_ValidateToken_Call struct {
	*mock.Call
}

// ValidateToken is a helper method to define mock.On call
//   - token string
func (_e *JWTService_Expecter) ValidateToken(token interface{}) *JWTService_ValidateToken_Call {
	return &JWTService_ValidateToken_Call{Call: _e.mock.On("ValidateToken", token)}
}

func (_c *JWTService_ValidateToken_Call) Run(run func(token string)) *JWTService_ValidateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *JWTService_ValidateToken_Call) Return(_a0 map[string]interface{}, _a1 error) *JWTService_ValidateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *JWTService_ValidateToken_Call) RunAndReturn(run func(string) (map[string]interface{}, error)) *JWTService_ValidateToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewJWTService creates a new instance of JWTService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTService(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTService {
	mock := &JWTService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
