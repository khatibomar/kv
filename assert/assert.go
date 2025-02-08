package assert

import (
	"reflect"
	"testing"
)

func Nil(t *testing.T, value any, msgAndArgs ...any) bool {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return true
	}

	t.Errorf("Expected nil, but got: %v", value)
	return false
}

func NotNil(t *testing.T, value any, msgAndArgs ...any) bool {
	if value != nil || !reflect.ValueOf(value).IsNil() {
		return true
	}
	t.Error("Expected value not to be nil")
	return false
}

func Equal(t *testing.T, expected, actual any, msgAndArgs ...any) bool {
	// Handle nil case first
	if expected == nil && actual == nil {
		return true
	}

	// Special handling for errors
	expectedErr, expectedIsErr := expected.(error)
	actualErr, actualIsErr := actual.(error)

	if expectedIsErr || actualIsErr {
		if !expectedIsErr || !actualIsErr {
			t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			return false
		}
		if expectedErr.Error() != actualErr.Error() {
			t.Errorf("Not equal: expected error: %v, actual error: %v", expectedErr, actualErr)
			return false
		}
		return true
	}

	// Special handling for []uint8 ([]byte)
	expectedBytes, expectedIsBytes := expected.([]uint8)
	actualBytes, actualIsBytes := actual.([]uint8)

	if expectedIsBytes || actualIsBytes {
		if !expectedIsBytes || !actualIsBytes {
			t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			return false
		}
		if len(expectedBytes) != len(actualBytes) {
			t.Errorf("Not equal: byte slices have different lengths. expected: %d, actual: %d", len(expectedBytes), len(actualBytes))
			return false
		}
		for i := range expectedBytes {
			if expectedBytes[i] != actualBytes[i] {
				t.Errorf("Not equal: byte mismatch at index %d: expected %d, got %d", i, expectedBytes[i], actualBytes[i])
				return false
			}
		}
		return true
	}

	// Special handling for maps
	expectedMap, expectedIsMap := expected.(map[string]any)
	actualMap, actualIsMap := actual.(map[string]any)

	if expectedIsMap || actualIsMap {
		if !expectedIsMap || !actualIsMap {
			t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			return false
		}
		if len(expectedMap) != len(actualMap) {
			t.Errorf("Not equal: maps have different lengths. expected: %d, actual: %d", len(expectedMap), len(actualMap))
			return false
		}
		for key, expectedVal := range expectedMap {
			actualVal, exists := actualMap[key]
			if !exists {
				t.Errorf("Not equal: key %q exists in expected but not in actual", key)
				return false
			}
			if !Equal(t, expectedVal, actualVal) {
				t.Errorf("Not equal: value mismatch for key %q", key)
				return false
			}
		}
		return true
	}

	// Special handling for slices
	expectedSlice, expectedIsSlice := expected.([]any)
	actualSlice, actualIsSlice := actual.([]any)

	if expectedIsSlice || actualIsSlice {
		if !expectedIsSlice || !actualIsSlice {
			t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			return false
		}
		if len(expectedSlice) != len(actualSlice) {
			t.Errorf("Not equal: slices have different lengths. expected: %d, actual: %d", len(expectedSlice), len(actualSlice))
			return false
		}
		for i := range expectedSlice {
			if !Equal(t, expectedSlice[i], actualSlice[i]) {
				t.Errorf("Not equal: value mismatch at index %d", i)
				return false
			}
		}
		return true
	}

	// Normal comparison for other types
	if expected == actual {
		return true
	}
	t.Errorf("Not equal: expected: %v, actual: %v", expected, actual)
	return false
}

func NotEqual(t *testing.T, expected, actual any, msgAndArgs ...any) bool {
	if expected != actual {
		return true
	}
	t.Errorf("Values should not be equal: %v", expected)
	return false
}

func False(t *testing.T, actual any) bool {
	if actual == false {
		return true
	}
	t.Errorf("Expected false, but got: %v", actual)
	return false
}

func True(t *testing.T, actual any) bool {
	if actual == true {
		return true
	}
	t.Errorf("Expected true, but got: %v", actual)
	return false
}

func NoError(t *testing.T, err error, msgAndArgs ...any) bool {
	if err == nil {
		return true
	}
	t.Errorf("Expected no error, but got: %v", err)
	return false
}

func EqualError(t *testing.T, err error, expectedError string, msgAndArgs ...any) bool {
	if err == nil {
		t.Errorf("Expected error message: %s but got nil", expectedError)
		return false
	}
	if err.Error() != expectedError {
		t.Errorf("Expected error message: %s but got: %s", expectedError, err.Error())
		return false
	}
	return true
}
