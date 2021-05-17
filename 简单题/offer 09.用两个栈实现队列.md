```python
from collections import deque
class CQueue:

    def __init__(self):
        self.stack1 = deque()
        self.stack2 = deque()

    def appendTail(self, value: int) -> None:
        self.stack1.append(value)

    def deleteHead(self) -> int:
        if len(self.stack1) == 0 and len(self.stack2) == 0:
            return -1
        elif len(self.stack2) > 0 :
            return self.stack2.pop()
        else:
            while self.stack1:
                self.stack2.append(self.stack1.pop())
            # for i in range(len(self.stack1)):
            #     self.stack2.append(self.stack1.pop())
            # print(self.stack2)
            return self.stack2.pop()
```

栈实现了先进后出，而队列是先进先出。

两个栈叠加使用就变成了：第一个栈先进后出，第二个栈先进入的就是第一个栈的后出的，也就实现了将第一个栈中的值进行倒置的排序。

例如：

输入值：1， 2， 3， 4， 5

第一个栈输入后：1<=2<=3<=4<=5

第二个栈将第一个栈中的值作为输入：5<=4<=3<=2<=1

将第二个栈中的值作为结果输出，就实现了先进先出，注意一点就是当第二个栈还有值，那么此时不能进行入第二个栈的操作，否则会使得后面输入的值在前面输入的值之前输出。例如：

输入值：1，2，3，删除栈中值

第一个栈：1<=2<=3

第二个栈：3<=2<=1

继续输入：4，5

第一个栈：4<=5

第二个栈：3<=2<=1<=5<=4       这是错误的

此时应该将第二个栈中的值输出完，再进行入栈操作。

## 复杂度分析

输入值：仅有一个入栈操作，时间复杂度：O（1）

输出值：需要进行出第一个栈和入第二个栈，总共加起来时间复杂度为：O（N）

空间复杂度：两个栈加起来有O（N）