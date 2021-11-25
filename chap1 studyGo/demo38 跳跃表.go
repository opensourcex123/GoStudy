package main

import (
	"fmt"
)

type NodeS struct {
	Data      int
	NextPoint *NodeS
	PrePoint  *NodeS
	NextLevel *NodeS
}

type LinkedListS struct {
	Head    *NodeS
	Current *NodeS
	Tail    *NodeS
	Length  int
	Level   int
}

type SkipList struct {
	List        LinkedListS
	FirstIndex  LinkedListS
	SecondIndex LinkedListS
}

func InitSkipList() {
	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 51, 62}
	s1 := SkipList{}
	s1.initSkip(data)
	s1.find(19)
	s1.add(11)
	showSkipList(s1)
}

func showSkipLinkedList(link LinkedListS, name int) {
	var currentNode *NodeS
	currentNode = link.Head
	for {
		i := 1
		fmt.Println(name, "-Node:", currentNode.Data)
		if currentNode.NextPoint == nil {
			break
		} else {
			currentNode = currentNode.NextPoint
		}
		if name == 1 {
			fmt.Println("------->")
		} else if name == 2 {
			for i <= 3 {
				fmt.Println("--------->")
				i++
			}
		} else {
			for i <= 7 {
				fmt.Println("-------->")
				i++
			}
		}
	}
	fmt.Println("")
}

func (s1 *SkipList) initSkip(list []int) {
	s1.List = LinkedListS{}
	s1.FirstIndex = LinkedListS{}
	s1.SecondIndex = LinkedListS{}
	var currentNode *NodeS
	for i := 0; i < len(list); i++ {
		currentNode = new(NodeS)
		currentNode.Data = list[i]
		addNode(s1, currentNode)
	}
	showSkipList(*s1)
}

func (s1 *SkipList) find(x int) {
	var current *NodeS
	current = s1.SecondIndex.Head
	if current.Data == x {
		fmt.Println(current.Data)
		return
	}
	if x < current.Data {
		panic("No exist in skipList")
		return
	}
	for {
		if x > current.Data {
			fmt.Println(current.Data)
			current = current.NextPoint
		} else if x < current.Data {
			// 下到下级索引
			fmt.Println(current.Data)
			current = current.PrePoint.NextLevel.NextPoint
		} else {
			fmt.Println(current.Data)
			return
		}
	}
}

func (s1 *SkipList) add(x int) {
	var current *NodeS
	current = s1.SecondIndex.Head
	if current.Data == x {
		panic("Had existed in skipList")
		return
	}
	if x < current.Data {
		newNode2 := new(NodeS)
		newNode2.Data = x
		newNode2.NextPoint = s1.SecondIndex.Head
		s1.SecondIndex.Head.PrePoint = newNode2
		s1.SecondIndex.Head = newNode2

		newNode1 := new(NodeS)
		newNode1.Data = x
		newNode1.NextPoint = s1.FirstIndex.Head
		s1.FirstIndex.Head.PrePoint = newNode1
		s1.FirstIndex.Head = newNode1

		newNode := new(NodeS)
		newNode.Data = x
		newNode.NextPoint = s1.List.Head
		s1.SecondIndex.Head.PrePoint = newNode
		s1.List.Head = newNode
	}

	for {
		if x > current.Data {
			if current.NextPoint == nil {
				if current.NextLevel != nil {
					current = current.NextLevel
				} else {
					// 插入
					newNode := new(NodeS)
					newNode.Data = x
					current.NextPoint = newNode
					newNode.PrePoint = current
					return
				}
			} else {
				fmt.Println(current.Data)
				current = current.NextPoint
			}
		} else if x < current.Data {
			// 向下去寻找第一个大于x的值
			fmt.Println(current.Data)
			if current.PrePoint.NextLevel != nil {
				current = current.PrePoint.NextLevel.NextPoint
			} else {
				// 插入
				newNode := new(NodeS)
				newNode.Data = x
				current.NextPoint = newNode
				newNode.PrePoint = current
				return
			}
		} else {
			fmt.Println(current.Data)
			return
		}
	}
}

func showSkipList(s1 SkipList) {
	showSkipLinkedList(s1.SecondIndex, 3)
	fmt.Println("")
	showSkipLinkedList(s1.FirstIndex, 2)
	fmt.Println("")
	showSkipLinkedList(s1.List, 1)
}

func addNode(skipList *SkipList, node *NodeS) {
	insertToLink(&skipList.List, node)
	if skipList.FirstIndex.Length == 0 || ((skipList.List.Length-1)%2 == 0 && skipList.List.Length > 2) {
		newNode := new(NodeS)
		newNode.Data = node.Data
		newNode.NextLevel = node
		insertToLink(&skipList.FirstIndex, newNode)
		if skipList.SecondIndex.Length == 0 || ((skipList.FirstIndex.Length-1)%2 == 0 && skipList.FirstIndex.Length > 2) {
			newNode2 := new(NodeS)
			newNode2.Data = node.Data
			newNode2.NextLevel = newNode
			insertToLink(&skipList.SecondIndex, newNode2)
		}
	}
}

func insertToLink(link *LinkedListS, node *NodeS) {
	if link.Head == nil {
		link.Head = node
		link.Tail = node
		link.Current = node
	} else {
		link.Tail.NextPoint = node
		node.PrePoint = link.Tail
		link.Tail = node
	}
	link.Length++
}

func main() {
	InitSkipList()
}
