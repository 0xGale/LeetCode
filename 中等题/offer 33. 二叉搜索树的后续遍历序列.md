# URL

[剑指 Offer 33. 二叉搜索树的后序遍历序列 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/)

# 代码

```go
func verifyPostorder(postorder []int) bool {
    return recur(postorder, 0, len(postorder) - 1)
}

func recur(postorder []int, i, j int) bool {
    if i >= j {
        return true
    }
    p := i 
    for postorder[p] < postorder[j] {
        p ++
    }
    m := p
    for postorder[p] > postorder[j] {
        p ++
    }
    return p == j && recur(postorder, i, m -1) && recur(postorder, m, j-1)
}
```

