package main

// 整数配列 `nums` と整数 `target` が与えられたとき、その合計が `target` になる2つの数値のインデックスを返してください。
// 各入力に対して必ず1つの解が存在し、同じ要素を2回以上使用することはできないものとします。
// 答えは任意の順序で返すことができます。
// https://leetcode.com/problems/two-sum/description/?envType=problem-list-v2&envId=plakya4j

func twoSum(nums []int, target int) []int {
	kv := make(map[int]int, len(nums))
	for i, n := range nums {
		diff := target - n
		if diffV, ok := kv[diff]; ok {
			return []int{diffV, i}
		}
		kv[n] = i
	}

	return nil
}
