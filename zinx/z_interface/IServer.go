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
	//提供connmng的get方法
	GetConnMgr()IConnManager

	//用户可自定义添加Hook钩子函数
	AddOnConnStart(hook func(conn IConnection))
	AddOnConnStop(hook func(conn IConnection))
	//调用hook函数的入口
	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)


}
