// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Rmarken5/file-broadcaster/file-listener (interfaces: IFileListener)

// Package mocks is a generated GoMock package.
package mocks

import (
	fs "io/fs"
	reflect "reflect"

	fsnotify "github.com/fsnotify/fsnotify"
	gomock "github.com/golang/mock/gomock"
)

// MockIFileListener is a mock of IFileListener interface.
type MockIFileListener struct {
	ctrl     *gomock.Controller
	recorder *MockIFileListenerMockRecorder
}

// MockIFileListenerMockRecorder is the mock recorder for MockIFileListener.
type MockIFileListenerMockRecorder struct {
	mock *MockIFileListener
}

// NewMockIFileListener creates a new mock instance.
func NewMockIFileListener(ctrl *gomock.Controller) *MockIFileListener {
	mock := &MockIFileListener{ctrl: ctrl}
	mock.recorder = &MockIFileListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileListener) EXPECT() *MockIFileListenerMockRecorder {
	return m.recorder
}

// ListenForFiles mocks base method.
func (m *MockIFileListener) ListenForFiles(arg0 string) (chan fsnotify.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenForFiles", arg0)
	ret0, _ := ret[0].(chan fsnotify.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenForFiles indicates an expected call of ListenForFiles.
func (mr *MockIFileListenerMockRecorder) ListenForFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenForFiles", reflect.TypeOf((*MockIFileListener)(nil).ListenForFiles), arg0)
}

// ReadDirectory mocks base method.
func (m *MockIFileListener) ReadDirectory(arg0 []fs.DirEntry) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDirectory", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDirectory indicates an expected call of ReadDirectory.
func (mr *MockIFileListenerMockRecorder) ReadDirectory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDirectory", reflect.TypeOf((*MockIFileListener)(nil).ReadDirectory), arg0)
}
