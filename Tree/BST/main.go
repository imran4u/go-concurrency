package main

import "fmt"

/*
Using the example tree:

        50
       /  \
     30    70
    / \    / \
   20 40  60 80

Inorder (Left → Root → Right): 20, 30, 40, 50, 60, 70, 80 (always sorted in a BST)
Preorder (Root → Left → Right): 50, 30, 20, 40, 70, 60, 80
Postorder (Left → Right → Root): 20, 40, 30, 60, 80, 70, 50
Level Order (BFS): 50, 30, 70, 20, 40, 60, 80

*/

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	root := &Node{value: 50}
	papulateTree(root)

	rootAlternate := createTree()
	fmt.Println("InOrder")
	root.InOrder()
	fmt.Println("Alternate root InOrder")
	rootAlternate.InOrder()

	//PreOrder
	fmt.Println("PreOrder")
	root.PreOrder()
	//PostOrder
	fmt.Println("PostOrder")
	root.PostOrder()

	//level order
	fmt.Println("Level Order")
	root.LevelOrder()

	// invert the tree
	invertTree(root)
}

func papulateTree(root *Node) {
	//hard coded value from the above example
	root.left = &Node{value: 30}
	root.left.left = &Node{value: 20}
	root.left.right = &Node{value: 40}

	root.right = &Node{value: 70}
	root.right.left = &Node{value: 60}
	root.right.right = &Node{value: 80}

}

func createTree() *Node {
	return &Node{
		value: 50,
		left: &Node{
			value: 30,
			left:  &Node{value: 20},
			right: &Node{value: 40},
		},
		right: &Node{
			value: 70,
			left:  &Node{value: 60},
			right: &Node{value: 80},
		},
	}
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}

	n.left.InOrder()
	fmt.Println(n.value)
	n.right.InOrder()

}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	n.left.PreOrder()
	n.right.PreOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}

	n.left.PostOrder()
	n.right.PostOrder()
	fmt.Println(n.value)

}

func (n *Node) LevelOrder() {
	if n == nil {
		return
	}
	queue := []*Node{n} // slice of node pointer and its first item is root.

	for len(queue) > 0 {
		//remove first element
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.value)
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
		//fmt.Println("q-len", len(queue))
	}

}

func invertTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Swap left and right
	root.left, root.right = root.right, root.left

	// Invert both subtrees
	invertTree(root.left)
	invertTree(root.right)

	return root
}
