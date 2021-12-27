package search

import (
	"testing"
)

func TestLinearSearchIntegers(t *testing.T) {
	for _, test := range searchTestsIntegers {
		actualValue, actualError := LinearSearchIntegers(test.data, test.key)
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

func BenchmarkLinearSearchIntegers(b *testing.B) {
	testCase := generateBenchmarkTestCaseIntegers()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = LinearSearchIntegers(testCase, i)
	}
}

func TestLinearSearchGeneral(t *testing.T) {
	generateSearchTestGeneral()
	for _, test := range searchTestsGeneral {
		actualValue, actualError := LinearSearchGeneral(test.data, test.key)
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

func BenchmarkLinearSearchGeneral(b *testing.B) {
	testCase := generateBenchmarkTestCaseGeneral()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = LinearSearchGeneral(testCase, testDataObj{key: i, val: "foo"})
	}
}
