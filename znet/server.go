package znet

import (
	"CanftIn/go-server/ziface"
	"errors"
	"fmt"
	"net"
)

/*
	iserver 实现
 */
type Server struct {
	// 名称
	Name string
	// ip版本
	IPVersion string
	// 监听的ip
	IP string
	// 监听的端口
	Port int
}

// 定义当前客户端连接的所绑定的HandleAPI(以后优化为用户自定义方法)
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// echo业务
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP: %s, Port %d, is starting. \n", s.IP, s.Port);

	go func() {
		// 获取 socket addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port));
		if err != nil {
			fmt.Printf("resolve tcp addr error: ", err)
			return
		}

		// 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listen ", s.IPVersion, " err ", err)
			return
		}

		fmt.Println("start zinx server succ, ", s.Name, " succ, Listenning...")
		var cid uint32
		cid = 0

		for {
			// 如果有客户端连接，阻塞返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			// 将处理新连接的业务方法和conn绑定，得到连接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			// 启动当前连接业务
			go dealConn.Start()
		}

	}()
}

func (s *Server) Stop() {
	// TODO 将资源、状态及连接信息回收
}

func (s *Server) Serve() {
	// 启动server的服务
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select {

	}
}

func NewServer(name string) ziface.IServer {
	s := &Server {
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 9999,
	}
	return s
}