// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mock

import (
	"context"

	proto "github.com/jdfalk/gcommon/pkg/health/proto"
	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// NewMockHealthServiceClient creates a new instance of MockHealthServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHealthServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHealthServiceClient {
	mock := &MockHealthServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockHealthServiceClient is an autogenerated mock type for the HealthServiceClient type
type MockHealthServiceClient struct {
	mock.Mock
}

type MockHealthServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHealthServiceClient) EXPECT() *MockHealthServiceClient_Expecter {
	return &MockHealthServiceClient_Expecter{mock: &_m.Mock}
}

// Check provides a mock function for the type MockHealthServiceClient
func (_mock *MockHealthServiceClient) Check(ctx context.Context, in *proto.HealthCheckRequest, opts ...grpc.CallOption) (*proto.HealthCheckResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Check")
	}

	var r0 *proto.HealthCheckResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HealthCheckRequest, ...grpc.CallOption) (*proto.HealthCheckResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HealthCheckRequest, ...grpc.CallOption) *proto.HealthCheckResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.HealthCheckResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.HealthCheckRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockHealthServiceClient_Check_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Check'
type MockHealthServiceClient_Check_Call struct {
	*mock.Call
}

// Check is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockHealthServiceClient_Expecter) Check(ctx interface{}, in interface{}, opts ...interface{}) *MockHealthServiceClient_Check_Call {
	return &MockHealthServiceClient_Check_Call{Call: _e.mock.On("Check",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockHealthServiceClient_Check_Call) Run(run func(ctx context.Context, in *proto.HealthCheckRequest, opts ...grpc.CallOption)) *MockHealthServiceClient_Check_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.HealthCheckRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockHealthServiceClient_Check_Call) Return(healthCheckResponse *proto.HealthCheckResponse, err error) *MockHealthServiceClient_Check_Call {
	_c.Call.Return(healthCheckResponse, err)
	return _c
}

func (_c *MockHealthServiceClient_Check_Call) RunAndReturn(run func(ctx context.Context, in *proto.HealthCheckRequest, opts ...grpc.CallOption) (*proto.HealthCheckResponse, error)) *MockHealthServiceClient_Check_Call {
	_c.Call.Return(run)
	return _c
}

// CheckAll provides a mock function for the type MockHealthServiceClient
func (_mock *MockHealthServiceClient) CheckAll(ctx context.Context, in *proto.HealthCheckAllRequest, opts ...grpc.CallOption) (*proto.HealthCheckAllResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for CheckAll")
	}

	var r0 *proto.HealthCheckAllResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HealthCheckAllRequest, ...grpc.CallOption) (*proto.HealthCheckAllResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HealthCheckAllRequest, ...grpc.CallOption) *proto.HealthCheckAllResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.HealthCheckAllResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.HealthCheckAllRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockHealthServiceClient_CheckAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckAll'
type MockHealthServiceClient_CheckAll_Call struct {
	*mock.Call
}

// CheckAll is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockHealthServiceClient_Expecter) CheckAll(ctx interface{}, in interface{}, opts ...interface{}) *MockHealthServiceClient_CheckAll_Call {
	return &MockHealthServiceClient_CheckAll_Call{Call: _e.mock.On("CheckAll",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockHealthServiceClient_CheckAll_Call) Run(run func(ctx context.Context, in *proto.HealthCheckAllRequest, opts ...grpc.CallOption)) *MockHealthServiceClient_CheckAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.HealthCheckAllRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockHealthServiceClient_CheckAll_Call) Return(healthCheckAllResponse *proto.HealthCheckAllResponse, err error) *MockHealthServiceClient_CheckAll_Call {
	_c.Call.Return(healthCheckAllResponse, err)
	return _c
}

func (_c *MockHealthServiceClient_CheckAll_Call) RunAndReturn(run func(ctx context.Context, in *proto.HealthCheckAllRequest, opts ...grpc.CallOption) (*proto.HealthCheckAllResponse, error)) *MockHealthServiceClient_CheckAll_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function for the type MockHealthServiceClient
func (_mock *MockHealthServiceClient) Watch(ctx context.Context, in *proto.HealthCheckRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[proto.HealthCheckResponse], error) {
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

	var r0 grpc.ServerStreamingClient[proto.HealthCheckResponse]
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HealthCheckRequest, ...grpc.CallOption) (grpc.ServerStreamingClient[proto.HealthCheckResponse], error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.HealthCheckRequest, ...grpc.CallOption) grpc.ServerStreamingClient[proto.HealthCheckResponse]); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.ServerStreamingClient[proto.HealthCheckResponse])
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.HealthCheckRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockHealthServiceClient_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockHealthServiceClient_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockHealthServiceClient_Expecter) Watch(ctx interface{}, in interface{}, opts ...interface{}) *MockHealthServiceClient_Watch_Call {
	return &MockHealthServiceClient_Watch_Call{Call: _e.mock.On("Watch",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockHealthServiceClient_Watch_Call) Run(run func(ctx context.Context, in *proto.HealthCheckRequest, opts ...grpc.CallOption)) *MockHealthServiceClient_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.HealthCheckRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockHealthServiceClient_Watch_Call) Return(serverStreamingClient grpc.ServerStreamingClient[proto.HealthCheckResponse], err error) *MockHealthServiceClient_Watch_Call {
	_c.Call.Return(serverStreamingClient, err)
	return _c
}

func (_c *MockHealthServiceClient_Watch_Call) RunAndReturn(run func(ctx context.Context, in *proto.HealthCheckRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[proto.HealthCheckResponse], error)) *MockHealthServiceClient_Watch_Call {
	_c.Call.Return(run)
	return _c
}
