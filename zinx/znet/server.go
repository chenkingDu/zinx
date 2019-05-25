package znet

import (
	"zinx/zinx/z_interface"
	"net"
	"fmt"
	"zinx/zinx/config"
)

type Server struct {
	//服务器ip
	IPVersion string
	IP string
	//服务器port
	Port int
	//服务器名称
	Name string
	//Router方法
	Router z_interface.IRouter
}

func NewServer(name string) z_interface.IServer {
	s := &Server{
		Name:config.GlobalObject.Name,
		IPVersion:"tcp4",
		IP:config.GlobalObject.Host,
		Port:config.GlobalObject.Port,
		Router:nil,
	}
	return s
}

//定义一个具有回显功能 针对type HandleFunc func(*net.TCPConn,[]byte,int)error
/*
func CallBackBusi(request z_interface.IRequest)error{
	fmt.Println("【conn Handle】 CallBack..")
	conn := request.GetConnection().GetTCPConnection()
	data := request.GetData()
	cnt := request.GetDataLen()
	if _,err := conn.Write(data[:cnt]);err != nil{
		fmt.Println("write back err:",err)
		return err
	}

	return nil
}
*/



//启动服务器
func (s *Server)Start(){
	fmt.Printf("[start] Server Linstenner at IP : %s, Port : %d, is Strating...\n",s.IP,s.Port)

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

	//生成id累加器
	var cid uint32
	cid = 0


	go func() {
		for  {
			conn ,err := listenner.AcceptTCP()
			if err != nil{
				fmt.Println("Accept err:",err)
				continue
			}


			//dealConn := NewConnection(conn,cid,CallBackBusi)
			dealConn := NewConnection(conn,cid,s.Router)
			cid ++

			go dealConn.Start()


			/*
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
			*/

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
//添加路由
func(s *Server)AddRouter(router z_interface.IRouter){
	s.Router = router
}