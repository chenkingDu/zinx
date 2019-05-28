package main

import (
	"zinx/zinx/znet"
	"zinx/zinx/z_interface"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

/*
func (this *PingRouter)PreHandle(request z_interface.IRequest){
	fmt.Println("Call Router PreHandlle....")

	_,err := request.GetConnection().GetTCPConnection().Write([]byte("Pre Ping.....\n"))
	if err != nil{
		fmt.Println("call back pre ping err:",err)
	}
}
*/

func (this *PingRouter)Handle(request z_interface.IRequest){
	fmt.Println("Call PingRouter Handlle....")
	/*
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("Ping.....\n"))
	if err != nil{
		fmt.Println("call back ping err:",err)
	}
	*/
	err := request.GetConnection().Send(1,[]byte("ping.....pong......ping.....pong.."))
	if err != nil{
		fmt.Println("send msg error : ",err)
		return
	}

}

/*
func (this *PingRouter)PostHandle(request z_interface.IRequest){
	fmt.Println("Call Router PreHandlle....")

	_,err := request.GetConnection().GetTCPConnection().Write([]byte("After Ping.....\n"))
	if err != nil{
		fmt.Println("call back after ping err:",err)
	}
}
*/

type HelloRouter struct {
	znet.BaseRouter
}

func (this *HelloRouter)Handle(request z_interface.IRequest){
	fmt.Println("Call HelloRouter Handlle....")
	/*
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("Ping.....\n"))
	if err != nil{
		fmt.Println("call back ping err:",err)
	}
	*/
	err := request.GetConnection().Send(2,[]byte("hello zinx!.."))
	if err != nil{
		fmt.Println("send msg error : ",err)
		return
	}

}

//创建链接之后执行的钩子函数
func DoConnectBegin(conn z_interface.IConnection){
	fmt.Println("===> DoConntionBegin  ....")

	//链接一旦创建成功 给用户返回一个消息
	if err := conn.Send(202, []byte("Hello welcome to zinx...")); err !=nil {
		fmt.Println(err)
	}

	//给链接绑定一些属性
	conn.SetProperty("Name","Go3")
	conn.SetProperty("Address","TBD")
	conn.SetProperty("Time","2019")

}


//销毁链接前执行的钩子函数
func DoConnectLost(conn z_interface.IConnection){
	fmt.Println("===> DoConntionLost  ....")
	fmt.Println("Conn id ", conn.GetConnID(), "is Lost!.")

	//开始获取属性
	if name,err := conn.GetProperty("Name");err == nil{
		fmt.Println("Name =", name)
	}
	if Address,err := conn.GetProperty("Address");err == nil{
		fmt.Println("Address =", Address)
	}
	if Time,err := conn.GetProperty("Time");err == nil{
		fmt.Println("Time =", Time)
	}


}

func main() {
	s := znet.NewServer("zinx v0.10")

	//注册路由
	s.AddMsgHandler(1,&PingRouter{})
	s.AddMsgHandler(2,&HelloRouter{})

	//注册钩子函数
	s.AddOnConnStart(DoConnectBegin)
	s.AddOnConnStop(DoConnectLost)

	s.Serve()

	return
}
