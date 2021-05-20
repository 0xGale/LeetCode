# URL
https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/submissions/
# 代码
```python
class Solution:
    def printNumbers(self, n: int) -> List[int]:
        max_num = 9
        while n > 1:
            max_num *= 10
            max_num += 9
            n -= 1
            
        return [x for x in range(1, max_num + 1)]
```
这题主要考察大数问题， Python中没有大数问题
