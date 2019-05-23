package main

import (
	"net"
	"fmt"
	"time"
)

func main(){
	conn ,err := net.Dial("tcp","0.0.0.0:8999")
	if err != nil{
		fmt.Println("Client Dial err: ",err)
		return
	}

	_,err = conn.Write([]byte("zinx client come in!"))
	if err != nil{
		fmt.Println("Client Write err: ",err)
		return
	}

	for {

		buf := make([]byte,512)
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("Client Read err: ",err)
			return
		}
		fmt.Printf(" servar call back : %s, cnt = %d\n", buf,n)

		time.Sleep(time.Second)
	}

}
