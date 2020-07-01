package ziface

/*
	服务器框架抽象层
 */
type IServer interface {
	Start()

	Stop()

	Serve()

	// 路由功能
	AddRouter(msgId uint32, router IRouter)
}