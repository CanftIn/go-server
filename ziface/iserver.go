package ziface

/*
	服务器框架抽象层
 */
type IServer interface {
	Start()

	Stop()

	Serve()

	// 路由功能，给当前的服务注册一个路由方法，供客户端连接使用
	AddRouter(router IRouter)
}