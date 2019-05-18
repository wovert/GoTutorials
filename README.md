# Go 语言

## Go 介绍

### 什么是 Go 语言

> 2009年由 Google 公司开发出的开源编程语言 GoLang。金花鼠(gordon)

### 为什么使用 Go 语言

- 脚本语言开发速度
- `C/C++` 运行速度
- 支持并发编程
- 为大数据、微服务，并发而生的通用编程语言

### Go feathers

- 垃圾回收
  - 内存自动回收，不需要 developer 管理内存
  - 专注业务实现
  - 只需要 `new` 分配内存，不需要释放
- 简洁、快速(开发效率和运行速度)、安全
- 并行(并发处理能力)、有趣、开源
  从语言层面支持并发，非常简单
  - `goroute`, 轻量级线程，创建成千上万个goroute成为可能
  - 基于 `CSP` (Communicating Sequential Process) 模型实现
- channel
  - 管道，雷士 类 `Unix pipe`
  - 多个 `goroute` 之间通过 `channel` 进行通信
  - 支持任何类型
  - `pipe := make(chan int, 3)` 定义3个int型管道
  - `pipe <- 1` 1 放到管理 pipe
  - `pipe <- 2` 2 放到管道 pipe
- 内存管理、数组安全、编译迅速

### GoLang 学习方向

- 区块链研发
- 游戏开发(数据通道)
- 服务器端开发(数据处理/大并发)
- 软件开发(分布式/云计算)
  - 盛大云 cdn
  - 京东云

### GoLang 应用领域

- 区块链应用(Blockchain technology)： 分布式账本技术，一种分布式数据库技术
  - 去中心化
  - 公开透明
  - 每个人可参与数据库记录
- 后端服务器应用
  - 美团后台流量支持程序
    - 计算能力：支撑主站后台流量（排序，推荐，搜索等）
    - 大并发能力
      - 提供负载均衡
      - cashe
      - 容错
      - 按条件分流
    - 统计运行指标(GQS, latency) 等功能
  - 仙侠道
    - 游戏服务器（通讯、逻辑、数据存储）
- 云计算/云服务后台应用
  - 盛大云CDN
    - [盛大云计算](http://www.grandcloud.cn/)
    - 应用范围：CDN的调度系统、分布系统、监控系统、短域名服务、CDN内部开放平台、运营报表系统以及其他一些小工具等
  - 京东消息推送云服务/京东分布式文件系统
    - 京东云
    - 应用范围：后台所有服务全部用 go 实现

### Go 目前的应用软件

- Docker
- Kubernet
- Caddy
- Codis(豆瓣)
- Glow(类似于 Hadoop)
- Cockroach(蟑螂) 数据库

### 使用 Go 语言公司

- Google
- AWS
- Cloudflare
- CoreOS
- 百度
- 阿里
- 滴滴
- 七牛云
- 小米
- 京东

### 计算机语言学习方法

1. 项目需求（局部刷新）
2. 解决方案
  - 传统方法 `iframe`
  - 使用新技术 `ajax`

3. 掌握新技术

  - 原理
  - 基本语法

4. 快速入门新技术
  - 安装环境
  - 基本使用

5. **深入新技术**
  - 使用规范
  - 使用陷阱
  - 使用注意点

6. 解决项目需求

## Go 语言环境搭建

### 安装方式

1. 源码安装
2. `Go` 标准包安装
3. 第三方工具安装 比如 `GVM`

[下载安装](https://studygolang.com/dl)

- IntelliJ IDEA 注册码：http://idea.lanyus.com/
- LiteIDE

> Go 语言开发的跨平台轻量级继承开发环境(IDE)

- GoLand(类似Pycharm), Eclipse, Intelli Idea等

### gopath 环境变量

> Go 语言依赖一个重要的环境变量：`$GOPATH`，不是安装目录

- $GOPATH 目录约定有三个子目录
  - src 存放源代码（比如：.go, .c, .h, .s等）
  - pkg 编译后生成的文件（比如：.a）
  - bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中）

- 当前进程设置 GOPATH 环境变量

``` cmd
> set GOPATH=c:/go_package
> echo %GOPATH%
> go get github.com/astaxie/beego
```

- `go get` 的本质就 `git` + `go install`
- go get github.com/beego/bee 映射至 $GOPATH/src/github.com/beego/bee

- 调试工具delve 安装
  - mac : `brew install go-delve/delve/delve`
  - linux&windows: `go get github.com/derekparker/delve/cmd/dlv`

- 可执行文件只有一个main 函数

## 程序

> 指令的集合

``` go
编译
# go build hello.go
编译运行
# go run hello.go
```

## 课程概览

- 基本语法
  - 变量
  - 选择/循环
  - 指针、数组、容器
- 面向接口
  - 结构体
  - duck typing的概念
  - 组合的思想
- 函数式编程
  - 闭包
  - 多样的例题
- 工程化
  - 资源管理、错误处理
  - 测试和文档
  - 性能调优
- 并发编程
  - goroutine 和 channel
  - 理解调度器
  - 多样的例题

## 班主任

- 考勤管理
- 转介绍
- 班委会
- 学院访谈
- 班级活动
- 就业指导

## 课堂纪律

- 上课时间：不要迟到，不要玩手机
- 上课氛围：积极听讲，有问题就问
- 课后练习：按时做完并提交
- 个人态度：主动学习并教授他人

## 基础语法

### 常用关键字

`break default func interface select case defer go map struct chan else goto package switch const fallthrought if range type continue for import return var`

### 标识符

> 字母和下划线开头，大小写敏感，后跟字母和数字和下划线

`_` 是特殊标识符，用来忽略结果

### 包管理

> 把相同功能的代码放到一个目录，称之为包

- 包可以别其他包引用
- main 包是用来生成可执行文件，每个程序只有一个 main包
- 包的主要用途是提高代码的可复用性

`package <pkgName>`

### var关键字

``` go
var num int

// 简写形式只能在函数体内使用
y,z := 100, "world"
fmt.Printf("%d, %s", y, z)
```

### const 常量

``` go
const PI = 3.14
const hello string = "wovert"
```

### 数据类型

- bool
- rune (符文，32bit,char类型)
- (u)int8 (u)int16 (u)int32 (u)int64 uintptr
- byte
- float32 float 64
- complex64 complex128
- string
- array slice
- map

### 变量定义

- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有 char, 只有 rune
- 原生支持复数类型

### map

make 用于内建类型(map, slice, channel)的内存分配

## file

资源关闭

```cgo
file.Open("file")
defer file.close()
```

## 异常处理

`panic("抛出异常")`

## 在线编辑 go语言

- golang.org

## init函数

> 每个源文件都可以包含一个 init 函数，这个init函数自动被 go 运行框架调用，main 函数调用之前执行

## 常量

``` go
const (
  a = 0
  b //1
  c //2
)
```

## 指针

> 指针是一个某个**内存地址的值**。这个内存地址是内存中存储的另一个变量的值的起始位置。Go语言对指针的支持介于 Java 语言和 C/C++ 语言之间，Go 语言没有向 Java 语言那样取消了对指针的直接操作的能力，也避免了 C/C++ 语言中由于对指针的滥用而造成的安全和可靠性问题。

### Go 语言的指针

> Go 语言保留了指针，但与 C 语言指针有所不同

- 默认值 `nil`
- 操作符 `&` 取变量地址，`*` 通过指针访问目标对象
- 不支持指针运算，不支持 `->` 运算符，直接用 `.` 访问目标成员

指针就是地址，指针变量就是存储地址的变量

*p: 解引用，间接引用