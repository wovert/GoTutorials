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

type Poster interface {
  Post(url string, form map[string]string)
}


func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster) {
  poster.Post("http://www.baidu.com", 
map[string]string {
  "name": "wovert",
  "course": "golang",
})
}

type RetrieverPoster interface {
  Retriever
  Poster
  // Connect(host string)
}

const url = "http://www.baiud.com"
func session(s RetrieverPoster) string {
  // s.Get()
  s.Post(url, map[string]string{
    "contents": "another faked wovert",
  })
  return s.Get(url)
}

func main() {
  var r Retriever
  r = wovert.Retriever{"this is wovert.com"}
  inspect(r)
  wovetRetriever := r.(wovert.Retriever)
  fmt.Println(wovetRetriever.Contents)

  fmt.Println("Try a session")
  fmt.Println(session(r))
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