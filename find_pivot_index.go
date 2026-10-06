package main

// 整数配列 nums が与えられたとき、左側の和と右側の和が等しくなる最も左の位置 i を返してください。
// i 自身はどちらにも含めません。そのような位置がなければ -1 を返してください。
// 左端では左の和は 0、右端では右の和は 0 です。
// Input: nums = [1,7,3,6,5,6]
// Output: 3
// Input: nums = [1,2,3]
// Output: -1
// Input: nums = [2,1,-1]
// Output: 0
// https://leetcode.com/problems/find-pivot-index/
// 考え方
// iの両隣が一致すること(i-1==i+1)を探索する。
// そのためにnumsを左からnums[i]+nums[i+1]、右からnums[len-1-i]+nums[len-2-i]を実施して=になるiを探す
func pivotIndex(nums []int) int {
	lacum := make([]int, len(nums))
	for i := range nums {
		l := 0
		if i == 0 { // 左端
			l = nums[i]
		} else {
			l = nums[i] + lacum[i-1]
		}
		lacum[i] = l
	}

	racum := make([]int, len(nums))
	for i := range nums {
		ri := len(nums) - 1 - i
		r := 0
		if ri == len(nums)-1 { // 右端
			r = nums[ri]
		} else {
			r = nums[ri] + racum[ri+1]
		}
		racum[ri] = r
	}

	for i := range nums {
		if racum[i] == lacum[i] {
			return i
		}
	}

	return -1
}
