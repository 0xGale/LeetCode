注意题目中，矩阵已经给出了，不需要自己输入，这里实际上，将矩阵看成一个二叉树，要搜索的元素，在从最左下角开始搜索，如果target小于该元素则向上寻找，如果target大于该元素则向右寻找。这样时间复杂度仅为O(M+N), 空间复杂度：O(1), 这样就利用了每行从左到右递增以及每列从上到下递增的特点

![Picture1.png](https://pic.leetcode-cn.com/6584ea93812d27112043d203ea90e4b0950117d45e0452d0c630fcb247fbc4af-Picture1.png)

```python
class Solution:
    def findNumberIn2DArray(self, matrix: List[List[int]], target: int) -> bool:
        i, j = len(matrix) - 1, 0
        while i >= 0 and j <= len(matrix[0]) - 1:
            if matrix[i][j] == target:
                return True
            elif matrix[i][j] > target:
                i -= 1
            else:
                j += 1
            
        return False
```

