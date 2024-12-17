package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head  *Node
	Tail  *Node
	Lengh int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	head.Left = tail
	return Queue{Head: head, Tail: tail}

}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

type Hash map[string]*Node

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove %s\n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Lengh--
	delete(c.Hash, n.Val)
	return n

}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add %s\n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Lengh++
	if c.Queue.Lengh > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Lengh)
	for i := 0; i < q.Lengh; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Lengh-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {

	fmt.Println("Start Cache")
	cache := NewCache()
	for _, word := range []string{"parrot", "dragonfruit", "tomato", "dog", "cat", "elephant"} {
		cache.Check(word)
		cache.Display()
	}

}
