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
	msghandler z_interface.IMsgHandler
	//连接管理
	connmanager z_interface.IConnManager
	//server创建连接后自动调用的函数
	onConnStart func(conn z_interface.IConnection)
	//server销毁连接前自动调用的函数
	onConnStop func(conn z_interface.IConnection)
}

func NewServer(name string) z_interface.IServer {
	s := &Server{
		Name:config.GlobalObject.Name,
		IPVersion:"tcp4",
		IP:config.GlobalObject.Host,
		Port:config.GlobalObject.Port,
		msghandler:NewMsgHandler(),
		connmanager:NewConnManager(),
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

	s.msghandler.StartWorkerPool()

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

			//先判断当前server是否达到规定的最大连接数
			//最大连接数由配置文件给出
			if s.connmanager.Len() >= config.GlobalObject.MaxConn{
				//如果已经达到最大连接数,就断开连接
				fmt.Println("has become Max Connection,piease connect later....")
				conn.Close()
				continue
			}

			//如果没有达到最大连接数,调用NewConnection
			//dealConn := NewConnection(conn,cid,CallBackBusi)
			//dealConn := NewConnection(conn,cid,s.Router)
			//dealConn := NewConnection(conn,cid,s.msghandler)
			dealConn := NewConnection(s,conn,cid,s.msghandler)

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
	//当服务器停止的时候,删除管理器中的所有连接
	s.connmanager.ClearConn()
}
//运行服务器
func (s *Server)Serve(){
	s.Start()
	select {}
}
//添加路由
/*
func(s *Server)AddRouter(router z_interface.IRouter){
	s.Router = router
}
*/
func(s *Server)AddMsgHandler(msgId uint32, router z_interface.IRouter){
	s.msghandler.AddRouter(msgId,router)
	fmt.Println("Add Router Success! msgId = ",msgId)
}

func(s *Server)GetConnMgr()z_interface.IConnManager{
	return s.connmanager
}

//用户可自定义添加Hook钩子函数
//注册OnConnStart()hook函数
func(s *Server)AddOnConnStart(hook func(conn z_interface.IConnection)){
	s.onConnStart = hook
}
//注册OnConnStop()钩子函数
func(s *Server)AddOnConnStop(hook func(conn z_interface.IConnection)){
	s.onConnStop = hook
}

//调用hook函数的入口
func(s *Server)CallOnConnStart(conn z_interface.IConnection){
	if s.onConnStart != nil{
		fmt.Println("---Call the OnConnStart ---")
		s.onConnStart(conn)
	}
}
func(s *Server)CallOnConnStop(conn z_interface.IConnection){
	if s.onConnStop != nil{
		fmt.Println("=== Call the OnConnStop ===")
		s.onConnStop(conn)
	}
}