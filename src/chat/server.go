package main

import (
	"../proto"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)


const (
	// 匹配状态
	STATUS_STOPING 		= 1;		// 停止匹配
	STATUS_FINDING 		= 2;		// 寻找匹配
	STATUS_CHATING 		= 3;		// 聊天中

	// 性别
	SEX_FEMALE 	= 0		// 女人
	SEX_MALE 		= 1		// 男人

	// 匹配类型
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
	MSG_LOGOUT = "logout"
)

// 命令类型
const (
	CMD_HB = 100								// 发送心跳
	CMD_HB_ACK_FAIL = 101 			// 心跳失败
	CMD_BIND = 200          		// 请求绑定
	CMD_BIND_OK = 201       		// 绑定成功
	CMD_BIND_FAIL = 202 				// 绑定失败
	CMD_TOKEN_FAIL = 300    		// token验证失败
	CMD_SSO = 301								// 单点登录
	CMD_MSG = 400								// 发送消息
	CMD_MSG_OK = 401 						// 发送消息确认成功
	CMD_MSG_READ = 402 					// 读取消息
	CMD_MSG_RETRY = 403 				// 重新发送消息
	CMD_MATCH = 500 						// 匿名匹配
	CMD_MATCH_OK = 501					// 匹配成功
	CMD_MATCH_FAIL = 502 				// 匹配失败
	CMD_MATCH_EXIT = 503 				// 退出匹配
	CMD_SYS_NOTICE = 600 				// 系统通知
	CMD_SYS_NOTICE_OK = 601 		// 系统通知接受成功
	CMD_SERVICE_NOTICE = 700		// 服务通知
	CMD_SERVICE_NOTICE_OK = 701 // 服务通知接受成功
	CMD_MSG_SCENE = 800 				// 场景消息
)

// 聊天类型
const (
	CHAT_PRIVATE = 1 				// 私聊
	CHAT_MATCH = 2 					// 匹配聊
)

// 内容类型
const (
	MEDIA_TYPE_TEXT = 1					// 文本类型
	MEDIA_TYPE_NEWS	= 2		  		// 图文类型
	MEDIA_TYPE_VOICE = 3		  	// 语音类型
	MEDIA_TYPE_IMAGE = 4		  	// 图片类型
	MEDIA_TYPE_REDPACKET = 5		// 红包类型
	MEDIA_TYPE_EMOJI	= 6				// 表情类型
	MEDIA_TYPE_LINK = 7					// 超链接类型
	MEDIA_TYPE_VIDEO = 8				// 视频类型
	MEDIA_TYPE_CONCAT = 9				// 名片类型
)

// 客户端信息
type Client struct {
	C chan string			// 用户发送数据的管道
	Name string				// 用户名
	Addr string				// 网络地址
	Key string				// 用户编号=md5(key+uid)
	Sex int8					// 性别
	Type int8					// 匹配类型
	Status int8				// 匹配状态
	Pay int8					// 付费类型
	Room int64				// 聊天室名称(纳秒)
	Order int8				// 匹配顺序
}


type Emoticon struct {
	Emoji		string `json:"emoj,omitempty" form:"emoj"` 				// 表情名
}

type RedPacket struct {
	Amount	float64 `json:"amount,omitempty" form:"amount"` 	// 金额
}

type ImageText struct {
	Title		string `json:"title,omitempty" form:"title"` 			// 图文标题
	Memo		string `json:"memo,omitempty" form:"memo"` 				// 图文描述
}

type Card struct {
	Avatar		string `json:"avatar,omitempty" form:"avatar"` 	// 头像
	Name 			string `json:"name,omitempty" form:"name"` 			// 姓名
	Key 			string `json:"key,omitempty" form:"key"` 				// 用户编号=md5(key+id)
}

type Data struct {
	Text		string `json:"text,omitempty" form:"text"` 				// 文本
	Preview string `json:"preview,omitempty" form:"preview"` 	// 预览(图片|视频)
	Url			string `json:"url,omitempty" form:"url"` 					// 服务的URL(图片|语音|视频|链接地址)
	Time    int    `json:"time,omitempty" form:"time"` 				// 语音时间
	Emoticon
	RedPacket
	ImageText
	Card
}

type Message struct {
	Id			int64  `json:"id,omitempty" form:"id"` 						// 消息ID
	Cmd     int    `json:"cmd,omitempty" form:"cmd"` 					// 命令类型
	From  	string `json:"from,omitempty" form:"from"` 				// 谁发的
	To   		string `json:"to,omitempty" form:"to"`						// 对端用户ID/群ID
	Chat   	int    `json:"chat,omitempty" form:"chat"` 				// 聊天类型(1:私聊, 2:匹配聊)
	Token   string `json:"token,omitempty" form:"token"`			// 用户登录令牌
	Os			string `json:"os,omitempty" form:"os"`						// 操作系统
	Media   int    `json:"media,omitempty" form:"media"` 			// 消息类型
	Data		Data
}

// 聊天室列表: roomName: [Addr2, Addr1]
var roomList map[int64][]string

// 在线用户列表
var onlineMap map[string]Client

