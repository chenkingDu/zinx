package znet

import (
	"zinx/zinx/z_interface"
	"sync"
	"fmt"
	"errors"
)

type ConnManager struct {
	//保存连接的集合map
	connections map[uint32]z_interface.IConnection
	//针对这个map的读写锁
	connLock sync.RWMutex

}

//初始化方法
func NewConnManager()z_interface.IConnManager{
	return &ConnManager{
		connections:make(map[uint32] z_interface.IConnection),
	}
}

//添加连接
func(cm *ConnManager)Add(connection z_interface.IConnection){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[connection.GetConnID()] = connection
	fmt.Println("Add connection Success!connId = ",connection.GetConnID())
}
//删除连接
func(cm *ConnManager)Remove(connId uint32){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections,connId)
	fmt.Println("Remove connId = ",connId," connection Success!")
}
//根据ID获取连接
func(cm *ConnManager)Get(connId uint32)(z_interface.IConnection,error){
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn,ok := cm.connections[connId];ok{
		//找到了
		return conn,nil
	}else {
		//没找到
		return nil,errors.New("connection not found!")
	}
}
//获取连接的总个数
func(cm *ConnManager)Len()uint32{
	return uint32(len(cm.connections))
}
//清空全部连接
func(cm *ConnManager)ClearConn(){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//遍历删除
	for connId,conn := range cm.connections{
		//先关闭连接
		conn.Stop()
		//再删除连接
		delete(cm.connections,connId)
	}

	fmt.Println("Clear All Connections Success!")

}