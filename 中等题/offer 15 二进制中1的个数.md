```go
func hammingWeight(num uint32) int {
    sum := 0
    for ;num > 0; {
        if num & 1 == 1 {
            sum ++
        }
        // sum += num & 1
        num >>= 1
    }
    return sum
}
```
输入的是十进制整数