// 用户发送消息管道
var message = make(chan string)

// 创建信息
func makeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + ": " + msg
	return
}

/**
 * 初始化在线用户，聊天室列表和给每个用户发送消息
 */
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

	matchFlag := false

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

			// 初始化聊天室
			roomList[room] = []string{cli.Addr, client.Addr}

			// 更新主动匹配用户信息
			cli.Room = room
			cli.Order = ORDER_SECOND // 相反值，从聊天室索引好查找对方
			cli.Status = STATUS_CHATING

			// 更新被动匹配用户信息
			client.Room = room
			client.Order = ORDER_FIRST
			client.Status = STATUS_CHATING

			// 向双方发送匹配信息
			cli.C <- makeMsg(cli, "与" + client.Key + "匹配成功")
			client.C <- makeMsg(cli, "与" + cli.Key + "匹配成功")
			//cli.C <- makeMsg(cli, "与" + strconv.FormatInt(client.Key, 10) + "匹配成功")
			//client.C <- makeMsg(cli, "与" + strconv.FormatInt(cli.Key, 10) + "匹配成功")

			// 广播匹配成功的用户信息
			//broadMsg := strconv.FormatInt(cli.Key, 10) + "与" + strconv.FormatInt(client.Key, 10) + "聊天中"
			broadMsg := cli.Key + "与" + client.Key + "聊天中"
			message <- makeMsg(cli, broadMsg)

			// 更新用户状态
			onlineMap[cli.Addr] = cli
			onlineMap[key] = client

			rwlocker.Unlock()
			matchFlag = true
			break
		}
		rwlocker.Unlock()
	}
	if !matchFlag {
		message <- makeMsg(client, "广播：" + client.Key + "等待匹配中")
	}
	//fmt.Println(client.Addr + "等待匹配中")
}

