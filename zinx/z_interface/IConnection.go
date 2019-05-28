package z_interface

import "net"

type IConnection interface {
	//启动连接
	Start()

	//停止连接
	Stop()

	//获取连接ID
	GetConnID() uint32

	//获取conn原生套接字
	GetTCPConnection() *net.TCPConn

	//获取远程客户端的ip
	GetRemoteAddr() net.Addr

	//发送数据给对方客户端
	Send(id uint32,data []byte) error

	//设置属性
	SetProperty(key string,value interface{})
	//查找属性
	GetProperty(key string)(interface{},error)
	//删除属性
	DelProperty(key string)

}

//业务处理方法，抽象定义
type HandleFunc func(request IRequest)error
