package ziface

/*
	服务器框架抽象层
 */
type IServer interface {
	Start()

	Stop()

	Serve()
}