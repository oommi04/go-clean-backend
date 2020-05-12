// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"
import time "time"

// BotService is an autogenerated mock type for the BotService type
type BotService struct {
	mock.Mock
}

// ReplyMessage provides a mock function with given fields: r, messageReply, timeout
func (_m *BotService) ReplyMessage(r *http.Request, messageReply string, timeout time.Duration) error {
	ret := _m.Called(r, messageReply, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Request, string, time.Duration) error); ok {
		r0 = rf(r, messageReply, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}