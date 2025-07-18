// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mock

import (
	"context"

	proto "github.com/jdfalk/gcommon/pkg/cache/proto"
	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// NewMockCacheServiceClient creates a new instance of MockCacheServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCacheServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCacheServiceClient {
	mock := &MockCacheServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockCacheServiceClient is an autogenerated mock type for the CacheServiceClient type
type MockCacheServiceClient struct {
	mock.Mock
}

type MockCacheServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCacheServiceClient) EXPECT() *MockCacheServiceClient_Expecter {
	return &MockCacheServiceClient_Expecter{mock: &_m.Mock}
}

// Clear provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Clear(ctx context.Context, in *proto.ClearRequest, opts ...grpc.CallOption) (*proto.ClearResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 *proto.ClearResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.ClearRequest, ...grpc.CallOption) (*proto.ClearResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.ClearRequest, ...grpc.CallOption) *proto.ClearResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ClearResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.ClearRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type MockCacheServiceClient_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Clear(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Clear_Call {
	return &MockCacheServiceClient_Clear_Call{Call: _e.mock.On("Clear",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Clear_Call) Run(run func(ctx context.Context, in *proto.ClearRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.ClearRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Clear_Call) Return(clearResponse *proto.ClearResponse, err error) *MockCacheServiceClient_Clear_Call {
	_c.Call.Return(clearResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Clear_Call) RunAndReturn(run func(ctx context.Context, in *proto.ClearRequest, opts ...grpc.CallOption) (*proto.ClearResponse, error)) *MockCacheServiceClient_Clear_Call {
	_c.Call.Return(run)
	return _c
}

// Decrement provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Decrement(ctx context.Context, in *proto.DecrementRequest, opts ...grpc.CallOption) (*proto.DecrementResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Decrement")
	}

	var r0 *proto.DecrementResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.DecrementRequest, ...grpc.CallOption) (*proto.DecrementResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.DecrementRequest, ...grpc.CallOption) *proto.DecrementResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.DecrementResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.DecrementRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_Decrement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrement'
type MockCacheServiceClient_Decrement_Call struct {
	*mock.Call
}

// Decrement is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Decrement(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Decrement_Call {
	return &MockCacheServiceClient_Decrement_Call{Call: _e.mock.On("Decrement",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Decrement_Call) Run(run func(ctx context.Context, in *proto.DecrementRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Decrement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.DecrementRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Decrement_Call) Return(decrementResponse *proto.DecrementResponse, err error) *MockCacheServiceClient_Decrement_Call {
	_c.Call.Return(decrementResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Decrement_Call) RunAndReturn(run func(ctx context.Context, in *proto.DecrementRequest, opts ...grpc.CallOption) (*proto.DecrementResponse, error)) *MockCacheServiceClient_Decrement_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Delete(ctx context.Context, in *proto.DeleteRequest, opts ...grpc.CallOption) (*proto.DeleteResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *proto.DeleteResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.DeleteRequest, ...grpc.CallOption) (*proto.DeleteResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.DeleteRequest, ...grpc.CallOption) *proto.DeleteResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.DeleteResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.DeleteRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCacheServiceClient_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Delete(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Delete_Call {
	return &MockCacheServiceClient_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Delete_Call) Run(run func(ctx context.Context, in *proto.DeleteRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.DeleteRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Delete_Call) Return(deleteResponse *proto.DeleteResponse, err error) *MockCacheServiceClient_Delete_Call {
	_c.Call.Return(deleteResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Delete_Call) RunAndReturn(run func(ctx context.Context, in *proto.DeleteRequest, opts ...grpc.CallOption) (*proto.DeleteResponse, error)) *MockCacheServiceClient_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteMulti provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) DeleteMulti(ctx context.Context, in *proto.DeleteMultiRequest, opts ...grpc.CallOption) (*proto.DeleteMultiResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for DeleteMulti")
	}

	var r0 *proto.DeleteMultiResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.DeleteMultiRequest, ...grpc.CallOption) (*proto.DeleteMultiResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.DeleteMultiRequest, ...grpc.CallOption) *proto.DeleteMultiResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.DeleteMultiResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.DeleteMultiRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_DeleteMulti_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteMulti'
type MockCacheServiceClient_DeleteMulti_Call struct {
	*mock.Call
}

// DeleteMulti is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) DeleteMulti(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_DeleteMulti_Call {
	return &MockCacheServiceClient_DeleteMulti_Call{Call: _e.mock.On("DeleteMulti",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_DeleteMulti_Call) Run(run func(ctx context.Context, in *proto.DeleteMultiRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_DeleteMulti_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.DeleteMultiRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_DeleteMulti_Call) Return(deleteMultiResponse *proto.DeleteMultiResponse, err error) *MockCacheServiceClient_DeleteMulti_Call {
	_c.Call.Return(deleteMultiResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_DeleteMulti_Call) RunAndReturn(run func(ctx context.Context, in *proto.DeleteMultiRequest, opts ...grpc.CallOption) (*proto.DeleteMultiResponse, error)) *MockCacheServiceClient_DeleteMulti_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Exists(ctx context.Context, in *proto.ExistsRequest, opts ...grpc.CallOption) (*proto.ExistsResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 *proto.ExistsResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.ExistsRequest, ...grpc.CallOption) (*proto.ExistsResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.ExistsRequest, ...grpc.CallOption) *proto.ExistsResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ExistsResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.ExistsRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockCacheServiceClient_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Exists(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Exists_Call {
	return &MockCacheServiceClient_Exists_Call{Call: _e.mock.On("Exists",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Exists_Call) Run(run func(ctx context.Context, in *proto.ExistsRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.ExistsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Exists_Call) Return(existsResponse *proto.ExistsResponse, err error) *MockCacheServiceClient_Exists_Call {
	_c.Call.Return(existsResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Exists_Call) RunAndReturn(run func(ctx context.Context, in *proto.ExistsRequest, opts ...grpc.CallOption) (*proto.ExistsResponse, error)) *MockCacheServiceClient_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Get(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption) (*proto.GetResponse, error) {
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

// MockCacheServiceClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCacheServiceClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Get(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Get_Call {
	return &MockCacheServiceClient_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Get_Call) Run(run func(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.GetRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Get_Call) Return(getResponse *proto.GetResponse, err error) *MockCacheServiceClient_Get_Call {
	_c.Call.Return(getResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Get_Call) RunAndReturn(run func(ctx context.Context, in *proto.GetRequest, opts ...grpc.CallOption) (*proto.GetResponse, error)) *MockCacheServiceClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetMulti provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) GetMulti(ctx context.Context, in *proto.GetMultiRequest, opts ...grpc.CallOption) (*proto.GetMultiResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for GetMulti")
	}

	var r0 *proto.GetMultiResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.GetMultiRequest, ...grpc.CallOption) (*proto.GetMultiResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.GetMultiRequest, ...grpc.CallOption) *proto.GetMultiResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetMultiResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.GetMultiRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_GetMulti_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMulti'
type MockCacheServiceClient_GetMulti_Call struct {
	*mock.Call
}

// GetMulti is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) GetMulti(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_GetMulti_Call {
	return &MockCacheServiceClient_GetMulti_Call{Call: _e.mock.On("GetMulti",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_GetMulti_Call) Run(run func(ctx context.Context, in *proto.GetMultiRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_GetMulti_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.GetMultiRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_GetMulti_Call) Return(getMultiResponse *proto.GetMultiResponse, err error) *MockCacheServiceClient_GetMulti_Call {
	_c.Call.Return(getMultiResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_GetMulti_Call) RunAndReturn(run func(ctx context.Context, in *proto.GetMultiRequest, opts ...grpc.CallOption) (*proto.GetMultiResponse, error)) *MockCacheServiceClient_GetMulti_Call {
	_c.Call.Return(run)
	return _c
}

// Increment provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Increment(ctx context.Context, in *proto.IncrementRequest, opts ...grpc.CallOption) (*proto.IncrementResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Increment")
	}

	var r0 *proto.IncrementResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.IncrementRequest, ...grpc.CallOption) (*proto.IncrementResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.IncrementRequest, ...grpc.CallOption) *proto.IncrementResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.IncrementResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.IncrementRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_Increment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Increment'
type MockCacheServiceClient_Increment_Call struct {
	*mock.Call
}

// Increment is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Increment(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Increment_Call {
	return &MockCacheServiceClient_Increment_Call{Call: _e.mock.On("Increment",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Increment_Call) Run(run func(ctx context.Context, in *proto.IncrementRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Increment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.IncrementRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Increment_Call) Return(incrementResponse *proto.IncrementResponse, err error) *MockCacheServiceClient_Increment_Call {
	_c.Call.Return(incrementResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Increment_Call) RunAndReturn(run func(ctx context.Context, in *proto.IncrementRequest, opts ...grpc.CallOption) (*proto.IncrementResponse, error)) *MockCacheServiceClient_Increment_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Set(ctx context.Context, in *proto.SetRequest, opts ...grpc.CallOption) (*proto.SetResponse, error) {
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

// MockCacheServiceClient_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockCacheServiceClient_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Set(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Set_Call {
	return &MockCacheServiceClient_Set_Call{Call: _e.mock.On("Set",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Set_Call) Run(run func(ctx context.Context, in *proto.SetRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.SetRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Set_Call) Return(setResponse *proto.SetResponse, err error) *MockCacheServiceClient_Set_Call {
	_c.Call.Return(setResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Set_Call) RunAndReturn(run func(ctx context.Context, in *proto.SetRequest, opts ...grpc.CallOption) (*proto.SetResponse, error)) *MockCacheServiceClient_Set_Call {
	_c.Call.Return(run)
	return _c
}

// SetMulti provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) SetMulti(ctx context.Context, in *proto.SetMultiRequest, opts ...grpc.CallOption) (*proto.SetMultiResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for SetMulti")
	}

	var r0 *proto.SetMultiResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.SetMultiRequest, ...grpc.CallOption) (*proto.SetMultiResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.SetMultiRequest, ...grpc.CallOption) *proto.SetMultiResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SetMultiResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.SetMultiRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_SetMulti_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMulti'
type MockCacheServiceClient_SetMulti_Call struct {
	*mock.Call
}

// SetMulti is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) SetMulti(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_SetMulti_Call {
	return &MockCacheServiceClient_SetMulti_Call{Call: _e.mock.On("SetMulti",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_SetMulti_Call) Run(run func(ctx context.Context, in *proto.SetMultiRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_SetMulti_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.SetMultiRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_SetMulti_Call) Return(setMultiResponse *proto.SetMultiResponse, err error) *MockCacheServiceClient_SetMulti_Call {
	_c.Call.Return(setMultiResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_SetMulti_Call) RunAndReturn(run func(ctx context.Context, in *proto.SetMultiRequest, opts ...grpc.CallOption) (*proto.SetMultiResponse, error)) *MockCacheServiceClient_SetMulti_Call {
	_c.Call.Return(run)
	return _c
}

// Stats provides a mock function for the type MockCacheServiceClient
func (_mock *MockCacheServiceClient) Stats(ctx context.Context, in *proto.StatsRequest, opts ...grpc.CallOption) (*proto.StatsResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Stats")
	}

	var r0 *proto.StatsResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.StatsRequest, ...grpc.CallOption) (*proto.StatsResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *proto.StatsRequest, ...grpc.CallOption) *proto.StatsResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.StatsResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *proto.StatsRequest, ...grpc.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCacheServiceClient_Stats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stats'
type MockCacheServiceClient_Stats_Call struct {
	*mock.Call
}

// Stats is a helper method to define mock.On call
//   - ctx
//   - in
//   - opts
func (_e *MockCacheServiceClient_Expecter) Stats(ctx interface{}, in interface{}, opts ...interface{}) *MockCacheServiceClient_Stats_Call {
	return &MockCacheServiceClient_Stats_Call{Call: _e.mock.On("Stats",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockCacheServiceClient_Stats_Call) Run(run func(ctx context.Context, in *proto.StatsRequest, opts ...grpc.CallOption)) *MockCacheServiceClient_Stats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[2].([]grpc.CallOption)
		run(args[0].(context.Context), args[1].(*proto.StatsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCacheServiceClient_Stats_Call) Return(statsResponse *proto.StatsResponse, err error) *MockCacheServiceClient_Stats_Call {
	_c.Call.Return(statsResponse, err)
	return _c
}

func (_c *MockCacheServiceClient_Stats_Call) RunAndReturn(run func(ctx context.Context, in *proto.StatsRequest, opts ...grpc.CallOption) (*proto.StatsResponse, error)) *MockCacheServiceClient_Stats_Call {
	_c.Call.Return(run)
	return _c
}
