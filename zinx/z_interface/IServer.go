package z_interface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//添加路由
	//AddRouter(router IRouter)
	AddMsgHandler(msgId uint32,router IRouter)

}
