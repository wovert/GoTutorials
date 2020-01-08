package main

import (
	"../proto"
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 匹配状态
const (
	STATUS_STOPING 		= 1;		// 停止匹配
	STATUS_FINDING 		= 2;		// 寻找匹配
	STATUS_CHATING 		= 3;		// 聊天中
)

// 性别
const (
	SEX_FEMALE 	= 0		// 女人
	SEX_MALE 		= 1		// 男人
)

// 匹配类型
const (
	TYPE_RANDOM 	= 1		// 随机
	TYPE_PAYMENT 	= 2		// 付费
)

// 付费类型
const (
	CONN_SEX = 1 		// 性别连接
	CONN_DIS = 2 		// 距离连接
)

// 匹配顺序
const (
	ORDER_FIRST 	= 0 		// 先入
	ORDER_SECOND 	= 1 		// 后入
)

// 消息
const (
	MSG_MATCH = "match"
	MSG_WHO 	= "who"
	MSG_ROOMS = "rooms"
)

// 客户端信息
type Client struct {
	C chan string			// 用户发送数据的管道
	Name string				// 用户名
	Addr string				// 网络地址
	Uid int64					// 用户编号
	Sex int8					// 性别
	Type int8					// 匹配类型
	Status int8				// 匹配状态
	Pay int8					// 付费类型
	Room int64				// 聊天室名称(纳秒)
	Order int8				// 匹配顺序
}

// 聊天室列表: roomName: [Addr2, Addr1]
var roomList map[int64][]string

// 在线用户
var onlineMap map[string]Client

// 用户发送消息管道
var message = make(chan string)

func makeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + ": " + msg
	return
}

func maneger() {
	// 给 聊天室列表 分配空间
	roomList = make(map[int64][]string)

	// 给 map 分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message // 没有消息前，这里会阻塞

		// 遍历onlineMap，给每个成员发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

// 发送信息给客户端
func writeMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		// 给当前客户端发送信息
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
		//conn.Write([]byte(msg + "\n"))
	}
}

// 读写锁
var rwlocker sync.RWMutex

// 查找在线用户是否与当前用户匹配
func searchMatch(key string)  {
	client := onlineMap[key]

	// 遍历onlineMap
	for _, cli := range onlineMap {
		// 是否是自己
		if (cli.Addr == client.Addr) {
			fmt.Println("匹配到自己跳过")
			continue
		}

		// 查找到在线用户，锁定找到的用户
		rwlocker.Lock()

		//fmt.Println("匹配状态：",cli.Status, cli.Status == STATUS_FINDING)
		//fmt.Println("性别：", cli.Sex == client.Sex)

		if (cli.Status == STATUS_FINDING && cli.Sex == client.Sex) {
			// 获取聊天室名称，纳秒时间戳
			room := time.Now().UnixNano()
			fmt.Println("创建聊天室名称：", room)
			//fmt.Println("主动匹配地址：", cli.Addr)
			//fmt.Println("被动匹配地址：", client.Addr)

			// 初始化聊天室
			roomList[room] = []string{cli.Addr, client.Addr}

			cli.Room = room
			cli.Order = ORDER_SECOND // 相反值，从聊天室索引好查找对方
			cli.Status = STATUS_CHATING

			client.Room = room
			client.Order = ORDER_FIRST
			client.Status = STATUS_CHATING

			// 向双方发送匹配信息
			cli.C <- makeMsg(cli, "与" + strconv.FormatInt(client.Uid, 10) + "匹配成功")
			client.C <- makeMsg(cli, "与" + strconv.FormatInt(cli.Uid, 10) + "匹配成功")

			// 广播匹配成功的用户信息
			broadMsg := strconv.FormatInt(cli.Uid, 10) + "与" + strconv.FormatInt(client.Uid, 10) + "聊天中"
			message <- makeMsg(cli, broadMsg)

			// 更新用户状态
			onlineMap[cli.Addr] = cli
			onlineMap[key] = client

			rwlocker.Unlock()
			break
		}
		rwlocker.Unlock()
	}
	message <- makeMsg(client, "广播：" + strconv.FormatInt(client.Uid, 10) + "等待匹配中")
	//fmt.Println(client.Addr + "等待匹配中")
}

