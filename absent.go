// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

var (
	// ErrNil is the error that returns when a value is not nil.
	ErrNil = NewError("validation_nil", "must be blank")
	// ErrEmpty is the error that returns when a not nil value is not empty.
	ErrEmpty = NewError("validation_empty", "must be blank")
)

type NilRule[T comparable] struct {
	condition bool
	err       Error
}

type EmptyRule[T comparable] struct {
	condition bool
	err       Error
}

// Validate checks if the given value is valid or not.
func (r EmptyRule[T]) Validate(value *T) error {
	if r.condition {
		return nil
	}

	if value == nil {
		return nil
	}

	var zero T
	if *value == zero {
		return nil
	}
	return r.getError(ErrEmpty)
}

// When sets the condition that determines if the validation should be performed.
func (r EmptyRule[T]) When(condition bool) EmptyRule[T] {
	r.condition = !condition
	return r
}

// Error sets the error message for the rule.
func (r EmptyRule[T]) Error(message string) EmptyRule[T] {
	if r.err == nil {
		r.err = ErrEmpty
	}
	r.err = r.err.SetMessage(message)
	return r
}

// ErrorObject sets the error struct for the rule.
func (r EmptyRule[T]) ErrorObject(err Error) EmptyRule[T] {
	r.err = err
	return r
}

// getError returns the custom error if set, otherwise returns the default error
func (r EmptyRule[T]) getError(defaultErr Error) error {
	if r.err != nil {
		return r.err
	}
	return defaultErr
}

// Validate checks if the given value is valid or not.
func (r NilRule[T]) Validate(value *T) error {
	if r.condition {
		return nil
	}

	if value == nil {
		return nil
	}

	return r.getError(ErrNil)
}

// When sets the condition that determines if the validation should be performed.
func (r NilRule[T]) When(condition bool) NilRule[T] {
	r.condition = !condition
	return r
}

// Error sets the error message for the rule.
func (r NilRule[T]) Error(message string) NilRule[T] {
	if r.err == nil {
		r.err = ErrNil
	}
	r.err = r.err.SetMessage(message)
	return r
}

// ErrorObject sets the error struct for the rule.
func (r NilRule[T]) ErrorObject(err Error) NilRule[T] {
	r.err = err
	return r
}

// getError returns the custom error if set, otherwise returns the default error
func (r NilRule[T]) getError(defaultErr Error) error {
	if r.err != nil {
		return r.err
	}
	return defaultErr
}
