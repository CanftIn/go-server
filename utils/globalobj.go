package utils

import (
	"CanftIn/go-server/ziface"
	"encoding/json"
	"io/ioutil"
)

/*
	存储一切框架的全局参数，供其他模块使用
	通过zinx.json由用户配置
 */
type GlobalObj struct {
	/*
		Server
	 */
	TcpServer ziface.IServer   // 当前全局的server对象
	Host string                // 当前服务器主机监听的IP
	TcpPort int                // 当前服务器主机监听的端口号
	Name string                // 当前服务器名称

	/*
		zinx
	 */
	Version string             // 版本号
	MaxConn int                // 当前服务器主机允许的最大连接数
	MaxPackageSize uint32      // 当前框架数据包最大值
}

/*
	定义一个全局的对外Globalobj
 */
var GlobalObject *GlobalObj


/*
	从zinx.json中加载自定义的参数
 */
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 将json文件解析到struct中
	json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}


/*
	提供一个init方法，初始化当前的GlobalObject
 */
func init() {
	// 如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name: "ZinxServer",
		Version: "v0.6",
		TcpPort: 9999,
		Host: "0.0.0.0",
		MaxConn: 1000,
		MaxPackageSize: 4096,
	}

	// 从conf/zinx.json中导入值
	GlobalObject.Reload()
}
