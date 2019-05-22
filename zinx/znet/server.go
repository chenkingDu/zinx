package znet

import (
	"zinx/zinx/z_interface"
	"net"
	"fmt"
)

type Server struct {
	//服务器ip
	IPVersion string
	IP string
	//服务器port
	Port int
	//服务器名称
	Name string

}

func NewServer(name string) z_interface.IServer {
	s := &Server{
		Name:name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:8999,
	}
	return s
}

//启动服务器
func (s *Server)Start(){
	addr ,err := net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err != nil{
		fmt.Println("ResolveTCPAddr err:",err)
		return
	}

	listenner ,err := net.ListenTCP(s.IPVersion,addr)
	if err != nil{
		fmt.Println("ListenTCP err:",err)
		return
	}

	go func() {
		for  {
			conn ,err := listenner.Accept()
			if err != nil{
				fmt.Println("Accept err:",err)
				continue
			}

			go func() {
				for {
					buf := make([]byte,512)
					n,err := conn.Read(buf)
					if err != nil{
						fmt.Println("read err:",err)
						break
					}

					_,err = conn.Write(buf[:n])
					if err != nil{
						fmt.Println("write err: ",err)
						continue
					}
				}
			}()
		}

	}()


}
//停止服务器
func (s *Server)Stop(){

}
//运行服务器
func (s *Server)Serve(){
	s.Start()
	select {}
}