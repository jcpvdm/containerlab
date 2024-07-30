// Code generated by MockGen. DO NOT EDIT.
// Source: nodes/default_node.go
//
// Generated by this command:
//
//	mockgen -package=mocknodes -source=nodes/default_node.go -destination=./mocks/mocknodes/default_node.go
//

// Package mocknodes is a generated GoMock package.
package mocknodes

import (
	context "context"
	reflect "reflect"

	exec "github.com/srl-labs/containerlab/clab/exec"
	runtime "github.com/srl-labs/containerlab/runtime"
	gomock "go.uber.org/mock/gomock"
)

// MockNodeOverwrites is a mock of NodeOverwrites interface.
type MockNodeOverwrites struct {
	ctrl     *gomock.Controller
	recorder *MockNodeOverwritesMockRecorder
}

// MockNodeOverwritesMockRecorder is the mock recorder for MockNodeOverwrites.
type MockNodeOverwritesMockRecorder struct {
	mock *MockNodeOverwrites
}

// NewMockNodeOverwrites creates a new mock instance.
func NewMockNodeOverwrites(ctrl *gomock.Controller) *MockNodeOverwrites {
	mock := &MockNodeOverwrites{ctrl: ctrl}
	mock.recorder = &MockNodeOverwritesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeOverwrites) EXPECT() *MockNodeOverwritesMockRecorder {
	return m.recorder
}

// CheckInterfaceName mocks base method.
func (m *MockNodeOverwrites) CheckInterfaceName() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckInterfaceName")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckInterfaceName indicates an expected call of CheckInterfaceName.
func (mr *MockNodeOverwritesMockRecorder) CheckInterfaceName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInterfaceName", reflect.TypeOf((*MockNodeOverwrites)(nil).CheckInterfaceName))
}

// GetContainerName mocks base method.
func (m *MockNodeOverwrites) GetContainerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContainerName indicates an expected call of GetContainerName.
func (mr *MockNodeOverwritesMockRecorder) GetContainerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerName", reflect.TypeOf((*MockNodeOverwrites)(nil).GetContainerName))
}

// GetContainers mocks base method.
func (m *MockNodeOverwrites) GetContainers(ctx context.Context) ([]runtime.GenericContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainers", ctx)
	ret0, _ := ret[0].([]runtime.GenericContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainers indicates an expected call of GetContainers.
func (mr *MockNodeOverwritesMockRecorder) GetContainers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainers", reflect.TypeOf((*MockNodeOverwrites)(nil).GetContainers), ctx)
}

// GetImages mocks base method.
func (m *MockNodeOverwrites) GetImages(ctx context.Context) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages", ctx)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetImages indicates an expected call of GetImages.
func (mr *MockNodeOverwritesMockRecorder) GetImages(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockNodeOverwrites)(nil).GetImages), ctx)
}

// PullImage mocks base method.
func (m *MockNodeOverwrites) PullImage(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullImage indicates an expected call of PullImage.
func (mr *MockNodeOverwritesMockRecorder) PullImage(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockNodeOverwrites)(nil).PullImage), ctx)
}

// RunExec mocks base method.
func (m *MockNodeOverwrites) RunExec(arg0 context.Context, arg1 *exec.ExecCmd) (*exec.ExecResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunExec", arg0, arg1)
	ret0, _ := ret[0].(*exec.ExecResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunExec indicates an expected call of RunExec.
func (mr *MockNodeOverwritesMockRecorder) RunExec(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunExec", reflect.TypeOf((*MockNodeOverwrites)(nil).RunExec), arg0, arg1)
}

// VerifyHostRequirements mocks base method.
func (m *MockNodeOverwrites) VerifyHostRequirements() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHostRequirements")
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHostRequirements indicates an expected call of VerifyHostRequirements.
func (mr *MockNodeOverwritesMockRecorder) VerifyHostRequirements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHostRequirements", reflect.TypeOf((*MockNodeOverwrites)(nil).VerifyHostRequirements))
}

// VerifyLicenseFileExists mocks base method.
func (m *MockNodeOverwrites) VerifyLicenseFileExists(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLicenseFileExists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyLicenseFileExists indicates an expected call of VerifyLicenseFileExists.
func (mr *MockNodeOverwritesMockRecorder) VerifyLicenseFileExists(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLicenseFileExists", reflect.TypeOf((*MockNodeOverwrites)(nil).VerifyLicenseFileExists), arg0)
}

// VerifyStartupConfig mocks base method.
func (m *MockNodeOverwrites) VerifyStartupConfig(topoDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyStartupConfig", topoDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyStartupConfig indicates an expected call of VerifyStartupConfig.
func (mr *MockNodeOverwritesMockRecorder) VerifyStartupConfig(topoDir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyStartupConfig", reflect.TypeOf((*MockNodeOverwrites)(nil).VerifyStartupConfig), topoDir)
}