# URL
https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
# 简单分析
1. n > 0
按照n计数， 循环n次
2. n = 0
返回1
3. n < 0
n = n * -1
按照n计数， 循环n次
# 优化分析
x^2 = x * x
x^3 = x^2 * x
x^4 = x^2 * x^2
x^5 = x^4 * x = x^2 * x^3 = x * x * x^2
x^6 = x^5 * x = x^3 * x^3
x^7 = x^(3+4) =x^3 * x^4
x^8 = x^(4+4) = x^4 * x^4 = x^2 * x^2 * x^2 * x^2 = 

8 = 4 *2 = 
# 错误代码
下面代码不正确
```python
import math
class Solution:
    def myPow(self, x: float, n: int) -> float:
        flag = False
        if n < 0:
            flag = True
            n *= -1
        pro_min_list = [1] * (int(n // 2) + 1)
        pro_min_list[1] = x
        pro = 1
        if n >= 2:
            for i in range(2, len(pro_min_list)):
                last = int(i // 2)
                if i % 2 == 0:
                    pro_min_list[i] = pro_min_list[last] * pro_min_list[last]
                    # print(pro_min_list[i])
                else:
                    pro_min_list[i] = pro_min_list[last] * pro_min_list[last + 1]
                    # print(pro_min_list[i])
        print(pro_min_list)
        if n % 2 == 0:
            pro = pro_min_list[-1] * pro_min_list[-1]
        else:
            pro = pro_min_list[-1] * pro_min_list[-2]

        if flag:
            pro = 1 / pro

        return pro
```
重新理一下代码的思路：
1. x^2 = x * x
2. x^3 = x * x^2
3. x^4 = x^2 * x^2
4. x^5 = x^2 * x^2 * x
5. if n%2==0: x^n = x^(n//2) * x^(n//2) , else: x^n = x^(n//2) * x^(n//2) * x
实际上这里的n//2是可以重复使用的。
这个代码： x *= x， 可以实现x^2的叠加:
1. x^1 = x
2. x^2
3. x^4
4. x^8
5. x^16
6. x^(log2(n))
代码：
```python
import math
class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        if n < 0:
            x = 1 / x
            n = -n
        res = 1
        while n:
            if n & 1:
                res *= x
            x *= x
            n >>= 1
        return res
```
这样时间复杂度就是：O(log(n))
空间复杂度:O(1)

