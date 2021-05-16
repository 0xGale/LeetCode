```python
class Solution:

    # def buildCore(self, preorder)

    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if len(preorder) == 0 | len(inorder) == 0:
            return None
        rootval = preorder[0]
        root = TreeNode(rootval)
        l_list = inorder[:inorder.index(rootval)]
        r_list = inorder[inorder.index(rootval)+1:]
        l_tree = self.buildTree(preorder[1:len(l_list)+1], l_list)
        root.left = l_tree
        r_tree = self.buildTree(preorder[len(l_list)+1:], r_list)
        root.right = r_tree
        return root
```

这里描述一下重建二叉树的过程，递归方法



