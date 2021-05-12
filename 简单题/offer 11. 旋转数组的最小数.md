[剑指 Offer 11. 旋转数组的最小数字 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/submissions/)

这题耗费了一些时间，因为总是用mid=(l + r) // 2和l来比较，最终发现这样比较还是不能成功，

这里引用官方的一张图：

![最小值](https://assets.leetcode-cn.com/solution-static/jianzhi_11/1.png)

这里所有的比较都是跟最右边的值进行比较，有如下三种情况：

1. nums[mid] < nums[r], 那么这个r的范围就要缩小道mid这里
2. nums[mid] > nums[r], 那么这个mid就是在最小值的左边，缩小l的范围， l = mid + 1, 这里加1 是因为mid已经比较过了，并且它一定小于nums[r]
3. nums[mid] == nums[r], 那么这里已经判断不了最小值是在哪一个范围了，这里要么r -= 1缩小一个值的范围，要么遍历l到r中所有的值，找出最小值。

![fig4](https://assets.leetcode-cn.com/solution-static/jianzhi_11/4.png)

这里最大的问题是不能比较mid和l，如果思路是nums[mid]> nums[l],那么就默认这是一个旋转至少一次的数组，也就是最小值一定不是最右边的值。

例如：

输入： 1， 2， 3，

l, r = 0, 2

mid = 1

nums[mid] > nums[l], l  = mid

上述代码就把最小值1给忽略了。

如果数组中遍历到这种子数组，那么就会导致这种情况发生，所以还是不要用mid和l比较。

Python:

```python
class Solution:
    def minArray(self, numbers: List[int]) -> int:
        l, r = 0, len(numbers) - 1
        mid = (l + r) >> 1
        while l < r :
            print(l, mid, r)
            if numbers[r] > numbers[mid]:
                r = mid
            elif numbers[r] < numbers[mid]:
                l = mid + 1
            else:
                minV = numbers[l]
                for i in range(l, r + 1):
                    if numbers[i] < minV:
                        minV = numbers[i]
                return minV
            mid = (l + r) >> 1
        
        return numbers[l]
```

Golang:

```GO
func minArray(numbers []int) int {
    l, r := 0, len(numbers) - 1
    for ;l < r; {
        
        mid := l + (r - l) // 2
        // fmt.Println(l, mid, r)
        if numbers[mid] < numbers[r] {
            r = mid
        } else if numbers[mid] > numbers[r] {
            l = mid + 1
        } else {
            minV := numbers[l]
            for i := l; i <= r; i ++ {
                if numbers[i] < minV {
                    minV = numbers[i]
                }
            } 
            return minV
        }
    }
    return numbers[l]
}
```

