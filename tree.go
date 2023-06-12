package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, ",")
	traverse(t.Right)
	fmt.Println()
}

func create(value int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*value; i++ {
		temp := rand.Intn(2 * value)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, value int) *Tree {
	if t == nil {
		return &Tree{nil, value, nil} //maintaing order of the structure
	}

	if value == t.Value {
		return t
	}

	if value < t.Value {
		t.Left = insert(t.Left, value)
		return t
	}
	t.Right = insert(t.Right, value)
	return t
}

func main() {

	tree := create(10)
	fmt.Println("Root of the tree ", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, 19)
	tree = insert(tree, -2)
	fmt.Println("Root changed?", tree.Value)
	traverse(tree)
}
