# 题目链接
https://leetcode-cn.com/problems/integer-break/
https://leetcode-cn.com/problems/jian-sheng-zi-lcof/


动态规划：
```python

class Solution:
    def cuttingRope(self, n: int) -> int:
        dp = [0] * (n + 1)
        for i in range(2, n + 1):
            for j in range(i):
                dp[i] = max(dp[i], j * (i - j), j * dp[i - j])
        return dp[n]
```
# 分析：
1. 大问题切分成小问题：n长度，切若干份，乘积最大， 那么n切分之后的长度，也是最大的， 例如：切分(n - 1)和1， 那么乘积是(n - 1) * 1, 最大就是n-1的乘积，由此可以看出大问题是由小问题组成的
2. 首先，上面切成1和(n - 1)是不对的，所以需要确定i, 也就是切分成i 和 (n - i), 乘积为i*(n - i)， i从1开始（也可以从0开始）到 n结束
3. 因为，大问题切分成小问题是会出现重复的情况，例如： f(6)*f(3), f(6)也可以切分成f(3)*f(3)， 所以使用一个数组进行保存
4. dp[i]表示有i个数时，乘积最大的值, 那么就是求dp[n]。
5. dp[i]的切分有两种情况： 1. dp[i] = j * (i - j) 2. dp[i] = j * dp[i - j] , 当i长度的绳子从j开始剪的时候， j长度， 另外一半是i - j, 这个i - j 有可能是i - j比较大，也有可能是dp[i - j]最大， 比如i = 3, j = 1, 1 * (3 - 1) = 2 和 1 * dp[2] = 1，dp[2] = 1，其中， j从0 到 i-1，dp[i]= max(j * (i - j), j * dp[i - j])
从大到小问题的计算过程中，会重复计算多个值， 但是从小到大问题，也就是从2开始计算，依次保留之前计算过的值，用到时，直接用即可，那么就可以减少重复计算。


贪心算法：
```python
class Solution:
    def cuttingRope(self, n: int) -> int:
        sum = 0
        result = 0
        if n == 2:
            return 1
        elif n == 3:
            return 2

        while n >= 3 :
            # print(n)
            n -= 3
            sum += 1
        
        if n == 2:
            result = pow(3, sum)
            result *= 2
        elif n == 1:
            sum -= 1
            result = pow(3, sum)
            result *= 4
        else:
            result = pow(3, sum)
        return result % 1000000007
```
n大于等于5时， n 按照每次切3的长度是最优解，那么切到n < 5时
1. n = 4时， 4 = 2 + 2， 2 * 2=4最大
2. n = 3时， 3 不切，最大 注意如果n只有3， 就需要切， 3 = 1 + 2， 1 * 2 = 2
3. n = 2时， 2 不切， 最大 注意如果n只有2， 就需要切， 2 = 1 + 1， 1*1 = 1

