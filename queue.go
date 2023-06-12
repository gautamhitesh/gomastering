package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func push(value int, queue *Node) *Node {
	temp := new(Node)
	temp.Value = value
	if queue == nil {
		temp.Next = nil
	} else {
		temp.Next = queue
	}
	return temp
}

func pop(queue *Node) (*Node, int) {
	if queue == nil {
		return nil, 0
	}

	temp := queue
	for queue.Next != nil {
		temp = queue
		queue = queue.Next
	}
	value := temp.Next.Value
	poppedNode := temp.Next
	fmt.Println("popped ", value)
	temp.Next = nil
	return poppedNode, value
}

func traverse(queue *Node) {
	if queue == nil {
		fmt.Println("Empty queue")
		return
	}
	for queue != nil {
		fmt.Printf("%d-->", queue.Value)
		queue = queue.Next
	}
	fmt.Println()
}

func main() {
	var q *Node
	for i := 0; i < 10; i++ {
		q = push(i, q)
	}
	traverse(q)
	temp, value := pop(q)
	fmt.Println("Popped item ", value, temp)
	traverse(q)

}
