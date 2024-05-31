// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockGitProjectProvider is an autogenerated mock type for the GitProjectProvider type
type MockGitProjectProvider struct {
	mock.Mock
}

type MockGitProjectProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGitProjectProvider) EXPECT() *MockGitProjectProvider_Expecter {
	return &MockGitProjectProvider_Expecter{mock: &_m.Mock}
}

// CreateProject provides a mock function with given fields: ctx, gitlabURL, token, fullPath
func (_m *MockGitProjectProvider) CreateProject(ctx context.Context, gitlabURL string, token string, fullPath string) error {
	ret := _m.Called(ctx, gitlabURL, token, fullPath)

	if len(ret) == 0 {
		panic("no return value specified for CreateProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, gitlabURL, token, fullPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGitProjectProvider_CreateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProject'
type MockGitProjectProvider_CreateProject_Call struct {
	*mock.Call
}

// CreateProject is a helper method to define mock.On call
//   - ctx context.Context
//   - gitlabURL string
//   - token string
//   - fullPath string
func (_e *MockGitProjectProvider_Expecter) CreateProject(ctx interface{}, gitlabURL interface{}, token interface{}, fullPath interface{}) *MockGitProjectProvider_CreateProject_Call {
	return &MockGitProjectProvider_CreateProject_Call{Call: _e.mock.On("CreateProject", ctx, gitlabURL, token, fullPath)}
}

func (_c *MockGitProjectProvider_CreateProject_Call) Run(run func(ctx context.Context, gitlabURL string, token string, fullPath string)) *MockGitProjectProvider_CreateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockGitProjectProvider_CreateProject_Call) Return(_a0 error) *MockGitProjectProvider_CreateProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGitProjectProvider_CreateProject_Call) RunAndReturn(run func(context.Context, string, string, string) error) *MockGitProjectProvider_CreateProject_Call {
	_c.Call.Return(run)
	return _c
}

// ProjectExists provides a mock function with given fields: ctx, gitlabURL, token, projectID
func (_m *MockGitProjectProvider) ProjectExists(ctx context.Context, gitlabURL string, token string, projectID string) (bool, error) {
	ret := _m.Called(ctx, gitlabURL, token, projectID)

	if len(ret) == 0 {
		panic("no return value specified for ProjectExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (bool, error)); ok {
		return rf(ctx, gitlabURL, token, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) bool); ok {
		r0 = rf(ctx, gitlabURL, token, projectID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, gitlabURL, token, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGitProjectProvider_ProjectExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProjectExists'
type MockGitProjectProvider_ProjectExists_Call struct {
	*mock.Call
}

// ProjectExists is a helper method to define mock.On call
//   - ctx context.Context
//   - gitlabURL string
//   - token string
//   - projectID string
func (_e *MockGitProjectProvider_Expecter) ProjectExists(ctx interface{}, gitlabURL interface{}, token interface{}, projectID interface{}) *MockGitProjectProvider_ProjectExists_Call {
	return &MockGitProjectProvider_ProjectExists_Call{Call: _e.mock.On("ProjectExists", ctx, gitlabURL, token, projectID)}
}

func (_c *MockGitProjectProvider_ProjectExists_Call) Run(run func(ctx context.Context, gitlabURL string, token string, projectID string)) *MockGitProjectProvider_ProjectExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockGitProjectProvider_ProjectExists_Call) Return(_a0 bool, _a1 error) *MockGitProjectProvider_ProjectExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGitProjectProvider_ProjectExists_Call) RunAndReturn(run func(context.Context, string, string, string) (bool, error)) *MockGitProjectProvider_ProjectExists_Call {
	_c.Call.Return(run)
	return _c
}

// SetDefaultBranch provides a mock function with given fields: ctx, githubURL, token, projectID, branch
func (_m *MockGitProjectProvider) SetDefaultBranch(ctx context.Context, githubURL string, token string, projectID string, branch string) error {
	ret := _m.Called(ctx, githubURL, token, projectID, branch)

	if len(ret) == 0 {
		panic("no return value specified for SetDefaultBranch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, githubURL, token, projectID, branch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGitProjectProvider_SetDefaultBranch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDefaultBranch'
type MockGitProjectProvider_SetDefaultBranch_Call struct {
	*mock.Call
}

// SetDefaultBranch is a helper method to define mock.On call
//   - ctx context.Context
//   - githubURL string
//   - token string
//   - projectID string
//   - branch string
func (_e *MockGitProjectProvider_Expecter) SetDefaultBranch(ctx interface{}, githubURL interface{}, token interface{}, projectID interface{}, branch interface{}) *MockGitProjectProvider_SetDefaultBranch_Call {
	return &MockGitProjectProvider_SetDefaultBranch_Call{Call: _e.mock.On("SetDefaultBranch", ctx, githubURL, token, projectID, branch)}
}

func (_c *MockGitProjectProvider_SetDefaultBranch_Call) Run(run func(ctx context.Context, githubURL string, token string, projectID string, branch string)) *MockGitProjectProvider_SetDefaultBranch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockGitProjectProvider_SetDefaultBranch_Call) Return(_a0 error) *MockGitProjectProvider_SetDefaultBranch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGitProjectProvider_SetDefaultBranch_Call) RunAndReturn(run func(context.Context, string, string, string, string) error) *MockGitProjectProvider_SetDefaultBranch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGitProjectProvider creates a new instance of MockGitProjectProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGitProjectProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGitProjectProvider {
	mock := &MockGitProjectProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
