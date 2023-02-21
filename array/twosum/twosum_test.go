package twosum

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/two-sum/description/

func TestTwoSumExample1(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("twoSum = %d", result)
	}
}

func TestTwoSumExample2(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("twoSum = %d", result)
	}
}

func TestTwoSumExample3(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("twoSum = %d", result)
	}
}

func TestTwoSumExample4(t *testing.T) {
	result := twoSum([]int{0, 4, 3, 0}, 0)
	if !reflect.DeepEqual(result, []int{0, 3}) {
		t.Errorf("twoSum = %d", result)
	}
}
