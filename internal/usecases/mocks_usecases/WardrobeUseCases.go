// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks_usecases

import (
	context "context"
	request "sagara_backend_test/internal/usecases/request"

	mock "github.com/stretchr/testify/mock"

	response "sagara_backend_test/internal/usecases/response"

	uuid "github.com/google/uuid"
)

// WardrobeUseCases is an autogenerated mock type for the WardrobeUseCases type
type WardrobeUseCases struct {
	mock.Mock
}

// AddStock provides a mock function with given fields: ctx, id, add
func (_m *WardrobeUseCases) AddStock(ctx context.Context, id *uuid.UUID, add int) (*response.WardrobeResponse, error) {
	ret := _m.Called(ctx, id, add)

	var r0 *response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, int) (*response.WardrobeResponse, error)); ok {
		return rf(ctx, id, add)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, int) *response.WardrobeResponse); ok {
		r0 = rf(ctx, id, add)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID, int) error); ok {
		r1 = rf(ctx, id, add)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWardrobe provides a mock function with given fields: ctx, id
func (_m *WardrobeUseCases) DeleteWardrobe(ctx context.Context, id *uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllWardrobe provides a mock function with given fields: ctx
func (_m *WardrobeUseCases) GetAllWardrobe(ctx context.Context) (*[]response.WardrobeResponse, error) {
	ret := _m.Called(ctx)

	var r0 *[]response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]response.WardrobeResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]response.WardrobeResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvailable provides a mock function with given fields: ctx
func (_m *WardrobeUseCases) GetAvailable(ctx context.Context) (*[]response.WardrobeResponse, error) {
	ret := _m.Called(ctx)

	var r0 *[]response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]response.WardrobeResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]response.WardrobeResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLessThan provides a mock function with given fields: ctx, amount
func (_m *WardrobeUseCases) GetLessThan(ctx context.Context, amount int) (*[]response.WardrobeResponse, error) {
	ret := _m.Called(ctx, amount)

	var r0 *[]response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*[]response.WardrobeResponse, error)); ok {
		return rf(ctx, amount)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *[]response.WardrobeResponse); ok {
		r0 = rf(ctx, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnavailable provides a mock function with given fields: ctx
func (_m *WardrobeUseCases) GetUnavailable(ctx context.Context) (*[]response.WardrobeResponse, error) {
	ret := _m.Called(ctx)

	var r0 *[]response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]response.WardrobeResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]response.WardrobeResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWardrobe provides a mock function with given fields: ctx, id
func (_m *WardrobeUseCases) GetWardrobe(ctx context.Context, id *uuid.UUID) (*response.WardrobeResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID) (*response.WardrobeResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID) *response.WardrobeResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertWardrobe provides a mock function with given fields: ctx, _a1
func (_m *WardrobeUseCases) InsertWardrobe(ctx context.Context, _a1 *request.WardrobeInsertRequest) (*response.WardrobeResponse, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.WardrobeInsertRequest) (*response.WardrobeResponse, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *request.WardrobeInsertRequest) *response.WardrobeResponse); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *request.WardrobeInsertRequest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, color, size
func (_m *WardrobeUseCases) Search(ctx context.Context, color string, size string) (*[]response.WardrobeResponse, error) {
	ret := _m.Called(ctx, color, size)

	var r0 *[]response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*[]response.WardrobeResponse, error)); ok {
		return rf(ctx, color, size)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *[]response.WardrobeResponse); ok {
		r0 = rf(ctx, color, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, color, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubStock provides a mock function with given fields: ctx, id, def
func (_m *WardrobeUseCases) SubStock(ctx context.Context, id *uuid.UUID, def int) (*response.WardrobeResponse, error) {
	ret := _m.Called(ctx, id, def)

	var r0 *response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, int) (*response.WardrobeResponse, error)); ok {
		return rf(ctx, id, def)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, int) *response.WardrobeResponse); ok {
		r0 = rf(ctx, id, def)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID, int) error); ok {
		r1 = rf(ctx, id, def)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWardrobe provides a mock function with given fields: ctx, id, _a2
func (_m *WardrobeUseCases) UpdateWardrobe(ctx context.Context, id *uuid.UUID, _a2 *request.WardrobeUpdateRequest) (*response.WardrobeResponse, error) {
	ret := _m.Called(ctx, id, _a2)

	var r0 *response.WardrobeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, *request.WardrobeUpdateRequest) (*response.WardrobeResponse, error)); ok {
		return rf(ctx, id, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID, *request.WardrobeUpdateRequest) *response.WardrobeResponse); ok {
		r0 = rf(ctx, id, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.WardrobeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID, *request.WardrobeUpdateRequest) error); ok {
		r1 = rf(ctx, id, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWardrobeUseCases creates a new instance of WardrobeUseCases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWardrobeUseCases(t interface {
	mock.TestingT
	Cleanup(func())
}) *WardrobeUseCases {
	mock := &WardrobeUseCases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
