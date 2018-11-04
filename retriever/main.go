package main

import (
	"fmt"
	"github.com/kingwangboss/golangpro/retriever/mock"
	real2 "github.com/kingwangboss/golangpro/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post("http://www.baiud.com", map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get("http://www.baidu.com")
}

func main() {
	fmt.Println(download(&mock.Retriever{Contents: "this is fake www.baidu.com"}))
	var r Retriever
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	fmt.Println(session(&mock.Retriever{Contents: "this is fake www.baidu.com"}))

}
