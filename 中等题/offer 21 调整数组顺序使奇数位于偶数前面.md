# URL

[剑指 Offer 21. 调整数组顺序使奇数位于偶数前面 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/submissions/)

# 代码

最先想到的一种方法: 设置两个数组,分别保存需要的两部分数据

```python
class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        n1 = []
        n2 = []
        for v in nums:
            if v % 2 :
                n1.append(v)
            else:
                n2.append(v)
        print(n1.extend(n2))
        return n1

```

## 还有常用的双指针算法

这种方法需要注意的是: 指针边界问题, Python中不像C语言那样即使指向数组外面依旧可以返回值,这里如果指到数组外面就会立即报错.

```python
class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        p1 = 0
        p2 = len(nums) - 1
        while p1 < p2:
            while p1 < p2 and nums[p1] % 2 == 1:
                p1 += 1
            while p1 < p2 and nums[p2] % 2 == 0:
                p2 -= 1
            if p1 < p2:
                nums[p1], nums[p2] = nums[p2], nums[p1]
        return nums
```

注意在and和or的机制可以使得判断前面的必须的步骤, 就不会再运行后面的数值, 就可以避免出现指针指向数组外面的情况



