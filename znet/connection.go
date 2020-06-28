package znet

import (
	"CanftIn/go-server/ziface"
	"net"
)

/*
	连接模块
 */
type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前的连接状态
	IsClosed bool

	// 当前连接所绑定的处理业务方法API
	HandleAPI ziface.HandleFunc

	// 告知当前连接已经退出 channel
	ExitChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		HandleAPI: callback,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}