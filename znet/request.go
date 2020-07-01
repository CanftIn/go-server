package znet

import "CanftIn/go-server/ziface"

type Request struct {
	// 已经和客户端建立好的连接
	conn ziface.IConnection

	// 客户端请求的数据
	msg ziface.IMessage
}

// 得到当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 得到当前的请求数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}