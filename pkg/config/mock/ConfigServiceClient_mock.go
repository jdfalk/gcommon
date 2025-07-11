// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mock

import (
	"context"

	proto "github.com/jdfalk/gcommon/pkg/config/proto"
	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// NewMockConfigServiceClient creates a new instance of MockConfigServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConfigServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConfigServiceClient {
	mock := &MockConfigServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockConfigServiceClient is an autogenerated mock type for the ConfigServiceClient type
type MockConfigServiceClient struct {
	mock.Mock
}

type MockConfigServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfigServiceClient) EXPECT() *MockConfigServiceClient_Expecter {
	return &MockConfigServiceClient_Expecter{mock: &_m.Mock}
}

// AllKeys provides a mock function for the type MockConfigServiceClient
func (_mock *MockConfigServiceClient) AllKeys(ctx context.Context, in *proto.AllKeysRequest, opts ...grpc.CallOption) (*proto.AllKeysResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for AllKeys")
	}

	var r0 *proto.AllKeysResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.AllKeysRequest, ...grpc.CallOption) (*proto.AllKeysResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.AllKeysRequest, ...grpc.CallOption) *proto.AllKeysResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.AllKeysResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.AllKeysRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockConfigServiceClient_AllKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllKeys'
type MockConfigServiceClient_AllKeys_Call struct {
	*mock.Call
}

// AllKeys is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockConfigServiceClient_Expecter) AllKeys(ctx interface{}, in interface{}, opts ...interface{}) *MockConfigServiceClient_AllKeys_Call {
	return &MockConfigServiceClient_AllKeys_Call{Call: _e.mock.On("AllKeys",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockConfigServiceClient_AllKeys_Call) Run(run func(ctx context.Context, in *proto.AllKeysRequest, opts ...grpc.CallOption)) *MockConfigServiceClient_AllKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.AllKeysRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockConfigServiceClient_AllKeys_Call) Return(allKeysResponse *proto.AllKeysResponse, err error) *MockConfigServiceClient_AllKeys_Call {
	_c.Call.Return(allKeysResponse, err)
	return _c
}

func (_c *MockConfigServiceClient_AllKeys_Call) RunAndReturn(run func(ctx context.Context, in *proto.AllKeysRequest, opts ...grpc.CallOption) (*proto.AllKeysResponse, error)) *MockConfigServiceClient_AllKeys_Call {
	_c.Call.Return(run)
	return _c
}

// AllSettings provides a mock function for the type MockConfigServiceClient
func (_mock *MockConfigServiceClient) AllSettings(ctx context.Context, in *proto.AllSettingsRequest, opts ...grpc.CallOption) (*proto.AllSettingsResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for AllSettings")
	}

	var r0 *proto.AllSettingsResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.AllSettingsRequest, ...grpc.CallOption) (*proto.AllSettingsResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.AllSettingsRequest, ...grpc.CallOption) *proto.AllSettingsResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.AllSettingsResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.AllSettingsRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockConfigServiceClient_AllSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllSettings'
type MockConfigServiceClient_AllSettings_Call struct {
	*mock.Call
}

// AllSettings is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockConfigServiceClient_Expecter) AllSettings(ctx interface{}, in interface{}, opts ...interface{}) *MockConfigServiceClient_AllSettings_Call {
	return &MockConfigServiceClient_AllSettings_Call{Call: _e.mock.On("AllSettings",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockConfigServiceClient_AllSettings_Call) Run(run func(ctx context.Context, in *proto.AllSettingsRequest, opts ...grpc.CallOption)) *MockConfigServiceClient_AllSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.AllSettingsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockConfigServiceClient_AllSettings_Call) Return(allSettingsResponse *proto.AllSettingsResponse, err error) *MockConfigServiceClient_AllSettings_Call {
	_c.Call.Return(allSettingsResponse, err)
	return _c
}

func (_c *MockConfigServiceClient_AllSettings_Call) RunAndReturn(run func(ctx context.Context, in *proto.AllSettingsRequest, opts ...grpc.CallOption) (*proto.AllSettingsResponse, error)) *MockConfigServiceClient_AllSettings_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function for the type MockConfigServiceClient
func (_mock *MockConfigServiceClient) Get(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption) (*proto.GetResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *proto.GetResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.GetRequest, ...grpc.CallOption) (*proto.GetResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.GetRequest, ...grpc.CallOption) *proto.GetResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.GetRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockConfigServiceClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockConfigServiceClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockConfigServiceClient_Expecter) Get(ctx interface{}, in interface{}, opts ...interface{}) *MockConfigServiceClient_Get_Call {
	return &MockConfigServiceClient_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockConfigServiceClient_Get_Call) Run(run func(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption)) *MockConfigServiceClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.GetRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockConfigServiceClient_Get_Call) Return(getResponse *proto.GetResponse, err error) *MockConfigServiceClient_Get_Call {
	_c.Call.Return(getResponse, err)
	return _c
}

