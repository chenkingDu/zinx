package znet

import (
	"zinx/zinx/z_interface"
	"fmt"
)

type MsgHandler struct {
	//存放路由表的map
	Apis map[uint32]z_interface.IRouter
}

//初始化方法
func NewMsgHandler()z_interface.IMsgHandler{
	return &MsgHandler{
		Apis:make(map[uint32]z_interface.IRouter),
	}
}

//添加路由到map集合中
func(mh *MsgHandler)AddRouter(msgId uint32, router z_interface.IRouter){
	if _,ok := mh.Apis[msgId];ok{
		//说明这个msgId已经注册过了
		fmt.Println("this Id has been register....")
		return
	}
	mh.Apis[msgId] = router
	fmt.Println("Add Api MsgId = ",msgId,"Success!")
}
//调度路由，根据request中的msgId
func(mh *MsgHandler)DoMsgHandler(request z_interface.IRequest){
	if router,ok := mh.Apis[request.GetMsg().GetMsgId()];ok{
		//调用相应的router
		router.PreHandle(request)
		router.Handle(request)
		router.PostHandle(request)
	}else {
		//说明这个msgId不存在
		fmt.Println("this Api no found")
		return
	}

}
