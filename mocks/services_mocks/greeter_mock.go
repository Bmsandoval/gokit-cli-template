// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bmsandoval/gokit-esque-cli-template/services/grpc_service (interfaces: GreeterClient)

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc_service "github.com/bmsandoval/gokit-esque-cli-template/services/grpc_service"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGreeterClient is a mock of GreeterClient interface
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterClientMockRecorder
}

// MockGreeterClientMockRecorder is the mock recorder for MockGreeterClient
type MockGreeterClientMockRecorder struct {
	mock *MockGreeterClient
}

// NewMockGreeterClient creates a new mock instance
func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &MockGreeterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGreeterClient) EXPECT() *MockGreeterClientMockRecorder {
	return m.recorder
}

// GetHello mocks base method
func (m *MockGreeterClient) GetHello(arg0 context.Context, arg1 *grpc_service.GetHelloRequest, arg2 ...grpc.CallOption) (*grpc_service.GetHelloReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHello", varargs...)
	ret0, _ := ret[0].(*grpc_service.GetHelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHello indicates an expected call of GetHello
func (mr *MockGreeterClientMockRecorder) GetHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHello", reflect.TypeOf((*MockGreeterClient)(nil).GetHello), varargs...)
}

// PostHello mocks base method
func (m *MockGreeterClient) PostHello(arg0 context.Context, arg1 *grpc_service.PostHelloRequest, arg2 ...grpc.CallOption) (*grpc_service.PostHelloReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostHello", varargs...)
	ret0, _ := ret[0].(*grpc_service.PostHelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostHello indicates an expected call of PostHello
func (mr *MockGreeterClientMockRecorder) PostHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHello", reflect.TypeOf((*MockGreeterClient)(nil).PostHello), varargs...)
}