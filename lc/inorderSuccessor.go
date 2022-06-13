package lc

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil {
		cur := p.Right
		for cur.Left != nil {
			cur = cur.Left
		}

		return cur
	}

	var target *TreeNode
	node := root
	for node != nil {
		if node.Val > p.Val {
			target = node
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return target
}
