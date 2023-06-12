package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func prepend(value int, ll *Node) *Node {
	if ll == nil {
		ll = new(Node)
		ll.Value = value
		ll.Next = nil
		return ll
	}

	if value == ll.Value {
		fmt.Println("Node already present")
		return ll
	}
	temp := new(Node)
	temp.Value = value
	temp.Next = ll
	return temp
}

func append(value int, ll *Node) *Node {
	if ll == nil {
		ll := new(Node)
		ll.Value = value
		ll.Next = nil
		return ll
	} else if ll.Next == nil {
		temp := new(Node)
		temp.Value = value
		temp.Next = nil
		ll.Next = temp
	} else {
		append(value, ll.Next)
	}
	return ll
}

func traverse(ll *Node) {
	for ll != nil {
		fmt.Printf("%d-->", ll.Value)
		ll = ll.Next
	}
	fmt.Println()
}

func search(value int, ll *Node) bool {
	for ll != nil {
		if ll.Value == value {
			return true
		}
		ll = ll.Next
	}
	return false
}

func insertAt(pos, value int, ll *Node) bool {
	temp := &Node{value, nil}
	if pos <= 0 || ll == nil {
		ll = temp
		return true
	}
	for i := 0; i < pos; i++ {
		if ll.Next != nil {
			ll = ll.Next
		} else {
			fmt.Println("Reached end of the list, could not add the element to the list!")
			return false
		}
	}
	temp.Next = ll.Next
	ll.Next = temp
	return true
}

func deleteFromBegining(ll *Node) *Node {
	if ll == nil {
		fmt.Println("Nothing to delete")
		return ll
	}
	//how else to delete struct
	//set ll as nil?
	return ll.Next
}

func deleteNFromBegining(n int, ll *Node) *Node {
	for i := 0; i < n; i++ {
		if ll != nil {
			ll = ll.Next
		} else {
			fmt.Printf("Cannot remove %d elements, reached the end of the list.\n", n)
		}
	}
	return ll
}

func len(ll *Node) int {
	counter := 0
	for ll != nil {
		counter++
		ll = ll.Next
	}
	return counter
}

func deleteNFromEnd(n int, ll *Node) bool {
	if n > len(ll) {
		fmt.Println("n greater than the length of the list!")
		return false
	}
	if n == len(ll) {
		fmt.Println("Emptying the whole list!")
		var temp Node
		*ll = temp
		return true
	}
	temp := ll
	for i := 0; i < len(ll)-n-1; i++ {
		temp = temp.Next
	}
	temp.Next = nil
	return true
}

// func reverseLL(ll *Node) *Node {
// 	if ll.Next == nil {

// 		return
// 	}
// }

func main() {

	//create a ll with 10 random elements

	var ll *Node //create a nil node
	traverse(ll)
	// rand.Seed(time.Now().Unix())
	for i := 1; i <= 10; i++ {
		// temp := rand.Intn(20)
		ll = prepend(i, ll)
	}
	traverse(ll)
	ll = deleteFromBegining(ll)
	traverse(ll)
	//delete 4 elements
	ll = deleteNFromBegining(4, ll)
	traverse(ll)
	//append at the end of the list
	ll = append(0, ll)
	traverse(ll)
	//append in a nil list
	var ll1 *Node
	ll1 = append(1, ll1)
	traverse(ll1)

	//search
	fmt.Println("Searching...", search(4, ll))
	//insert at pos 2
	insertAt(2, -3, ll)
	traverse(ll)
	fmt.Println("Inserting at last pos using insert at", insertAt(7, -1, ll))
	traverse(ll)
	fmt.Println("Length of the linkedlist", len(ll))
	fmt.Println("Delete from end", deleteNFromEnd(1, ll))
	traverse(ll)

}
