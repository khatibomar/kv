// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func abcValidation(val string) bool {
	return val == "abc"
}

func TestWhen(t *testing.T) {
	abcRule := NewStringRule(abcValidation, "wrong_abc")
	validateMeRule := NewStringRule(validateMe, "wrong_me")

	tests := []struct {
		tag       string
		condition bool
		value     any
		rules     []Rule[any]
		elseRules []Rule[any]
		err       string
	}{
		// True condition
		{"t1.1", true, nil, []Rule[any]{}, []Rule[any]{}, ""},
		{"t1.2", true, "", []Rule[any]{}, []Rule[any]{}, ""},
		{"t1.3", true, "", []Rule[any]{abcRule}, []Rule[any]{}, ""},
		{"t1.4", true, 12, []Rule[any]{Required}, []Rule[any]{}, ""},
		{"t1.5", true, nil, []Rule[any]{Required}, []Rule[any]{}, "cannot be blank"},
		{"t1.6", true, "123", []Rule[any]{abcRule}, []Rule[any]{}, "wrong_abc"},
		{"t1.7", true, "abc", []Rule[any]{abcRule}, []Rule[any]{}, ""},
		{"t1.8", true, "abc", []Rule[any]{abcRule, abcRule}, []Rule[any]{}, ""},
		{"t1.9", true, "abc", []Rule[any]{abcRule, validateMeRule}, []Rule[any]{}, "wrong_me"},
		{"t1.10", true, "me", []Rule[any]{abcRule, validateMeRule}, []Rule[any]{}, "wrong_abc"},
		{"t1.11", true, "me", []Rule[any]{}, []Rule[any]{abcRule}, ""},

		// False condition
		{"t2.1", false, "", []Rule[any]{}, []Rule[any]{}, ""},
		{"t2.2", false, "", []Rule[any]{abcRule}, []Rule[any]{}, ""},
		{"t2.3", false, "abc", []Rule[any]{abcRule}, []Rule[any]{}, ""},
		{"t2.4", false, "abc", []Rule[any]{abcRule, abcRule}, []Rule[any]{}, ""},
		{"t2.5", false, "abc", []Rule[any]{abcRule, validateMeRule}, []Rule[any]{}, ""},
		{"t2.6", false, "me", []Rule[any]{abcRule, validateMeRule}, []Rule[any]{}, ""},
		{"t2.7", false, "", []Rule[any]{abcRule, validateMeRule}, []Rule[any]{}, ""},
		{"t2.8", false, "me", []Rule[any]{}, []Rule[any]{abcRule, validateMeRule}, "wrong_abc"},
	}

	for _, test := range tests {
		err := Validate(test.value, When(test.condition, test.rules...).Else(test.elseRules...))
		assertError(t, test.err, err, test.tag)
	}
}

type ctxKey int

const (
	contains ctxKey = iota
)

func TestWhenWithContext(t *testing.T) {
	rule := WithContext(func(ctx context.Context, value any) error {
		if !strings.Contains(value.(string), ctx.Value(contains).(string)) {
			return errors.New("unexpected value")
		}
		return nil
	})
	ctx1 := context.WithValue(context.Background(), contains, "abc")
	ctx2 := context.WithValue(context.Background(), contains, "xyz")

	tests := []struct {
		tag       string
		condition bool
		value     any
		ctx       context.Context
		err       string
	}{
		// True condition
		{"t1.1", true, "abc", ctx1, ""},
		{"t1.2", true, "abc", ctx2, "unexpected value"},
		{"t1.3", true, "xyz", ctx1, "unexpected value"},
		{"t1.4", true, "xyz", ctx2, ""},

		// False condition
		{"t2.1", false, "abc", ctx1, ""},
		{"t2.2", false, "abc", ctx2, "unexpected value"},
		{"t2.3", false, "xyz", ctx1, "unexpected value"},
		{"t2.4", false, "xyz", ctx2, ""},
	}

	for _, test := range tests {
		err := ValidateWithContext(test.ctx, test.value, When(test.condition, rule).Else(rule))
		assertError(t, test.err, err, test.tag)
	}
}
