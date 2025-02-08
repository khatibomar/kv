// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

import (
	"testing"
	"time"

	"github.com/khatibomar/kv/internal/assert"
)

func TestMinInt(t *testing.T) {
	tests := []struct {
		tag       string
		threshold int
		exclusive bool
		value     int
		err       string
	}{
		{"test 1", 1, false, 1, ""},
		{"test 2", 1, false, 2, ""},
		{"test 3", 1, false, -1, "must be no less than 1"},
		{"test 4", 1, false, 0, ""},
		{"test 5", 1, true, 1, "must be greater than 1"},
	}

	for _, test := range tests {
		r := Min(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMinUint(t *testing.T) {
	tests := []struct {
		tag       string
		threshold uint
		exclusive bool
		value     uint
		err       string
	}{
		{"test 1", uint(2), false, uint(2), ""},
		{"test 2", uint(2), false, uint(3), ""},
		{"test 3", uint(2), false, uint(1), "must be no less than 2"},
		{"test 4", uint(2), false, uint(0), ""},
		{"test 5", uint(2), true, uint(2), "must be greater than 2"},
	}

	for _, test := range tests {
		r := Min(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMinFloat64(t *testing.T) {
	tests := []struct {
		tag       string
		threshold float64
		exclusive bool
		value     float64
		err       string
	}{
		{"test 1", float64(2), false, float64(2), ""},
		{"test 2", float64(2), false, float64(3), ""},
		{"test 3", float64(2), false, float64(1), "must be no less than 2"},
		{"test 4", float64(2), false, float64(0), ""},
		{"test 5", float64(2), true, float64(2), "must be greater than 2"},
	}

	for _, test := range tests {
		r := Min(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMinTime(t *testing.T) {
	date0 := time.Time{}
	date20000101 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date20001201 := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	date20000601 := time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		tag       string
		threshold time.Time
		exclusive bool
		value     time.Time
		err       string
	}{
		{"test 1", date20000601, false, date20000601, ""},
		{"test 2", date20000601, false, date20001201, ""},
		{"test 3", date20000601, false, date20000101, "must be no less than 2000-06-01 00:00:00 +0000 UTC"},
		{"test 4", date20000601, false, date0, ""},
		{"test 5", date20000601, true, date20000601, "must be greater than 2000-06-01 00:00:00 +0000 UTC"},
		{"test 6", date0, false, date20000601, ""},
	}

	for _, test := range tests {
		r := MinTime(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMinError(t *testing.T) {
	r := Min(10)
	assert.Equal(t, "must be no less than 10", r.Validate(9).Error())

	r = r.Error("123")
	assert.Equal(t, "123", r.err.Message())
}

func TestMaxInt(t *testing.T) {
	tests := []struct {
		tag       string
		threshold int
		exclusive bool
		value     int
		err       string
	}{
		{"test 1", 2, false, 2, ""},
		{"test 2", 2, false, 1, ""},
		{"test 3", 2, false, 3, "must be no greater than 2"},
		{"test 4", 2, false, 0, ""},
		{"test 5", 2, true, 2, "must be less than 2"},
	}

	for _, test := range tests {
		r := Max(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMaxUint(t *testing.T) {
	tests := []struct {
		tag       string
		threshold uint
		exclusive bool
		value     uint
		err       string
	}{
		{"test 1", uint(2), false, uint(2), ""},
		{"test 2", uint(2), false, uint(1), ""},
		{"test 3", uint(2), false, uint(3), "must be no greater than 2"},
		{"test 4", uint(2), false, uint(0), ""},
		{"test 5", uint(2), true, uint(2), "must be less than 2"},
	}

	for _, test := range tests {
		r := Max(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMaxFloat64(t *testing.T) {
	tests := []struct {
		tag       string
		threshold float64
		exclusive bool
		value     float64
		err       string
	}{
		{"test 1", float64(2), false, float64(2), ""},
		{"test 2", float64(2), false, float64(1), ""},
		{"test 3", float64(2), false, float64(3), "must be no greater than 2"},
		{"test 4", float64(2), false, float64(0), ""},
		{"test 5", float64(2), true, float64(2), "must be less than 2"},
	}

	for _, test := range tests {
		r := Max(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMaxTime(t *testing.T) {
	date0 := time.Time{}
	date20000101 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date20001201 := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	date20000601 := time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		tag       string
		threshold time.Time
		exclusive bool
		value     time.Time
		err       string
	}{
		{"test 1", date20000601, false, date20000601, ""},
		{"test 2", date20000601, false, date20000101, ""},
		{"test 3", date20000601, false, date20001201, "must be no greater than 2000-06-01 00:00:00 +0000 UTC"},
		{"test 4", date20000601, false, date0, ""},
		{"test 5", date20000601, true, date20000601, "must be less than 2000-06-01 00:00:00 +0000 UTC"},
	}

	for _, test := range tests {
		r := MaxTime(test.threshold)
		if test.exclusive {
			r = r.Exclusive()
		}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestMaxError(t *testing.T) {
	r := Max(10)
	assert.Equal(t, "must be no greater than 10", r.Validate(11).Error())

	r = r.Error("123")
	assert.Equal(t, "123", r.err.Message())
}

func TestThresholdRule_ErrorObject(t *testing.T) {
	r := Max(10)
	err := NewError("code", "abc")
	r = r.ErrorObject(err)

	assert.Equal(t, err, r.err)
	assert.Equal(t, err.Code(), r.err.Code())
	assert.Equal(t, err.Message(), r.err.Message())
}
