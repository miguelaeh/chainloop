// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gax "github.com/googleapis/gax-go/v2"

	mock "github.com/stretchr/testify/mock"

	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// SecretsManagerInterface is an autogenerated mock type for the SecretsManagerInterface type
type SecretsManagerInterface struct {
	mock.Mock
}

// AccessSecretVersion provides a mock function with given fields: ctx, req, opts
func (_m *SecretsManagerInterface) AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *secretmanagerpb.AccessSecretVersionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.AccessSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.AccessSecretVersionRequest, ...gax.CallOption) *secretmanagerpb.AccessSecretVersionResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretmanagerpb.AccessSecretVersionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretmanagerpb.AccessSecretVersionRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddSecretVersion provides a mock function with given fields: ctx, req, opts
func (_m *SecretsManagerInterface) AddSecretVersion(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *secretmanagerpb.SecretVersion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.AddSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.AddSecretVersionRequest, ...gax.CallOption) *secretmanagerpb.SecretVersion); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretmanagerpb.SecretVersion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretmanagerpb.AddSecretVersionRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSecret provides a mock function with given fields: ctx, req, opts
func (_m *SecretsManagerInterface) CreateSecret(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *secretmanagerpb.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.CreateSecretRequest, ...gax.CallOption) (*secretmanagerpb.Secret, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.CreateSecretRequest, ...gax.CallOption) *secretmanagerpb.Secret); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretmanagerpb.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretmanagerpb.CreateSecretRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSecret provides a mock function with given fields: ctx, req, opts
func (_m *SecretsManagerInterface) DeleteSecret(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...gax.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.DeleteSecretRequest, ...gax.CallOption) error); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSecret provides a mock function with given fields: ctx, req, opts
func (_m *SecretsManagerInterface) GetSecret(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *secretmanagerpb.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.GetSecretRequest, ...gax.CallOption) (*secretmanagerpb.Secret, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretmanagerpb.GetSecretRequest, ...gax.CallOption) *secretmanagerpb.Secret); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretmanagerpb.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretmanagerpb.GetSecretRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSecretsManagerInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSecretsManagerInterface creates a new instance of SecretsManagerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSecretsManagerInterface(t mockConstructorTestingTNewSecretsManagerInterface) *SecretsManagerInterface {
	mock := &SecretsManagerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
