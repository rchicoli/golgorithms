package bst

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

// Lookup return true if a node
// with the target data is found in the tree. Recurs
// down the tree, chooses the left or right
// branch by comparing the target to each node.
func Lookup(node *Node, target int) bool {
	// 1. Base case == empty tree
	// in that case, the target is not found so return false
	if node == nil {
		return false
	}
	// 2. see if found here
	if target == node.data {
		return true
	}

	// 3. otherwise recur down the correct subtree
	if target < node.data {
		return Lookup(node.left, target)
	}
	return Lookup(node.right, target)
}

// NewNode allocates a new node
// with the given data and nil left and right
// pointers.
func NewNode(data int) *Node {
	return &Node{
		left:  nil,
		right: nil,
		data:  data,
	}
}

// Insert a new node
// with the given number in the correct place in the tree.
// Returns the new root pointer which the caller should
// then use (the standard trick to avoid using reference
// parameters).
func Insert(node *Node, data int) *Node {
	// 1. If the tree is empty, return a new, single node
	if node == nil {
		return NewNode(data)
	}

	// 2. Otherwise, recur down the tree
	if data <= node.data {
		node.left = Insert(node.left, data)
		return node
	}

	node.right = Insert(node.right, data)

	return node
}

// BuildTree call newNode() three times
func BuildTree() *Node {
	var root = NewNode(2)
	root.left = NewNode(5)
	root.right = NewNode(3)

	return root
}

// Size computes the number of nodes in a tree.
func Size(node *Node) int {
	if node == nil {
		return 0
	}
	return Size(node.left) + 1 + Size(node.right)
}

