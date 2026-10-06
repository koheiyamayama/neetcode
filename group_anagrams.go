package main

// 文字列の配列 `strs` が与えられたとき、アナグラム同士をグループ化してください。結果は任意の順序で返して構いません。
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 考え方
// アナグラム = 26文字のバリエーションと個数を数えた時、一致すること。
// 26文字を格納する配列を作り、0:a, 1:b....z: 26とする。要素にカウント数値を入れれば、その文字列のバリエーションと個数が分かる。
// この配列が一致する場合、同じグループとみなせる。
// どうやってグルーピングするか？
// 1. インデックスをグルーピングする
//   例) [[5], [4, 2], [0, 1, 3]]
// 2. 文字列をグルーピングする
// 1を採用しつつ、配列を文字列化してキーとしたマップを作ればグルーピングが楽そう
// 計算量は
// 0. strs 100, strの長さが最大10**4
// 1. 100個の文字列を配列にマップする。この時に100 * 10 ** 4 = 100万回計算が必要。O(N*M)と言えそう。
// 2. 100 * 100の計算量で全配列の探索を終えられる。

func groupAnagrams(strs []string) [][]string {
	alphabets := make([][26]int, len(strs))
	group := make(map[[26]int][]int, len(strs))

	for i, str := range strs {
		for _, c := range str {
			alphabets[i][c-97]++
		}
		group[alphabets[i]] = append(group[alphabets[i]], i)
	}

	ans := make([][]string, len(group))
	cnt := 0
	for _, v := range group {
		for _, n := range v {
			ans[cnt] = append(ans[cnt], strs[n])
		}
		cnt++
	}

	return ans

}
