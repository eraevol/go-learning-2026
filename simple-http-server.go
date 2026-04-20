package main

import (
	"net/http"
	"time"
)

// healthHandle健康检查接口，返回ok表示服务正常
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// timeHandler返回当前服务器时间
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().String()
	w.Write([]byte(currentTime))
}

func main() {
	//注册路由
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	//启动服务，监听8080窗口
	http.ListenAndServe(":8080", nil)
}
