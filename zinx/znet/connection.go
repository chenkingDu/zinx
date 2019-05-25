package znet

import (
	"net"
	"zinx/zinx/z_interface"
	"fmt"
	//"zinx/zinx/config"
	"io"
	"errors"
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
	//交给用户写自己的回调函数
	//handleAPI z_interface.HandleFunc

	//Router成员
	//Router z_interface.IRouter
	msgHandler z_interface.IMsgHandler

}

//初始化连接方法
func NewConnection(conn *net.TCPConn,connID uint32,handler z_interface.IMsgHandler)z_interface.IConnection{
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		//handleAPI:callback_api,
		isClosed:false,
		//Router:router
		msgHandler:handler,
	}

	return c
}

func(c *Connection)StartRead(){
	fmt.Println("Reader go is starting ....")
	defer fmt.Println("ConnID = ",c.ConnID,"Reader is exit,remote addr is = ",c.GetRemoteAddr().String())
	defer c.Stop()

	//buf := make([]byte,config.GlobalObject.MaxPackageSize)
	for{
		/*
		cnt ,err := c.Conn.Read(buf)
		if err != nil{
			fmt.Println("read buf err :",err)
			continue
		}
		*/

		//开始拆包
		//创建datapack
		dp := NewDataPack()

		//读取头部信息
		headdata := make([]byte,dp.HeadLen())
		_,err := io.ReadFull(c.Conn,headdata)
		if err != nil{
			fmt.Println("read headdata error : ",err)
			return
		}

		//根据头部进行第二次读取
		headmsg,err := dp.UnPack(headdata)
		if err != nil{
			fmt.Println("unpack error : ",err)
			return
		}

		data := make([]byte,headmsg.GetMsgLen())
		if headmsg.GetMsgLen() > 0{
			_,err := io.ReadFull(c.Conn,data)
			if err != nil{
				fmt.Println("read msg data error : ",err)
				return
			}
			headmsg.SetMsgData(data)
		}

		msg := headmsg.(*Message)

		req := NewRequest(c,msg)

		//将数据传递给定义好的handle回调
		//if err := c.handleAPI(c.Conn,buf,cnt);err != nil{
		/*
		//抽离请求
		if err := c.handleAPI(req);err != nil{
				fmt.Println("ConnID:",c.ConnID,"Handle is err : ",err)
			break
		}
		*/

		/*
		go func() {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()
		*/
		go c.msgHandler.DoMsgHandler(req)

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
func(c *Connection)Send(id uint32,data []byte)error{
	/*
	if _,err := c.Conn.Write(data[:cnt]);err != nil{
		fmt.Println("send buf error")
		return err
	}
	return nil
	*/

	//先检测连接是否关闭
	if c.isClosed == true{
		return errors.New("Cinnection is Closed..")
	}

	//先封包再发送
	msg := NewMessage(id,data)

	dp := NewDataPack()
	binarymsg,err := dp.Pack(msg)
	if err != nil{
		fmt.Println("pack msg error : ",err)
		return err
	}

	_,err = c.Conn.Write(binarymsg)
	if err != nil{
		fmt.Println("write msg error : ",err)
		return err
	}

	return nil


}


