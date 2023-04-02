package lc692

import (
	"container/heap"
	"sort"
)

type wordFreq struct {
	word  string
	freq  int
	index int
}

type wordHeap []*wordFreq

func (h wordHeap) Len() int { return len(h) }
func (h wordHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq < h[j].freq
	}
	return h[i].word > h[j].word
}
func (h wordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}
func (h *wordHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*wordFreq)
	item.index = n
	*h = append(*h, item)
}
func (h *wordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1
	*h = old[0 : n-1]
	return item
}

func TopKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	h := &wordHeap{}
	heap.Init(h)

	for word, f := range freq {
		heap.Push(h, &wordFreq{word: word, freq: f})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(*wordFreq).word
	}

	return result
}

func TopKFrequent2(words []string, k int) []string {
	cnt := map[string]int{}
	for _, v := range words {
		cnt[v]++
	}
	ans := []string{}
	for v := range cnt {
		ans = append(ans, v)
	}
	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return cnt[a] > cnt[b] || cnt[a] == cnt[b] && a < b
	})
	return ans[:k]
}
