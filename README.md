# Go 语言

## Go 介绍

### Go 语言的背景

- **Ken Tompsom** (B/Unix/C 发明人)
- **Rob Pike**(罗布·派克) Unix退队，Plan 9 OS 成员，创造 UTF-8 字符编码, 1980奥运会射箭银牌，天文学家，演讲家
- **Rober Griesemer**: 协助 Java HotSpot 编译器，Chrome 浏览器的 JavaScript 引擎 V8

### 为什么创造 Go 语言

- 计算机硬件技术更新频繁，性能提高很快。目前主流的编程语言发展明显落后于硬件，不能合理利用**多核多CPU**的优势提升软件系统性能。现有的编程语言：1. 风格不统一；2. 计算能力不够；3. 处理大并发不够好
- **软件复杂度越来越高，维护成本越来越高**，目前缺乏一个足够简洁高效的编程语言
- 企业运行维护很多 C/C++ 项目，C/C++ 程序运行速度虽然很快，但是**编译速度很慢**，同时还存在**内存泄露**的一系列困扰需要解决

### Go 语言发展史

- 2007，三大牛人开始设计
- 2009-11-10, 开源 Go 语言源码
- 2015-9-19, Go 1.5 版发布，移除了最后参与的C代码
- 2017-2-17，Go 1.8
- 2017-8-24, Go 1.9
- 2018-2-16, Go 1.10

### 什么是 Go 语言

> 2009年由 Google 公司开发出的开源编程语言 GoLang。金花鼠(gordon, Rob Pike老婆设计的)

- 高性能
- 开发效率

### 为什么使用 Go 语言

- 脚本语言开发速度
- `C/C++` 运行速度
- 支持并发编程
- 为**大数据、微服务，高并发**而生的通用编程语言

### Go 特性

- 静态类型语言的**安全和性能**，又是动态语言**开发维护**的高效率
  - Go = C + Python
  - C 静态语言编程的**运行速度**，又能达到 Python 动态语言的**快速开发**
  - 继承 C 语言的理念：**表达式语法，控制结构，基本数据类型，调用参数值，指针**等等，也保留了和 C 语言一样的编译执行方式及弱化的指针。
  - 引入包的概念，用于组织程序结构，Go 语言的**一个文件都要归属于一个包**，而不能单独存在
- **垃圾回收**
  - 内存自动回收，不需要 developer 管理内存
  - 专注业务实现
  - 只需要 `new` 分配内存，不需要释放
- **不需要依赖库**
  - C 依赖库：`ldd hello_c`列出hello_c文件依赖哭
- **并行**(并发处理能力)、**开源**
  - 从语言层面支持并发，非常简单
  - `goroute`, 轻量级线程，创建成千上万个 `goroute` 成为可能
  - 基于 `CSP` (Communicating Sequential Process) 模型实现
- 管理通信机制，形成 Go 语言特有的管道 **channel**
  - 管道，类似 `Unix pipe`
  - 多个 `goroute` 之间通过 `channel` 进行通信
  - 支持任何类型
  - `pipe := make(chan int, 3)` 定义3个int型管道
  - `pipe <- 1` 1 放到管理 pipe
  - `pipe <- 2` 2 放到管道 pipe
- 内存管理、**数组安全、编译迅速**

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

### Go SDK

> Go SDK: Go 语言开发工具包。包括编译、运行和API等

