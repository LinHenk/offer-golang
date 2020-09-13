package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// p39
func main() {
	testArr := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(testArr))

}

// 所有的数字都是 0 ~ n-1，如果这个数组中没有重复的数字，那么数组下标i下的值必然是i
// 从头到尾遍历数组
// 判断nums[i]是否等于i
// 等于 则表明i，在数组下标i下
// 不等于
// 		m = nums[i]
//      判断 m 和nums[m]的值是否相等
//      	相等，则表明已经存在m已经在nums[m]下了，则m是重复元素
// 			不等，则交换nums[i]和nums[nums[i]]的值
// p39
func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			// 第i个位置的数据，和第m个位置上的数相等
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			//不等，则将第i个位置上的数m，放到第m的位置上
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// 不修改数组找出重复的数字
// 长度为n+1的数组，里面的数字是1-n，所以数组肯定有一个重复的数字
// 1. 创建一个新的n+1数组，在位置i上防元素i，如果i位置上已经存在元素i则表示存在，但是这样需要使用O(n)的空间
// 2. 把1-n的数字从中间的数字m拆分为两部分，前面一半是1-m，后面一半是m+1-n
//    计算前半部分的元素出现的次数是否大于m，大于m则重复数字在前半部分，否则在后半部分
//	  一直循环以上的防范，指导start和end重合
// p41
func findRepeatNumber2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (end-start)/2 + start
		//计算数组在前半部分元素出现的次数
		count := countRange(nums, start, mid)

		// 如果两个位置重合了，并且元素出现的次数大于1，那么这个数就是重复元素
		if end == start {
			if count > 1 {
				return start
			}
		}

		if count > (mid - start + 1) {
			//在前半部分
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

// 计算在某个范围出现的次数
func countRange(nums []int, start, end int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}
