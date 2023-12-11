/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go

// Package mock_scalesetvms is a generated GoMock package.
package mock_scalesetvms

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	gomock "go.uber.org/mock/gomock"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// Mockclient is a mock of client interface.
type Mockclient struct {
	ctrl     *gomock.Controller
	recorder *MockclientMockRecorder
}

// MockclientMockRecorder is the mock recorder for Mockclient.
type MockclientMockRecorder struct {
	mock *Mockclient
}

// NewMockclient creates a new mock instance.
func NewMockclient(ctrl *gomock.Controller) *Mockclient {
	mock := &Mockclient{ctrl: ctrl}
	mock.recorder = &MockclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockclient) EXPECT() *MockclientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAsync mocks base method.
func (m *Mockclient) CreateOrUpdateAsync(arg0 context.Context, arg1 azure.ResourceSpecGetter, arg2 string, arg3 interface{}) (interface{}, *runtime.Poller[armcompute.VirtualMachineScaleSetVMsClientUpdateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAsync", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*runtime.Poller[armcompute.VirtualMachineScaleSetVMsClientUpdateResponse])
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateOrUpdateAsync indicates an expected call of CreateOrUpdateAsync.
func (mr *MockclientMockRecorder) CreateOrUpdateAsync(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAsync", reflect.TypeOf((*Mockclient)(nil).CreateOrUpdateAsync), arg0, arg1, arg2, arg3)
}

// DeleteAsync mocks base method.
func (m *Mockclient) DeleteAsync(arg0 context.Context, arg1 azure.ResourceSpecGetter, arg2 string) (*runtime.Poller[armcompute.VirtualMachineScaleSetVMsClientDeleteResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsync", arg0, arg1, arg2)
	ret0, _ := ret[0].(*runtime.Poller[armcompute.VirtualMachineScaleSetVMsClientDeleteResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAsync indicates an expected call of DeleteAsync.
func (mr *MockclientMockRecorder) DeleteAsync(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsync", reflect.TypeOf((*Mockclient)(nil).DeleteAsync), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *Mockclient) Get(arg0 context.Context, arg1 azure.ResourceSpecGetter) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockclientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mockclient)(nil).Get), arg0, arg1)
}