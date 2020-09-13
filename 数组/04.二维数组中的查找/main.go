package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
// 二维数组中的查找
// p44
func main() {
	testArr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	fmt.Println(findNumberIn2DArray(testArr, 19))
}

// 1.取右上角的值m
// 2. target大于m，则值肯定在下一行
// 3. target小于m, 则值肯定在左一列
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	i := 0
	j := len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if target == matrix[i][j] {
			return true
		}

		if target > matrix[i][j] {
			i++
		} else {
			j--
		}
	}
	return false
}
