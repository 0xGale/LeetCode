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

```Go
func levelOrder(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var res []int
    list := []*TreeNode{root}
    for i := 0; i < len(list); i++ {
        node := list[i]
        res = append(res, node.Val)
        if node.Left != nil {
            list = append(list, node.Left)
        }
        if node.Right != nil {
            list = append(list, node.Right)
        }
    }
    return res 
}
```



# 从上到下打印二叉树 II

```GO
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var res [][]int // 记录结果
    list := []*TreeNode{root} // 记录所有遍历的节点
    for len(list) != 0 {
        // 记录下一层中所有的节点
        tmp := []*TreeNode{}
        // 记录当前层中所有节点的数值
        var r []int
        for _, v := range list {  // 对当前层中所有的节点遍历
            r = append(r, v.Val)
            if v.Left != nil {
                tmp = append(tmp, v.Left)
            }
            if v.Right != nil {
                tmp = append(tmp, v.Right)
            }
        }
        res = append(res, r)
        list = tmp  // 上一层遍历完, 准备遍历下一层
    }
    return res
    
}
```



# 从上到下打印二叉树 III

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var res [][]int  // 结果
    list := []*TreeNode{root}  // 存放每一层所有的节点
    for len(list) != 0 {
        var tmp []*TreeNode  // 存放到当前节点层的下一层所有节点
        r := make([]int, len(list))  // 存放所有计算的结果
        for i := 0; i < len(list); i ++ {  // 开始遍历当前节点层的所有节点, 广度优先
            if len(res) % 2 != 0 {  // 利用res判断那一层需要反过来记录数值
                r[len(list) - i - 1] = list[i].Val
            } else {
                r[i] = list[i].Val
            }
            if list[i].Left != nil {  // 其他正常
                tmp = append(tmp, list[i].Left)
            }
            if list[i].Right != nil {
                tmp = append(tmp, list[i].Right)
            }
        }
        list = tmp  // 下一层的节点赋值到当前层
        res = apepend(res, r)  // 
    }
    return res
}
```



# 分析

数的广度优先搜索遍历

使用一个队列记录每次即将访问的元素.

