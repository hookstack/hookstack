// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/license/license.go
//
// Generated by this command:
//
//	mockgen --source internal/pkg/license/license.go --destination mocks/license.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLicenser is a mock of Licenser interface.
type MockLicenser struct {
	ctrl     *gomock.Controller
	recorder *MockLicenserMockRecorder
}

// MockLicenserMockRecorder is the mock recorder for MockLicenser.
type MockLicenserMockRecorder struct {
	mock *MockLicenser
}

// NewMockLicenser creates a new mock instance.
func NewMockLicenser(ctrl *gomock.Controller) *MockLicenser {
	mock := &MockLicenser{ctrl: ctrl}
	mock.recorder = &MockLicenserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLicenser) EXPECT() *MockLicenserMockRecorder {
	return m.recorder
}

// AddEnabledProject mocks base method.
func (m *MockLicenser) AddEnabledProject(projectID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEnabledProject", projectID)
}

// AddEnabledProject indicates an expected call of AddEnabledProject.
func (mr *MockLicenserMockRecorder) AddEnabledProject(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEnabledProject", reflect.TypeOf((*MockLicenser)(nil).AddEnabledProject), projectID)
}

// AdvancedEndpointMgmt mocks base method.
func (m *MockLicenser) AdvancedEndpointMgmt() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvancedEndpointMgmt")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AdvancedEndpointMgmt indicates an expected call of AdvancedEndpointMgmt.
func (mr *MockLicenserMockRecorder) AdvancedEndpointMgmt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvancedEndpointMgmt", reflect.TypeOf((*MockLicenser)(nil).AdvancedEndpointMgmt))
}

// AdvancedMsgBroker mocks base method.
func (m *MockLicenser) AdvancedMsgBroker() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvancedMsgBroker")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AdvancedMsgBroker indicates an expected call of AdvancedMsgBroker.
func (mr *MockLicenserMockRecorder) AdvancedMsgBroker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvancedMsgBroker", reflect.TypeOf((*MockLicenser)(nil).AdvancedMsgBroker))
}

// AdvancedRetentionPolicy mocks base method.
func (m *MockLicenser) AdvancedRetentionPolicy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvancedRetentionPolicy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AdvancedRetentionPolicy indicates an expected call of AdvancedRetentionPolicy.
func (mr *MockLicenserMockRecorder) AdvancedRetentionPolicy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvancedRetentionPolicy", reflect.TypeOf((*MockLicenser)(nil).AdvancedRetentionPolicy))
}

// AdvancedSubscriptions mocks base method.
func (m *MockLicenser) AdvancedSubscriptions() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvancedSubscriptions")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AdvancedSubscriptions indicates an expected call of AdvancedSubscriptions.
func (mr *MockLicenserMockRecorder) AdvancedSubscriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvancedSubscriptions", reflect.TypeOf((*MockLicenser)(nil).AdvancedSubscriptions))
}

// AdvancedWebhookFiltering mocks base method.
func (m *MockLicenser) AdvancedWebhookFiltering() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvancedWebhookFiltering")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AdvancedWebhookFiltering indicates an expected call of AdvancedWebhookFiltering.
func (mr *MockLicenserMockRecorder) AdvancedWebhookFiltering() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvancedWebhookFiltering", reflect.TypeOf((*MockLicenser)(nil).AdvancedWebhookFiltering))
}

// AgentExecutionMode mocks base method.
func (m *MockLicenser) AgentExecutionMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentExecutionMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AgentExecutionMode indicates an expected call of AgentExecutionMode.
func (mr *MockLicenserMockRecorder) AgentExecutionMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentExecutionMode", reflect.TypeOf((*MockLicenser)(nil).AgentExecutionMode))
}

// AsynqMonitoring mocks base method.
func (m *MockLicenser) AsynqMonitoring() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsynqMonitoring")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AsynqMonitoring indicates an expected call of AsynqMonitoring.
func (mr *MockLicenserMockRecorder) AsynqMonitoring() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsynqMonitoring", reflect.TypeOf((*MockLicenser)(nil).AsynqMonitoring))
}

// CanExportPrometheusMetrics mocks base method.
func (m *MockLicenser) CanExportPrometheusMetrics() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanExportPrometheusMetrics")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanExportPrometheusMetrics indicates an expected call of CanExportPrometheusMetrics.
func (mr *MockLicenserMockRecorder) CanExportPrometheusMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanExportPrometheusMetrics", reflect.TypeOf((*MockLicenser)(nil).CanExportPrometheusMetrics))
}

// CircuitBreaking mocks base method.
func (m *MockLicenser) CircuitBreaking() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CircuitBreaking")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CircuitBreaking indicates an expected call of CircuitBreaking.
func (mr *MockLicenserMockRecorder) CircuitBreaking() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CircuitBreaking", reflect.TypeOf((*MockLicenser)(nil).CircuitBreaking))
}

// ConsumerPoolTuning mocks base method.
func (m *MockLicenser) ConsumerPoolTuning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumerPoolTuning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ConsumerPoolTuning indicates an expected call of ConsumerPoolTuning.
func (mr *MockLicenserMockRecorder) ConsumerPoolTuning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerPoolTuning", reflect.TypeOf((*MockLicenser)(nil).ConsumerPoolTuning))
}

