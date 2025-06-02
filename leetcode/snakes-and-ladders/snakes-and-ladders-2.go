package snakes_and_ladders

import "fmt"

type QueueNode struct {
	idx  int
	step int
}

type Queue struct {
	nodes []*QueueNode
	first int
	last  int
}

func NewQueue() *Queue {
	return &Queue{
		nodes: make([]*QueueNode, 1000),
		first: 0,
		last:  0,
	}
}

func (q *Queue) First() *QueueNode {
	if q.IsEmpty() {
		return nil
	}
	return q.nodes[q.first]
}

func (q *Queue) Add(node *QueueNode) {
	nextLast := (q.last + 1) % 1000
	if nextLast == q.first {
		panic("Queue full")
	}
	q.nodes[q.last] = node
	q.last++
}

func (q *Queue) Remove() *QueueNode {
	if q.IsEmpty() {
		panic("Queue empty")
	}
	node := q.nodes[q.first]
	q.first = (q.first + 1) % 1000
	return node
}

func (q *Queue) IsEmpty() bool {
	return q.last == q.first
}

func rc2Idx(n, r, c int) int {
	if (n-1-r)%2 == 0 {
		return (n-1-r)*n + c + 1
	} else {
		return (n-1-r)*n + (n - 1 - c) + 1
	}
}

func bfs(q *Queue, qMap map[int]*QueueNode, blocks []int, target int) int {
	for !q.IsEmpty() {
		current := q.Remove()
		fmt.Println(current.idx)
		if current.idx == target {
			return current.step
		}

		for i := 1; i <= 6; i++ {
			nextIdx := current.idx + i
			if nextIdx > target {
				continue
			}
			if blocks[nextIdx] != -1 {
				nextIdx = blocks[nextIdx]
			}
			if _, ok := qMap[nextIdx]; !ok {
				node := &QueueNode{
					idx:  nextIdx,
					step: current.step + 1,
				}
				q.Add(node)
				qMap[nextIdx] = node
			}
		}
	}
	return -1
}

func snakesAndLadders(board [][]int) int {

	n := len(board)

	blocks := make([]int, n*n+1)

	for i := range n {
		for j := range n {
			v := board[i][j]
			idx := rc2Idx(n, i, j)
			blocks[idx] = v
		}
	}

	fmt.Println(blocks)

	queue := NewQueue()
	queue.Add(&QueueNode{
		idx:  1,
		step: 0,
	})

	myMap := make(map[int]*QueueNode)

	return bfs(queue, myMap, blocks, n*n)
}