[下载安装 Go SDK](https://studygolang.com/dl)

- go1.9.2 darwin-amd64.pkg: Mac 下的 SDK 图像化安装包
- go1.9.2 darwin-amd64.pkg

- go1.9.2.freebsd-amd64.tar.gz: Unix 下的 SDK

- go1.9.2.linux-amd64.tar.gz: Linux 下的 SDK
- go1.9.2.src.tar.gz: 源码

- go1.9.2.windows-amd64.zip: Windows 下的 SDK

### 安装方式 

1. 源码安装
2. `Go` 标准包安装
3. 第三方工具安装，比如 `GVM`

### gopath 环境变量

> GOPATH：为我们开发常用的目录，建议不要和 Go 的安装目录一致

- `GOROOT`: SDK 安装目录
- `Path`: 添加 SDK 的 `/bin` 目录
- `$GOPATH` 工作目录（Go项目的工作路径）：目录约定有三个子目录
  - `src` 存放源代码（比如：.go, .c, .h, .s等）
  - `pkg` 编译后生成的文件(.a文件)（非main函数的文件在go install后生成）
  - `bin` 存放编译后生成的可执行文件（为了方便，可以把此目录加入到 `$PATH` 变量中）

- `GOBIN`：是 `GOPATH` 下的 `bin` 目录
- `Path`：环境变量，需要 `go-bin` 目录加入到 `path` 路径下，生成可执行文件就可以直接运行了

**可执行文件只有一个`main` 函数的文件**

### windows 安装及配置

1. 下载 go1.9.2.windows-amd64.zip
2. 安装路径不要有中文或者特殊符号如空格等
3. SDK 安装目录建议：`E:/usr_local`
4. 安装时，基本上是傻瓜式安装，解压就可以使用

- go
  - bin
    - go.exe 编译和运行 go 源码
    - godoc.exe
    - gofmt.exe

```cmd
检测 SDK 是否安装成功
E:usr_local\go\bin> go version
```

#### Windows 下配置环境变量

> 当前进程设置 GOPATH 环境变量

```cmd
d:\> go version
'go' 不是内部或外部命令，也不是可运行的程序或批处理文件
```

错误分析：当前执行的程序在当前目录下不存在，Windows OS 会在系统中已有的一个名 `path` 环境变量指定的目录中查找。如果仍未找到，会出现以上的错误提示。所以进入到 `go` 安装路径 `\bin` 目录下，执行 `go`，会看到 `go` 参数的提示信息

1. 我的电脑 -> 属性 -> 高级系统设置 -> 环境变量 -> 系统变量 -> 编辑 Path
2. 拷贝 Go SDK 安装目录 `E:\usr_local\go` 到 `GOROOT` 环境变量中
2. `%GOROOT%\bin`添加到 `PATH` 环境变量中
3. go 工作目录 `E:\goProject` 添加到 `GOPATH` 环境变量中

测试环境变量
``` cmd
> echo %GOPATH%
> go get github.com/astaxie/beego
```

- `go get` 的本质就 `git` + `go install`
- `go get github.com/beego/bee`
  - `go get github.com/beego/bee` 映射至 `$GOPATH/src/github.com/beego/bee`

### Linux 安装及配置

1. 查看 Linux 位数版本 `uname -a`
2. 下载 SDK go1.9.2.linux-adm64.tar.gz
3. 复制到 `/opt` 目录下
4. 解压文件 `tar -zxvf go1.9.2.linux-adm64.tar.gz`
5. 配置环境变量

```sh
# vim /etc/proile.d/go.sh
  export GOPATH=/opt/go
  export PATH=$PATH:$GOPATH/bin
  export GOPATH=$HOME/goProjects/
# source /etc/profile.d/go.sh
# go version
```

### mac 安装及配置

1. 下载 go1.9.2.darwin-amd64.tar.gz
2. 复制到 `用户目录/go_dev/go`
3. 剩下与 Linux 安装方式一样

```sh
$ brew install go
$ go version
$ vim ~/.bash_profile
  我的源码库没有跟安装目录放在一起
  1）单源码库环境变量配置
  export GOPATH=/Applications/MAMP/htdocs/go
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:$GOBIN
  2）多源码库环境变量配置
  export GOPATH=/Applications/MAMP/htdocs/go(:自由添加目录，其他不变)
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:${GOPATH//://bin:}/bin
$ source ~/.bash_profile
```

### Go 语言的开发工具

- VS Code
- IntelliJ IDEA 注册码：http://idea.lanyus.com/
- LiteIDE
  - Go 语言开发的跨平台轻量级继承开发环境(IDE)
- GoLand (类似 Pycharm), Eclipse, Intelli Idea等

#### 在 Linux OS 安装 vscode

1. 下載 code-stable-codexxxx.amd64.tar.gz
2. 文件复制到 `/opt` 目录下
3. 解压文件 `tar -zxvf code-stable-codexxxx.amd64.tar.gz`
4. 进入解压后的目录运行vscode 程序 `./code`

#### 在 Mac OS 安装 vscode

```sh
启动 sshd 服务
$ sudo launchctl loadl -w /System/Library/LaunchDaemons/ssh.plist

停止 sshd 服务
$ sudo launchctl unload -w /System/Library/LaunchDaemons/ssh.plist

查看是否启动
$ sudo launchctl list | grep ssh

显示 0 com.openssh.sshd 代表启动成功了

解压
$ unzip VScode-darwin-stable.zip
```

### 调试工具delve 安装

- mac : `brew install go-delve/delve/delve`
- linux&windows: `go get github.com/derekparker/delve/cmd/dlv`

## 程序

> 指令的集合

文件编码必须是 `utf-8`

``` shell
编译
$ go build hello.go

编译运行
$ go run hello.go
```

### go 命令行

- go get : 获取远程包（需提前安装 git 或 hg）
- go run : 直接运行程序
- go build : 测试编译、检查是否有编译错误
- go fmt : 格式化源码
- go install : 编译包文件并编译整个程序
- go test : 运行测试文件
- go doc : 查看文档 

## 快速入门

### 开发步骤

1. 目录结构

- E:\goProject 工作目录
  - src 源码
    - go_code
      - project1
        - main
        - package
      - project1

2. 创建 demo 文件 `vim project1/main/hello.go`

```go
package main // 文件归属 main 包
import "ftm"
func main() {
  fmt.Println("hello World")
}
```

3. 通过 `go build hello.go` 命令对 go 文件进行编译，生成 .exe 文件

**生产环境必须先编译**

```sh
# cd project1/main
# go build hello.go
# hello.exe
```

4. 通过 `go run hello.go` 命令编译执行程序

**用于开发环境**

```sh
# cd project1/main
# go run hello.go
```

### GoLang 执行流程分析

- 源码编译，在执行，Go执行流程
  - .go文件 ---- `go build`(编译) ---> 可执行文件 ---- 运行 ----> 结果

- 源码直接执行 `go run`
  - .go文件 ---- `go run`(编译和运行) -----> 结果

#### 执行流程区别

1. 先编译生成可执行文件，可以将可执行文件复制到没有 go 开发环境的机器上可以运行
2. 直接 `go run 源码文件`，如果在另外一个机器上运行，需要 go 开发环境，否则无法执行
3. 在编译时，编译器会将程序运行依赖的库文件包含在可执行文件中，所以可执行文件变大了很多

### 什么是编译

1. 通过编译器将其编译成机器(CPU指令)可以识别的二进制码文件
2. 指定编译输出的文件：`go build -o hi.exe hello.go`

### go 开发注意事项

- 源文件必须以`.go` 扩展名
- 应用程序的执行入口是 `main()`方法
- 严格区分大小写
- 每个语句后不需要分号（自动加分号(`;`)）
- Go 编译器逐行进行编译的，因为每行写一条语句，不能把多条语句写在同一个，否则报错
- 定义的变量或者 `import` 的包没有使用，那么代码不能编译通过
- 大括号都是成对儿出现的，缺一不可

### 注释

- `// 行注释`
- `/* 块注释 */`

### 规范代码风格

- Shift + Tab
- `gofmt` 进行格式化
  - `gofmt -w hello.go`

## 基础语法

### 常用关键字

`break default func interface select case defer go map struct chan else goto package switch const fallthrought if range type continue for import return var`

### 标识符

> 变量、方法和函数等命名时使用的字符序列称为标识符。字母和下划线开头，大小写敏感，后跟字母和数字和下划线

`_` 是特殊标识符，用来忽略结果

- 标识符规范
  - 包名与目录名保持一致
  - 变量名、函数名、常量名采用驼峰法
  - 变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用

### 包管理

> 把相同功能的代码放到一个目录，称之为包

- 包可以别其他包引用
- `main` 包是用来生成可执行文件，每个程序只有一个 `main` 包
- 包的主要用途是提高代码的可复用性

`package <pkgName>`

### var 关键字

``` go
var num int

// 简写形式只能在函数体内使用: 自动推到类型（同一个变量只是使用一次，用于初始化那次）
y,z := 100, "world"
fmt.Printf("%d, %s", y, z)
```

### const 常量

``` go
const PI = 3.14
const hello string = "wovert"
```

### 变量

> 内存中一个数据存储空间

1. 声明变量
2. 变量赋值
3. 使用变量

- 声明变量之后，没有赋值会使用默认值(`int=0, string=""`)
- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有 `char`, 只有 `rune`
- 原生支持复数类型
- 该数据内类的值可以在同一类型范围内不断变化
- 变量在同一个作用域内不能重名
- 变量 = 变量名 + 值 + 数据类型

```go
var (
  name = "wovert"
  age = 30
)
```

## 数据类型

### 基础数据类型

| 类型 ｜ 名称 ｜ 长度 ｜ 默认值 ｜ 说明 ｜
｜ --- ｜ --- ｜ ---- ｜ ----  ｜ --- ｜
| bool ｜ 布尔类型 ｜ 1 ｜ false ｜ true/false 不可以用数字代表 ｜
| byte ｜ 字节型 ｜ 1 ｜ 0 ｜ uint8 别名 ｜
| rune ｜ 字符型 ｜ 4 ｜ 0 ｜ 专用于存储 unicode 编码，等价于uint32 ｜
| int, uint ｜ 整型 ｜ 4 or 8 ｜ 0 ｜ 32 bit or 64 bit ｜
| int8, uint8 ｜ 整型 ｜ 1 ｜ 0 ｜ -128 ~ 127, 0 ~ 255 ｜
| int16, uint16 ｜ 整型 ｜ 2 ｜ 0 ｜ -32768 ~ 32767, 0 ~ 65535 ｜
| int32, uint32 ｜ 整型 ｜ 4 ｜ 0 ｜ -21亿 ~ 21亿, 0 ~ 42亿 ｜
| int64, uint64 ｜ 整型 ｜ 8 ｜ 0 ｜ ｜
| float32 ｜ 浮点型 ｜ 4 ｜ 0.0 ｜ 小数位精确到 7 位 ｜
| float64 ｜ 浮点型 ｜ 8 ｜ 0.0 ｜ 小数位精确到 15 位 ｜
| complex64 ｜ 符合类型 ｜ 8 ｜  ｜  ｜
| complex128 ｜ 符合类型 ｜ 16 ｜  ｜  ｜
| uintptr ｜ 整型 ｜ 4 or 8 ｜  ｜ 存储指针 uint32 or uint64 整数  ｜
| string ｜ 字符串 ｜  ｜ "" ｜ utf-8 字符串 ｜

### fmt 包的格式化输出输入

| 格式 | 含义 |
| ---- | ---- |
| %% | %字面量 |
| %b | 二进制整数值（基数为2）或者是科学计数法表示的指数为2的浮点数 |
| %d | 十进制数值（基数为10） |
| %e | 科学计数法表示的浮点数或者复数值 |
| %E | 科学计数法表示的浮点数或者复数值 |
| %f |浮点数或者复数值 |
| %g | %e或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出 |

| %G | %E或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出 |
| %o | 八进制整数值 |
| %p | 十六进制表示的一个值的地址。前缀为0x后跟小写的a-f表示 |
| %s | 字符串 |
| %t | true/false |
| %T | 值得类型 |
| %v | 默认格式输出的内置或者自动移类型的值，或者是使用其类型的 String()方式输出的自定义值，如果该防范存在的话 |
| %U | Unicode 表示整数码点，默认值为4个数字字符 |
| %x | 16进制整数值 a-f |
| %X | 16进制整数值 A-F |

### 类型转换

- 不允许隐式转换，所有类型必须显式声明，而且转换只能两个相互兼容的类型之间
- bool 与 整型不能相互转换
- 0就是假，非0就是真
- 字符类型本质就是整型

### 类型别名

```cgo
// int64 类型改名为bigint
type bigint int64
var x bigint = 100

type (
  myint int // int改名为int
  mystr string // string改名为 mystr
)
```

### 整数值

- 默认声明未 int 类型

### 浮点型

- 浮点数 = 符号位 + 指数位 + 尾数位
- 有固定的范围和字段长度，不受具体 OS 的影响
- 默认声明未 float64 类型

### 字符类型

- 字符型存储到计算机中，需要将字符对应的码值（整数）找出来
  - 存储：字符->对应码值->二进制->存储
  - 读取：二进制->码值->字符->读取
- Go 语言的编码都统一使用 `utf-8`

### 布尔类型

- 占一个字节
- bool 取值范围：`true | false`，不能使用0或非0

### 字符串类型

- 使用 UTF-8 编码标识 Unicode 文本
- 字符串一旦赋值，就不能修改。字符串是不可变的
- "会识别转义字符"
- `反引号，已字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果`
- 字符串拼接换行时，`+` 保留在行尾


#### 基本数据类型转换

- 必须显示转换，不能自动转换
T(v)

```go
var i int32 = 257
var n1 float32 = float32(i)
var n2 int8 = int8(i)
var n3 int64 = int64(i)
var n4 byte = byte(i)
fmt.Printf("i=%v n1=%v n2=%v n3=%v n4=%v\n", i, n1, n2, n3, n4)
fmt.Printf("i=%T \n", i)
```

- 范围小的可以转换为范围大的，也可以范围大的转换为范围小的
- 被转换的是变量存储的数据，变量本身的数据类型并没有变化
- `int64` 转换成 `int8`，编译时不会拨错，只是转换的结果是按**溢出**处理

### 复合类型

| 类型 ｜ 名称 ｜ 长度 ｜ 默认值 ｜ 说明 ｜
｜ --- ｜ --- ｜ ---- ｜ ----  ｜ --- ｜
| pointer ｜ 指针 ｜  ｜ nil ｜  ｜
| array ｜ 数组 ｜  ｜ 0 ｜ ｜
| slice ｜ 切片 ｜  ｜ nil ｜ 引用类型 ｜
| map ｜ 字典 ｜  ｜ nil ｜ 引用类型 ｜
| struct ｜ 结构体 ｜  ｜ nil ｜ ｜


- 派生/复杂数据类型
  - 指针
  - 数组
  - 结构体
  - 管道
  - 函数
  - 切片
  - 接口
  - map


## map

`make` 用于内建类型(`map, slice, channel`)的内存分配

### 创建方式

1. `var m1 map[int]string` 不能存储数据
2. `m2 := map[int]string{}` 能存储数据
3. `m3 := make(map[int]string)` 默认len=0
4. `m4 := make(map[int]string, 10)`

### 初始化：
  
- 1. `var m map[int] string = map[int]string{1:"name",2:"age"}`
- 2. `m := map[int]string{1:"name",2:"age"}`

### 赋值：
  
- m := make(map[int]string, 1)
- m[700] = "nami"
- m[20] = "hello"
- m[3] = "world"
- m[3] = " world"


新map元素key与原map元素相同，则覆盖


### 遍历元素

```
if v, has = m[1]; has {
  println("has")
} else {
  println("no has")
} 
```

### 删除元素

`delete(map, key)`

## file

1. os.Create 文件不存在创建，文件存在，将文件内容清空
2. open 打开文件，以只读方式打开文件
3. OpenFile,以只读、只写、读写方式打开文件

资源关闭

``` go
file.Open("file")
defer file.close()
```

### 写文件

- 按字符串写
- 按位置写
- 按字节写

### 读文件

- 按行读
  - 1.创建一个带有缓冲区的 Reader(读写)
  - 2.从reader的缓冲区中，读取指定长度的数据。数据长度取决于参数delime
- 按字节写

### 缓冲区



## 异常处理

`panic("抛出异常")`

## 在线编辑 go 语言

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

- 值类型：`int, float, bool, string, array, struct`
- 引用类型：`pointer, slice, map, chan`

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

## 运算符

- 算符运算符: +(正号, 加, 字符串相加), -(符号, 减), *, /(取整数), %, i++, i--
  - `10 / 4 == 2`
  - `10.0 / 4 = 2.5`
  - `10%-3 == 1` 与除数有关系
  - i++, i-- 必须独立使用
- 赋值运算符: `=, +=, -=, *=, /=, %=`
- 比较运算符/关系运算符: `==, !=, <, >, <=, >=`
- 逻辑运算符: `&&, ||, !`
- 位运算符
- 其他运算符: `&, *`

### 优先级

- 括号, ++, --
- 单目运算符：+,-,!,~, (type), *, &, sizeof **右到左**
- 算符
- 移位
- 关系
- 位
- 逻辑
- 赋值 **右到左**
- 逗号

## 键盘输入语句

- fmt.Scanln()
- fmt.Scanf()

## 流程控制语句

- 顺序控制语句
- 分支控制语句
  - 单分支
  - 双分支
  - 多分支
- 循环控制

- `if a := 1; a < 100 {`
- `switch num := 100; num {`
  - `fallthrough` 与 `break` 相反

```go
switch num :=1; num {
  case 1:
    ...
  case 2, 3, 4:
    ...
  default:
    ...
}

score := 85
switch {
  case socre > 90:
    ...
  case score > 80:
    ...
  default:
    ...
}
```

- range str: 元素索引，元素值

```go
for i := 1; i <= n; i++ {
  ...
}
for i <= n {
  ...
  i++
}
str := "abc"
for i, data := range str {
  // i 元素索引
  // data 元素值
}
// 值返回原素值
for _, data := range str {
  ...
}
// 仅返回元素索引
for i := range str {
  ...
}
```

- break and goto

```go
label:
for ...
  for ...
    break label
```

## 包

- `import` 包时，路径从 `$GOPATH` 环境变量的 `src` 目录开始，编译器自动从 `src` 下开始引入
- 同一个包下，不能有相同的函数名，否则报重复定义

```go
import (
  // 别名
  util "go/tools/utils"
)
```

编译后生成一个有默认名的可执行文件，在 $GOPATH 目录下，可以指定名字和目录 `go build -o bin/my.exe go_code/project/main`

- `pkg/windows_amd64/go_code/project/function/utils.a` 库文件

## 函数

```go
func 函数名(形参列表)(返回值列表) {
  ...
  return 返回值列表
}
```

### 函数的特点

- 不支持重载
- 函数是一等公民，函数可以赋值给变量
- 匿名函数
- 多返回值

无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效，而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

- `map, slice, chan, pointer，interface` 默认以引用方式传递

- 命令返回值的名字

### 参数的传递

- 值传递：基本数据类型 int系列, float系列， bool, string, 数组和结构体
- 引用传递: 指针，slice 切片, map, 管道 chan, interface 等都是引用类型

不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是地址的拷贝。地址拷贝效率高，因为数据量小，而值拷贝决定拷贝的数据大小，数据越大，效率越低

### 可变参数

``` go
// 0个或多个参数
func add(arg... int) int {}

// 1个或多个参数
func add(a int, arg... int) int {
}

// 2个货多个参数
func add(a int, b int, arg... int) int {}
```

其中 `arg` 是一个`slice`, 通过`arg[index]`一次访问所有参数通过`len(arg)`来判断传递参数的个数

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

- `close`: 关闭 channel
- `len`: 长度，比如 string, array, slice, map, channel
- `new`: 分配内存，主要用来分配值类型，比如int, struct, 返回的是指针
- `make`: 分配内存，主要分配引用类型，比如chan, map, slice
- `append`: 追加元素到数组，slice中
- `panic`和`recover`, 用来做错误处理

### 函数是数据类型

- 自定义数据类型
  - `type 自定义数据类型名 数据类型`
  - `type myInt int` 给int取了别名，myInt和int都是 int类型，但是go认为是两个不同类型
  - `type mySum func(int, int) int`

- 支持返回值命名

```go
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
  sum = n1 + n2
  sub = n1 - n2
  return
}
```

### 匿名函数和闭包

```go
sum := func (a int, b int) func () int {
  a *= a
  b *= b
  return func () int {
    return a + b
  }
}
```

- `range str`: 元素索引，元素值
### defer

在函数中需要创建资源（数据库连接、文件句柄、锁等），为了在函数执行完毕后，及时的释放资源，Go 的设计者提供defer(延时机制)

```go
func sum(n1 int, n2 int) int {
  defer fmt.Println("n1=", n1) // 压栈
  defer fmt.Println("n2=", n2) // 压栈
  res := n1 + n2
  fmt.Println("res=", res)
  return res
}
```

- 遇到defer 语句时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈中
- 当函数执行完毕之后，再从 defer 栈，**先入后出**的方式出栈（一次从栈顶去除语句执行）
- 在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈

### 变量作用域

- 局部变量：函数内声明/定义的变量，作用域**仅限于函数内部**
- 全局变量：函数外声明/定义的变量，作用域在**整个包**都有效，如果其**首字母为大写**，则作用域在**整个程序有效**
- 代码块：在 for/if 中，作用域在该代码块中

### 注释事项

- 函数外不能使用 `Name := "tome"`，因为等价 `var Name string Name="tom"`，**赋值语句不能在函数体外**

## 指针

> 指针是一个某个**内存地址的值**。这个内存地址是内存中存储的另一个变量的值的起始位置。Go语言对指针的支持介于 Java 语言和 C/C++ 语言之间，Go 语言没有向 Java 语言那样取消了对指针的直接操作的能力，也避免了 C/C++ 语言中由于对指针的滥用而造成的安全和可靠性问题。

### 栈帧的内存布局

- 栈帧：用来给函数运行提供内存空间，取内存上 stack 上。当函数调用时，产生栈帧，函数调研结束，释放栈帧。
- 栈帧存储
  - 1)局部变量
  - 2)形参（形参与局部变量存储地位等同）
  - 3)内存字段描述值

