// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/repository/interface.go
//
// Generated by this command:
//
//	mockgen -destination=../internal/repository/mocks.go -package=repository -source=../internal/repository/interface.go
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"
	actor "vk-film-library/internal/entity/actor"
	transaction "vk-film-library/internal/transaction"

	gomock "go.uber.org/mock/gomock"
)

// MockActors is a mock of Actors interface.
type MockActors struct {
	ctrl     *gomock.Controller
	recorder *MockActorsMockRecorder
}

// MockActorsMockRecorder is the mock recorder for MockActors.
type MockActorsMockRecorder struct {
	mock *MockActors
}

// NewMockActors creates a new mock instance.
func NewMockActors(ctrl *gomock.Controller) *MockActors {
	mock := &MockActors{ctrl: ctrl}
	mock.recorder = &MockActorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActors) EXPECT() *MockActorsMockRecorder {
	return m.recorder
}

// CreateActor mocks base method.
func (m *MockActors) CreateActor(ts transaction.Session, p actor.CreateActorParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActor", ts, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateActor indicates an expected call of CreateActor.
func (mr *MockActorsMockRecorder) CreateActor(ts, p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActor", reflect.TypeOf((*MockActors)(nil).CreateActor), ts, p)
}
