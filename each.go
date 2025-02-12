// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

import (
	"context"
	"errors"
	"reflect"
	"strconv"
)

// Each returns a validation rule that loops through an iterable (map, slice or array)
// and validates each value inside with the provided rules.
// An empty iterable is considered valid. Use the Required rule to make sure the iterable is not empty.
func Each(rules ...Rule[any]) EachRule {
	return EachRule{
		rules: rules,
	}
}

// EachRule is a validation rule that validates elements in a map/slice/array using the specified list of rules.
type EachRule struct {
	rules []Rule[any]
}

// Validate loops through the given iterable and calls the KV Validate() method for each value.
func (r EachRule) Validate(value any) error {
	return r.ValidateWithContext(context.TODO(), value)
}

// ValidateWithContext loops through the given iterable and calls the KV ValidateWithContext() method for each value.
func (r EachRule) ValidateWithContext(ctx context.Context, value any) error {
	errs := Errors{}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			val := r.getInterface(v.MapIndex(k))
			var err error
			if ctx == nil {
				err = Validate(val, r.rules...)
			} else {
				err = ValidateWithContext(ctx, val, r.rules...)
			}
			if err != nil {
				errs[r.getString(k)] = err
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			val := r.getInterface(v.Index(i))
			var err error
			if ctx == nil {
				err = Validate(val, r.rules...)
			} else {
				err = ValidateWithContext(ctx, val, r.rules...)
			}
			if err != nil {
				errs[strconv.Itoa(i)] = err
			}
		}
	default:
		return errors.New("must be an iterable (map, slice or array)")
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (r EachRule) getInterface(value reflect.Value) any {
	switch value.Kind() {
	case reflect.Ptr, reflect.Interface:
		if value.IsNil() {
			return nil
		}
		return value.Elem().Interface()
	default:
		return value.Interface()
	}
}

func (r EachRule) getString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Ptr, reflect.Interface:
		if value.IsNil() {
			return ""
		}
		return value.Elem().String()
	default:
		return value.String()
	}
}
