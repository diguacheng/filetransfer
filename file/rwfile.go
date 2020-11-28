package file

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"ffile/compress"
	"ffile/crypt"
	"fmt"
	"io"
	"log"
	"os"
)


type Fileinfo struct {
	FileName string
	Hash string
	Size int64
	IsCompress bool 
	IsCrpyt bool 
	Code []byte // 用于生成密码 

}

var bufferSize=1024*10
var FileStream=make(chan []byte,10)


func Getfileinfo(filename string,isCompress,isCrpyt bool)(*Fileinfo,error){
	fileinfo:=new(Fileinfo)
	info,err:=os.Stat(filename)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fileinfo.FileName=info.Name()
	fileinfo.Size=info.Size()
	fileinfo.Hash,_=Hash(filename)
	fileinfo.IsCompress=isCompress
	fileinfo.IsCrpyt=isCrpyt
	fileinfo.Code=crypt.Code
	return  fileinfo,nil
}

func Hash(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	md5 := md5.New()
	if _, err := io.Copy(md5, file); err != nil {
		log.Println(err.Error())
	}
	//md5.Write([]byte("123"))
	md5bytes := md5.Sum(nil)

	return hex.EncodeToString(md5bytes), nil
}


func Readfile(finfo *Fileinfo) {
	file, err := os.Open(finfo.FileName)
	if err != nil {
		log.Println(err.Error())
	}
	data := make([]byte, bufferSize/2)
	off := int64(0)

	for {
		n, err := file.ReadAt(data, off)
		offheader := &bytes.Buffer{}
		binary.Write(offheader, binary.LittleEndian, off) // 写入偏移量  8 个字节 
		data := append(offheader.Bytes(), data[:n]...)
		if finfo.IsCompress{
			data=compress.Compress(data[:n+8])
		}
		if finfo.IsCrpyt{
			data,_=crypt.Encrypt(crypt.Key,data)
		}
		//log.Println("send",len(data))
		//
		FileStream <- data
		//log.Println("red",n)
		if err != nil {
			//log.Println(err.Error())
			if err == io.EOF {
				log.Println("文件发送完毕")
			}
			break
		}
		off += int64(n)
	}
	_=file.Close()
	close(FileStream)
}


func WriteFile(finfo Fileinfo){

	file,err:=os.Create(finfo.FileName)
	if err!=nil{
		log.Println(err.Error())
	}
	var off int64
	var key []byte
	if finfo.IsCrpyt{
		key=crypt.GetKey(finfo.Code)
	}

	for data:=range FileStream {
		if finfo.IsCrpyt{
			data,_=crypt.Decrypt(key,data)
		}
		// 
		if finfo.IsCompress{
			data=compress.Decompress(data)
		}

		header:=bytes.NewReader(data[:8])
		binary.Read(header, binary.LittleEndian, &off)
		_,err:=file.WriteAt(data[8:],int64(off))
		
		if err!=nil{
			log.Println(err.Error())
		}
	}
	_= file.Close()
}
