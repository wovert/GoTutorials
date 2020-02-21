```sh
#编译myproto.proto文件
$ protoc --go_out=./ *.proto #不加grpc插件
$ protoc --go_out=plugins=grpc:./ *.proto #添加grpc插件
#对比发现内容增加
#得到 myproto.pb.go文件
```