// 用户处理连接
func handleConn(conn net.Conn) {
	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	// 创建一个结构体，默认用户名和网络地址一样
	cli := Client{
		C: make(chan string),
		Name: cliAddr,
		Addr: cliAddr,
		Status: STATUS_STOPING,
	}

	// 把结构体添加到 map
	onlineMap[cliAddr] = cli

	// 新开一个协程，专门给当前客户端发送信息
	go writeMsgToClient(cli, conn)

	// 广播某个人在线
	//message <- makeMsg(cli, "login")

	// 提示我是谁
	cli.C <- makeMsg(cli, "上线")

	isQuit := make(chan bool) // 对方是否主动退出
	hasData := make(chan bool) // 对方是否数据发送

	// 新建一个协程，接受用户发送过来的数据
	go func(){
		defer conn.Close()

		reader := bufio.NewReader(conn)

		for {
			msg, err := proto.Decode(reader)
			if err == io.EOF { // io.EOF 在网络编程中表示对端把链接关闭了
				isQuit <- true
				fmt.Println("对端把链接关闭")
				return
			}
			if err != nil {
				fmt.Println("decode msg failed, err:", err)
				isQuit <- true
				return
			}
			if len(msg) == 1 {
				fmt.Println("客户端传递空值")
				continue
				//return
			}
			msg = msg[:len(msg)-1]
			//fmt.Println([]byte(msg[:len(msg)-1]))

			fmt.Println("[" , conn.RemoteAddr().String() , "]发送的数据：", msg)

			// todo 根据msg对逻辑进行处理
			switch msg {
				// 匹配连接
				case MSG_MATCH:
					cli = onlineMap[cliAddr]
					message <- makeMsg(cli, strconv.FormatInt(cli.Uid, 10) + "的匹配状态" + strconv.Itoa(int(cli.Status)))
					// 是否已经连接匹配中
					if (cli.Status == STATUS_STOPING) {
						go match(cliAddr)
						msg := strconv.FormatInt(cli.Uid, 10) + "=>寻找匹配中"
						message <- makeMsg(cli, msg)

					// 正在聊天中
					} else if cli.Status == STATUS_CHATING {

						message <- makeMsg(cli, strconv.FormatInt(cli.Uid, 10) + "=>重新连接")

						// 退出聊天室并建立新的聊天
						go exitRoom(cliAddr, func(){
							message <- makeMsg(cli, strconv.FormatInt(cli.Uid, 10) + "=>重新匹配")
							match(cliAddr)
							msg := strconv.FormatInt(cli.Uid, 10) + "=>寻找匹配中"
							message <- makeMsg(cli, msg)
						})
					}
				// 在线用户列表
				case MSG_WHO:
					//rwlocker.Lock()
					for _, cli := range onlineMap {
						msg := "[" + strconv.FormatInt(cli.Uid,10) + "]" + ", 状态:" +
							strconv.Itoa(int(cli.Status)) + ", 房号：" +
							strconv.Itoa(int(cli.Room)) + ", 性别:" +
							strconv.Itoa(int(cli.Sex)) + ", 对方索引：" + strconv.Itoa(int(cli.Order))
						fmt.Println(msg)
						message <- makeMsg(cli, msg)
					}
					//rwlocker.Unlock()
				case MSG_ROOMS:
					for _, room := range roomList {
						msg := strconv.FormatInt(onlineMap[room[0]].Uid, 10) + "与" +
							strconv.FormatInt(onlineMap[room[1]].Uid, 10) + "聊天中"
						fmt.Println(msg)
						message <- makeMsg(cli, msg)
					}
			}
			//fmt.Println("login|1长度：", len(msg))
			if len(msg) >= 7 && msg[:5] == "login" {
				// rename|1
				fmt.Println(msg, strings.Contains(msg, "="))
				if !strings.Contains(msg, "=") {
					fmt.Println("不包含用户UID")
				} else {
					uid := strings.Split(msg, "=")[1]
					id, err := strconv.ParseInt(uid, 10, 64)
					if err != nil {
						fmt.Println("登录编号错误:", id)
					} else {
						cli.Uid = id
						onlineMap[cliAddr] = cli
						message <- makeMsg(cli, strconv.FormatInt(id, 10) + "=>登录成功")
					}
				}
			}

			// 断开匹配

			// 转发此内容
			//message <- makeMsg(cli, msg)
			hasData <- true // 有数据
		}
	}()

	for {
		// 通过select检测channel的流动
		select {
		case <-isQuit:
			//delete(onlineMap, cliAddr) // 当前用户从 map 移除
			//message <- makeMsg(cli, "logout") // 告诉所有在线用户，谁退出了
			return
		case <- hasData:
		case <- time.After(30 * time.Second): // 30s 后退出聊天室
			go exitRoom(cliAddr, func(){
				exitProcess(cliAddr)
			})
			return
		}
	}
}