- 栈基指针（调用函数开始） VS 栈顶指针（调用函数结束）

- stack(默认1M, Windows OS 可以扩展到8M, Linux 可以扩展到 16M)
- heap
- .bss(未初始化数据区) RW
- .data(初始化的数据区）RW
- .rodata(只读数据区) RO
- .text(代码区) RO

### Go 语言的指针

> Go 语言保留了指针，但与 C 语言指针有所不同

- 默认值 `nil`
- 操作符 `&` 取变量地址，`*` 通过指针访问目标对象
- 不支持指针运算，不支持 `->` 运算符，直接用 `.` 访问目标成员

指针就是地址，指针变量就是存储地址的变量

*p: 解引用，间接引用

## xorm 操作数据库

```
go get github.com/go-xorm/xorm
go get github.com/go-sql-driver/mysql
```

## go get golang.org/x 包失败

- gopm 来获取无法下载的包
```sh
$ go get -v github.com/gpmgo/gopm
$ gopm get -g -v golang.org/x/tools/cmd/goimports
```


Go version >= 1.11
```sh
export GO111MODULE=on
export GOPROXY=https://goproxy.io
go mod ...
Windows 用户，可以在 PowerShell 中设置：

$env:GOPROXY = "https://goproxy.io"
```


// string到int
int, err := strconv.Atoi(string)

// string到int64
int64, err := strconv.ParseInt(string, 10, 64)

// int到string
string := strconv.Itoa(int)

// int64到string
string := strconv.FormatInt(int64,10)

## leetcode.com


### 空指针和野指针

- 空指针：为初始化的指针 `var p *int` `*p --> error`
- 野指针：被一片无效的空间初始化 `var p *int = 0`

- `0x000-0xff` 内存地址给操作系统使用

heap 上申请一片内存地址空间

```go
var pStr *string
pStr = new(string)
```

- 变量存储
  - 等号左边的变量，是变量所指向的内存空间(写)
  - 等号右边的变量，是变量内存空间存储的数据值(读)

- 指针的函数传递
  - 传地址（传引用）：将形参的地址值作为函数参数传递
    - 传引用：在A栈帧内部，修改B栈帧中的变量值
  - 传值（数据值）：将实参的值拷贝一份给形参
  

## 字符串函

- 字符串长度(字节): `len(str)`
  - 一个 ascii 占用一个字节
  - 一个汉字占用3个字节
- `str1 := []rune(str)`
- 字符串转整数：`n, err = strconv.Atoi("12")`
- 整数转字符串：`str = strconv.Itoa(1234)`
- 字符串转 []byte: `var bytes = []byte("hello World")`
- []byte转字符串：`str = string([]byte{97,98,99})`
- 10进制转2,8,16进制：`str = strconv.FormatInt(123, 2) // 2->8,16`
- 查找子串是否在指定的字符串中：`bool strings.Contains("seafood", "foo")`
- 统计一个字符串有几个指定的子串：`strings.Count('ceheese', 'e')`
- 字符串比较（不区分大小写）: `bool strings.EqualFold("abc", "Abc")`
  - 区分大小写：`abc == Abc`
- 返回子串在字符串第一次出现的 `index` 值，没有返回 -1: `strings.Index("NLT_abc", "abc")`
- 返回子串在字符串最后一次出现的 `index` 值，没有返回 -1: `strings.LastIndex("NLT_abc", "abc")`
- 子串替换成另外一个子串(n:替换几个,-1替换所有)： `strings.Replace("go go hello", "go", "go 语言", -n)`
- 大小写转换：`strings.ToLower/ToUpper()`
- 字符串分割成字符串数组: `int strings.Split("hello World", ",")`
- 字符串左右两边空格去掉：`strings.TrimSpace(" tn a lone sss...  ")`
- 字符串左右两边指定的字符去掉：`strings.Trim("!Hello!", "!")`
- 字符串左边指定的字符去掉：`strings.TrimLeft("!Hello!", "!")`
- 字符串右边指定的字符去掉：`strings.TrimRight("!Hello!", "!")`
- 是否以某个字符串开头：`strings.HasPrefix("http://www.wovert.com", "http")`
- 是否以某个字符串结尾：`strings.HasSuffix("http://www.wovert.com", ".com")`

## 日期时间相关函数

> time 包

- now := time.Now()
- now.Year()
- now.Month()
- int(now.Day())
- now.Hour()
- now.Minute()
- now.Second()

- 格式化日期时间
  - `dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d \n", now.Year(),now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())`
  - `fmt.Println(now.Format("2006/01/02 15:04:05"))`

- 时间常量
- 休眠时间: time.Sleep(time.Second)
- 随机时间
  - now.Unix()
  - now.UnixNano()
- 执行代码时间
  - start := time.Now().Unix()
  - test()
  - end := time.Now().Unix()

## 错误处理机制

```go
func test() {
  // defer + recover 来捕获处理异常
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("err=", err)	
      fmt.Println("发送邮件给admin@amdin.com")
    }
  }()
  num1 := 10
  num2 := 0
  res := num1 / num2
  fmt.Println("res=", res)
}
```

### 自定义错误

```go
// 读取配置文件的init.conf 的信息
// 文件传入错误，返回自定义错误
func readConf(name string) (err error) {
  if name == "config.ini" {
    return nil
  } else {
    return errors.New("读取文件错误")
  }
}

func test2() {
  err := readConf("config.ini0")
  if err != nil {
    // 读取文件发送错误，输出错误，并终止程序
    panic(err)
  }

  fmt.Println("test2继续执行剩余代码...")
}
```

## 数组

> 存放多个同一类型数据，是值类型


### 数组初始化

```go
var a [5]int = [5]int{1,2,3,4,5}

// 自动推倒类型
b := [5]int{1,2,3,4,5}

// 部分初始化
c := [5]int{1,2,3} // [1 2 3 0 0]

// 指定某个元素初始化
d := [5]int{2: 10, 4:20} // [0 0 10 0 20]

```

### 数组使用注意事项

- 声明/定义数组之后，长度固定，不能动态变化
- 数组是值类型
- 指定的数组元素个数必须是常量
- 下标可以是变量或常量

### 数组比较和赋值

- 支持 == 或 !=
  - 比较数组类型和元素是否一样

## 切片

### 为什么使用切片

1. 数组的容量固定，不能自动扩展
2. 值传递，数组作为函数参数时，将整个数组值拷贝一份给形参

Go语言中，几乎可以在所有的场景中，使用切片替换数组使用

### 什么是切片

> 切片并不是数组或数组指针，一种数据结构体，用来操作数组内部元素。它通过内部指针和相关属性引用数组片段，以实现变长方案。

- 切片时引用类型，引用传递
- 动态变化数组
- `var sliceName [这里不需要长度]类型`

```go
// runtime/slice.go
type slice struct {
  ptr *[2]int // low 指针地址
  len int // 元素长度
  cap int // 容量大小
}
```

### 如何使用切片

`切片名称 [low:high:max]`

- low: 起始下标位置
- high: 结束下标位置(不包含index of high)
- cap(容量): max-low (容量是动态变化，必须大于等于切片元素数量)

1. 定义一个切片，然后让切片引用一个已经创建好的数组：切片或数组都可以访问元素

```go
intArr := [...]int{1,2,3,4,5,6}
slice := intArr[0:3:4] // index:1-3(不包含3)
```

2. 通过 make 来创建切片: 只能通过 slice 下边访问元素

`var sliceName []int = make([]int, len, cap)`

3. 直接指定具体数组

`var slice []int = []int {1,3,4}`

- 总结
  - 通过 make 方式创建切片可以指定切片的大小和容量
  - 如果没有给切片各个元素赋值，那个就会使用默认值：0, "", false
  - 通过 make 方式创建的切片对应的数组是 由 make 底层维护，对外不可见，即智能通过 slice 去访问各个元素

### 数组和切片定义区别

- 创建数组 [] 指定数组长度
- 创建切片时, [] 为空 或者 ...

截取数组，初始化切片时，切片容量跟随原数组

- `append(slice, apendElement)`
  - 智能的底层数组的容量增长，一旦超过底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的数据
- `copy(targetSlice, sourceSlice)`
  - 源元素覆盖目的元素

## struct

> 结构体就是自定义数据类型

### 普通变量定义和初始化

```cgo
type Person struct {
  name string
  sex uint8
  age int
}

```

1. 顺序初始化

```
var man Person = Person{"wovert", 1, 20}
```

2. 指定成员初始化

```cgo
// 未初始化的成员变量，取该数据类型对应的默认值
man := Person{name:"wovert", age:18}
```

### 结构体变量的比较

1. 只能使用 `==` or `!=`
2. 相同结构体类型(成员变量的类型、个数、顺序一样)变量之间可以直接赋值

### 结构体地址

> 结构体变量的地址 = 结构体首个元素的地址



### 结构体传参

- `Unsafe.Sizeof(变量名)` 变量类所占用的内存空间大小
- 将结构体变量的值拷贝一份，传递。内存消耗大，效率低。


### 指针变量定义和初始化

1. 顺序初始化：依次将结构体内部所有成员初始化

`var man *Person = &Person{"name", 'm', 20}`

2. `new(Person)`


### 结构体指针地址

> 结构体指针变量的值 = 结构体首个元素的地址


### 结构体指针传参

- `Unsafe.Sizeof(指针)` 64bit OS，大小未8字节
- 将结构体变量地址拷贝一份，传递。使用频率非常高～


**结构体指针做函数返回值：不能返回局部变量的地址值。局部变量保存在栈帧上，函数调用结束后，栈帧释放，局部变量的地址，不再受系统保护，随时可能分配给其他程序**

## oop

### go 接口

> 一个类只需要实现了接口要求的所有函数，这个类实现的接口

- 接口赋值
  - 将对象实例赋值给接口
  - 将一个接口赋值给另一个接口

## 协程 Coroutine

> 轻量级 “线程”

协程可以创建百万个而不会导致系统资源衰竭，而线程和进程通常多也不能超过1万个。这也是协程也叫做轻量级线程的原因

- **非抢占式**多任务处理，由协程主动交出控制权
- **编译器**/解释器/虚拟机层面的多任务
- 多个协程可能在一个或多个线程上运行

go + 函数名：启动一个协程执行函数体

### 并发

- 宏观：用户体验上，程序并行执行
- 微观：多个计划任务，顺序执行。快速的切换，轮换使用CPU时间轮片

- 进程并发
  - 程序：编译成功得到的二进制文件。占用磁盘空间
  - 进程：运行的程序，占用系统资源（内存）
  
- 进程状态：初始态、就绪态、运行态、挂起（阻塞）态、终止（停止）态

- 线程并发：
  - 线程：LWP 轻量级的进程。最小的执行单位。—— CPU分配时间轮片的对象
  - 进程：最小的系统资源分配单位
  
- 同步
  - 协同不调：按照预定的先后次序运行
  - 线程同步：一个线程发出某一功能调用时，在没有得到结果之前，该调用不返回。同时其他线程为保证数据一致性，不能调用该功能。
  - 线程同步机制：
    - 互斥锁（互斥量）：建议帧。拿到锁以后，才能访问数据，没有拿到锁的线程，阻塞等待。等到拿锁的线程释放帧
    - 读写锁：一把锁（读属性、写属性）。写独占、读共享，写锁优先级最高
    - 信号量
    - 条件变量

- 协程并发：Python、Lua、Rust

- 进程：稳定性强（生产线）
- 线程：节省资源（工人）
- 协程：效率高（利用闲暇时间做任务）


### goroutine 特性

- 主 goroutine 退出后，其他的工作 goroutine 也会自动退出


### runtime 包

- Goshed: 出让当前go程所占用的 CPU 时间片
- Goexit: 退出当前go程，也就是调用go协程函数
  - return: 返回当前函数调用处，return之前的 defer 注册生效
- Gomaxprocs: 设置并行计算的CPU核数的最大值，并返回之前的CPU核数量值


## channel

> 对应一个管道（通道FIFO）

`make(chan 在通道中传递的数据类型, 容量)`

- 容量 0: 无缓冲channel
- 容量 > 0: 有缓冲channel

`make(chan int) = make(chan string, 0)`

- 写端（传入端） `chan <-` 写端写数据，读端不再读，阻塞
- 读端 (传出端) `<- chan` 读端读数据，同时写端不再写，读端阻塞

读端和写端必须同时满足条件，才在栈上进行数据流动，否则阻塞。

### 无缓冲的 channel

> 无缓冲的通道（unbuffered channel）是指在接受前没有能力保存任何值的通道

通道容量为0， len=0

无缓冲的channel 应用于两个协程中，一个读，一个写

具备同步的能力。读、写同步

这种类型的通道要求发送 goroutine 和接受 goroutine 同时准备好，才能完成发送和接受操作。否则，通道会导致先执行发送或接受操作的goroutine 阻塞等待。

这种对通道进行发送和接受的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在。

- **阻塞**：由于某种原因数据没有到达，当前协程（线程）持续处于等待状态，知道条件满足，才能解除阻塞。
- **同步**：在两个或多个协程（线程）间，保持数据内容一致性的机制

[无缓冲的逻辑](./images/1.gif)


### 有缓冲的 channel

> 通道容量为非0，channel中剩余未读取数据个数。cap(ch), 通道的容量

channel 应用于两个协程中，一个读，一个写

缓冲区可以进行数据存储，存储至容量上限，阻塞。具备**异步**能力。


### 关闭channel

> `close(ch)`

确定不再相对发送、接受数据。关闭 channel。使用 close(ch) 关闭 channel
对端可以判断 channel 是否关闭

- 总结
  - 1.数据没有发送完，不应该关闭channel
  - 2.已经关闭的channel，不能向其写数据。报错 "panic: send on closed channel"
  - 3.写端已经关闭channel, 可以从中读取数据。
    - 读无缓冲channel: 读到 0/false/"" —— 说明：写端已经关闭
    - 读有缓冲channel: 如果缓冲区内有数据，先读数据，读完数据后，可以继续读，但读到0。

```cgo
// 如果对端已经关闭，ok 返回 false, num 无数据
// 如果对端没有关闭，ok 返回 true, num 保存读到的数据
if num, ok := <-ch; ok == true {
    // 读取数据
} else {
    // 没有数据
}

// range代替ok
for num := range ch { // 不能替换为 <-ch

}
```


len: 未读数据的长度, 查看channel中元素的个数

### 单向channel

- 默认channel是双向:  `var ch chain int; ch := make(chan int)`
- 单向写channel: `var sendCh chan <- int; sendChd := make(chan <-int)` 不能读操作
- 单向读channel: `var recvCh <-chan int; recvCh := make(<-chan int)`

- 转换
  - 双向channel可以隐式转换为任意一种单向channel `sendCh = ch`
  - 单向channel 不能转换为双向channel `ch = sendCh/recvCh error!!!`
  
- 传参：传引用

## 生产者消费者模型

- 写入数据（生产者）----> 缓冲区（公共区）----> 读取数据（消费者）
  - 生产者：发送数据端
  - 消费者：接受数据端
  - 缓冲区
    - 1. 解耦（降低生产者和消费者之间的耦合度）
    - 2. 并发（生产者消费者数量不对等时，依然能保持正常通信）
    - 3. 缓存（生产者和消费者数据处理速度不一致时，暂存数据）

- 缓冲区：channel
  - 有缓冲：异步通信
  - 无缓冲：同步通信
  
- 模拟订单
  - 下单（生产者）
  - 处理订单（消费者）

## timer 定时器

```cgo
type Tier struct {
    C <-chan Time,
    r runtimeTimer
}
// 定时时间到达以前，没有数据写入timer.C会一直阻塞，直到定时时间到，系统会自动向 timer.C 这个 channel中写入当前时间，阻塞即被解除
```

1. 创建定时器 myTime := time.NewTimer(2*time.Second)
2. 停止：myTimer.Stop - 将定时器归零 <-myTimer.C 会阻塞
3. 重置：myTimer.Reset(time.Second)

### 周期定时器

1. 创建周期定时器 myTicker := time.NewTicker(time.Seconds)

- 定时时长到达后，系统会自动向 Ticker 的 C中写入系统当前时间
- 并且，每个一个定时时长后，循环写入系统当前时间

2. 在子协程循环读取C，获取系统写入的时间

## select

> 监听 channel 上的数据流动，读？写？
> 每个case 语句必须是一个IO操作

1. 监听case中，没有满足监听条件，阻塞
2. 监听饿case中，有多个满足监听条件，任选一个执行
3. 可以使用default来处理所有case都不满足监听条件的状况。通常不用（会产生忙轮询）
4. select 自身不带有循环机制，需借助外层 for 来循环监听
5. break 只能跳出 select, 类似于 switch 中用法

```cgo
select {
  case <-chan1:
    // 如果chan1成功读到数据，则进行该case处理语句
  case chan2<-1:
    // 如果成功向chan2写入数据，则进行该case 处理语句
  default: // 忙轮询，不推荐使用
    // 上面都没有成功，即上面都阻塞，则进入default 处理流程
}
```

### 死锁

> 错误使用锁的现象

1. 单协程自己死锁；channel至少2个以上的协程进行通信，否则死锁；
2. 协程间 channel 访问顺序导致死锁；使用channel一端读（写），要保证另一端写（读）擦坐，同时有机会执行，否则死锁
3. 多个协程，多 channel 交叉死锁；A go程，掌握M的同时，尝试拿N； B go程，掌握N的同时尝试拿M。 
4. 在go语言中，尽量不要将 互斥锁、读写锁 与 channel 混用 ——  隐性死锁

### 互斥锁：（互斥量）

A 、B go程 共同访问共享数据。 由于cpu调度随机，需要对 共享数据访问顺序加以限定（同步）。

创建 mutex（互斥锁），访问共享数据之前，加锁，访问结束，解锁。 在A go程加锁期间，B go程加锁会失败——阻塞。

直至 A go程 解说mutex，B 从阻塞处。恢复执行。


### 读写锁：

读时共享，写时独占。写锁优先级比读锁高。



### 条件变量

本身不是锁！！！ 但经常与锁结合使用！！

使用流程：

1. 创建 条件变量： var cond    sync.Cond
2. 指定条件变量用的 锁：  cond.L = new(sync.Mutex)
3. cond.L.Lock()	给公共区加锁（互斥量）
4. 判断是否到达 阻塞条件（缓冲区满/空）	—— for 循环判断

`for  len(ch) == cap(ch) {   cond.Wait()` —— 1) 阻塞 2) 解锁 3) 加锁

