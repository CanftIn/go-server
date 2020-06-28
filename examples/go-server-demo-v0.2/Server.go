package main

import "CanftIn/go-server/znet"

func main() {
	// 创建一个server句柄，使用zinx api
	s := znet.NewServer("[go-server V0.2]")

	// 启动server
	s.Serve()
}
