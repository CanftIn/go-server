package znet

import (
	"CanftIn/go-server/utils"
	"CanftIn/go-server/ziface"
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
	// 当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	msgHandler ziface.IMsgHandle
}


func (s *Server) Start() {
	fmt.Printf("[zinx] server Name: %s, listenner at IP: %s, Port: %d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[zinx] Version %s, MaxConn: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)

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
			dealConn := NewConnection(conn, cid, s.msgHandler)
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

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router Succ!")
}

func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandle(),
	}
	return s
}