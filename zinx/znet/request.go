package znet

import "zinx/zinx/z_interface"

type Request struct {
	conn z_interface.IConnection
	//data []byte
	//len int
	msg z_interface.IMessage
}


func(r *Request)GetConnection()z_interface.IConnection{
	return r.conn
}


/*
func(r *Request)GetData()[]byte{
	return r.data
}


func(r *Request)GetDataLen()int{
	return r.len
}
*/

func(r *Request)GetMsg()z_interface.IMessage{
	return r.msg
}

func NewRequest(conn z_interface.IConnection,msg z_interface.IMessage)z_interface.IRequest{
	req := &Request{
		conn : conn,
		msg:msg,
	}

	return req
}