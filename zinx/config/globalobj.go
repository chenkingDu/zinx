package config

import (
	"io/ioutil"
	"encoding/json"
)

type GlobalObj struct {
	Host string		//当前监听的IP
	Port int 		//当前监听的Port
	Name string		//当前zinx_server的名称
	Version string	//当前框架的版本号
	MaxPackageSize uint32	//定义读写缓冲区buf的大小
}

//定义一个全局的对外的配置对象
var GlobalObject *GlobalObj

func(g *GlobalObj)LoadConfig(){
	//这个url是针对使用zinx框架开发的服务器的main主进程所在的路径的相对路径
	data,err := ioutil.ReadFile("conf/zinx.json")
	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(data,&GlobalObject)
	if err != nil{
		panic(err)
	}
}


func init(){
	//默认值
	GlobalObject = &GlobalObj{
		Host:"0.0.0.0",
		Port:8999,
		Name:"ZinxServerApp",
		Version:"v0.4",
		MaxPackageSize:512,
	}
	//配置文件读取
	//加载文件
	GlobalObject.LoadConfig()
}