package main

import (
	"github.com/micro/go-micro/v2/web"
	"net/http"
)

func main() {
	server := web.NewService(web.Address(":8080"))
	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!"))
	})
	server.Run()

}