func (_c *MockConfigServiceClient_Get_Call) RunAndReturn(run func(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption) (*proto.GetResponse, error)) *MockConfigServiceClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function for the type MockConfigServiceClient
func (_mock *MockConfigServiceClient) Has(ctx context.Context, in *proto.HasRequest, opts ...grpc.CallOption) (*proto.HasResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Has")
	}

	var r0 *proto.HasResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HasRequest, ...grpc.CallOption) (*proto.HasResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HasRequest, ...grpc.CallOption) *proto.HasResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.HasResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.HasRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockConfigServiceClient_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type MockConfigServiceClient_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockConfigServiceClient_Expecter) Has(ctx interface{}, in interface{}, opts ...interface{}) *MockConfigServiceClient_Has_Call {
	return &MockConfigServiceClient_Has_Call{Call: _e.mock.On("Has",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockConfigServiceClient_Has_Call) Run(run func(ctx context.Context, in *proto.HasRequest, opts ...grpc.CallOption)) *MockConfigServiceClient_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.HasRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockConfigServiceClient_Has_Call) Return(hasResponse *proto.HasResponse, err error) *MockConfigServiceClient_Has_Call {
	_c.Call.Return(hasResponse, err)
	return _c
}

func (_c *MockConfigServiceClient_Has_Call) RunAndReturn(run func(ctx context.Context, in *proto.HasRequest, opts ...grpc.CallOption) (*proto.HasResponse, error)) *MockConfigServiceClient_Has_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function for the type MockConfigServiceClient
func (_mock *MockConfigServiceClient) Set(ctx context.Context, in *proto.SetRequest, opts ...grpc.CallOption) (*proto.SetResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 *proto.SetResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.SetRequest, ...grpc.CallOption) (*proto.SetResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.SetRequest, ...grpc.CallOption) *proto.SetResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SetResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.SetRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockConfigServiceClient_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockConfigServiceClient_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockConfigServiceClient_Expecter) Set(ctx interface{}, in interface{}, opts ...interface{}) *MockConfigServiceClient_Set_Call {
	return &MockConfigServiceClient_Set_Call{Call: _e.mock.On("Set",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockConfigServiceClient_Set_Call) Run(run func(ctx context.Context, in *proto.SetRequest, opts ...grpc.CallOption)) *MockConfigServiceClient_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.SetRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockConfigServiceClient_Set_Call) Return(setResponse *proto.SetResponse, err error) *MockConfigServiceClient_Set_Call {
	_c.Call.Return(setResponse, err)
	return _c
}

func (_c *MockConfigServiceClient_Set_Call) RunAndReturn(run func(ctx context.Context, in *proto.SetRequest, opts ...grpc.CallOption) (*proto.SetResponse, error)) *MockConfigServiceClient_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function for the type MockConfigServiceClient
func (_mock *MockConfigServiceClient) Watch(ctx context.Context, in *proto.WatchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[proto.WatchResponse], error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Watch")
	}

	var r0 grpc.ServerStreamingClient[proto.WatchResponse]
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.WatchRequest, ...grpc.CallOption) (grpc.ServerStreamingClient[proto.WatchResponse], error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.WatchRequest, ...grpc.CallOption) grpc.ServerStreamingClient[proto.WatchResponse]); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.ServerStreamingClient[proto.WatchResponse])
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.WatchRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockConfigServiceClient_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockConfigServiceClient_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockConfigServiceClient_Expecter) Watch(ctx interface{}, in interface{}, opts ...interface{}) *MockConfigServiceClient_Watch_Call {
	return &MockConfigServiceClient_Watch_Call{Call: _e.mock.On("Watch",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockConfigServiceClient_Watch_Call) Run(run func(ctx context.Context, in *proto.WatchRequest, opts ...grpc.CallOption)) *MockConfigServiceClient_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.WatchRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockConfigServiceClient_Watch_Call) Return(serverStreamingClient grpc.ServerStreamingClient[proto.WatchResponse], err error) *MockConfigServiceClient_Watch_Call {
	_c.Call.Return(serverStreamingClient, err)
	return _c
}

func (_c *MockConfigServiceClient_Watch_Call) RunAndReturn(run func(ctx context.Context, in *proto.WatchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[proto.WatchResponse], error)) *MockConfigServiceClient_Watch_Call {
	_c.Call.Return(run)
	return _c
}