// 聊天入口
func main () {
	listener, err := net.Listen("tcp", "192.168.3.12:8989")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	// 新开一个协程，转发消息，只要有消息来了，遍历onlineMap, 给map每个成员发送消息
	go maneger()

	// 主协程, 循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}
		go handleConn(conn) // 处理用户连接
	}
}

/**
 * 匹配用户
 */
func match(key string) {
	cli := onlineMap[key]
	cli.Status = STATUS_FINDING
	cli.Type = TYPE_RANDOM
	cli.Sex = SEX_MALE 			// 这里需要根据当前客户端用户查找响应性别数据
	onlineMap[key] = cli
	fmt.Println("cli.Status=", cli.Status)

	go searchMatch(key) // 搜索匹配

}

/**
 * 退出聊天室
 */
func exitRoom(cliAddr string, callback func()) {

	// 当前用户退出聊天室
	cli := onlineMap[cliAddr]

	room := cli.Room 			// 房间号
	index := cli.Order 		// 对方索引号

	cli.Status = STATUS_STOPING
	cli.Room = 0
	cli.Order = 0
	// 更新当前用户信息
	onlineMap[cliAddr] = cli

	//fmt.Println("房间号：", room)
	//fmt.Println("对方索引号：", index)
	//fmt.Println("几个聊天室", len(roomList))

	cli.C <- makeMsg(cli, strconv.FormatInt(cli.Uid, 10) + "退出聊天室")

	// 房间号为空，说明对方已经退出聊天室
	if room == 0 {
		fmt.Println("聊天室已经移除")
		go callback()
		//message <- makeMsg(cli, "已经退出聊天室")
		return
	}

	// 更新另一个用户信息
	if len(roomList[room]) == 2 {
		// 聊天对象另一个用户信息
		key := roomList[room][index]
		matchUser := onlineMap[key]
		matchUser.Status = STATUS_STOPING
		matchUser.Room = 0

		// 更新聊天对象用户信息
		onlineMap[key] = matchUser

		// 移除当前聊天室
		delete(roomList, room)
		fmt.Println("移除聊天室：", room)

		// 告诉所有在线用户，谁退出了
		message <- makeMsg(cli, "广播：" + strconv.FormatInt(cli.Uid, 10) + "退出聊天室")		// 自己退出
		message <- makeMsg(matchUser, "广播：" + strconv.FormatInt(matchUser.Uid, 10) + "退出聊天室") // 对方退出

		//cli.C <- makeMsg(cli, "=======聊天对象:" + strconv.FormatInt(matchUser.Uid, 10) + "退出=======")
		//matchUser.C <- makeMsg(matchUser, "=======聊天对象:" + strconv.FormatInt(cli.Uid, 10) + "退出=======")

	}
	callback()
}

// 退出连接
func exitProcess(cliAddr string) {
	fmt.Println("客户端【", cliAddr, "]退出")
	// 当前过期用户
	cli := onlineMap[cliAddr]

	// 告诉所有在线用户，谁退出了
	message <- makeMsg(cli, strconv.FormatInt(cli.Uid, 10) + "关闭连接")

	fmt.Println("删除在线用户：", strconv.FormatInt(cli.Uid, 10))
	// 当前用户从map移除
	delete(onlineMap, cliAddr)

}