package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
//uint32
var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{-1, 2, 3},
	{4, 5, 9},
	{120, 1, 119},
}

func TestSumuUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//int32
var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//int64
var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//float64
var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%b, %b) returned %b, expected %b", v.n1, v.n2, val, v.expected)
		}
	}
}
