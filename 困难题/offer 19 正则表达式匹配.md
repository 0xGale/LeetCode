# URL
https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
# 代码
```python
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        m, n = len(s), len(p)

        def matches(i:int, j: int) -> bool:
            if i == 0:
                return False
            if p[j - 1] == '.':
                return True
            return s[i - 1] == p[j - 1]
        
        f = [[False] * (n + 1) for _ in range(m + 1)]
        f[0][0] = True   # 都为空串，则为True
        for i in range(m + 1):
            for j in range(n + 1):
                if p[j - 1] == '*':  # 当p中为*时， 分两种情况：1. c*丢弃不要了，2. c跟s中上一个字符再进行匹配， 
                    f[i][j] |= f[i][j - 2]
                    if matches(i, j - 1):  # 第2种情况， 就是p中*号前面的字符跟s中的字符相匹配， 那么再继续和s中上一个字符进行匹配
                        f[i][j] |= f[i - 1][j]  # 这里p中*号前面的字符跟s中字符相匹配， 那么就看前面的是否相等了。
                else:
                    if matches(i, j):  # 如果匹配的话
                        f[i][j] |= f[i - 1][j - 1]  # 就看前面的是否相等了
        return f[m][n]
```

