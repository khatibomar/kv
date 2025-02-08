package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func formatMessage(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	if len(msgAndArgs) == 1 {
		return fmt.Sprintf("%v", msgAndArgs[0])
	}
	if format, ok := msgAndArgs[0].(string); ok {
		return fmt.Sprintf(format, msgAndArgs[1:]...)
	}
	return fmt.Sprint(msgAndArgs...)
}

func Nil(t *testing.T, value any, msgAndArgs ...any) bool {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Expected nil, but got: %v", msg, value)
	} else {
		t.Errorf("Expected nil, but got: %v", value)
	}
	return false
}

func NotNil(t *testing.T, value any, msgAndArgs ...any) bool {
	if value != nil || !reflect.ValueOf(value).IsNil() {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Expected value not to be nil", msg)
	} else {
		t.Error("Expected value not to be nil")
	}
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
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: expected type: %T, actual type: %T", msg, expected, actual)
			} else {
				t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			}
			return false
		}
		if expectedErr.Error() != actualErr.Error() {
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: expected error: %v, actual error: %v", msg, expectedErr, actualErr)
			} else {
				t.Errorf("Not equal: expected error: %v, actual error: %v", expectedErr, actualErr)
			}
			return false
		}
		return true
	}

	// Special handling for []uint8 ([]byte)
	expectedBytes, expectedIsBytes := expected.([]uint8)
	actualBytes, actualIsBytes := actual.([]uint8)

	if expectedIsBytes || actualIsBytes {
		if !expectedIsBytes || !actualIsBytes {
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: expected type: %T, actual type: %T", msg, expected, actual)
			} else {
				t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			}
			return false
		}
		if len(expectedBytes) != len(actualBytes) {
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: byte slices have different lengths. expected: %d, actual: %d", msg, len(expectedBytes), len(actualBytes))
			} else {
				t.Errorf("Not equal: byte slices have different lengths. expected: %d, actual: %d", len(expectedBytes), len(actualBytes))
			}
			return false
		}
		for i := range expectedBytes {
			if expectedBytes[i] != actualBytes[i] {
				if msg := formatMessage(msgAndArgs...); msg != "" {
					t.Errorf("%s: Not equal: byte mismatch at index %d: expected %d, got %d", msg, i, expectedBytes[i], actualBytes[i])
				} else {
					t.Errorf("Not equal: byte mismatch at index %d: expected %d, got %d", i, expectedBytes[i], actualBytes[i])
				}
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
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: expected type: %T, actual type: %T", msg, expected, actual)
			} else {
				t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			}
			return false
		}
		if len(expectedMap) != len(actualMap) {
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: maps have different lengths. expected: %d, actual: %d", msg, len(expectedMap), len(actualMap))
			} else {
				t.Errorf("Not equal: maps have different lengths. expected: %d, actual: %d", len(expectedMap), len(actualMap))
			}
			return false
		}
		for key, expectedVal := range expectedMap {
			actualVal, exists := actualMap[key]
			if !exists {
				if msg := formatMessage(msgAndArgs...); msg != "" {
					t.Errorf("%s: Not equal: key %q exists in expected but not in actual", msg, key)
				} else {
					t.Errorf("Not equal: key %q exists in expected but not in actual", key)
				}
				return false
			}
			if !Equal(t, expectedVal, actualVal) {
				if msg := formatMessage(msgAndArgs...); msg != "" {
					t.Errorf("%s: Not equal: value mismatch for key %q", msg, key)
				} else {
					t.Errorf("Not equal: value mismatch for key %q", key)
				}
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
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: expected type: %T, actual type: %T", msg, expected, actual)
			} else {
				t.Errorf("Not equal: expected type: %T, actual type: %T", expected, actual)
			}
			return false
		}
		if len(expectedSlice) != len(actualSlice) {
			if msg := formatMessage(msgAndArgs...); msg != "" {
				t.Errorf("%s: Not equal: slices have different lengths. expected: %d, actual: %d", msg, len(expectedSlice), len(actualSlice))
			} else {
				t.Errorf("Not equal: slices have different lengths. expected: %d, actual: %d", len(expectedSlice), len(actualSlice))
			}
			return false
		}
		for i := range expectedSlice {
			if !Equal(t, expectedSlice[i], actualSlice[i]) {
				if msg := formatMessage(msgAndArgs...); msg != "" {
					t.Errorf("%s: Not equal: value mismatch at index %d", msg, i)
				} else {
					t.Errorf("Not equal: value mismatch at index %d", i)
				}
				return false
			}
		}
		return true
	}

	// Normal comparison for other types
	if expected == actual {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Not equal: expected: %v, actual: %v", msg, expected, actual)
	} else {
		t.Errorf("Not equal: expected: %v, actual: %v", expected, actual)
	}
	return false
}

func NotEqual(t *testing.T, expected, actual any, msgAndArgs ...any) bool {
	if expected != actual {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Values should not be equal: %v", msg, expected)
	} else {
		t.Errorf("Values should not be equal: %v", expected)
	}
	return false
}

func False(t *testing.T, actual any, msgAndArgs ...any) bool {
	if actual == false {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Expected false, but got: %v", msg, actual)
	} else {
		t.Errorf("Expected false, but got: %v", actual)
	}
	return false
}

func True(t *testing.T, actual any, msgAndArgs ...any) bool {
	if actual == true {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Expected true, but got: %v", msg, actual)
	} else {
		t.Errorf("Expected true, but got: %v", actual)
	}
	return false
}

func NoError(t *testing.T, err error, msgAndArgs ...any) bool {
	if err == nil {
		return true
	}
	if msg := formatMessage(msgAndArgs...); msg != "" {
		t.Errorf("%s: Expected no error, but got: %v", msg, err)
	} else {
		t.Errorf("Expected no error, but got: %v", err)
	}
	return false
}

func EqualError(t *testing.T, err error, expectedError string, msgAndArgs ...any) bool {
	if err == nil {
		if msg := formatMessage(msgAndArgs...); msg != "" {
			t.Errorf("%s: Expected error message: %s but got nil", msg, expectedError)
		} else {
			t.Errorf("Expected error message: %s but got nil", expectedError)
		}
		return false
	}
	if err.Error() != expectedError {
		if msg := formatMessage(msgAndArgs...); msg != "" {
			t.Errorf("%s: Expected error message: %s but got: %s", msg, expectedError, err.Error())
		} else {
			t.Errorf("Expected error message: %s but got: %s", expectedError, err.Error())
		}
		return false
	}
	return true
}
