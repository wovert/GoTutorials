# Go 语言

## 什么是 Go 语言

> 2009年由Google公司开发出的开源编程语言

## 为什么使用 Go 语言

- 脚本语言开发速度
- C/C++运行速度
- 支持并发编程
- 为大数据、微服务，并发而生的通用编程语言

## Go feathers

- 垃圾回收
  - 内存自动回收，不需要developer管理内存
  - 专注业务实现
  - 只需要new分配内存，不需要释放
- 简洁、快速(开发效率和运行速度)、安全
- 并行(并发处理能力)、有趣、开源
  从语言层面支持并发，非常简单
  - goroute, 轻量级线程，创建成千上万个goroute成为可能
  - 基于 CSP (Communicating Sequential Process) 模型实现
- channel
  - 管道，雷士 类Unix pipe
  - 多个goroute 之间通过 channel 进行通信
  - 支持任何类型
  - `pipe := make(chan int, 3)` 定义3个int型管道
  - `pipe <- 1` 1 放到管理 pipe
  - `pipe <- 2` 2 放到管道 pipe
- 内存管理、数组安全、编译迅速

## Go 目前的应用

- Docker
- Kubernet
- Caddy
- Codis(豆瓣)
- Glow(类似于 Hadoop)
- Cockroach(蟑螂) 数据库

## 使用 Go 语言公司

- 百度
- 阿里
- 滴滴
- 七牛云

## Go 语言环境搭建

### 安装方式

1. 源码安装
2. Go 标准包安装
3. 第三方工具安装 比如 GVM

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

## helloworld.go

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

var (
  name ="wovert"
  age = 30
)

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
const idenfifier [type] = value
const (
  a = 0
  b //1
  c //2
)
```

## 值类型和引用类型

- 值类型：int, float, bool, string, array, struct
- 引用类型：pointer, slice, map, chan

``` go
import (
  _ "add" // 会执行add包的init函数
)
```

## time

- time.Duration 纳秒1000 = 1 Millisecond => Second => Minute => Hour => 
- time.Now()
- time.Now().Format("02/1/2006 15:04:05")
- time.Now().Format("2006/1/02 15:04")
- time.Now().Format("2006/1/02")

## 函数

### 函数的特点

- 不支持重载
- 函数是一等公民，函数可以赋值给变量
- 匿名函数
- 多返回值

无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效，而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

- map, slice, chan, 指针，interface 默认以引用方式传递

- 命令返回值的名字

### 可变参数

``` go
// 0个货多个参数
func add(arg... int) int {}

// 1个货多个参数
func add(a int, arg... int) int {
}

// 2个货多个参数
func add(a int, b int, arg... int) int {}
```

其中arg是一个`slice`, 通过`arg[index]`一次访问所有参数通过`len(arg)`来判断传递参数的个数

``` go
func read() {
  mc.Lock()
  defer mc.Unlock()
}

func read2() {
  conn, err := openConn()

  defer func() {
    if err != nil {
      conn.Close()
    }
  }()

}
```

### 内置函数

- close: 关闭channel
- len: 长度，比如 string, array, slice, map, channel
- new: 分配内存，主要用来分配值类型，比如int, struct, 返回的是指针
- make: 分配内存，主要分配引用类型，比如chan, map, slice
- append: 追加元素到数组，slice中
- panic和recover, 用来做错误处理