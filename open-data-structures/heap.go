package main

// heapを配列を使って実装する

type heap []int

func (h *heap) Add(n int) (ok bool) {
	return ok

}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}
