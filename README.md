# HEAPLY is a friendly generic contract golang heap

Now *heaply* provides simple binary heap implementation and generic contracts. ```heaply.NewHeap``` - is not thread-safe

Usages example:
```
// int max heap
maxHeap := heaply.NewHeap(func(a, b int) int { return a - b })

_, ok := maxHeap.Top() // ok - false because maxHeap is empty

maxHeap.Push(10)
maxHeap.Push(5)
maxHeap.Push(20)
maxHeap.Push(1)
maxHeap.Push(15)
maxHeap.Push(8)

max, ok := maxHeap.Pop()
// ok == true
// max == 20
max, ok = maxHeap.Top()
// ok == true
// max == 15
```