package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func push(value int, stack *Node) *Node {
	if stack == nil {
		fmt.Println("Empty stack. Pushing first value.")
		stack = &Node{value, nil}
		return stack
	}
	temp := &Node{value, nil}
	temp.Next = stack
	stack = temp
	return stack
}

func pop(stack *Node) *Node {
	return stack.Next
}

func traverse(stack *Node) {
	for stack != nil {
		fmt.Printf("%d-->", stack.Value)
		stack = stack.Next
	}
	fmt.Println()
}

func main() {
	var stack *Node
	for i := 0; i < 10; i++ {
		stack = push(i+1, stack)
	}
	traverse(stack)
	stack = pop(stack)
	traverse(stack)
}
