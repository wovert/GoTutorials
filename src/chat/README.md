# chat

## 聊天服务器原理分析

- map 保存在线用户 `key => value`
  - `key = Client IP + Client Port`
  - 用户名
  - 地址
  - channel

```cgo
type Client struct {
    C chan string // 用于广播信息
    Name string
    Addr string // key一样
}

var onlineMap[string]Client
```

- 两个用户对应结构体
  - 127.0.0.1:2333 => Client
  - 127.0.0.1:1323 => Client

- 主要协程：处理用户连接
  - 1. 将用户加入到 map
  - 2. 告诉所有在线用户，谁上线了
    - `message<- 某个人上线了`

- 新开一个用户在在线协程

```cgo
for{
  msg := <- message // 如果有内容

  // 遍历map,有多少成员
  for _, cli := range onlineMap {
    cli.C <-msg
  }
}
```

- 新开一个发送内容协程

```cgo
// 传递参数cli
for msg := range cli.C {
  Write(msg)
}
```

- 新的协程，接好用户的请求，用户的数据进行转发

```cgo
// 用户发过来的数据时 buf
message <- buf

对方下限，把当前用户map里移除
```