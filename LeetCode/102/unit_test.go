package leetcode_102

import (
	"testing"
	"fmt"
)

func buildTree(input []int) *TreeNode {
	var root *TreeNode
	nodes := []*TreeNode{root}
	input_idx := 0
	input_len := len(input)
	for len(nodes) > 0 {
		currNode := nodes[0]
		nodes = nodes[1:]

		if(input_idx >= input_len) {break}
		currNode.Val = input[input_idx]
		input_idx += 1

		if(input_idx >= input_len) {break}
		nodes = append(nodes, &TreeNode{input[input_idx], nil, nil})
		input_idx += 1

		if(input_idx >= input_len) {break}
		input_idx += 1
		nodes = append(nodes, &TreeNode{input[input_idx], nil, nil})
	}
	return root
}

func TestSomething(t *testing.T) {
	// tree := buildTree([]int{3,9,20,nil,nil,15,7})
	// fmt.Println(tree)
	// fmt.Println([]string{"3","9","20","","","15","7"})
}


