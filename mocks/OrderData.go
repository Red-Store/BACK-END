// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	order "MyEcommerce/features/order"

	mock "github.com/stretchr/testify/mock"
)

// OrderData is an autogenerated mock type for the OrderDataInterface type
type OrderData struct {
	mock.Mock
}

// CancleOrder provides a mock function with given fields: userIdLogin, orderId, orderCore
func (_m *OrderData) CancleOrder(userIdLogin int, orderId string, orderCore order.OrderCore) error {
	ret := _m.Called(userIdLogin, orderId, orderCore)

	if len(ret) == 0 {
		panic("no return value specified for CancleOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, order.OrderCore) error); ok {
		r0 = rf(userIdLogin, orderId, orderCore)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertOrder provides a mock function with given fields: userIdLogin, cartIds, inputOrder
func (_m *OrderData) InsertOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore) (*order.OrderCore, error) {
	ret := _m.Called(userIdLogin, cartIds, inputOrder)

	if len(ret) == 0 {
		panic("no return value specified for InsertOrder")
	}

	var r0 *order.OrderCore
	var r1 error
	if rf, ok := ret.Get(0).(func(int, []uint, order.OrderCore) (*order.OrderCore, error)); ok {
		return rf(userIdLogin, cartIds, inputOrder)
	}
	if rf, ok := ret.Get(0).(func(int, []uint, order.OrderCore) *order.OrderCore); ok {
		r0 = rf(userIdLogin, cartIds, inputOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.OrderCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, []uint, order.OrderCore) error); ok {
		r1 = rf(userIdLogin, cartIds, inputOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectOrderAdmin provides a mock function with given fields: userIdLogin, page, limit
func (_m *OrderData) SelectOrderAdmin(userIdLogin int, page int, limit int) ([]order.OrderItemCore, int, error) {
	ret := _m.Called(userIdLogin, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for SelectOrderAdmin")
	}

	var r0 []order.OrderItemCore
if ret.Get(0) != nil {
    if rf, ok := ret.Get(0).(func(int, int, int) []order.OrderItemCore); ok {
        r0 = rf(userIdLogin, page, limit)
    } else {
        r0 = ret.Get(0).([]order.OrderItemCore)
    }
}

var r1 int
if ret.Get(1) != nil {
    if rf, ok := ret.Get(1).(func(int, int, int) int); ok {
        r1 = rf(userIdLogin, page, limit)
    } else {
        r1 = ret.Get(1).(int)
    }
}

var r2 error
if ret.Get(2) != nil {
    if rf, ok := ret.Get(2).(func(int, int, int) error); ok {
        r2 = rf(userIdLogin, page, limit)
    } else {
        r2 = ret.Error(2)
    }
}

return r0, r1, r2

}

// SelectOrderUser provides a mock function with given fields: userIdLogin
func (_m *OrderData) SelectOrderUser(userIdLogin int) ([]order.OrderItemCore, error) {
	ret := _m.Called(userIdLogin)

	if len(ret) == 0 {
		panic("no return value specified for SelectOrderUser")
	}

	var r0 []order.OrderItemCore
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]order.OrderItemCore, error)); ok {
		return rf(userIdLogin)
	}
	if rf, ok := ret.Get(0).(func(int) []order.OrderItemCore); ok {
		r0 = rf(userIdLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]order.OrderItemCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userIdLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebhoocksData provides a mock function with given fields: reqNotif
func (_m *OrderData) WebhoocksData(reqNotif order.OrderCore) error {
	ret := _m.Called(reqNotif)

	if len(ret) == 0 {
		panic("no return value specified for WebhoocksData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(order.OrderCore) error); ok {
		r0 = rf(reqNotif)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderData creates a new instance of OrderData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderData(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderData {
	mock := &OrderData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
