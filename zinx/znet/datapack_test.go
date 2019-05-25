package znet

import (
	"testing"
	"fmt"
	"net"
	"io"
)

func TestDataPack(t *testing.T){
	fmt.Println("this is TsetDataPack")

	//模拟一个server
	listenner ,err := net.Listen("tcp","192.168.136.233:7777")
	if err != nil{
		fmt.Println("listen err : ",err)
		return
	}

	//并发连接
	go func() {
		for {
			conn,err := listenner.Accept()
			if err != nil{
				fmt.Println("accept err : ",err)
				return
			}

			//读写业务
			go func(conn *net.Conn) {
				//拆包
				dp := NewDataPack()
				for {
					//把规定好的头部取出来
					headData := make([]byte,dp.HeadLen())
					//ReadFull，直到headData填充满，才会返回，否则阻塞
					_,err := io.ReadFull(*conn,headData)
					if err != nil{
						fmt.Println("read head err : ",err)
						break
					}

					msghead,err := dp.UnPack(headData)
					if err != nil{
						fmt.Println("拆包 err : ",err)
						break
					}

					if msghead.GetMsgLen()>0{
						//继续读数据放入data中
						msg := msghead.(*Message)
						msg.Data = make([]byte,msg.GetMsgLen())
						//再次read
						_,err := io.ReadFull(*conn,msg.Data)
						if err != nil{
							fmt.Println("read data err : ",err)
							break
						}
						fmt.Println("---Recv  MsgId = ",msg.Id,"MsgLen = ",msg.Len,"MsgData = ",string(msg.Data))
					}

				}
			}(&conn)
		}
	}()


	//模拟一个client
	conn,err := net.Dial("tcp","192.168.136.233:7777")
	if err != nil {
		fmt.Println("client dail err: ", err)
		return
	}


	dp := NewDataPack()
	//模拟粘包
	msg1 := &Message{
		Id:1,
		Len:4,
		Data: []byte{'z','i','n','x'},
	}
	//封包
	send1,err := dp.Pack(msg1)
	if err != nil{
		fmt.Println("client send data1 error")
		return
	}
	msg2 := &Message{
		Id:2,
		Len:5,
		Data: []byte{'h','e','l','l','o'},
	}
	//封包
	send2,err := dp.Pack(msg2)
	if err != nil{
		fmt.Println("client send data1 error")
		return
	}

	//将包粘在一起
	send1 = append(send1,send2...)
	//发送
	conn.Write(send1)

	select {}
}
