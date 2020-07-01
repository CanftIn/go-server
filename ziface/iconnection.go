package ziface

import "net"

/*
	定义连接模块的抽象层
 */
type IConnection interface {
	// 启动连接
	Start()

	// 停止连接
	Stop()

	// 获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn

	// 获取当前连接的模块的连接ID
	GetConnID() uint32

	// 获取远程客户端的TCP状态的IP port
	GetRemoteAddr() net.Addr

	// 直接将Message数据发送数据给远程的TCP客户端
	SendMsg(msgId uint32, data []byte) error
}

// 定义一个处理连接业务的方法类型
type HandleFunc func(*net.TCPConn, []byte, int) error