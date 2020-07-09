// Code generated by mockery v1.0.0. DO NOT EDIT.

package checks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mockAlertRegistry is an autogenerated mock type for the alertRegistry type
type mockAlertRegistry struct {
	mock.Mock
}

// CreateAlert provides a mock function with given fields: id, labels, annotations, delayFor
func (_m *mockAlertRegistry) CreateAlert(id string, labels map[string]string, annotations map[string]string, delayFor time.Duration) {
	_m.Called(id, labels, annotations, delayFor)
}

// RemovePrefix provides a mock function with given fields: prefix, keepIDs
func (_m *mockAlertRegistry) RemovePrefix(prefix string, keepIDs map[string]struct{}) {
	_m.Called(prefix, keepIDs)
}