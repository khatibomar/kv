// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

import (
	"testing"

	"github.com/khatibomar/kv/internal/assert"
)

func TestIn(t *testing.T) {
	var v = 1
	var v2 *int
	tests := []struct {
		tag    string
		values []any
		value  any
		err    string
	}{
		{"t0", []any{1, 2}, 0, ""},
		{"t1", []any{1, 2}, 1, ""},
		{"t2", []any{1, 2}, 2, ""},
		{"t3", []any{1, 2}, 3, "must be a valid value"},
		{"t4", []any{}, 3, "must be a valid value"},
		{"t5", []any{1, 2}, "1", "must be a valid value"},
		{"t6", []any{1, 2}, &v, ""},
		{"t7", []any{1, 2}, v2, ""},
		{"t8", []any{[]byte{1}, 1, 2}, []byte{1}, ""},
	}

	for _, test := range tests {
		r := In(test.values...)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func Test_InRule_Error(t *testing.T) {
	r := In(1, 2, 3)
	val := 4
	assert.Equal(t, "must be a valid value", r.Validate(&val).Error())
	r = r.Error("123")
	assert.Equal(t, "123", r.err.Message())
}

func TestInRule_ErrorObject(t *testing.T) {
	r := In(1, 2, 3)

	err := NewError("code", "abc")
	r = r.ErrorObject(err)

	assert.Equal(t, err, r.err)
	assert.Equal(t, err.Code(), r.err.Code())
	assert.Equal(t, err.Message(), r.err.Message())
}
