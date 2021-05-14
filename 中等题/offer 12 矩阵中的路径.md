这里主要使用的是DFS 深度优先搜索算法和剪枝
其中, 这里在剑指Offer书上是使用的是回溯法
```python
class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        def dfs(i, j, k):
            if not 0 <= i < len(board) or not 0 <= j < len(board[0]) or board[i][j] != word[k]:
                return False
            if len(word) - 1 == k:
                return True
            # print(board[i][j])
            board[i][j] = ''
            res = dfs(i + 1, j, k + 1) or dfs(i, j + 1, k + 1) or dfs(i - 1, j, k + 1) or dfs(i, j - 1, k + 1)
            board[i][j] = word[k]  # 当第一个选择的值不能够确定是否为True时,还会遍历后面的,若是不还原就会导致,之前访问的值不能遍历了
            return res
        
        for i in range(len(board)):
             for j in range(len(board[0])):
                if dfs(i, j, 0): 
                    return True
        return False
                
```

其中, 注意边界的检查. 这里主要不能越界, 注意i 和len(board)之间的关系

```GO
func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i ++ {
        for j := 0; j < len(board[0]); j ++ {
            if dfs(i, j, 0, board, word) {
                return true
            }
        }
    }
    return false
}

func dfs(i, j, k int, board [][]byte, word string) bool {
    if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
        return false
    }
    if k == len(word) -1 {
        return true
    }

    board[i][j] = '0'
    res := dfs(i - 1, j, k + 1, board, word) || dfs(i + 1, j, k + 1, board, word) || dfs(i, j - 1, k + 1, board, word) || dfs(i, j + 1, k + 1, board, word) 
    board[i][j] = word[k]

    return res

}
```

