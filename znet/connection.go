package znet

import (
	"CanftIn/go-server/ziface"
	"golang.org/x/exp/errors/fmt"
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

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("c.connID = ", c.ConnID, " Reader is exit, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}

		// 调用当前连接所绑定的HandleAPI
		if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID ", c.ConnID, " handle is error ", err)
			break
		}
	}
}

// 启动连接
func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID = ", c.ConnID)

	// 启动从当前连接的读数据的业务
	go c.StartReader()
}

// 停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	// 如果当前连接已经关闭
	if c.IsClosed == true {
		return
	}

	c.IsClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

// 获取当前连接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接的模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态的IP port
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
