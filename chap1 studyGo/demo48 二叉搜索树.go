package main

import (
	"fmt"
	"math"
)

type NodeT struct {
	value  int
	left   *NodeT
	right  *NodeT
	parent *NodeT
}

type Tree struct {
	root   *NodeT
	length int
}

func main() {
	createBTree()
}

func createBTree() {
	arrList := []int{14, 2, 5, 7, 23, 35, 12, 17, 31}
	myTree := Tree{}
	for i := 0; i < len(arrList); i++ {
		myTree = insertBNode(myTree, arrList[i])
		myTree.length++
	}

	fmt.Println(myTree)

	TreeHeight(myTree)
}

func TreeHeight(tree Tree) {
	var hLeft = 1
	if tree.root.left != nil {
		hLeft = heightMax(tree.root.left, hLeft)
	}
	var hRight = 1
	if tree.root.right != nil {
		hRight = heightMax(tree.root.right, hRight)
	}

	fmt.Println(hLeft, hRight)
	fmt.Println("Tree height is ", int(math.Max(float64(hLeft), float64(hRight))))
}

func heightMax(node *NodeT, h int) int {
	var hL = h
	var hR = h
	if node.left == nil && node.right == nil {
		fmt.Println(node)
		return h
	}
	if node.left != nil {
		h++
		hL = heightMax(node.left, h)
	}
	if node.right != nil {
		h++
		hR = heightMax(node.right, h)
	}

	return int(math.Max(float64(hL), float64(hR)))
}

// 中序遍历
func LDR(tree Tree) {
	readList := make(map[int]int)
	i := 0
	var currentNode *NodeT
	currentNode = tree.root
	for {
		//fmt.Println(currentNode)
		if i == tree.length {
			//fmt.Println(currentNode.value)
			break
		}
		if currentNode.left == nil {
			if readList[currentNode.value] == 1 {
				if readList[currentNode.right.value] == 1 {
					currentNode = currentNode.parent
					continue
				} else {
					currentNode = currentNode.right
					continue
				}
			} else {
				fmt.Println(currentNode.value)
				readList[currentNode.value] = 1
				i++
				if currentNode.right == nil {
					currentNode = currentNode.parent
					continue
				} else {
					if readList[currentNode.right.value] == 1 {
						currentNode = currentNode.parent
						continue
					} else {
						currentNode = currentNode.right
						continue
					}
				}
			}
		} else {
			if readList[currentNode.left.value] == 1 {
				if readList[currentNode.value] == 1 {
					currentNode = currentNode.right
					continue
				} else {
					fmt.Println(currentNode.value)
					readList[currentNode.value] = 1
					i++
					if currentNode.right == nil {
						currentNode = currentNode.parent
						continue
					} else {
						if readList[currentNode.right.value] == 1 {
							currentNode = currentNode.parent
							continue
						} else {
							currentNode = currentNode.right
							continue
						}

					}
				}
			} else {
				currentNode = currentNode.left
				continue
			}

		}

	}
}

func insertBNode(tree Tree, insertValue int) Tree {
	var currentNode *NodeT
	var tmp *NodeT
	i := 0
	if tree.length == 0 {
		currentNode = new(NodeT)
		currentNode.value = insertValue
		tree.root = currentNode
		return tree
	} else {
		currentNode = tree.root
	}

	for {
		// fmt.Println(currentNode.value)
		if currentNode.value < insertValue {
			if currentNode.right == nil {
				tmp = new(NodeT)
				tmp.value = insertValue
				currentNode.right = tmp
				tmp.parent = currentNode
				break
			} else {
				currentNode = currentNode.right
				continue
			}
		} else {
			if currentNode.left == nil {
				tmp = new(NodeT)
				tmp.value = insertValue
				currentNode.left = tmp
				tmp.parent = currentNode
				break
			} else {
				currentNode = currentNode.left
				continue
			}
		}
		i++
	}
	return tree
}