// 用户处理连接
func handleConn(conn net.Conn) {
	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	// 创建一个结构体，默认用户名和网络地址一样
	cli := Client {
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
	cli.C <- makeMsg(cli, "上线\n")

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

			fmt.Println("[" , conn.RemoteAddr().String() , "]发送的消息：", msg)
			//msg = msg[:len(msg)-1]
			//fmt.Println([]byte(msg[:len(msg)-1]))
			//fmt.Println("[" , conn.RemoteAddr().String() , "]发送的数据：", msg)

			// 对msg进行解码

			// todo 解析msg为message
			msgMap := Message{}
			err = json.Unmarshal([]byte(msg), &msgMap)
			if err != nil {
				log.Println("解析JSON错误：", err.Error())
				return
			}

			// 接受命令格式，媒介类型
			fmt.Println(msgMap.Cmd, msgMap.Media)

			cli = onlineMap[cliAddr]

			// 验证发送用户
			// .........


			// 命令类型
			switch msgMap.Cmd {
				case CMD_MSG: // 发送消息 400

				  // 没有绑定前不能发送消息
					if cli.Key  == "" {
						message <- makeMsg(cli, "请求绑定之前不能发送信息")
						break
					}

					switch msgMap.Media {
						case MEDIA_TYPE_TEXT: 			// 文本消息 1
							message <- makeMsg(cli, "发送文本消息=>" + msgMap.Data.Text)

							content := msgMap.Data.Text
							var destCli Client

							switch msgMap.Chat {
								case CHAT_MATCH: // 匹配聊
									room := cli.Room
									if string(room) != msgMap.To {
										message <- makeMsg(cli, "聊天室名不正确")
										//break
									}
									index := cli.Order
									destCli = onlineMap[roomList[room][index]]
									destCli.C <- makeMsg(destCli, cli.Key + "发送匹配聊：" + content)

								case CHAT_PRIVATE: // 私聊
									destCli = onlineMap[msgMap.To]
									destCli.C <- makeMsg(destCli, cli.Key + "发送私聊：" + content)
							}

						case MEDIA_TYPE_NEWS: 			// 图文消息 2
							message <- makeMsg(cli, "发送图文消息=>" + msgMap.Data.Title)

						case MEDIA_TYPE_VOICE: 			// 语音消息 3
							message <- makeMsg(cli, "发送语音消息=>" + msgMap.Data.Url)

						case MEDIA_TYPE_IMAGE: 			// 图片消息
							message <- makeMsg(cli, "发送图片消息=>" + msgMap.Data.Preview)

						case MEDIA_TYPE_REDPACKET: 	// 红包消息
							message <- makeMsg(cli, "发送红包消息=>" + strconv.FormatFloat(msgMap.Data.Amount, 'E', -1, 64))

						case MEDIA_TYPE_EMOJI: 			// 表情消息
							message <- makeMsg(cli, "发送表情消息=>" + msgMap.Data.Emoji)

						case MEDIA_TYPE_LINK: 			// 链接消息
							message <- makeMsg(cli, "发送链接消息=>" + msgMap.Data.Url)

						case MEDIA_TYPE_VIDEO:
							message <- makeMsg(cli, "发送视频消息=>" + msgMap.Data.Preview)

						case MEDIA_TYPE_CONCAT:
							message <- makeMsg(cli, "发送名片消息=>" + msgMap.Data.Name)

					}
				case CMD_BIND: // 绑定请求 200
					cli.Key = msgMap.From
					onlineMap[cliAddr] = cli
					message <- makeMsg(cli, "登录成功:" + msgMap.From)

				case CMD_MATCH: // 灵魂匹配 500
					cli = onlineMap[cliAddr]
					message <- makeMsg(cli, cli.Key + "的匹配状态" + strconv.Itoa(int(cli.Status)))

					// 是否已经连接匹配中
					if (cli.Status == STATUS_STOPING) {
						go match(cliAddr)
						msg := cli.Key + "=>寻找匹配中"
						message <- makeMsg(cli, msg)

						// 正在聊天中
					} else if cli.Status == STATUS_CHATING {

						message <- makeMsg(cli, cli.Key + "=>重新连接")

						// 退出聊天室并建立新的聊天
						go exitRoom(cliAddr, func(){
							message <- makeMsg(cli, cli.Key + "=>重新匹配")
							match(cliAddr)
							msg := cli.Key + "=>寻找匹配中"
							message <- makeMsg(cli, msg)
						})
					}
			}


			// todo 根据msg对逻辑进行处理
			switch msg {
				// 匹配连接
				case MSG_MATCH:
					cli = onlineMap[cliAddr]
					message <- makeMsg(cli, cli.Key + "的匹配状态" + strconv.Itoa(int(cli.Status)))
					// 是否已经连接匹配中
					if (cli.Status == STATUS_STOPING) {
						go match(cliAddr)
						msg := cli.Key + "=>寻找匹配中"
						message <- makeMsg(cli, msg)

						// 正在聊天中
					} else if cli.Status == STATUS_CHATING {

						message <- makeMsg(cli, cli.Key + "=>重新连接")

						// 退出聊天室并建立新的聊天
						go exitRoom(cliAddr, func(){
							message <- makeMsg(cli, cli.Key + "=>重新匹配")
							match(cliAddr)
							msg := cli.Key + "=>寻找匹配中"
							message <- makeMsg(cli, msg)
						})
					}

				// 在线用户列表
				case MSG_WHO:
					//rwlocker.Lock()
					for _, cli := range onlineMap {
						msg := "[" + cli.Key + "]" + ", 状态:" +
							strconv.Itoa(int(cli.Status)) + ", 房号：" +
							strconv.Itoa(int(cli.Room)) + ", 性别:" +
							strconv.Itoa(int(cli.Sex)) + ", 对方索引：" + strconv.Itoa(int(cli.Order))
						fmt.Println(msg)
						message <- makeMsg(cli, msg)
					}
					//rwlocker.Unlock()

				// 显示所有聊天室
				case MSG_ROOMS:
					for _, room := range roomList {
						msg := onlineMap[room[0]].Key + "与" + onlineMap[room[1]].Key + "聊天中"
						fmt.Println(msg)
						message <- makeMsg(cli, msg)
					}

				// 退出聊天室
				case MSG_LOGOUT:
					go exitRoom(cliAddr, func() {
						msg := cli.Key + "=>手动退出聊天室"
						message <- makeMsg(cli, msg)
					})
				}
				//fmt.Println("login|1长度：", len(msg))

				// 登录操作
				if len(msg) >= 7 && msg[:5] == "login" {
					// rename|1
					fmt.Println(msg, strings.Contains(msg, "="))
					if !strings.Contains(msg, "=") {
						fmt.Println("不包含用户UID")
					} else {
						uid := strings.Split(msg, "=")[1]
						cli.Key = uid
						onlineMap[cliAddr] = cli
						message <- makeMsg(cli, uid + "=>登录成功")
					}
				}

				// 私聊发送信息
				if len(msg) >= 6 && msg[:4] == "send" {
					cli = onlineMap[cliAddr]
					// send=1
					//fmt.Println(msg, strings.Contains(msg, "="))

					if !strings.Contains(msg, "=") {
						fmt.Println("不包含用户发送内容")
					} else {
						content := strings.Split(msg, "=")[1]
						room := cli.Room
						index := cli.Order
						destCli := onlineMap[roomList[room][index]]
						destCli.C <- makeMsg(destCli, cli.Key + "发送：" + content)
					}
				}

			// 转发此内容
			//message <- makeMsg(cli, msg)
			hasData <- true // 有数据
		}
	}()

	for {
		// 通过select检测channel的流动
		select {
		case <-isQuit:
			go exitRoom(cliAddr, func(){
				exitProcess(cliAddr)
			})
			return
		case <- hasData:
		case <- time.After(300 * time.Second): // 30s 后退出聊天室
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

	cli.C <- makeMsg(cli, cli.Key + "退出聊天室")

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
		message <- makeMsg(cli, "广播：" + cli.Key + "退出聊天室")		// 自己退出
		message <- makeMsg(matchUser, "广播：" + matchUser.Key + "退出聊天室") // 对方退出

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
	message <- makeMsg(cli, cli.Key + "关闭连接")

	fmt.Println("删除在线用户：", cli.Key)
	// 当前用户从map移除
	delete(onlineMap, cliAddr)

}