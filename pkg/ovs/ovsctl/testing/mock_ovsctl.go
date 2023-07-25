// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/ovs/ovsctl (interfaces: OVSCtlClient)

// Package testing is a generated GoMock package.
package testing

import (
	ovsctl "antrea.io/antrea/pkg/ovs/ovsctl"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOVSCtlClient is a mock of OVSCtlClient interface
type MockOVSCtlClient struct {
	ctrl     *gomock.Controller
	recorder *MockOVSCtlClientMockRecorder
}

// MockOVSCtlClientMockRecorder is the mock recorder for MockOVSCtlClient
type MockOVSCtlClientMockRecorder struct {
	mock *MockOVSCtlClient
}

// NewMockOVSCtlClient creates a new mock instance
func NewMockOVSCtlClient(ctrl *gomock.Controller) *MockOVSCtlClient {
	mock := &MockOVSCtlClient{ctrl: ctrl}
	mock.recorder = &MockOVSCtlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOVSCtlClient) EXPECT() *MockOVSCtlClientMockRecorder {
	return m.recorder
}

// DeleteDPInterface mocks base method
func (m *MockOVSCtlClient) DeleteDPInterface(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDPInterface", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDPInterface indicates an expected call of DeleteDPInterface
func (mr *MockOVSCtlClientMockRecorder) DeleteDPInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDPInterface", reflect.TypeOf((*MockOVSCtlClient)(nil).DeleteDPInterface), arg0)
}

// DumpFlows mocks base method
func (m *MockOVSCtlClient) DumpFlows(arg0 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DumpFlows", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpFlows indicates an expected call of DumpFlows
func (mr *MockOVSCtlClientMockRecorder) DumpFlows(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpFlows", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpFlows), arg0...)
}

// DumpFlowsWithoutTableNames mocks base method
func (m *MockOVSCtlClient) DumpFlowsWithoutTableNames(arg0 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DumpFlowsWithoutTableNames", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpFlowsWithoutTableNames indicates an expected call of DumpFlowsWithoutTableNames
func (mr *MockOVSCtlClientMockRecorder) DumpFlowsWithoutTableNames(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpFlowsWithoutTableNames", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpFlowsWithoutTableNames), arg0...)
}

// DumpGroup mocks base method
func (m *MockOVSCtlClient) DumpGroup(arg0 uint32) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpGroup", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpGroup indicates an expected call of DumpGroup
func (mr *MockOVSCtlClientMockRecorder) DumpGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpGroup", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpGroup), arg0)
}

// DumpGroups mocks base method
func (m *MockOVSCtlClient) DumpGroups() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpGroups")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpGroups indicates an expected call of DumpGroups
func (mr *MockOVSCtlClientMockRecorder) DumpGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpGroups", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpGroups))
}

// DumpMatchedFlow mocks base method
func (m *MockOVSCtlClient) DumpMatchedFlow(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpMatchedFlow", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpMatchedFlow indicates an expected call of DumpMatchedFlow
func (mr *MockOVSCtlClientMockRecorder) DumpMatchedFlow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpMatchedFlow", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpMatchedFlow), arg0)
}

// DumpPortsDesc mocks base method
func (m *MockOVSCtlClient) DumpPortsDesc() ([][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpPortsDesc")
	ret0, _ := ret[0].([][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpPortsDesc indicates an expected call of DumpPortsDesc
func (mr *MockOVSCtlClientMockRecorder) DumpPortsDesc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpPortsDesc", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpPortsDesc))
}

// DumpTableFlows mocks base method
func (m *MockOVSCtlClient) DumpTableFlows(arg0 byte) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpTableFlows", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpTableFlows indicates an expected call of DumpTableFlows
func (mr *MockOVSCtlClientMockRecorder) DumpTableFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpTableFlows", reflect.TypeOf((*MockOVSCtlClient)(nil).DumpTableFlows), arg0)
}

// GetDPFeatures mocks base method
func (m *MockOVSCtlClient) GetDPFeatures() (map[ovsctl.DPFeature]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDPFeatures")
	ret0, _ := ret[0].(map[ovsctl.DPFeature]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDPFeatures indicates an expected call of GetDPFeatures
func (mr *MockOVSCtlClientMockRecorder) GetDPFeatures() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDPFeatures", reflect.TypeOf((*MockOVSCtlClient)(nil).GetDPFeatures))
}

// RunAppctlCmd mocks base method
func (m *MockOVSCtlClient) RunAppctlCmd(arg0 string, arg1 bool, arg2 ...string) ([]byte, *ovsctl.ExecError) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAppctlCmd", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*ovsctl.ExecError)
	return ret0, ret1
}

// RunAppctlCmd indicates an expected call of RunAppctlCmd
func (mr *MockOVSCtlClientMockRecorder) RunAppctlCmd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAppctlCmd", reflect.TypeOf((*MockOVSCtlClient)(nil).RunAppctlCmd), varargs...)
}

// RunOfctlCmd mocks base method
func (m *MockOVSCtlClient) RunOfctlCmd(arg0 string, arg1 ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunOfctlCmd", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunOfctlCmd indicates an expected call of RunOfctlCmd
func (mr *MockOVSCtlClientMockRecorder) RunOfctlCmd(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOfctlCmd", reflect.TypeOf((*MockOVSCtlClient)(nil).RunOfctlCmd), varargs...)
}

// SetPortNoFlood mocks base method
func (m *MockOVSCtlClient) SetPortNoFlood(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPortNoFlood", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPortNoFlood indicates an expected call of SetPortNoFlood
func (mr *MockOVSCtlClientMockRecorder) SetPortNoFlood(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPortNoFlood", reflect.TypeOf((*MockOVSCtlClient)(nil).SetPortNoFlood), arg0)
}

// Trace mocks base method
func (m *MockOVSCtlClient) Trace(arg0 *ovsctl.TracingRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trace", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Trace indicates an expected call of Trace
func (mr *MockOVSCtlClientMockRecorder) Trace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockOVSCtlClient)(nil).Trace), arg0)
}
