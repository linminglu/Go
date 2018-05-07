// Automatically generated by MockGen. DO NOT EDIT!
// Source: message_sender.go

package interfaces

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	gate_rpc "steve/structs/proto/gate_rpc"
)

// Mock of MessageSender interface
type MockMessageSender struct {
	ctrl     *gomock.Controller
	recorder *_MockMessageSenderRecorder
}

// Recorder for MockMessageSender (not exported)
type _MockMessageSenderRecorder struct {
	mock *MockMessageSender
}

func NewMockMessageSender(ctrl *gomock.Controller) *MockMessageSender {
	mock := &MockMessageSender{ctrl: ctrl}
	mock.recorder = &_MockMessageSenderRecorder{mock}
	return mock
}

func (_m *MockMessageSender) EXPECT() *_MockMessageSenderRecorder {
	return _m.recorder
}

func (_m *MockMessageSender) SendPackage(clientID uint64, head *gate_rpc.Header, body proto.Message) error {
	ret := _m.ctrl.Call(_m, "SendPackage", clientID, head, body)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageSenderRecorder) SendPackage(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendPackage", arg0, arg1, arg2)
}

func (_m *MockMessageSender) BroadcastPackage(clientIDs []uint64, head *gate_rpc.Header, body proto.Message) error {
	ret := _m.ctrl.Call(_m, "BroadcastPackage", clientIDs, head, body)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageSenderRecorder) BroadcastPackage(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BroadcastPackage", arg0, arg1, arg2)
}

func (_m *MockMessageSender) SendPackageBare(clientID uint64, head *gate_rpc.Header, bodyData []byte) error {
	ret := _m.ctrl.Call(_m, "SendPackageBare", clientID, head, bodyData)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageSenderRecorder) SendPackageBare(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendPackageBare", arg0, arg1, arg2)
}

func (_m *MockMessageSender) BroadcastPackageBare(clientIDs []uint64, head *gate_rpc.Header, bodyData []byte) error {
	ret := _m.ctrl.Call(_m, "BroadcastPackageBare", clientIDs, head, bodyData)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageSenderRecorder) BroadcastPackageBare(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BroadcastPackageBare", arg0, arg1, arg2)
}
