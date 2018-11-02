# Go 语言

## 什么是 Go 语言

> 2009年由Google公司开发出的开源编程语言

## 为什么使用 Go 语言

- 脚本语言开发速度
- C/C++运行速度
- 支持并发编程
- 为大数据、微服务，并发而生的通用编程语言

## Go feathers

- 简洁、快速(开发效率和运行速度)、安全
- 并行(并发处理能力)、有趣、开源
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

- http://idea.iblue.me

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

- `go ge`t 的本质就 `git` + `go install`
- go get github.com/beego/bee 映射至 $GOPATH/src/github.com/beego/bee

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

### 包管理

`package <pkgName>`

`package main`

main.main() 函数是每一个独立的课运行程序的入口点

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
cosnt hello string = "wovert"
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

