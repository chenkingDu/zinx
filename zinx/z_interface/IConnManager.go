package z_interface

type IConnManager interface {
	//添加连接
	Add(connection IConnection)
	//删除连接
	Remove(connId uint32)
	//根据ID获取连接
	Get(connId uint32)(IConnection,error)
	//获取连接的总个数
	Len()uint32
	//清空全部连接
	ClearConn()

}

