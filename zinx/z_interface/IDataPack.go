package z_interface

type IDataPack interface {
	//获取头部长度  事先规定好协议，确定头部的长度
	HeadLen()uint32
	//封包方法，将要发送的Message 封装程协议规定的头部+内容的格式
	Pack(msg IMessage)([]byte,error)
	//拆包方法，将接收到的二进制字节流，按照协议拆分，获取准确的消息内容
	UnPack([]byte)(IMessage,error)
}
