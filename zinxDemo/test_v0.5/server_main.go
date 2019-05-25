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
	fmt.Println("Call Router Handlle....")
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

func main() {
	s := znet.NewServer("zinx v0.1")

	s.AddRouter(&PingRouter{})

	s.Serve()

	return
}
