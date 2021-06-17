# URL

[剑指 Offer 31. 栈的压入、弹出序列 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/submissions/)

# 代码

```python
class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        l = []
        popi = 0
        for v in pushed:
            l.append(v)
            while l and l[-1] == popped[popi]:
                l.pop()
                popi += 1
        return not l
```

```go
func validateStackSequences(pushed []int, popped []int) bool {
    stack := make([]int, 0)  // 这里初始化只能是0,不能是len(stack), 否则第一次运行嵌套的for循环时,会出错
    index := 0
    for _, v := range pushed {
        stack = append(stack, v)
         for len(stack) != 0 && stack[len(stack)-1] == popped[index]{
            stack = stack[:len(stack) - 1]
            if len(stack) == 0 {
                break
            }
            index ++
        }
    }
    return len(stack) == 0
}
```

类似于蜘蛛纸牌, 按照入栈A的序列,  每次比较当前入栈的元素是否跟要出栈序列的元素进行比较,如果一样就出栈, 注意:这里每个元素都是不同的

