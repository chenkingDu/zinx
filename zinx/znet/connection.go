package znet

import (
	"net"
	"zinx/zinx/z_interface"
)

//具体的TCP 连接模块
type Connection struct {
	//当前连接的原生套接字
	Conn *net.TCPConn

	//连接ID
	ConnID uint32

	//当前连接状态
	isClosed bool

	//当前连接所绑定的业务处理方法
	handleAPI z_interface.HandleFunc

}

//初始化连接方法
func NewConnection(conn *net.TCPConn,connID uint32,callback_api z_interface.HandleFunc)z_interface.IConnection{
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		handleAPI:callback_api,
		isClosed:false,
	}

	return c
}

func(c *Connection)Start(){
	
}

//停止连接
func(c *Connection)Stop(){

}

//获取连接ID
func(c *Connection)GetConnID()uint32{
	return 0
}

//获取conn的原生socket套接字
func(c *Connection)GetTCPConnection()*net.TCPConn{
	return nil
}

//获取远程客户端的ip地址
func(c *Connection)GetRemoteAddr()net.Addr{
	return nil
}

//发送数据给对方客户端
func(c *Connection)Send(data []byte)error{
	return nil
}
