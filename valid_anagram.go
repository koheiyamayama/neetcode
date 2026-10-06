// 2つの文字列 s と t が与えられたとき、t が s のアナグラムであれば true を、そうでない場合は false を返す関数を作成してください。
// https://leetcode.com/problems/valid-anagram/description/

package main

func isAnagram(s string, t string) bool {
	// 同じ文字の個数が等しい
	//
	indicies := make(map[rune]int)
	for _, v := range s {
		indicies[v]++
	}

	for _, c := range t {
		indicies[c]--
	}

	for _, v := range indicies {
		if v != 0 {
			return false
		}
	}

	return true
}
