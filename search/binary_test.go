package search

import (
	"testing"
)

func TestBinarySearchIterativeIntegers(t *testing.T) {
	for _, test := range searchTestsIntegers {
		actualValue, actualError := BinarySearchIterativeIntegers(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'",
				test.name, test.data, test.key, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'",
				test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func BenchmarkBinarySearchIterativeIntegers(b *testing.B) {
	testCase := generateBenchmarkTestCaseIntegers()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = BinarySearchIterativeIntegers(testCase, i)
	}
}

func TestBinarySearchRecursiveIntegers(t *testing.T) {
	for _, test := range searchTestsIntegers {
		actualValue, actualError := BinarySearchRecursiveIntegers(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'",
				test.name, test.data, test.key, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'",
				test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func BenchmarkBinarySearchRecursiveIntegers(b *testing.B) {
	testCase := generateBenchmarkTestCaseIntegers()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = BinarySearchRecursiveIntegers(testCase, i)
	}
}
