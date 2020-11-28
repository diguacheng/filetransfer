package trans

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"ffile/file"
	"log"
	"net"
)

type connection struct {
	conn *net.TCPConn
	Fileinfo  *file.Fileinfo
}

type Message struct {
	Type string  `json:"t,omitempty"`// start data end
	Data []byte 	`json:"d,omitempty"`
	//Done bool		`json:"b,omitempty"`
}



func NewConnection(addr string,filename string,isCompress,isCrpyt bool)(*connection,error){
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err!=nil{
		log.Println(err.Error())
	}
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	if err!=nil{
		log.Println(err.Error())
	}
	info,_:=file.Getfileinfo(filename,isCompress,isCrpyt)
	return &connection{conn:conn,Fileinfo:info},nil
}


func (C *connection)SendFileInfo()error{
	m:=new(Message)
	m.Type="fileinfo"
	m.Data,_=json.Marshal(C.Fileinfo)
	message,err:=json.Marshal(m)
	if err!=nil{
		log.Println(err.Error())
	}
	l:=len(message)
	header:=make([]byte,2)
	binary.LittleEndian.PutUint16(header,uint16(l))//加头部长度
	sdata:=append(header,message...)

	_,err=C.conn.Write(sdata)
	data:=make([]byte,1024)
	n, err := C.conn.Read(data)
	if err!=nil{
		log.Println(err.Error())
	}
	if string(data[:n])=="ok"{
		return nil
	}
	return errors.New("err in step 1 send fileinfo")
}

func (C *connection)SendData()error{
	m:=new(Message)
	m.Type="data"
	for data:= range file.FileStream{
		m.Data=data
		m,_:=json.Marshal(m)
		l:=len(m)
		header:=make([]byte,2)
		binary.LittleEndian.PutUint16(header,uint16(l))
		sdata:=append(header,m...)
		_,err:=C.conn.Write(sdata)
		if err!=nil{
			return err
		}
	}
	return nil
}

func (C *connection)SendClose()error{
	m:=new(Message)
	m.Type="done"
	message,_:=json.Marshal(m)
	l:=len(message)
	header:=make([]byte,2)
	binary.LittleEndian.PutUint16(header,uint16(l))//加头部长度
	sdata:=append(header,message...)
	_,err:=C.conn.Write(sdata)
	if err!=nil{
		log.Println(err.Error())
	}
	data:=make([]byte,1024)
	n,_:=C.conn.Read(data)
	if string(data[:n])=="true"{
		return nil
	}
	return  err

}

