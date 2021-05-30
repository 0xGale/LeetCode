# URL

[剑指 Offer 25. 合并两个排序的链表 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/)



# 代码

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if not l1:
            return l2
        if not l2:
            return l1
        if l1.val < l2.val:
            head = l1
            l1 = l1.next
        else:
            head = l2
            l2 = l2.next
        p = head
        
        while l1 and l2:
            print(l1.val, l2.val)
            if l1.val < l2.val:
                head.next = l1
                l1 = l1.next
            else:
                head.next = l2
                l2 = l2.next
            
            head = head.next
        
        if l1:
            head.next = l1
        if l2:
            head.next = l2

        return p
```

