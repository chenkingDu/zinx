package znet

import (
	"zinx/zinx/z_interface"
	"fmt"
	"zinx/zinx/config"
)

type MsgHandler struct {
	//存放路由表的map
	Apis map[uint32]z_interface.IRouter

	//负责Worker取任务的消息队列 ，一个worker对应一个任务队列
	TaskQueue []chan z_interface.IRequest

	//工作池的worker数
	WorkerPoolSize uint32
}

//初始化方法
func NewMsgHandler() z_interface.IMsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]z_interface.IRouter),
		WorkerPoolSize: config.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan z_interface.IRequest, config.GlobalObject.MaxWorkerTaskLen),
	}
}

//添加路由到map集合中
func (mh *MsgHandler) AddRouter(msgId uint32, router z_interface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		//说明这个msgId已经注册过了
		fmt.Println("this Id has been register....")
		return
	}
	mh.Apis[msgId] = router
	fmt.Println("Add Api MsgId = ", msgId, "Success!")
}

//调度路由，根据request中的msgId
func (mh *MsgHandler) DoMsgHandler(request z_interface.IRequest) {
	if router, ok := mh.Apis[request.GetMsg().GetMsgId()]; ok {
		//调用相应的router
		router.PreHandle(request)
		router.Handle(request)
		router.PostHandle(request)
	} else {
		//说明这个msgId不存在
		fmt.Println("this Api no found")
		return
	}

}

//启动Worker工作池(在整个server服务中只能启动一次)
func (mh *MsgHandler) StartWorkerPool() {
	fmt.Println("WorkerPool is Started.....")

	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//开启一个worker的Goroutine
		//给当前的Worker绑定消息channal对象 ，一个worker对应一个channal
		mh.TaskQueue[i] = make(chan z_interface.IRequest, config.GlobalObject.MaxWorkerTaskLen)

		//启动一个Worker，阻塞等待消息从对应的管道
		go mh.startOneWorker(i, mh.TaskQueue[i])

	}
}

//启动一个处理业务的Worker
func (mh *MsgHandler) startOneWorker(workerId int, taskQueue chan z_interface.IRequest) {
	fmt.Println("worker ID = ", workerId, "is starting ....")

	//不断地从管道读数据
	for {
		select {
		case req := <-taskQueue:
			//调度路由
			mh.DoMsgHandler(req)
		}
	}
}

//将消息添加到工作池中
//Reader调用
func (mh *MsgHandler) SendMsgToTaskQueue(request z_interface.IRequest) {
	//将消息平均分给worker，以确定当前的request到底要给那个worker来处理
	workerId := request.GetConnection().GetConnID()%mh.WorkerPoolSize

	//直接将消息发送给对应的worker的taskqueue
	mh.TaskQueue[workerId] <- request

}
