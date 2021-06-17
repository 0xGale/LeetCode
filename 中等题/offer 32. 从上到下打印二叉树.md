# URL

[剑指 Offer 32 - I. 从上到下打印二叉树 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/)



# 代码

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def levelOrder(self, root: TreeNode) -> List[int]:
        if not root: return []
        queue = collections.deque()
        l = []
        queue.append(root)
        while queue:
            node = queue.popleft()
            l.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        return l
```

# 分析

数的广度优先搜索遍历

使用一个队列记录每次即将访问的元素.

