package main

// import (
// 	"fmt"
// )

// type dataStructure interface {
// 	pop() *node
// 	push(*node)
// 	structSize() int
// }

// type node struct {
// 	Value int
// }

// func (n *node) String() string {
// 	return fmt.Sprint(n.Value)
// }

// // NewStack returns a new stack.
// func NewStack() *stack {
// 	return &stack{}
// }

// // stack is a basic LIFO stack that resizes as needed.
// type stack struct {
// 	nodes []*node
// 	count int
// }

// // push adds a node to the stack.
// func (s *stack) push(n *node) {
// 	s.nodes = append(s.nodes[:s.count], n)
// 	s.count++
// }

// // pop removes and returns a node from the stack in last to first order.
// func (s *stack) pop() *node {
// 	if s.count == 0 {
// 		return nil
// 	}
// 	s.count--
// 	return s.nodes[s.count]
// }

// // structSize returns the size or the count of the elements in the stack
// func (s *stack) structSize() int {
// 	return s.count
// }

// // NewQueue returns a new queue with the given initial size.
// func NewQueue(size int) *queue {
// 	return &queue{
// 		nodes: make([]*node, size),
// 		size:  size,
// 	}
// }

// // queue is a basic FIFO queue based on a circular list that resizes as needed.
// type queue struct {
// 	nodes []*node
// 	size  int
// 	head  int
// 	tail  int
// 	count int
// }

// // Push adds a node to the queue.
// func (q *queue) push(n *node) {
// 	if q.head == q.tail && q.count > 0 {
// 		nodes := make([]*node, len(q.nodes)+q.size)
// 		copy(nodes, q.nodes[q.head:])
// 		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
// 		q.head = 0
// 		q.tail = len(q.nodes)
// 		q.nodes = nodes
// 	}
// 	q.nodes[q.tail] = n
// 	q.tail = (q.tail + 1) % len(q.nodes)
// 	q.count++
// }

// // Pop removes and returns a node from the queue in first to last order.
// func (q *queue) pop() *node {
// 	if q.count == 0 {
// 		return nil
// 	}
// 	node := q.nodes[q.head]
// 	q.head = (q.head + 1) % len(q.nodes)
// 	q.count--
// 	return node
// }

// // structSize returns the size or the count of the elements in the queue
// func (q *queue) structSize() int {
// 	return q.count
// }

// func Insert(ds dataStructure, val int) {
// 	// fmt.Println("Adding value: %s", string(val))
// 	ds.push(&node{val})
// }

// func Remove(ds dataStructure) string {
// 	val := ds.pop().String()
// 	// fmt.Println("Removing value: %s", string(val))
// 	return val
// }

// func Count(ds dataStructure) int {
// 	return ds.structSize()
// }

// // Output:
// //
// //  Stack
// //  ======
// // No of Elements: %s 4
// // 9 7 5 3
// // No of Elements: %s 0
// //
// //  Queue
// //  ======
// // No of Elements: %s 4
// // 1 2 4
// // No of Elements: %s 1
// // 8
// // No of Elements: %s 0

// func main() {
// 	s := NewStack()
// 	Insert(s, 3)
// 	Insert(s, 5)
// 	Insert(s, 7)
// 	Insert(s, 9)

// 	fmt.Printf("\n Stack \n ====== \n")
// 	fmt.Println("No of Elements: %s", Count(s))
// 	fmt.Println(Remove(s), Remove(s), Remove(s), Remove(s))
// 	fmt.Println("No of Elements: %s", Count(s))

// 	q := NewQueue(1)
// 	Insert(q, 1)
// 	Insert(q, 2)
// 	Insert(q, 4)
// 	Insert(q, 8)

// 	fmt.Printf("\n Queue \n ====== \n")
// 	fmt.Println("No of Elements: %s", Count(q))
// 	fmt.Println(Remove(q), Remove(q), Remove(q))
// 	fmt.Println("No of Elements: %s", Count(q))
// 	fmt.Println(Remove(q))
// 	fmt.Println("No of Elements: %s", Count(q))
// }