// MaxDepth compute the "maxDepth" of a tree -- the number of nodes along
// the longest path from the root node down to the farthest leaf node.
func MaxDepth(node *Node) int {
	if node == nil {
		return 0
	}
	// compute the depth of each subtree
	var lDepth = MaxDepth(node.left)
	var rDepth = MaxDepth(node.right)

	// use the larger one
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

// MinValue return the minimum data value found in that tree.
// Note that the entire tree does not need to be searched.
func MinValue(node *Node) int {
	var current = node

	// loop down to find the leftmost leaf
	for current.left != nil {
		current = current.left
	}

	return current.data
}

// MaxValue return the maximum data value found in that tree.
func MaxValue(node *Node) int {
	var current = node

	// loop down to find the leftmost leaf
	for current.right != nil {
		current = current.right
	}

	return current.data
}

// PrintTree print out its data elements in increasing
// sorted order.
func PrintTree(node *Node) {
	if node == nil {
		return
	}

	PrintTree(node.left)
	fmt.Printf("%d ", node.data)
	PrintTree(node.right)

}

// PrintPostOrder print its nodes according to the "bottom-up"
//  postorder traversal.
func PrintPostOrder(node *Node) {
	if node == nil {
		return
	}
	// first recur on both subtrees
	PrintTree(node.left)
	PrintTree(node.right)

	// then deal with the node
	fmt.Printf("%d ", node.data)
}

// HasPathSum return true if there is a path from the root
// down to a leaf, such that adding up all the values along the path
// equals the given sum.
// Strategy: subtract the node value from the sum when recurring down,
// and check to see if the sum is 0 when you run out of tree.
func HasPathSum(node *Node, sum int) bool {
	// return true if we run out of tree and sum==0
	if node == nil {
		return sum == 0
	}

	// otherwise check both subtrees
	var subSum = sum - node.data
	return HasPathSum(node.left, subSum) || HasPathSum(node.right, subSum)
}

// PrintPaths print out all of its root-to-leaf
//  paths, one per line. Uses a recursive helper to do the work.
func PrintPaths(node *Node) {
	path := make([]int, 1000)

	printPathsRecur(node, path, 0)
}

// printPathsRecur Recursive helper function -- given a node, and an array containing
// the path from the root node up to but not including this node,
// print out all the root-leaf paths.
func printPathsRecur(node *Node, path []int, pathLen int) {
	if node == nil {
		return
	}

	// append this node to the path array
	path[pathLen] = node.data
	pathLen++

	// it's a leaf, so print the path that led to here
	if node.left == nil && node.right == nil {
		printArray(path, pathLen)
	} else {
		// otherwise try both subtrees
		printPathsRecur(node.left, path, pathLen)
		printPathsRecur(node.right, path, pathLen)
	}
}

// printArray prints out an array on a line.
func printArray(ints []int, len int) {
	for i := 0; i < len; i++ {
		fmt.Printf("%d ", ints[i])
	}
	fmt.Printf("\n")
}

// Mirror changes a tree so that the roles of the
// left and right pointers are swapped at every node.
//
// So the tree...
//        4
//       / \
//      2   5
//     / \
//    1   3
//
// is changed to...
//        4
//       / \
//      5   2
//         / \
//        3   1
//
func Mirror(node *Node) {
	if node == nil {
		return
	}
	var temp *Node

	// do the subtrees
	Mirror(node.left)
	Mirror(node.right)

	// swap the pointers in this node
	temp = node.left
	node.left = node.right
	node.right = temp
}

// DoubleTree for each node in a binary search tree,
//  create a new duplicate node, and insert
//  the duplicate as the left child of the original node.
//  The resulting tree should still be a binary search tree.
//
//  So the tree...
//     2
//    / \
//   1   3
//
//  Is changed to...
//        2
//       / \
//      2   3
//     /   /
//    1   3
//   /
//  1
//
func DoubleTree(node *Node) {
	var oldLeft *Node

	if node == nil {
		return
	}

	// do the subtrees
	DoubleTree(node.left)
	DoubleTree(node.right)

	// duplicate this node to its left
	oldLeft = node.left
	node.left = NewNode(node.data)
	node.left.left = oldLeft
}

// SameTree return true if they are structurally identical.
func SameTree(a, b *Node) bool {
	// 1. both empty -> true
	if a == nil && b == nil {
		return true
	}

	// 2. both non-empty -> compare them
	if a != nil && b != nil {
		return (a.data == b.data &&
			SameTree(a.left, b.left) &&
			SameTree(a.right, b.right))
	}

	// 3. one empty, one not -> false
	return false
}

// CountTrees For the key values 1...numKeys
// return how many structurally unique
// binary search trees are possible that store those keys.
// Strategy: consider that each value could be the root.
// Recursively find the size of the left and right subtrees.
func CountTrees(numKeys int) int {

	if numKeys <= 1 {
		return 1
	}

	// there will be one value at the root, with whatever remains
	// on the left and right each forming their own subtrees.
	// Iterate through all the values that could be the root...
	var sum = 0
	var left, right int

	for root := 1; root <= numKeys; root++ {
		left = CountTrees(root - 1)
		right = CountTrees(numKeys - root)

		// number of possible trees with this root == left*right
		sum += left * right
	}

	return (sum)
}

// IsBST returns true if a binary tree is a binary search tree.
func IsBST(node *Node) bool {
	if node == nil {
		return true
	}

	// false if the max of the left is > than us

	// (bug -- an earlier version had min/max backwards here)
	if node.left != nil && MaxValue(node.left) > node.data {
		return false
	}

	// false if the min of the right is <= than us
	if node.right != nil && MinValue(node.right) <= node.data {
		return false
	}
	// false if, recursively, the left or right is not a BST
	if !IsBST(node.left) || !IsBST(node.right) {
		return false
	}

	return true
}

// IsBST2 returns true if the given tree is a binary search tree
//  (efficient version).
func IsBST2(node *Node, min, max int) bool {
	return isBSTUtil(node, min, max)
}

// isBSTUtil Returns true if the given tree is a BST and its
// values are >= min and <= max.
func isBSTUtil(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	// false if this node violates the min/max constraint
	if node.data < min || node.data > max {
		return false
	}

	// otherwise check the subtrees recursively,
	// tightening the min or max constraint
	return (isBSTUtil(node.left, min, node.data) &&
		isBSTUtil(node.right, node.data+1, max))
}
