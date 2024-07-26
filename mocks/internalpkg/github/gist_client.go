// Code generated by mockery v2.42.1. DO NOT EDIT.

package github_mocks

import (
	context "context"

	github "github.com/google/go-github/v58/github"

	mock "github.com/stretchr/testify/mock"
)

// GistClient is an autogenerated mock type for the GistClient type
type GistClient struct {
	mock.Mock
}

type GistClient_Expecter struct {
	mock *mock.Mock
}

func (_m *GistClient) EXPECT() *GistClient_Expecter {
	return &GistClient_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, gist
func (_m *GistClient) Create(ctx context.Context, gist *github.Gist) (*github.Gist, *github.Response, error) {
	ret := _m.Called(ctx, gist)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *github.Gist
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *github.Gist) (*github.Gist, *github.Response, error)); ok {
		return rf(ctx, gist)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *github.Gist) *github.Gist); ok {
		r0 = rf(ctx, gist)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Gist)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *github.Gist) *github.Response); ok {
		r1 = rf(ctx, gist)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *github.Gist) error); ok {
		r2 = rf(ctx, gist)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GistClient_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type GistClient_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - gist *github.Gist
func (_e *GistClient_Expecter) Create(ctx interface{}, gist interface{}) *GistClient_Create_Call {
	return &GistClient_Create_Call{Call: _e.mock.On("Create", ctx, gist)}
}

func (_c *GistClient_Create_Call) Run(run func(ctx context.Context, gist *github.Gist)) *GistClient_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*github.Gist))
	})
	return _c
}

func (_c *GistClient_Create_Call) Return(_a0 *github.Gist, _a1 *github.Response, _a2 error) *GistClient_Create_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GistClient_Create_Call) RunAndReturn(run func(context.Context, *github.Gist) (*github.Gist, *github.Response, error)) *GistClient_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Edit provides a mock function with given fields: ctx, id, gist
func (_m *GistClient) Edit(ctx context.Context, id string, gist *github.Gist) (*github.Gist, *github.Response, error) {
	ret := _m.Called(ctx, id, gist)

	if len(ret) == 0 {
		panic("no return value specified for Edit")
	}

	var r0 *github.Gist
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.Gist) (*github.Gist, *github.Response, error)); ok {
		return rf(ctx, id, gist)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *github.Gist) *github.Gist); ok {
		r0 = rf(ctx, id, gist)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Gist)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *github.Gist) *github.Response); ok {
		r1 = rf(ctx, id, gist)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *github.Gist) error); ok {
		r2 = rf(ctx, id, gist)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GistClient_Edit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Edit'
type GistClient_Edit_Call struct {
	*mock.Call
}

// Edit is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - gist *github.Gist
func (_e *GistClient_Expecter) Edit(ctx interface{}, id interface{}, gist interface{}) *GistClient_Edit_Call {
	return &GistClient_Edit_Call{Call: _e.mock.On("Edit", ctx, id, gist)}
}

func (_c *GistClient_Edit_Call) Run(run func(ctx context.Context, id string, gist *github.Gist)) *GistClient_Edit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*github.Gist))
	})
	return _c
}

func (_c *GistClient_Edit_Call) Return(_a0 *github.Gist, _a1 *github.Response, _a2 error) *GistClient_Edit_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GistClient_Edit_Call) RunAndReturn(run func(context.Context, string, *github.Gist) (*github.Gist, *github.Response, error)) *GistClient_Edit_Call {
	_c.Call.Return(run)
	return _c
}

// NewGistClient creates a new instance of GistClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGistClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *GistClient {
	mock := &GistClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
