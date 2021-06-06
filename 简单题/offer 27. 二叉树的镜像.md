# URL

[剑指 Offer 27. 二叉树的镜像 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

# 代码

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def mirrorTree(self, root: TreeNode) -> TreeNode:
        if not root:
            return None
        root.right, root.left = root.left, root.right
        self.mirrorTree(root.right)
        self.mirrorTree(root.left)
        return root
```

```GO
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = root.Right, root.Left
    mirrorTree(root.Left)
    mirrorTree(root.Right)
    return root
}
```

这里注意循环实现的方式.

