package main

import (
	"fmt"
	"net/http"

	"CampingNow/pkg/setting"
	"CampingNow/routers"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

// main v2 has implemented go Hot Restart
// 很好，因为 win 系统缺少变量，endless 不能在 win 上使用，有解决方案，但是解决方案后果不清楚，保留
//func main() {
//	endless.DefaultReadTimeOut = setting.ReadTimeout
//	endless.DefaultWriteTimeOut = setting.WriteTimeout
//	endless.DefaultMaxHeaderBytes = 1 << 20
//	endPoint := fmt.Sprintf(".%d", setting.HTTPPort)
//
//	server := endless.NewServer(endPoint, routers.InitRouter())
//	server.BeforeBegin = func(add string) {
//		log.Printf("Actual pid is %d", syscall.Getpid())
//	}
//
//	err := server.ListenAndServe()
//	if err != nil {
//		log.Printf("Server err: %v", err)
//	}
//
//}
