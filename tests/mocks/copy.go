package mocks

import "github.com/stretchr/testify/mock"

// FakeCopy is an autogenerated mock type for the FakeCopy type
type FakeCopy struct {
	mock.Mock
}

// Execute provides a mock function with given fields: source, fields
func (_m *FakeCopy) Execute(source any, fields ...string) any {
	_va := make([]any, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []any
	_ca = append(_ca, source)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 any
	if rf, ok := ret.Get(0).(func(any, ...string) any); ok {
		r0 = rf(source, fields...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0)
		}
	}

	return r0
}

type mockConstructorTestingTNewCopy interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeCopy creates a new instance of FakeCopy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeCopy(t mockConstructorTestingTNewCopy) *FakeCopy {
	fakeCopy := &FakeCopy{}
	fakeCopy.Mock.Test(t)

	t.Cleanup(func() { fakeCopy.AssertExpectations(t) })

	return fakeCopy
}