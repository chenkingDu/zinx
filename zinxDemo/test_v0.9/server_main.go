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


func main() {
	s := znet.NewServer("zinx v0.1")

	//注册路由
	s.AddMsgHandler(1,&PingRouter{})
	s.AddMsgHandler(2,&HelloRouter{})

	s.Serve()

	return
}
