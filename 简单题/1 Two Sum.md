# 1. 两数之和

```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, n in enumerate(nums):
            num = target - nums[i]
            if num in nums:  
                if i != nums.index(num):
                    return [i, nums.index(num)]
```

这个值不能是原来那个元素, 那么就有两个判断: 1. 这个得先在list中, 然后这个值通过index索引不是索引到的第一个值
下面是最初写错的:

```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_to_index = {}
        for i, n in enumerate(nums):
            nums_to_index[n] = i
        for i, n in enumerate(nums):
            if target - n in nums and target - n != n:
                l = [i, nums_to_index[target-n]]
                # print('[', i, nums_to_index[target-n], ']')
                return l
```

这里我考虑的是同一个值不使用两次, 而题目中要求的是同一个元素不使用两次,也就是如果一个值出现了多次那么是可以使用的, 即同一个位置不能使用多次

```go
func twoSum(nums []int, target int) []int {
    result := []int{}
    m := make(map[int]int)
    for i,k := range nums {      
        if value,exist := m[target-k];exist {
            print(exist)
            result = append(result,value)
            result = append(result,i)
        }
        m[k] = i
    }
    return result
}
```

上面通过将值和索引映射成: 由value->index的方法类似于创建的hash表, 这种是以空间来换时间的方式.

