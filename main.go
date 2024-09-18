package main

import (
	"fmt"
)

const size = 5

type Node struct {
	Value string
	Right *Node
	Left  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type HashMap map[string]*Node

type Cache struct {
	Queue Queue
	Hash  HashMap
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: HashMap{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(word string) {

	var node *Node
	if val, ok := c.Hash[word]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: word}
	}
	c.Add(node)
	c.Hash[word] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Removing %s..\n", n.Value)
	left := n.Left
	right := n.Right
	right.Left = left
	left.Right = right
	c.Queue.Length--
	delete(c.Hash, n.Value)
	c.Display()
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Adding %s..\n", n.Value)
	right := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = right
	right.Left = n
	c.Queue.Length++

	if c.Queue.Length > size {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	fmt.Printf("Length of Queue is %d\n", c.Queue.Length)
	// return
	Node := c.Queue.Head.Right
	for i := 0; i < c.Queue.Length; i++ {
		fmt.Printf("{%s}", Node.Value)
		if Node != nil {
			Node = Node.Right
		} else {
			break
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Starting Cache...")
	cache := NewCache()
	for _, word := range []string{"cars", "fruits", "bikes", "aeroplanes", "pets", "cars", "laptops"} {
		cache.Check(word)
		cache.Display()
	}
}
