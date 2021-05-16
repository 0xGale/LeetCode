# URL
https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
# 分析
终止条件
1. 0<= i <= m - 1, 0 <= j <= n - 1, 边界情况
2. 逻辑判断情况， i%10+i//10+j%10+j//10 <= k

逻辑过程
1. 计算当前方格是否可以移动
2. 计算当前方格[i-1][j], [i+1][j], [i][j-1], [i][j+1]是否可以移动
还有个问题：
- 计算是否会出现重复计算的情况。
  - 处理办法： 已移动过的做标记为None，方格为None时不计算
  
```python
class Solution:
    def movingCount(self, m: int, n: int, k: int) -> int:
        matrix = [[False] * n for _ in range(m)]
        self.dfs(m, n, 0, 0, k, matrix)
        sum = 0
        for i in range(m):
            for j in range(n):
                if matrix[i][j]:
                    sum += 1
        return sum
    
    def dfs(self, m, n, i, j, k, matrix):
        if not 0 <= i < m or not 0 <= j < n or matrix[i][j] or (i % 10 + i // 10 + j %10 + j // 10) > k:
            return
        
        matrix[i][j] = True

        self.dfs(m, n, i + 1, j, k, matrix)
        self.dfs(m, n, i, j + 1, k, matrix)

```
这样写的话，空间复杂度比较高，递归并且每次都包含了一个m*n的矩阵。
稍微简化一下，使用集合保存
```python
    def movingCount(self, m: int, n: int, k: int) -> int:
        ma = set()
        self.dfs(m, n, 0, 0, k, ma)
        return len(ma)
    
    def dfs(self, m, n, i, j, k, ma):
        if not 0 <= i < m or not 0 <= j < n or (i, j) in ma or (i % 10 + i // 10 + j %10 + j // 10) > k:
            return
        print(i, j)
        ma.add((i, j))

        self.dfs(m, n, i + 1, j, k, ma)
        self.dfs(m, n, i, j + 1, k, ma)
```
时间复杂度和空间复杂度都为O(mn)

