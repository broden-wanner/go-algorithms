package search

// Struct for testing general linear search
type testDataObj struct {
	key int
	val string
}

// EqualTo method for the testDataObj
func (curr testDataObj) EqualTo(other interface{}) bool {
	if otherTdo, ok := other.(testDataObj); ok {
		return otherTdo.key == curr.key
	}
	return false
}

type searchTestInt struct {
	data          []int
	key           int
	expected      int
	expectedError error
	name          string
}

type searchTestGeneral struct {
	data          []Equaler
	key           Equaler
	expected      int
	expectedError error
	name          string
}

var searchTestsIntegers = []searchTestInt{
	// Sanity
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, "Sanity"},
	// Absent
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, ErrorNotFound, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrorNotFound, "Absent"},
	// Empty slice
	{[]int{}, 2, -1, ErrorNotFound, "Empty"},
}

// This function generate consistent testcase for benchmark test.
func generateBenchmarkTestCaseIntegers() []int {
	var testCase []int
	for i := 0; i < 1000; i++ {
		testCase = append(testCase, i)
	}
	return testCase
}

var searchTestsGeneral []searchTestGeneral

func generateSearchTestGeneral() {
	var testArray []Equaler
	for i := 0; i < 20; i++ {
		testArray = append(testArray, testDataObj{key: i + 1, val: "foo"})
	}
	// Generate success cases
	for i := 0; i < 20; i++ {
		testCase := searchTestGeneral{
			data:          testArray,
			key:           testDataObj{key: i + 1, val: "bar"},
			expected:      i,
			expectedError: nil,
			name:          "Sanity",
		}
		searchTestsGeneral = append(searchTestsGeneral, testCase)
	}
	// Generate failure cases
	searchTestsGeneral = append(searchTestsGeneral,
		searchTestGeneral{
			data:          testArray,
			key:           testDataObj{key: 1000, val: "bar"},
			expected:      -1,
			expectedError: ErrorNotFound,
			name:          "Absent",
		},
	)
	searchTestsGeneral = append(searchTestsGeneral,
		searchTestGeneral{
			data:          testArray,
			key:           testDataObj{key: -1000, val: "bar"},
			expected:      -1,
			expectedError: ErrorNotFound,
			name:          "Absent",
		},
	)
	// Empty slice
	searchTestsGeneral = append(searchTestsGeneral,
		searchTestGeneral{
			data:          []Equaler{},
			key:           testDataObj{key: -1000, val: "bar"},
			expected:      -1,
			expectedError: ErrorNotFound,
			name:          "Empty",
		},
	)
}

// This function generate consistent testcase for benchmark test.
func generateBenchmarkTestCaseGeneral() []Equaler {
	var testCase []Equaler
	for i := 0; i < 1000; i++ {
		obj := testDataObj{key: i, val: "foo"}
		testCase = append(testCase, obj)
	}
	return testCase
}
