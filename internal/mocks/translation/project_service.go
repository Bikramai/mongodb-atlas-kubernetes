// Code generated by mockery. DO NOT EDIT.

package translation

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	project "github.com/mongodb/mongodb-atlas-kubernetes/v2/internal/translation/project"
)

// ProjectServiceMock is an autogenerated mock type for the ProjectService type
type ProjectServiceMock struct {
	mock.Mock
}

type ProjectServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectServiceMock) EXPECT() *ProjectServiceMock_Expecter {
	return &ProjectServiceMock_Expecter{mock: &_m.Mock}
}

// CreateProject provides a mock function with given fields: ctx, _a1
func (_m *ProjectServiceMock) CreateProject(ctx context.Context, _a1 *project.Project) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *project.Project) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectServiceMock_CreateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProject'
type ProjectServiceMock_CreateProject_Call struct {
	*mock.Call
}

// CreateProject is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *project.Project
func (_e *ProjectServiceMock_Expecter) CreateProject(ctx interface{}, _a1 interface{}) *ProjectServiceMock_CreateProject_Call {
	return &ProjectServiceMock_CreateProject_Call{Call: _e.mock.On("CreateProject", ctx, _a1)}
}

func (_c *ProjectServiceMock_CreateProject_Call) Run(run func(ctx context.Context, _a1 *project.Project)) *ProjectServiceMock_CreateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*project.Project))
	})
	return _c
}

func (_c *ProjectServiceMock_CreateProject_Call) Return(_a0 error) *ProjectServiceMock_CreateProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectServiceMock_CreateProject_Call) RunAndReturn(run func(context.Context, *project.Project) error) *ProjectServiceMock_CreateProject_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProject provides a mock function with given fields: ctx, _a1
func (_m *ProjectServiceMock) DeleteProject(ctx context.Context, _a1 *project.Project) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *project.Project) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectServiceMock_DeleteProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProject'
type ProjectServiceMock_DeleteProject_Call struct {
	*mock.Call
}

// DeleteProject is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *project.Project
func (_e *ProjectServiceMock_Expecter) DeleteProject(ctx interface{}, _a1 interface{}) *ProjectServiceMock_DeleteProject_Call {
	return &ProjectServiceMock_DeleteProject_Call{Call: _e.mock.On("DeleteProject", ctx, _a1)}
}

func (_c *ProjectServiceMock_DeleteProject_Call) Run(run func(ctx context.Context, _a1 *project.Project)) *ProjectServiceMock_DeleteProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*project.Project))
	})
	return _c
}

func (_c *ProjectServiceMock_DeleteProject_Call) Return(_a0 error) *ProjectServiceMock_DeleteProject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectServiceMock_DeleteProject_Call) RunAndReturn(run func(context.Context, *project.Project) error) *ProjectServiceMock_DeleteProject_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectByName provides a mock function with given fields: ctx, name
func (_m *ProjectServiceMock) GetProjectByName(ctx context.Context, name string) (*project.Project, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectByName")
	}

	var r0 *project.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*project.Project, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *project.Project); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectServiceMock_GetProjectByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectByName'
type ProjectServiceMock_GetProjectByName_Call struct {
	*mock.Call
}

// GetProjectByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *ProjectServiceMock_Expecter) GetProjectByName(ctx interface{}, name interface{}) *ProjectServiceMock_GetProjectByName_Call {
	return &ProjectServiceMock_GetProjectByName_Call{Call: _e.mock.On("GetProjectByName", ctx, name)}
}

func (_c *ProjectServiceMock_GetProjectByName_Call) Run(run func(ctx context.Context, name string)) *ProjectServiceMock_GetProjectByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectServiceMock_GetProjectByName_Call) Return(_a0 *project.Project, _a1 error) *ProjectServiceMock_GetProjectByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectServiceMock_GetProjectByName_Call) RunAndReturn(run func(context.Context, string) (*project.Project, error)) *ProjectServiceMock_GetProjectByName_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectServiceMock creates a new instance of ProjectServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectServiceMock {
	mock := &ProjectServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