5. 访问公共区 —— 读、写数据、打印 
6. 解锁条件变量用的 锁  `cond.L.Unlock()`
7. 唤醒阻塞在条件变量上的 对端。 `signal()  Broadcast()`



## socket 编程

> 网络通信过程中，socket一定是成对儿出现

- client : server
- "" : listen(ip+port)
- dial : accept
- write : read
- read : write
- close: close

### TCP 并发服务器

1. 创建监听套接字 `listener := net.Listener("tcp", "SERVER_IP:PORT")`
2. `defer listener.Close()`
3. for 循环阻塞监听客户端连续事件 `conn, err := listener.Accept()`
4. 创建子协程对应每一个客户端进行数据通信 `go HandlerConnect()`
5. 实现 HandlerConnect

```cgo
Handler(conn net.Conn){
    defer conn.Close()
    获取成功连接的客户端Add conn.RemoteAddr()
    for循环读取客户读发送的数据 conn.Read(buf)
    处理数据
    回写转化后的数据
}
```

服务端判断关闭：Read读客户端数据返回0 - 对端关闭

### UDP

> 无连接的，不可靠的报文传递

#### UDP服务器

1. 创建 server端地址结构(IP+PORT) `net.ResolveUDPAddr()`
2. 创建用于通信的socket, 绑定地址结构 `conn = net.ListenUDP()`
3. `defer conn.Close()`
4. 读取客户端发送的数据 `n, clientAddr, err := conn.ReadFromUDP()`
5. 写数据给客户端 WriteToUDP("数据", clientAddr)

