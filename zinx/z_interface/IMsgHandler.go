package z_interface


type IMsgHandler interface {
	//添加路由到map集合中
	AddRouter(msgId uint32, router IRouter)
	//调度路由，根据request中的msgId
	DoMsgHandler(request IRequest)
}
