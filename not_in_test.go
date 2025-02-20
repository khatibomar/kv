// Copyright 2016 Qiang Xue, Google LLC. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

import (
	"testing"

	"github.com/khatibomar/kv/internal/assert"
)

func TestNotIn(t *testing.T) {
	v := 1
	var v2 *int
	var tests = []struct {
		tag    string
		values []any
		value  any
		err    string
	}{
		{"t0", []any{1, 2}, 0, ""},
		{"t1", []any{1, 2}, 1, "must not be in list"},
		{"t2", []any{1, 2}, 2, "must not be in list"},
		{"t3", []any{1, 2}, 3, ""},
		{"t4", []any{}, 3, ""},
		{"t5", []any{1, 2}, "1", ""},
		{"t6", []any{1, 2}, &v, "must not be in list"},
		{"t7", []any{1, 2}, v2, ""},
	}

	for _, test := range tests {
		r := NotIn(test.values...)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func Test_NotInRule_Error(t *testing.T) {
	r := NotIn(1, 2, 3)
	assert.Equal(t, "must not be in list", r.Validate(1).Error())
	r = r.Error("123")
	assert.Equal(t, "123", r.err.Message())
}

func TestNotInRule_ErrorObject(t *testing.T) {
	r := NotIn(1, 2, 3)

	err := NewError("code", "abc")
	r = r.ErrorObject(err)

	assert.Equal(t, err, r.err)
	assert.Equal(t, err.Code(), r.err.Code())
	assert.Equal(t, err.Message(), r.err.Message())
}
