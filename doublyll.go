package main

import "fmt"

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func traverseLR(dll *Node) {
	if dll == nil {
		fmt.Println("Empty list")
		return
	}
	for dll != nil {
		fmt.Printf("%d-->", dll.Value)
		dll = dll.Right
	}
	fmt.Println()
	return
}

func traverseRL(dll *Node) {
	if dll == nil {
		fmt.Println("Empty list")
		return
	}
	for dll.Right != nil {
		dll = dll.Right
	}
	fmt.Println("Reached tail!")
	for dll != nil {
		fmt.Printf("%d-->", dll.Value)
		dll = dll.Left
	}
	fmt.Println()
	return
}

func prepend(value int, dll *Node) *Node {
	temp := new(Node)
	temp.Value = value
	if dll == nil {
		temp.Left = nil
		temp.Right = nil
	} else {
		temp.Left = nil
		temp.Right = dll
		dll.Left = temp
	}
	return temp
}

func append(value int, dll *Node) {
	temp := new(Node)
	temp.Value = value
	if dll == nil {
		temp.Left = nil
		temp.Right = nil
		dll = temp
		return
	}
	ptr := dll
	for ptr.Right != nil {
		ptr = ptr.Right
	}
	temp.Left = ptr
	ptr.Right = temp
}

func main() {
	var dll *Node
	// fmt.Println("Appending 10 elements in a Doubly LinkedList")
	dll = prepend(0, dll)
	for i := 1; i < 10; i++ {
		dll = prepend(i, dll)
	}
	traverseLR(dll)
	traverseRL(dll)
	append(-1, dll)
	traverseLR(dll)
	traverseRL(dll)

}
