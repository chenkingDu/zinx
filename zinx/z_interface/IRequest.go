package z_interface

type IRequest interface {
	//得到当前的请求连接
	GetConnection() IConnection

	//得到数据内容
	GetData() []byte

	//得到数据长度
	GetDataLen() int

}

