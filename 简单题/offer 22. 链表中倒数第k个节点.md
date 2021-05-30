# URL

[剑指 Offer 22. 链表中倒数第k个节点 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

# 代码

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        p1 = head
        p2 = head
        for _ in range(k - 1):
            p2 = p2.next
        while p2.next:
            p1 = p1.next
            p2 = p2.next
        return p1
```

这里的思路是:

两个指针相距k-1距离, 当后面的指针指向末尾时,那么p1正好距离最后一个节点有k个节点