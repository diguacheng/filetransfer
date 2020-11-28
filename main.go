package main

import (
	"encoding/binary"
	"encoding/json"
	"ffile/file"
	"ffile/trans"
	"ffile/util"
	"fmt"
	"log"
	"net"
	"time"
)

func revice(port string) {
	listener, err := net.Listen("tcp4", "127.0.0.1:"+port)
	if err != nil {
		log.Println(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		var header []byte
		numByts := 2
		for numByts > len(header) {
			temp := make([]byte, numByts-len(header))
			n, err := conn.Read(temp)
			if err != nil {
				log.Println(err.Error())
			}
			header = append(header, temp[:n]...)
		}
		l := int(binary.LittleEndian.Uint16(header))
		message := make([]byte, 0)
		for l > len(message) {
			temp := make([]byte, l-len(message))
			n, err := conn.Read(temp)
			if err != nil {
				log.Println(err.Error())
			}
			message = append(message, temp[:n]...)
		}
		var m trans.Message
		json.Unmarshal(message, &m)
		var fileinfo file.Fileinfo
		if m.Type == "fileinfo" {
			json.Unmarshal(m.Data, &fileinfo)
			log.Printf("%+v \n", fileinfo)
		} else {
			log.Println("接受文件描述信息失败")
		}
		conn.Write([]byte("ok"))
		//log.Println("go 协程",fileinfo.FileName)
		go file.WriteFile(fileinfo)
		log.Println("接受中")
		starttime := time.Now()
		for {
			var header []byte
			numByts := 2
			for numByts > len(header) {
				temp := make([]byte, numByts-len(header))
				n, err := conn.Read(temp)
				if err != nil {
					log.Println(err.Error())
				}
				header = append(header, temp[:n]...)
			}
			l := int(binary.LittleEndian.Uint16(header))
			message := make([]byte, 0)
			for l > len(message) {
				temp := make([]byte, l-len(message))
				n, err := conn.Read(temp)
				if err != nil {
					log.Println(err.Error())
				}
				message = append(message, temp[:n]...)
			}
			m := &trans.Message{}
			err = json.Unmarshal(message, &m)
			if err != nil {
				log.Println("解析错误", err.Error())
				log.Printf("%+v \n", string(message))
			}
			if m.Type == "done" {
				break
			}
			file.FileStream <- m.Data

		}
		close(file.FileStream) // 收完
		durtime := time.Now().Sub(starttime)
		fmt.Println("接收速度：", util.Get(float64(fileinfo.Size), durtime.Seconds()))
		log.Println("收完")
		//time.Sleep(1)
		str, _ := file.Hash(fileinfo.FileName)
		if str == fileinfo.Hash {
			conn.Write([]byte("true"))
			log.Println(str)
		} else {
			log.Println(err.Error())
		}
		return
	}
}

func send(addr string, filePath string) {

	c, _ := trans.NewConnection(addr, filePath, true, true)
	//fmt.Printf("%v \n",c.Fileinfo)
	starttime := time.Now()
	go file.Readfile(c.Fileinfo)
	c.SendFileInfo()
	c.SendData()
	c.SendClose()
	dur := time.Now().Sub(starttime)
	fmt.Println("发送速度：", util.Get(float64(c.Fileinfo.Size), dur.Seconds()))
}

// 发送端 
func main() {
	remoteaddr := "ip:port" // 接收端的ip:port
	filePath := "test.txt"  // 发送的文件
	send(remoteaddr, filePath)
}

// 接收端
// func main() {
// 	port:="12345" // 本机用于接收的端口号
// 	revice(port)
// }
