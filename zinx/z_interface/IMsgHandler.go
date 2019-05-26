package z_interface


type IMsgHandler interface {
	//添加路由到map集合中
	AddRouter(msgId uint32, router IRouter)
	//调度路由，根据request中的msgId
	DoMsgHandler(request IRequest)
	//启动Worker工作池
	StartWorkerPool()
	//将消息添加到工作池中
	SendMsgToTaskQueue(request IRequest)

}