### 聊天室主模块

- 主协程：创建监听套接字，循环Accept()客户端连接 —— 启动子协程 HandlerConnect
- HandlerConnect: 创建用户结构体对象，存入onlineMap, 发送用户登录广播，聊天信息，处理查询在线用户、改名、下线和超时剔除。
- Manager: 监听全局 channel message，将读到的消息广播给onlineMap 中的所有用户
- WriteMsgToClient: 读取每个用户自带 channel C上消息（由Manager发送该消息），回写给用户
- 全局数据模块：
  - 用户结构体：Client{C,Name,Addr string}
  - 在线用户列表：onlineMap[string]Client key:客户端IP+PORT value:Client
  - 消息通道：message

### 广播用户在线

1. 主协程，创建监听套接字并defer
2. for 循环监听客户端连接请求
3. 有一个客户端连接请求，创建子协程处理客户端数据 HandlerConnect(conn)并defer
4. 定义全局结构体类型 C, Name, Addr
5. 创建全局 map, channel
6. 实现HandlerConnect函数，获取客户端IP+PORT —— RemoteAddr()；初始化新用户结构体信息，name==Addr
7. 创建 Mananger管理子协程—— Accept() 之间
8. 实现 Manager，初始化在用用户map，循环读取全局channel，如果无数据，阻塞，有数据，遍历在线用户map，将数据写到用户通道C中
9. 新用户添加到在线用户map中，key==IP+PORT value=新用户结构体
10. 创建 WriteMsgToClient 子协程里，专门给当前用户写数据 —— 来源于用户自带的 C 中
11. 创建 WriteMsgToClient(client, conn)，遍历自带的C，读数据，conn.Write 到客户端
12. HandlerConnect 中，结束位置，组织用户上线信息，将用户上线信息写到全局 channel —— Manager的读被激活（原来一直阻塞） 
13. HandlerConnect 中，结尾加 for {;}

