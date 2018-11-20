// Automatically generated by MockGen. DO NOT EDIT!
// Source: common.go

package network

import (
	net "net"

	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"

	v1 "kubevirt.io/kubevirt/pkg/api/v1"
)

// Mock of NetworkHandler interface
type MockNetworkHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkHandlerRecorder
}

// Recorder for MockNetworkHandler (not exported)
type _MockNetworkHandlerRecorder struct {
	mock *MockNetworkHandler
}

func NewMockNetworkHandler(ctrl *gomock.Controller) *MockNetworkHandler {
	mock := &MockNetworkHandler{ctrl: ctrl}
	mock.recorder = &_MockNetworkHandlerRecorder{mock}
	return mock
}

func (_m *MockNetworkHandler) EXPECT() *_MockNetworkHandlerRecorder {
	return _m.recorder
}

func (_m *MockNetworkHandler) LinkByName(name string) (netlink.Link, error) {
	ret := _m.ctrl.Call(_m, "LinkByName", name)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkByName", arg0)
}

func (_m *MockNetworkHandler) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "AddrList", link, family)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrList", arg0, arg1)
}

func (_m *MockNetworkHandler) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	ret := _m.ctrl.Call(_m, "RouteList", link, family)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteList", arg0, arg1)
}

func (_m *MockNetworkHandler) AddrDel(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.ctrl.Call(_m, "AddrDel", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) AddrDel(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrDel", arg0, arg1)
}

func (_m *MockNetworkHandler) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.ctrl.Call(_m, "AddrAdd", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrAdd", arg0, arg1)
}

func (_m *MockNetworkHandler) LinkSetDown(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetDown", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetDown(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetDown", arg0)
}

func (_m *MockNetworkHandler) LinkSetUp(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetUp", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetUp", arg0)
}

func (_m *MockNetworkHandler) LinkAdd(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkAdd", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkAdd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkAdd", arg0)
}

func (_m *MockNetworkHandler) LinkSetLearningOff(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetLearningOff", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetLearningOff(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetLearningOff", arg0)
}

func (_m *MockNetworkHandler) ParseAddr(s string) (*netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "ParseAddr", s)
	ret0, _ := ret[0].(*netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) ParseAddr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseAddr", arg0)
}

func (_m *MockNetworkHandler) SetRandomMac(iface string) (net.HardwareAddr, error) {
	ret := _m.ctrl.Call(_m, "SetRandomMac", iface)
	ret0, _ := ret[0].(net.HardwareAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) SetRandomMac(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRandomMac", arg0)
}

func (_m *MockNetworkHandler) GetMacDetails(iface string) (net.HardwareAddr, error) {
	ret := _m.ctrl.Call(_m, "GetMacDetails", iface)
	ret0, _ := ret[0].(net.HardwareAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) GetMacDetails(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMacDetails", arg0)
}

func (_m *MockNetworkHandler) StartDHCP(nic *VIF, serverAddr *netlink.Addr, bridgeInterfaceName string, dhcpOptions v1.DhcpOptions) {
	_m.ctrl.Call(_m, "StartDHCP", nic, serverAddr, bridgeInterfaceName, dhcpOptions)
}

func (_mr *_MockNetworkHandlerRecorder) StartDHCP(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartDHCP", arg0, arg1, arg2, arg3)
}
