URL:https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    row, columns := 0, len(matrix[0]) - 1
    for row < len(matrix) && columns >= 0 {
        fmt.Println(row, columns)
        if matrix[row][columns] == target {
            return true
        } else if matrix[row][columns] < target {
            row ++
        } else {
            columns --
        }
    }
    return false
}
```