### 广播用户消息

1. 疯长函数 MakeMsg() 处理广播、用户消息
2. HandlerConnect 中，创建匿名子协程，读取用户套接字上发送来的聊天内容。写到全局channel
3. 循环 conn.Read n == 0, err != nil
4. 写给全局 message —— 后续的事，原来广播用户上线模块完成。(Manager, WriteMsgToClient)

### 查询在线用户

1. 将读到的用户消息 msg 结尾 "\r\n" 或 "\n" 去掉
2. 判断是否是 "who"命令
3. 如果是，遍历在线用户列表，组织显示信息，写到socket中
4. 如果不是，写给全局 message

### 修改用户名

1. 将读取到用户消息 msg 判断是否包含 "rename"
2. 提取 "|" 后面的字符串，存入到Client的Name成员中
3. 更新在线用户列表。onlineMap. key —— IP+PORT
4. 提示用户更新完成。conn.Write()

### 用户退出

1. 在用户成功登录之后，创建监听用户退出的 channel - isQuit
2. 当conn.Read == 0, isQuit <- true
3. 在HandlerConnect 结尾中，添加 select 监听 isQuit
4. 条件满足，将用户从在线用户列表移除。组织用户下线消息，写入message(广播)

### 超时处理

1. 在select中监听定时器，time.After()，计时到达。将用户从在线列表移除，组织用户下线消息，写入message(广播)
2. 创建监听用户活跃的channel - isOnline
3. 只要用户执行：聊天、改名、who、任意一个操作，isOnline <- true
4. 在 select 中添加监听 <-isOnline，条件满足，不做任何事情。目的是重置计数器 

