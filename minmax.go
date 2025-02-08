// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kv

import "time"

var (
	// ErrMinGreaterEqualThanRequired is the error that returns when a value is less than a specified threshold.
	ErrMinGreaterEqualThanRequired = NewError("validation_min_greater_equal_than_required", "must be no less than {{.threshold}}")
	// ErrMaxLessEqualThanRequired is the error that returns when a value is greater than a specified threshold.
	ErrMaxLessEqualThanRequired = NewError("validation_max_less_equal_than_required", "must be no greater than {{.threshold}}")
	// ErrMinGreaterThanRequired is the error that returns when a value is less than or equal to a specified threshold.
	ErrMinGreaterThanRequired = NewError("validation_min_greater_than_required", "must be greater than {{.threshold}}")
	// ErrMaxLessThanRequired is the error that returns when a value is greater than or equal to a specified threshold.
	ErrMaxLessThanRequired = NewError("validation_max_less_than_required", "must be less than {{.threshold}}")
)

// ThresholdRule is a validation rule that checks if a value satisfies the specified threshold requirement.
type ThresholdRule[T Ordered] struct {
	threshold T
	operator  int
	err       Error
}

type TimeThresholdRule struct {
	threshold time.Time
	operator  int
	err       Error
}

const (
	greaterThan = iota
	greaterEqualThan
	lessThan
	lessEqualThan
)

// Min returns a validation rule that checks if a value is greater or equal than the specified value.
// By calling Exclusive, the rule will check if the value is strictly greater than the specified value.
func Min[T Ordered](min T) ThresholdRule[T] {
	return ThresholdRule[T]{
		threshold: min,
		operator:  greaterEqualThan,
		err:       ErrMinGreaterEqualThanRequired,
	}
}

// MinTime returns a validation rule that checks if a time is greater or equal than the specified time.
// By calling Exclusive, the rule will check if the time is strictly greater than the specified time.
func MinTime(min time.Time) TimeThresholdRule {
	return TimeThresholdRule{
		threshold: min,
		operator:  greaterEqualThan,
		err:       ErrMinGreaterEqualThanRequired,
	}
}

// Max returns a validation rule that checks if a value is less or equal than the specified value.
// By calling Exclusive, the rule will check if the value is strictly less than the specified value.
func Max[T Ordered](max T) ThresholdRule[T] {
	return ThresholdRule[T]{
		threshold: max,
		operator:  lessEqualThan,
		err:       ErrMaxLessEqualThanRequired,
	}
}

// MaxTime returns a validation rule that checks if a time is less or equal than the specified time.
// By calling Exclusive, the rule will check if the time is strictly less than the specified time.
func MaxTime(max time.Time) TimeThresholdRule {
	return TimeThresholdRule{
		threshold: max,
		operator:  lessEqualThan,
		err:       ErrMaxLessEqualThanRequired,
	}
}

// Exclusive sets the comparison to exclude the boundary value.
func (r ThresholdRule[T]) Exclusive() ThresholdRule[T] {
	if r.operator == greaterEqualThan {
		r.operator = greaterThan
		r.err = ErrMinGreaterThanRequired
	} else if r.operator == lessEqualThan {
		r.operator = lessThan
		r.err = ErrMaxLessThanRequired
	}
	return r
}

// Exclusive sets the comparison to exclude the boundary value.
func (r TimeThresholdRule) Exclusive() TimeThresholdRule {
	if r.operator == greaterEqualThan {
		r.operator = greaterThan
		r.err = ErrMinGreaterThanRequired
	} else if r.operator == lessEqualThan {
		r.operator = lessThan
		r.err = ErrMaxLessThanRequired
	}
	return r
}

// Validate checks if the given value is valid or not.
func (r ThresholdRule[T]) Validate(value T) error {
	if IsEmpty(value) {
		return nil
	}

	valid := false
	switch r.operator {
	case greaterThan:
		valid = value > r.threshold
	case greaterEqualThan:
		valid = value >= r.threshold
	case lessThan:
		valid = value < r.threshold
	case lessEqualThan:
		valid = value <= r.threshold
	}

	if !valid {
		return r.err.SetParams(map[string]any{"threshold": r.threshold})
	}
	return nil
}

// Validate checks if the given value is valid or not.
func (r TimeThresholdRule) Validate(value time.Time) error {
	if IsEmpty(value) {
		return nil
	}

	valid := false
	switch r.operator {
	case greaterThan:
		valid = value.After(r.threshold)
	case greaterEqualThan:
		valid = value.After(r.threshold) || value.Equal(r.threshold)
	case lessThan:
		valid = value.Before(r.threshold)
	case lessEqualThan:
		valid = value.Before(r.threshold) || value.Equal(r.threshold)
	}

	if !valid {
		return r.err.SetParams(map[string]any{"threshold": r.threshold})
	}
	return nil
}

// Error sets the error message for the rule.
func (r ThresholdRule[T]) Error(message string) ThresholdRule[T] {
	r.err = r.err.SetMessage(message)
	return r
}

// Error sets the error message for the rule.
func (r TimeThresholdRule) Error(message string) TimeThresholdRule {
	r.err = r.err.SetMessage(message)
	return r
}

// ErrorObject sets the error struct for the rule.
func (r ThresholdRule[T]) ErrorObject(err Error) ThresholdRule[T] {
	r.err = err
	return r
}

// ErrorObject sets the error struct for the rule.
func (r TimeThresholdRule) ErrorObject(err Error) TimeThresholdRule {
	r.err = err
	return r
}
