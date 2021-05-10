基础实现：

```python
class Solution:
    def fib(self, n: int) -> int:
        if n == 0:
            return 0
        elif n == 1:
            return 1
        l = [0] * (n + 1)
        l[0] = 0
        l[1] = 1
        for i in range(2, n+1):
            l[i] =  l[i-1] + l[i-2]
        return int(l[n] % (1e9 + 7))
```

简单优化：

```python
class Solution:
    def fib(self, n: int) -> int:
        if n == 0:
            return 0
        elif n == 1:
            return 1
        l = [0, 1]
        for i in range(2, n+1):
            l.append(l[i -1] + l[i - 2])
        return l[n] % (1000000007)
```

优化空间：

```python
class Solution:
    def fib(self, n: int) -> int:
        if n == 0:
            return 0
        elif n == 1:
            return 1
        a, b = 0, 1
        for _ in range(2, n+1):
            a, b = b, a + b
        return b % (1000000007)
```

这里明明使用的空间比上面两个代码少，但是在提交的时候，空间使用的跟上面差不多，

还有在Python中取模运算，使用 b%(1000000007)和b %(1e9 + 7)结果是不一样的，1e9+7是一个浮点数，浮点数取模结果不太一样，这里还不太清为什么

。。。