package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://daily.zhihu.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	defer resp.Body.Close()
}
