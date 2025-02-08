package heaply

type heap[T any] struct {
	mem []T
	cmp func(T, T) int
}

func NewHeap[T any](compare func(T, T) int) Heap[T] {
	return &heap[T]{
		cmp: compare,
	}
}

func (h *heap[T]) Top() (T, bool) {
	var ret T
	if len(h.mem) == 0 {
		return ret, false
	}
	return h.mem[0], true
}

func (h *heap[T]) Push(elem T) {
	h.mem = append(h.mem, elem)
	curI := len(h.mem) - 1
	for curI > 0 {
		parI := (curI - 1) / 2
		parent := h.mem[parI]
		if h.cmp(parent, elem) >= 0 {
			return
		}
		h.mem[curI] = parent
		h.mem[parI] = elem
		curI = parI
	}
}

func (h *heap[T]) Pop() (T, bool) {
	var ret T
	if len(h.mem) == 0 {
		return ret, false
	}
	ret = h.mem[0]

	if len(h.mem) == 1 {
		h.mem = nil
		return ret, true
	}

	h.mem[0] = h.mem[len(h.mem)-1]
	h.mem = h.mem[:len(h.mem)-1]

	curI := 0
	leftI := curI*2 + 1
	rightI := curI*2 + 2
	var left T
	var right T
	var curent T
	for leftI < len(h.mem)-1 {
		left = h.mem[leftI]
		right = h.mem[rightI]
		curent = h.mem[curI]
		if h.cmp(left, curent) < 0 && h.cmp(right, curent) < 0 {
			return ret, true
		}
		if h.cmp(left, right) > 0 {
			h.mem[curI] = left
			h.mem[leftI] = curent
			curI = leftI
		} else {
			h.mem[curI] = right
			h.mem[rightI] = curent
			curI = rightI
		}
		leftI = curI*2 + 1
		rightI = curI*2 + 2
	}

	if leftI == len(h.mem)-1 && h.cmp(h.mem[leftI], h.mem[curI]) > 0 {
		curent = h.mem[curI]
		h.mem[curI] = h.mem[leftI]
		h.mem[leftI] = curent
	}

	return ret, true
}

func (h *heap[T]) Size() int {
	return len(h.mem)
}
