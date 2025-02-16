package kv

import (
	"testing"
	"time"

	"github.com/khatibomar/kv/internal/assert"
)

func TestNilString(t *testing.T) {
	s1 := "123"
	s2 := ""
	tests := []struct {
		tag   string
		value *string
		err   string
	}{
		{"t3", &s1, "must be blank"},
		{"t4", &s2, "must be blank"},
		{"t5", nil, ""},
	}

	for _, test := range tests {
		r := NilRule[string]{}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestNilTime(t *testing.T) {
	time1 := time.Now()
	tests := []struct {
		tag   string
		value *time.Time
		err   string
	}{
		{"t6", &time1, "must be blank"},
		{"t7", nil, ""},
	}

	for _, test := range tests {
		r := NilRule[time.Time]{}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestEmptyString(t *testing.T) {
	s1 := "123"
	s2 := ""
	tests := []struct {
		tag   string
		value *string
		err   string
	}{
		{"t3", &s1, "must be blank"},
		{"t4", &s2, ""},
		{"t5", nil, ""},
	}

	for _, test := range tests {
		r := EmptyRule[string]{}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestEmptyTime(t *testing.T) {
	time1 := time.Now()
	var time2 time.Time
	tests := []struct {
		tag   string
		value *time.Time
		err   string
	}{
		{"t6", &time1, "must be blank"},
		{"t7", &time2, ""},
		{"t8", nil, ""},
	}

	for _, test := range tests {
		r := EmptyRule[time.Time]{}
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestAbsentRule_When(t *testing.T) {
	val := 42
	r := NilRule[int]{}.When(false)
	err := r.Validate(&val)
	assert.Nil(t, err)

	r = NilRule[int]{}.When(true)
	err = r.Validate(&val)
	assert.Equal(t, ErrNil, err)
}

func Test_absentRule_Error(t *testing.T) {
	val := "42"
	r := NilRule[string]{}
	assert.Equal(t, "must be blank", r.Validate(&val).Error())
	r2 := r.Error("123")
	assert.Equal(t, "must be blank", r.Validate(&val).Error())
	assert.Equal(t, "123", r2.err.Message())

	e := EmptyRule[string]{}
	assert.Equal(t, "must be blank", e.Validate(&val).Error())
	r2 = r.Error("123")
	assert.Equal(t, "must be blank", e.Validate(&val).Error())
	assert.Equal(t, "123", r2.err.Message())
}

func TestAbsentRule_ErrorObject(t *testing.T) {
	r := NilRule[string]{}

	err := NewError("code", "abc")
	r = r.ErrorObject(err)

	assert.Equal(t, err, r.err)
	assert.Equal(t, err.Code(), r.err.Code())
	assert.Equal(t, err.Message(), r.err.Message())
	assert.NotEqual(t, err, NilRule[string]{}.err)
}
