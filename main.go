package main

import (
	"log"
	"net/http"
)

func main() {
	//程序入口，一个项目只能有一个入口
	//Web程序，Http协议 ip prot
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
