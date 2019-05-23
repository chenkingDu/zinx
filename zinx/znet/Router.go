package znet

import "zinx/zinx/z_interface"

type BaseRouter struct {
}

func(r *BaseRouter)PreHandle(request z_interface.IRequest){}
func(r *BaseRouter)Handle(request z_interface.IRequest){}
func(r *BaseRouter)PostHandle(request z_interface.IRequest){}