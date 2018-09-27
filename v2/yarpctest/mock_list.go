// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/v2 (interfaces: Chooser,List,ChooserList)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package yarpctest is a generated GoMock package.
package yarpctest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v2 "go.uber.org/yarpc/v2"
	reflect "reflect"
)

// MockChooser is a mock of Chooser interface
type MockChooser struct {
	ctrl     *gomock.Controller
	recorder *MockChooserMockRecorder
}

// MockChooserMockRecorder is the mock recorder for MockChooser
type MockChooserMockRecorder struct {
	mock *MockChooser
}

// NewMockChooser creates a new mock instance
func NewMockChooser(ctrl *gomock.Controller) *MockChooser {
	mock := &MockChooser{ctrl: ctrl}
	mock.recorder = &MockChooserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChooser) EXPECT() *MockChooserMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockChooser) Choose(arg0 context.Context, arg1 *v2.Request) (v2.Peer, func(error), error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(v2.Peer)
	ret1, _ := ret[1].(func(error))
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Choose indicates an expected call of Choose
func (mr *MockChooserMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockChooser)(nil).Choose), arg0, arg1)
}

// MockList is a mock of List interface
type MockList struct {
	ctrl     *gomock.Controller
	recorder *MockListMockRecorder
}

// MockListMockRecorder is the mock recorder for MockList
type MockListMockRecorder struct {
	mock *MockList
}

// NewMockList creates a new mock instance
func NewMockList(ctrl *gomock.Controller) *MockList {
	mock := &MockList{ctrl: ctrl}
	mock.recorder = &MockListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockList) EXPECT() *MockListMockRecorder {
	return m.recorder
}

// Update mocks base method
func (m *MockList) Update(arg0 v2.ListUpdates) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockListMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockList)(nil).Update), arg0)
}

// MockChooserList is a mock of ChooserList interface
type MockChooserList struct {
	ctrl     *gomock.Controller
	recorder *MockChooserListMockRecorder
}

// MockChooserListMockRecorder is the mock recorder for MockChooserList
type MockChooserListMockRecorder struct {
	mock *MockChooserList
}

// NewMockChooserList creates a new mock instance
func NewMockChooserList(ctrl *gomock.Controller) *MockChooserList {
	mock := &MockChooserList{ctrl: ctrl}
	mock.recorder = &MockChooserListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChooserList) EXPECT() *MockChooserListMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockChooserList) Choose(arg0 context.Context, arg1 *v2.Request) (v2.Peer, func(error), error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(v2.Peer)
	ret1, _ := ret[1].(func(error))
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Choose indicates an expected call of Choose
func (mr *MockChooserListMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockChooserList)(nil).Choose), arg0, arg1)
}

// Update mocks base method
func (m *MockChooserList) Update(arg0 v2.ListUpdates) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockChooserListMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChooserList)(nil).Update), arg0)
}
