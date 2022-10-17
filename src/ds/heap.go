package ds

type __heap struct {
	Data []int
}

type (
	LHeap __heap
	RHeap __heap
)

func (heap *LHeap) Push(val int) {
	heap.Data = append(heap.Data, val)
	idx := len(heap.Data) - 1
	heap.reorder(idx)
}

func (heap *LHeap) reorder(idx int) {
	for {
		if idx == 0 {
			return
		}
		next := (idx - 1) / 2
		if heap.Data[next] <= heap.Data[idx] {
			return
		}
		heap.Data[next], heap.Data[idx] = heap.Data[idx], heap.Data[next]
		idx = next
	}
}

func (heap *LHeap) Pop() (int, bool) {
	if len(heap.Data) <= 0 {
		return 0, false // ? todo
	}
	dest := heap.Data[0]
	max := len(heap.Data) - 1
	//__heap.Data[0] = __heap.Data[max]
	//__heap.Data = __heap.Data[:max]
	idx := 0
	for {
		next := idx*2 + 1
		if next >= max {
			heap.Data[idx] = heap.Data[max]
			heap.reorder(idx)
			break
		}

		if heap.Data[next] >= heap.Data[next+1] {
			heap.Data[idx] = heap.Data[next+1]
			idx = next + 1
			continue
		}
		heap.Data[idx] = heap.Data[next]
		idx = next
	}
	heap.Data = heap.Data[:max]
	return dest, true
}

func (heap *RHeap) Push(val int) {
	heap.Data = append(heap.Data, val)
	idx := len(heap.Data) - 1
	heap.reorder(idx)
}

func (heap *RHeap) reorder(idx int) {
	for {
		if idx == 0 {
			return
		}
		next := (idx - 1) / 2
		if heap.Data[next] >= heap.Data[idx] {
			return
		}
		heap.Data[next], heap.Data[idx] = heap.Data[idx], heap.Data[next]
		idx = next
	}
}

func (heap *RHeap) Pop() (int, bool) {
	if len(heap.Data) <= 0 {
		return 0, false // ? todo
	}
	dest := heap.Data[0]
	max := len(heap.Data) - 1
	//__heap.Data[0] = __heap.Data[max]
	//__heap.Data = __heap.Data[:max]
	idx := 0
	for {
		next := idx*2 + 1
		if next >= max {
			heap.Data[idx] = heap.Data[max]
			heap.reorder(idx)
			break
		}
		if next == max {
			heap.Data[idx] = heap.Data[next]
			break
		}
		if heap.Data[next] <= heap.Data[next+1] {
			heap.Data[idx] = heap.Data[next+1]
			idx = next + 1
			continue
		}
		heap.Data[idx] = heap.Data[next]
		idx = next
	}
	heap.Data = heap.Data[:max]
	return dest, true
}
