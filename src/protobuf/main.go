package main

import (
	"../prototext"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	text := &prototext.PandaRequest {
		Name: "熊猫",
		Height: 130,
		Weight: []int32{120, 130, 140, 160},
	}

	fmt.Println(text)

	// proto编码
	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Println("编码失败")
	}
	fmt.Println("编码后打印：", data)

	// 解码
	newtext := &prototext.PandaRequest{}
	err = proto.Unmarshal(data, newtext)
	if err != nil {
		fmt.Println("解码失败")
	}
	fmt.Println("解码后打印：", newtext)

	fmt.Println(newtext.String())

	fmt.Println("名字：", newtext.Name)
	fmt.Println("身高：", newtext.Height)
	fmt.Println("体重：", newtext.Weight)
}