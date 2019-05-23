package znet

import "zinx/zinx/z_interface"

type Request struct {
	conn z_interface.IConnection
	data []byte
	len int
}


func(r *Request)GetConnection()z_interface.IConnection{
	return r.conn
}


func(r *Request)GetData()[]byte{
	return r.data
}


func(r *Request)GetDataLen()int{
	return r.len
}


func NewRequest(conn z_interface.IConnection,data []byte,len int)z_interface.IRequest{
	req := &Request{
		conn : conn,
		data : data,
		len : len,
	}

	return req
}