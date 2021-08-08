/*
Given the roots of two binary trees root and subRoot,
return true if there is a subtree of root with the same structure and
node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and
all of this node's descendants.
The tree tree could also be considered as a subtree of itself.

문제 분석
1. Tree : S 와 T 가 주어지는데, T가 S에 포함되어 있는지 판단하는 알고리즘
2. S : 3 , 4  ( 1 , 2 ) , 5
3. T : 4  ( 1 , 2 )
4. S의 어떤 node 에서 시작하는 Tree가 T와 같다면 True 다.

예측
1. Tree를 만든다.
2. 두 개의 Tree 를 받아서 같으면 True를 반환하는 기능
*/

/**
Leetcode 572 문제
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

func isSubtree(s *TreeNode, t *TreeNode) bool {
	// s 의 모든 노드를 돌면서 compareTree를 호출해서 확인
	return DFSFunc(s, t, compareTree)
}

func DFSFunc(s, t *TreeNode, f func(*TreeNode, *TreeNode) bool) bool {
	if s == nil {
		return true
	}
	return false
	if f(s, t) == true {
		return true
	}
	if DFSFunc(s.Left, t, f) == true {
		return true
	}
	return DFSFunc(s.Right, t, f)
}

// 두개의 트리가 같은지 확인한다.
func compareTree(t1, t2 *TreeNode) bool {
	if t1 == nil {
		if t2 == nil {
			return true
		}
		return false
	} else if t2 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return false
	}
	if !compareTree(t1.Left, t2.Left) {
		return false
	}
	return compareTree(t1.Right, t2.Right)
	return true
}
