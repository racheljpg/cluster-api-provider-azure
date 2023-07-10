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
// Source: ../scalesetvms.go

// Package mock_scalesetvms is a generated GoMock package.
package mock_scalesetvms

import (
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockScaleSetVMScope is a mock of ScaleSetVMScope interface.
type MockScaleSetVMScope struct {
	ctrl     *gomock.Controller
	recorder *MockScaleSetVMScopeMockRecorder
}

// MockScaleSetVMScopeMockRecorder is the mock recorder for MockScaleSetVMScope.
type MockScaleSetVMScopeMockRecorder struct {
	mock *MockScaleSetVMScope
}

// NewMockScaleSetVMScope creates a new mock instance.
func NewMockScaleSetVMScope(ctrl *gomock.Controller) *MockScaleSetVMScope {
	mock := &MockScaleSetVMScope{ctrl: ctrl}
	mock.recorder = &MockScaleSetVMScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScaleSetVMScope) EXPECT() *MockScaleSetVMScopeMockRecorder {
	return m.recorder
}

// AdditionalTags mocks base method.
func (m *MockScaleSetVMScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockScaleSetVMScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockScaleSetVMScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockScaleSetVMScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockScaleSetVMScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockScaleSetVMScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockScaleSetVMScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockScaleSetVMScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockScaleSetVMScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockScaleSetVMScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockScaleSetVMScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockScaleSetVMScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockScaleSetVMScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockScaleSetVMScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockScaleSetVMScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockScaleSetVMScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockScaleSetVMScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockScaleSetVMScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockScaleSetVMScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockScaleSetVMScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockScaleSetVMScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockScaleSetVMScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockScaleSetVMScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockScaleSetVMScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockScaleSetVMScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockScaleSetVMScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockScaleSetVMScope)(nil).ClusterName))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockScaleSetVMScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockScaleSetVMScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockScaleSetVMScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// ExtendedLocation mocks base method.
func (m *MockScaleSetVMScope) ExtendedLocation() *v1beta1.ExtendedLocationSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocation")
	ret0, _ := ret[0].(*v1beta1.ExtendedLocationSpec)
	return ret0
}

// ExtendedLocation indicates an expected call of ExtendedLocation.
func (mr *MockScaleSetVMScopeMockRecorder) ExtendedLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocation", reflect.TypeOf((*MockScaleSetVMScope)(nil).ExtendedLocation))
}

// ExtendedLocationName mocks base method.
func (m *MockScaleSetVMScope) ExtendedLocationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationName indicates an expected call of ExtendedLocationName.
func (mr *MockScaleSetVMScopeMockRecorder) ExtendedLocationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationName", reflect.TypeOf((*MockScaleSetVMScope)(nil).ExtendedLocationName))
}

// ExtendedLocationType mocks base method.
func (m *MockScaleSetVMScope) ExtendedLocationType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationType indicates an expected call of ExtendedLocationType.
func (mr *MockScaleSetVMScopeMockRecorder) ExtendedLocationType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationType", reflect.TypeOf((*MockScaleSetVMScope)(nil).ExtendedLocationType))
}

// FailureDomains mocks base method.
func (m *MockScaleSetVMScope) FailureDomains() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockScaleSetVMScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockScaleSetVMScope)(nil).FailureDomains))
}

// GetLongRunningOperationState mocks base method.
func (m *MockScaleSetVMScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockScaleSetVMScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockScaleSetVMScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// HashKey mocks base method.
func (m *MockScaleSetVMScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockScaleSetVMScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockScaleSetVMScope)(nil).HashKey))
}

// InstanceID mocks base method.
func (m *MockScaleSetVMScope) InstanceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceID indicates an expected call of InstanceID.
func (mr *MockScaleSetVMScopeMockRecorder) InstanceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockScaleSetVMScope)(nil).InstanceID))
}

// Location mocks base method.
func (m *MockScaleSetVMScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockScaleSetVMScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockScaleSetVMScope)(nil).Location))
}

// OrchestrationMode mocks base method.
func (m *MockScaleSetVMScope) OrchestrationMode() v1beta1.OrchestrationModeType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrchestrationMode")
	ret0, _ := ret[0].(v1beta1.OrchestrationModeType)
	return ret0
}

// OrchestrationMode indicates an expected call of OrchestrationMode.
func (mr *MockScaleSetVMScopeMockRecorder) OrchestrationMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrchestrationMode", reflect.TypeOf((*MockScaleSetVMScope)(nil).OrchestrationMode))
}

// ProviderID mocks base method.
func (m *MockScaleSetVMScope) ProviderID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProviderID indicates an expected call of ProviderID.
func (mr *MockScaleSetVMScopeMockRecorder) ProviderID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderID", reflect.TypeOf((*MockScaleSetVMScope)(nil).ProviderID))
}

// ResourceGroup mocks base method.
func (m *MockScaleSetVMScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockScaleSetVMScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockScaleSetVMScope)(nil).ResourceGroup))
}

// ScaleSetName mocks base method.
func (m *MockScaleSetVMScope) ScaleSetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleSetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ScaleSetName indicates an expected call of ScaleSetName.
func (mr *MockScaleSetVMScopeMockRecorder) ScaleSetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleSetName", reflect.TypeOf((*MockScaleSetVMScope)(nil).ScaleSetName))
}

// SetLongRunningOperationState mocks base method.
func (m *MockScaleSetVMScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockScaleSetVMScopeMockRecorder) SetLongRunningOperationState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockScaleSetVMScope)(nil).SetLongRunningOperationState), arg0)
}

// SetVMSSVM mocks base method.
func (m *MockScaleSetVMScope) SetVMSSVM(vmssvm *azure.VMSSVM) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVMSSVM", vmssvm)
}

// SetVMSSVM indicates an expected call of SetVMSSVM.
func (mr *MockScaleSetVMScopeMockRecorder) SetVMSSVM(vmssvm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVMSSVM", reflect.TypeOf((*MockScaleSetVMScope)(nil).SetVMSSVM), vmssvm)
}

// SubscriptionID mocks base method.
func (m *MockScaleSetVMScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockScaleSetVMScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockScaleSetVMScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockScaleSetVMScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockScaleSetVMScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockScaleSetVMScope)(nil).TenantID))
}

// UpdateDeleteStatus mocks base method.
func (m *MockScaleSetVMScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockScaleSetVMScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockScaleSetVMScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockScaleSetVMScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockScaleSetVMScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockScaleSetVMScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockScaleSetVMScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockScaleSetVMScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockScaleSetVMScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}
