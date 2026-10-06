package main

// 整数配列 nums と整数 k が与えられたとき、出現頻度が上位 k 個の要素を返してください。答えの順序は問いません。
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// https://leetcode.com/problems/top-k-frequent-elements/
// 考え方
// 出現する数値の個数をカウントする。
// 降順で並べて上位k件を取得する。
// 実装方針1
// mapで集計する O(n)
// mapをループしてk個の要素を配列に入れていく。入れる条件は以下。
// 配列を走査して小さい値がある場合、最も小さい値と入れ替える。 O(n**2)
// mapを全走査し終えたら、終了する

type node struct {
	cnt int
	num int
}

// func topKFrequent(nums []int, k int) []int {
// 	group := make(map[int]*node, len(nums))
// 	for _, n := range nums {
// 		if no, ok := group[n]; ok {
// 			no.cnt++
// 		} else {
// 			group[n] = &node{cnt: 1, num: n}
// 		}
// 	}

// 	sorted := make([]*node, len(group))

// 	i := 0
// 	for _, v := range group {
// 		sorted[i] = v
// 		i++
// 	}

// 	sort.Slice(sorted, func(i, j int) bool {
// 		return sorted[i].cnt > sorted[j].cnt
// 	})

// 	ans := make([]int, k)
// 	for i, v := range sorted[0:k] {
// 		ans[i] = v.num
// 	}
// 	return ans
// }

func topKFrequent(nums []int, k int) []int {
	kv := make(map[int]int, len(nums))
	for _, n := range nums {
		kv[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, cnt := range kv {
		bucket[cnt] = append(bucket[cnt], key)
	}

	ans := make([]int, 0)
	for i := range len(bucket) - 1 {
		i = len(bucket) - i - 1
		for _, n := range bucket[i] {
			if len(ans) == k {
				break
			}
			ans = append(ans, n)
		}
		if len(ans) == k {
			break
		}
	}

	return ans
}
