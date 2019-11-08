package main

import (
	"fmt"
	// "github.com/astaxie/goredis"
	"github.com/go-redis/redis"
)

func main() {
  var redisClient *redis.Client // redis 客户端连接
	// todo redis服务器连接
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "passwd",
		DB:       2,
	})

	// todo 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	// sta, err := redisClient.Set("name", "hello world", 0).Result()
	// if err != nil {
	//   panic(err)
	// }

	val, err := redisClient.Get("name").Result()
	if err != nil {
	  panic(err)
	}

	fmt.Printf("%v\n", val)


  //   var client goredis.Client
//   client.Addr = "47.95.192.23:6379"
//   client.Password = "$sbdnl$@)!("
//   client.Db = 2

//   vals := []string{"a", "b", "c", "d", "e"}
//   for _, v := range vals {
// 	client.Rpush("l", []byte(v))
//   }
//   dbvals,_ := client.Lrange("l", 0, 4)
//   for i, v := range dbvals {
// 	fmt.Println(i,":",string(v))
//   }


//   err := client.Set("abc", []byte("hello world"))
//   if err != nil {
// 	fmt.Printf("失败：%v\n", err)
// 	// panic(err2)
//   } 
//   val, err2 := client.Get("abc")
//   if err2 != nil {
// 	fmt.Printf("%v\n", err2)
// 	// panic(err2)
//   }
//   fmt.Println(string(val))
//   client.Del("a")
}