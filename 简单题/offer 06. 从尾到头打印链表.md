url: https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }
    s := []int{}  // 使用数组保存数字， 
    for cur := head; cur != nil; cur = cur.Next {
        s = append(s, cur.Val)
    }
    for i, j := 0, len(s) - 1; i < j; {  // 按顺序保存数字之后，倒序输出数组中的值
        s[i], s[j] = s[j], s[i]
        i ++
        j --
    }
    return s
}
```

