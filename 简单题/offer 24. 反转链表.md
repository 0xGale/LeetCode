# URL

[剑指 Offer 24. 反转链表 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)



# 代码

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if not head:
            return None
        l = head
        r = head.next
        l.next = None
        while r:
            t = r.next
            # print(l.val)
            r.next = l
            l = r
            r = t
        return l
```

 