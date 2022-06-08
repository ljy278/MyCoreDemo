package main

import (
	"net/http"
	"coredemo/framework"
)

func main() {
	server := &http.Server{
		// 自定义的请求处理函数
		Handler: framework.NewCore(),
		// 监听地址
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
