先排序，然后看相邻元素是否相同， 有直接返回， 时间O(nlog n)空间O（1）

```python
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        nums.sort()
        pre = nums[0]
        for i in range(1, len(nums)):
            if pre == nums[i]:
                return pre
            pre = nums[i]
```

使用哈希表, 时间O（n），空间O（n）

```python
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        repeat = {}
        for v in nums:
            if v not in repeat:
                repeat[v] = 1
            else:
                return v
```

原地交换：

原地交换的意思是：由于所有的元素都小于n, 那么将每个元素放到对应的位置上，例如：元素1就放到位置1上，如果发现某个元素和该元素对应的位置上的元素相同，则说明该元素重复。

```python
class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            if i != nums[i]:
                if nums[i] == nums[nums[i]]:
                    return nums[i]
                temp = nums[i]
                nums[i], nums[temp] = nums[temp], nums[i]
```

![Picture0.png](https://pic.leetcode-cn.com/1618146573-bOieFQ-Picture0.png)

Python 中， a, b = c, da,b=c,d 操作的原理是先暂存元组 (c, d)(c,d) ，然后 “按左右顺序” 赋值给 a 和 b 。
因此，若写为 nums[i], nums[nums[i]] = nums[nums[i]], nums[i]nums[i],nums[nums[i]]=nums[nums[i]],nums[i] ，则 nums[i]nums[i] 会先被赋值，之后 nums[nums[i]]nums[nums[i]] 指向的元素则会出错。

作者：jyd
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/mian-shi-ti-03-shu-zu-zhong-zhong-fu-de-shu-zi-yua/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

这里要加上一个temp的原因。