## http

http 请求报文格式
```
GET / HTTP/1.1 —— 请求行：请求方法+空格+URL+空格+协议版本+回车符(\r)+换行符(\n)
Host: 127.0.0.1:8900 —— 请求头部：头部字段名称+冒号(:)+空格+值+回车符(\r)+换行符(\n)
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36
Sec-Fetch-User: ?1
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9,ko;q=0.8,en;q=0.7
                —— 回车符(\r)+换行符(\n)，表示请求头结束
请求报体：请求方法对应
```

回调函数：函数指针，通过地址，在某一特定位置，调用函数。

在程序中，定义一个函数，但不显示调用。当某一条件满足时，该函数有操作系统系统自动调用。

### HTTP 请求头部



## 算法和数据结构

### 排序

- 排序：将一组数据，依指定的顺序进行排列的过程
- 排序的分类
  - 内部排序：将所有的数据加载到内部存储器中进行排序（**交换式排序法、选择式排序法、插入式排序法**）
    - 交换式排序：冒泡排序(Bubble sort)、快速排序(Quick sort)
  - 外部排序：数据量过大，无法全部加载到内存中，需要借助外部存储进行排序（**合并排序法和直接合并排序法**）

### 查找

- 顺序查找
- 二分查找（必须是有序数组）

## select

```go
select {
  case <-chan1: // chan1成功读到数据，则进行该case处理语句
  case chan2<-1: // 成功向chan2写入数据，则进行该case处理语句
  default: // 上面没有成功，则进入default处理流程
}
```

## beego

```sh
$ go get -u github.com/astaxie/beego

bee工具
$ go get -u github.com/beego/bee
```

安装之后，bee 可执行文件默认存放在$GOPATH/bin里面，所以需要把`$GOPATH/bin`添加到环境变量中
```
$ echo 'export PATH="$PATH/bin:$PATH"' >> ~/.bashrc
$ source .bashrc
```

- `bee new 项目名` 创建一个项目；通过 bee 创建的项目代码都是在`$GOPATH/src`目录下面
- `bee api`


dyld: malformed mach-o image: segment __DWARF has vmsize < filesize

升级到1.13.4后问题被修复

brew升级golang: `brew upgrade go`



## 函数式编程

- 不可变性：不能有状态，只有常量和函数
- 函数只能有一个参数