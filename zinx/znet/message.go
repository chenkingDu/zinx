package znet

import "zinx/zinx/z_interface"

type Message struct {
	Data []byte
	Id uint32
	Len uint32
}

func NewMessage(id uint32,data []byte)z_interface.IMessage{
	return &Message{
		Id:id,
		Data:data,
		Len:uint32(len(data)),
	}
}

func(m *Message)GetMsgData()[]byte{
	return m.Data
}
func(m *Message)GetMsgId()uint32{
	return m.Id
}
func(m *Message)GetMsgLen()uint32{
	return m.Len
}

func(m *Message)SetMsgData(buf []byte){
	m.Data = buf
}
func(m *Message)SetMsgId(id uint32){
	m.Id = id
}
func(m *Message)SetMsgLen(len uint32){
	m.Len = len
}

