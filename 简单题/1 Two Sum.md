# 1. 两数之和

```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, n in enumerate(nums):
            num = target - nums[i]
            if num in nums:  
                if i != nums.index(num):  # index索引给定值的第一个出现的位置， 这句也就是找到不是第一个出现的元素
                    return [i, nums.index(num)]
```

这个值不能是原来那个元素, 那么就有两个判断: 1. 这个得先在list中, 然后这个值通过index索引不是索引到的第一个值
下面是我最初写错的:

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

下面是Go语言的实现

```go
func twoSum(nums []int, target int) []int {
    result := []int{}
    m := make(map[int]int)
    for i,k := range nums {      
        if value,exist := m[target-k];exist {  // m数组是由nums中的值到键的索引， 如果一个target-k值存在，那么在m中一定有对应的索引
            print(exist)
            result = append(result,value)  // 这个value是m中的值，所以这个是一个索引
            result = append(result,i)
        }
        m[k] = i // 将已经遍历过的值保存到m数组中， 之后再遍历就会有该值
    }
    return result
}
```

上面通过将值和索引映射成: 由value->index的方法类似于创建的hash表, 这种是以空间来换时间的方式.

