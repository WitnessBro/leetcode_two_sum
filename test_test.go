package leetcodetwosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var nums = []int{3, 2, 4}
	var expectedOutput = []int{1, 2}
	var target = 6
	var k = twoSum(nums, target)

	assert.Equal(t, k, expectedOutput)
}
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		secondNumber := target - num

		if index, ok := numMap[secondNumber]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{}
}
