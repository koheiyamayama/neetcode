package main

// 整数配列 nums が与えられたとき、配列内に少なくとも2回以上出現する値があれば true を返し、すべての要素が一意であれば false を返す関数を作成してください。
// https: //leetcode.com/problems/contains-duplicate/description/

func containsDuplicate(nums []int) bool {
	kv := make(map[int]int, len(nums))

	for _, n := range nums {
		kv[n]++
		v := kv[n]
		if v >= 2 {
			return true
		}
	}

	return false
}
