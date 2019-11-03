# slice

## 向 Slice 添加元素



- 添加元素时如果超越cap, 系统会重新分配更大的底层数组
- 由于值传递的关系，必须接受append的返回值

`newSlice = append(slice, value)`
