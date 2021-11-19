package main

import "fmt"

type Node struct {
	Data      int
	NextPoint *Node
	PrePoint  *Node
}

type LinkedList struct {
	head    *Node
	current *Node
	tail    *Node
}

func CreatedLinkedList() {
	data := []int{1, 21, 31, 51, 62, 2, 3, 42, 33, 12, 12}
	link := LinkedList{}
	var currentNode *Node

	for i := 0; i < len(data); i++ {
		currentNode = new(Node)
		currentNode.Data = data[i]
		insertNode(&link, currentNode)
	}
	showLinkedList(link)
}

func showLinkedList(link LinkedList) {
	var currentNode *Node
	currentNode = link.head
	for currentNode.NextPoint != nil {
		fmt.Println("Node", currentNode.Data)
		currentNode = currentNode.NextPoint
	}
}

func insertNode(link *LinkedList, node *Node) {
	if link.head == nil {
		link.head = node
		link.tail = node
		link.current = node
	} else {
		link.tail.NextPoint = node
		node.PrePoint = link.tail
		link.tail = node
	}
}

func main() {
	CreatedLinkedList()
}