// CreateOrg mocks base method.
func (m *MockLicenser) CreateOrg(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrg", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrg indicates an expected call of CreateOrg.
func (mr *MockLicenserMockRecorder) CreateOrg(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrg", reflect.TypeOf((*MockLicenser)(nil).CreateOrg), ctx)
}

// CreateProject mocks base method.
func (m *MockLicenser) CreateProject(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockLicenserMockRecorder) CreateProject(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockLicenser)(nil).CreateProject), ctx)
}

// CreateUser mocks base method.
func (m *MockLicenser) CreateUser(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockLicenserMockRecorder) CreateUser(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockLicenser)(nil).CreateUser), ctx)
}

// FeatureListJSON mocks base method.
func (m *MockLicenser) FeatureListJSON(ctx context.Context) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeatureListJSON", ctx)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FeatureListJSON indicates an expected call of FeatureListJSON.
func (mr *MockLicenserMockRecorder) FeatureListJSON(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeatureListJSON", reflect.TypeOf((*MockLicenser)(nil).FeatureListJSON), ctx)
}

// HADeployment mocks base method.
func (m *MockLicenser) HADeployment() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HADeployment")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HADeployment indicates an expected call of HADeployment.
func (mr *MockLicenserMockRecorder) HADeployment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HADeployment", reflect.TypeOf((*MockLicenser)(nil).HADeployment))
}

// IngestRate mocks base method.
func (m *MockLicenser) IngestRate() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestRate")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IngestRate indicates an expected call of IngestRate.
func (mr *MockLicenserMockRecorder) IngestRate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestRate", reflect.TypeOf((*MockLicenser)(nil).IngestRate))
}

// IpRules mocks base method.
func (m *MockLicenser) IpRules() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpRules")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IpRules indicates an expected call of IpRules.
func (mr *MockLicenserMockRecorder) IpRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpRules", reflect.TypeOf((*MockLicenser)(nil).IpRules))
}

// MultiPlayerMode mocks base method.
func (m *MockLicenser) MultiPlayerMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiPlayerMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// MultiPlayerMode indicates an expected call of MultiPlayerMode.
func (mr *MockLicenserMockRecorder) MultiPlayerMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiPlayerMode", reflect.TypeOf((*MockLicenser)(nil).MultiPlayerMode))
}

// MutualTLS mocks base method.
func (m *MockLicenser) MutualTLS() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutualTLS")
	ret0, _ := ret[0].(bool)
	return ret0
}

// MutualTLS indicates an expected call of MutualTLS.
func (mr *MockLicenserMockRecorder) MutualTLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutualTLS", reflect.TypeOf((*MockLicenser)(nil).MutualTLS))
}

// PortalLinks mocks base method.
func (m *MockLicenser) PortalLinks() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PortalLinks")
	ret0, _ := ret[0].(bool)
	return ret0
}

// PortalLinks indicates an expected call of PortalLinks.
func (mr *MockLicenserMockRecorder) PortalLinks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PortalLinks", reflect.TypeOf((*MockLicenser)(nil).PortalLinks))
}

// ProjectEnabled mocks base method.
func (m *MockLicenser) ProjectEnabled(projectID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectEnabled", projectID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ProjectEnabled indicates an expected call of ProjectEnabled.
func (mr *MockLicenserMockRecorder) ProjectEnabled(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectEnabled", reflect.TypeOf((*MockLicenser)(nil).ProjectEnabled), projectID)
}

// RemoveEnabledProject mocks base method.
func (m *MockLicenser) RemoveEnabledProject(projectID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveEnabledProject", projectID)
}

// RemoveEnabledProject indicates an expected call of RemoveEnabledProject.
func (mr *MockLicenserMockRecorder) RemoveEnabledProject(projectID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEnabledProject", reflect.TypeOf((*MockLicenser)(nil).RemoveEnabledProject), projectID)
}

// SynchronousWebhooks mocks base method.
func (m *MockLicenser) SynchronousWebhooks() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SynchronousWebhooks")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SynchronousWebhooks indicates an expected call of SynchronousWebhooks.
func (mr *MockLicenserMockRecorder) SynchronousWebhooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SynchronousWebhooks", reflect.TypeOf((*MockLicenser)(nil).SynchronousWebhooks))
}

// Transformations mocks base method.
func (m *MockLicenser) Transformations() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transformations")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Transformations indicates an expected call of Transformations.
func (mr *MockLicenserMockRecorder) Transformations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transformations", reflect.TypeOf((*MockLicenser)(nil).Transformations))
}

// UseForwardProxy mocks base method.
func (m *MockLicenser) UseForwardProxy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseForwardProxy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseForwardProxy indicates an expected call of UseForwardProxy.
func (mr *MockLicenserMockRecorder) UseForwardProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseForwardProxy", reflect.TypeOf((*MockLicenser)(nil).UseForwardProxy))
}

// WebhookAnalytics mocks base method.
func (m *MockLicenser) WebhookAnalytics() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebhookAnalytics")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WebhookAnalytics indicates an expected call of WebhookAnalytics.
func (mr *MockLicenserMockRecorder) WebhookAnalytics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebhookAnalytics", reflect.TypeOf((*MockLicenser)(nil).WebhookAnalytics))
}
