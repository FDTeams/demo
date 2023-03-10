package quartz

import "container/heap"

// item is the priorityQueue item.
type item struct {
	Job      Job
	Trigger  Trigger
	priority int64 // item priority, backed by the next run time.
	index    int   // maintained by the heap.Interface methods.
}

// priorityQueue implements the heap.Interface.
type priorityQueue []*item

// Len returns the priorityQueue length.
func (pq priorityQueue) Len() int { return len(pq) }

// Less is the items less comparator.
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap exchanges the indexes of the items.
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push implements the heap.Interface.Push.
// Adds x as element Len().
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop implements the heap.Interface.Pop.
// Removes and returns element Len() - 1.
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Head returns the first item of the priorityQueue without removing it.
func (pq *priorityQueue) Head() *item {
	return (*pq)[0]
}

// Remove removes and returns the element at index i from the priorityQueue.
func (pq *priorityQueue) Remove(i int) interface{} {
	return heap.Remove(pq, i)
}
