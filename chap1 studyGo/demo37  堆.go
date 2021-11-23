package main

import "fmt"

type NodeH struct {
	Value int
	Key   string
}

type Heap struct {
	list   []*NodeH
	length int
}

func main() {
	CreateHeap()
}

// 创建堆
func CreateHeap() {
	arrList := []int{1, 2, 11, 3, 7, 8, 4, 5}
	var myHeap Heap
	myHeap.list = append(myHeap.list, &NodeH{})
	for _, value := range arrList {
		tmp := NodeH{}
		tmp.Value = value
		myHeap.InsertHeap(&tmp)
	}
	for {
		node := myHeap.GetTopHeap()
		fmt.Println(node)
	}
	myHeap.SortHeap(myHeap.list)
	heapShow(myHeap.list)
}

// 插入堆
func (h *Heap) InsertHeap(one *NodeH) {
	h.list = append(h.list, one)
	length := len(h.list)
	h.length = length - 1
	h.AdjustHeap(h.length)
}

// 堆排序
func (h *Heap) SortHeap(heaps []*NodeH) {
	length := len(heaps)
	length = length - 1
	if length == 1 {
		return
	}
	if length == 2 {
		h.AdjustHeap(length - 1)
	}
	for length > 0 {
		h.SliceNodeSwap(1, length)
		length--
		h.Heapfiy(length, 1)
	}

	minPos := 1
	maxPos := h.length
	for minPos < maxPos {
		h.SliceNodeSwap(minPos, maxPos)
		minPos++
		maxPos--
	}
}

// 自下而上调整
func (h *Heap) AdjustHeap(length int) {
	if length < 1 {
		return
	}
	if length == 2 {
		if h.list[length].Value > h.list[length-1].Value {
			h.SliceNodeSwap(length, length-1)
		}
		return
	}
	i := length
	for i/2 > 0 && h.list[i].Value > h.list[i/2].Value {
		h.SliceNodeSwap(i, i/2)
		i = i / 2
	}
	return
}

// 输出heap
func heapShow(heaps []*NodeH) {
	for index, value := range heaps {
		fmt.Println(index, value)
	}
}

// node slice交换
func (h *Heap) SliceNodeSwap(i int, j int) {
	x := h.list[i]
	h.list[i] = h.list[j]
	h.list[j] = x
}

// 自上向下堆化
func (h *Heap) Heapfiy(length int, pos int) {
	for {
		maxPos := pos
		if pos*2 < length && h.list[pos].Value < h.list[pos*2].Value {
			maxPos = pos * 2
		}
		if pos*2+1 < length && h.list[maxPos].Value < h.list[pos*2+1].Value {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.SliceNodeSwap(pos, maxPos)
		pos = maxPos
	}
}

// 获取堆顶
func (h *Heap) GetTopHeap() *NodeH {
	if h.length == 0 {
		panic("Heap is empty")
	}
	top := h.list[1]

	h.SliceNodeSwap(1, len(h.list)-1)
	length := len(h.list) - 2
	fmt.Println(length)
	h.Heapfiy(length, 1)
	heapShow(h.list)
	h.list = append(h.list[:length+1], h.list[length+2:]...)
	h.length--
	return top
}
