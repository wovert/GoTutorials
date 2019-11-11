package main

import (
	"fmt"
	"time"
	"./wovert"
	"./real"
)


type Retriever interface {
	Get(url string) string
}


func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
  var r Retriever
  r = wovert.Retriever{"this is wovert.com"}
  inspect(r)
  wovetRetriever := r.(wovert.Retriever)
  fmt.Println(wovetRetriever.Contents)

//   fmt.Println(download(r))

  r = &real.Retriever{
	UserAgent: "Mozillar/5.0",
	TimeOut: time.Minute,
  }
  inspect(r)

  // Type assertion
  realRetriever := r.(*real.Retriever)
  fmt.Println(realRetriever.TimeOut)
//   fmt.Println(download(r))

  
}

func inspect(r Retriever) {
  fmt.Printf("%T %v\n", r, r)
  switch v := r.(type) {
	case wovert.Retriever:
	  fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
	  fmt.Println("UserAgent:", v.UserAgent)
  }
}