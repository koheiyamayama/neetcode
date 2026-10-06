package main

// 整数配列 nums が与えられたとき、answer[i] が「nums[i] を除く全要素の積」となる配列 answer を返してください。
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// https://leetcode.com/problems/product-of-array-except-self/
// 考え方
// 普通に考えると、二重ループで全探索するのが良さそう。
// だが、そうすると計算量がO(n**2)となってTLEしてしまう。
//
// func productExceptSelf(nums []int) []int {
// 	ans := make([]int, len(nums))
// 	for i1, _ := range nums {
// 		ans[i1] = 1
// 		for i2, n2 := range nums {
// 			if i1 == i2 {
// 				continue
// 			}
// 			ans[i1] *= n2
// 		}
// 	}

// 	return ans
// }

// 考え方2
// 配列内の要素を全て掛ける。
// 掛けた値をarray[i]で割ると答えになる？
// 0を考慮しづらい
func productExceptSelf(nums []int) []int {
	length := len(nums)
	left := make([]int, length)
	right := make([]int, length)
	ans := make([]int, length)
	for i := range nums {
		left[i] = 1
		if i == 0 { // 左端
		} else {
			left[i] = nums[i-1] * left[i-1]
		}

		ri := length - 1 - i

		right[ri] = 1
		if ri == length-1 { // 右端
		} else {
			right[ri] = nums[ri+1] * right[ri+1]
		}
	}

	for i := range length {
		ans[i] = left[i] * right[i]
	}

	return ans
}
