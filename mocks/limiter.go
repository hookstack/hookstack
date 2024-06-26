// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/limiter/limiter.go
//
// Generated by this command:
//
//	mockgen --source internal/pkg/limiter/limiter.go --destination mocks/limiter.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRateLimiter is a mock of RateLimiter interface.
type MockRateLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimiterMockRecorder
}

// MockRateLimiterMockRecorder is the mock recorder for MockRateLimiter.
type MockRateLimiterMockRecorder struct {
	mock *MockRateLimiter
}

// NewMockRateLimiter creates a new mock instance.
func NewMockRateLimiter(ctrl *gomock.Controller) *MockRateLimiter {
	mock := &MockRateLimiter{ctrl: ctrl}
	mock.recorder = &MockRateLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimiter) EXPECT() *MockRateLimiterMockRecorder {
	return m.recorder
}

// Allow mocks base method.
func (m *MockRateLimiter) Allow(ctx context.Context, key string, rate int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allow", ctx, key, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Allow indicates an expected call of Allow.
func (mr *MockRateLimiterMockRecorder) Allow(ctx, key, rate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allow", reflect.TypeOf((*MockRateLimiter)(nil).Allow), ctx, key, rate)
}

// AllowWithDuration mocks base method.
func (m *MockRateLimiter) AllowWithDuration(ctx context.Context, key string, rate, duration int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowWithDuration", ctx, key, rate, duration)
	ret0, _ := ret[0].(error)
	return ret0
}

// AllowWithDuration indicates an expected call of AllowWithDuration.
func (mr *MockRateLimiterMockRecorder) AllowWithDuration(ctx, key, rate, duration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowWithDuration", reflect.TypeOf((*MockRateLimiter)(nil).AllowWithDuration), ctx, key, rate, duration)
}
