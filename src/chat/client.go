package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 服务器IP地址和端口，创建通信套接字
	conn, err := net.Dial("tcp", "192.168.3.12:8989")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	// 获取用户键盘输入（stdin），将输入数据发送给服务器
	go func(){
		for {
			str := make([]byte, 4096)
			n, err := os.Stdin.Read(str)

			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}

			//写给服务器，读多少，写多少
			//conn.Write(str[:n])

			//编码处理
			data, err := Encode(string(str[:n]))

			if err != nil {
				fmt.Println("encode msg failed, err:", err)
				return
			}

			conn.Write(data)
		}

		// 测试粘包和半包
		//for i := 0; i < 20; i++ {
		//	msg := `Hello, Hello. How are you?`
		//	data, err := Encode(msg)
		//	if err != nil {
		//		fmt.Println("encode msg failed, err:", err)
		//		return
		//	}
		//	conn.Write(data)
		//}
	}()

	// 回显服务器回发的大些数据
	//buf := make([]byte, 4096)

	//n, err := conn.Read(buf)
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF { //io.EOF在网络编程中表示对端把链接关闭了
			fmt.Println("检测到服务器已经关闭")
			return
		}
		//fmt.Println("err != nil:", err != nil)
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("server发来的数据：", msg)
	}
	for {
		;
	}
}

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
