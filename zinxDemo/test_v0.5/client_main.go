package main

import (
	"net"
	"fmt"
	"time"
	"zinx/zinx/znet"
	"io"
)

func main(){
	conn ,err := net.Dial("tcp","0.0.0.0:7777")
	if err != nil{
		fmt.Println("Client Dial err: ",err)
		return
	}

	for {
		/*
		_,err = conn.Write([]byte("zinx client come in!"))
		if err != nil{
			fmt.Println("Client Write err: ",err)
			return
		}

		buf := make([]byte,512)
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("Client Read err: ",err)
			return
		}
		fmt.Printf(" servar call back : %s, cnt = %d\n", buf,n)
		*/


		//消息的发送还是要先封包
		msg := znet.NewMessage(1,[]byte("zinx client come in!"))

		dp := znet.NewDataPack()

		binarydata ,err := dp.Pack(msg)
		if err != nil{
			fmt.Println("lient pack error : ",err)
			return
		}

		_,err = conn.Write(binarydata)
		if err != nil{
			fmt.Println("lient write error : ",err)
			return
		}

		//消息的接收也是要分两次接收
		headMsg := make([]byte,dp.HeadLen())
		_,err = io.ReadFull(conn,headMsg)
		if err != nil{
			fmt.Println("client read head error : ",err)
			return
		}
		//二次读取
		recvmsg,err := dp.UnPack(headMsg)
		if err != nil{
			fmt.Println("拆包失败 ： ",err)
			return
		}
		if recvmsg.GetMsgLen()>0{
			data := make([]byte,recvmsg.GetMsgLen())

			_,err := io.ReadFull(conn,data)
			if err != nil{
				fmt.Println("client read msg data error : ",err)
				return
			}

			recvmsg.SetMsgData(data)

			fmt.Println("---> Recv Server Msg : id = ",recvmsg.GetMsgId(), "len = ", recvmsg.GetMsgLen(), " data = ", string(recvmsg.GetMsgData()))

		}


		time.Sleep(time.Second)
	}

}
