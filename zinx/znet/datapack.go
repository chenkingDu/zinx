package znet

import (
	"zinx/zinx/z_interface"
	"bytes"
	"encoding/binary"
	"fmt"
)

type DataPack struct {
}

func NewDataPack()*DataPack{
	return &DataPack{}
}

//规定头部大小
func(dp *DataPack)HeadLen()uint32{
	return 8
}


//封包，将message封装成/len/id/data/的格式
func(dp *DataPack)Pack(message z_interface.IMessage)([]byte,error){
	//创建一个存放二进制流的缓冲区  bytes.NewBuffer([]byte{})会自动扩容
	databuf := bytes.NewBuffer([]byte{})

	//将len写入databuf
	//binary.Write(存放的缓冲区，大小端法，要写如的数据)
	if err := binary.Write(databuf,binary.LittleEndian,message.GetMsgLen());err != nil{
		fmt.Println("写入len出错：",err)
		return nil,err
	}
	//将id写入databuf
	if err := binary.Write(databuf,binary.LittleEndian,message.GetMsgId());err != nil{
		fmt.Println("写入id出错：",err)
		return nil,err
	}

	//将data写入databuf
	if err := binary.Write(databuf,binary.LittleEndian,message.GetMsgData());err != nil{
		fmt.Println("写入data出错：",err)
		return nil,err
	}

	//写入成功后，返回databuf
	return databuf.Bytes(),nil

}


//拆包 ，把len和id作为头部读出来进行拆开，data的读取要根据len的长度
func(dp *DataPack)UnPack(HeadData []byte)(z_interface.IMessage,error){
	msghead := &Message{}

	//创建一个读取二进制数据流的io.Reader
	readbuf := bytes.NewReader(HeadData)

	//从readbuf读数据存入len
	if err := binary.Read(readbuf,binary.LittleEndian,&msghead.Len);err != nil{
		fmt.Println("读取头部len失败：",err)
		return nil,err
	}

	//从readbuf读数据存入id
	if err := binary.Read(readbuf,binary.LittleEndian,&msghead.Id);err != nil{
		fmt.Println("读取头部id失败：",err)
		return nil,err
	}

	return msghead,nil
}