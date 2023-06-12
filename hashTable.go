package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

const SIZE = 15

type HashTable struct {
	Size  int
	Table map[int]*Node
}

func hashFunction(value, size int) int {
	return value % size
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)                // find the index of the incoming element
	element := Node{Value: value, Next: hash.Table[index]} //create the memory for the element and it should point to the existing node
	hash.Table[index] = &element                           // point the index to the incoming element like adding in front of the linkedlist
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d-->", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookUp(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE) //make a map of 15 size of type NODE
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("SIZE of the hashtable", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
	fmt.Println(lookUp(hash, 6))
	fmt.Println(lookUp(hash, 47))

}
