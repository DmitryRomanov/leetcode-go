package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("twoSum = %d", result)
	}
}

func TestTwoSumSix(t *testing.T) {
	result := twoSum([]int{3, 24, 50, 79, 88, 150, 345}, 200)
	if !reflect.DeepEqual(result, []int{3, 6}) {
		t.Errorf("twoSum = %d", result)
	}
}
