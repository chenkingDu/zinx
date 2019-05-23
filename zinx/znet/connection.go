package znet

import (
	"net"
	"zinx/zinx/z_interface"
	"fmt"
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

func(c *Connection)StartRead(){
	fmt.Println("Reader go is starting ....")
	defer fmt.Println("ConnID = ",c.ConnID,"Reader is exit,remote addr is = ",c.GetRemoteAddr().String())
	defer c.Stop()

	buf := make([]byte,512)
	for{
		cnt ,err := c.Conn.Read(buf)
		if err != nil{
			fmt.Println("read buf err :",err)
			continue
		}

		//将数据传递给定义好的handle回调
		if err := c.handleAPI(c.Conn,buf,cnt);err != nil{
			fmt.Println("ConnID:",c.ConnID,"Handle is err : ",err)
			break
		}
	}



}

//启动连接
func(c *Connection)Start(){
	fmt.Println("Conn Start() ....id =  ",c.ConnID)

	go c.StartRead()


}

//停止连接
func(c *Connection)Stop(){
	fmt.Println("c.Stop() ....")
	if c.isClosed == true{
		return
	}
	c.isClosed = true

	c.Conn.Close()
}

//获取连接ID
func(c *Connection)GetConnID()uint32{
	return c.ConnID
}

//获取conn的原生socket套接字
func(c *Connection)GetTCPConnection()*net.TCPConn{
	return c.Conn
}

//获取远程客户端的ip地址
func(c *Connection)GetRemoteAddr()net.Addr{
	return c.Conn.RemoteAddr()
}

//发送数据给对方客户端
func(c *Connection)Send(data []byte,cnt int)error{
	if _,err := c.Conn.Write(data[:cnt]);err != nil{
		fmt.Println("send buf error")
		return err
	}
	return nil
}

