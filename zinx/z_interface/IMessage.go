package z_interface

type IMessage interface {
	GetMsgData()[]byte
	GetMsgId()uint32
	GetMsgLen()uint32

	SetMsgData([]byte)
	SetMsgId(uint32)
	SetMsgLen(uint32)
}
