package compress

import (
	"bytes"
	"compress/flate"
	"fmt"

	"log"
)


func Compress(data []byte)[]byte {
	var b bytes.Buffer
	defer b.Reset()
	compressor,err := flate.NewWriter(&b,flate.BestCompression)
	if err!=nil{
		log.Println(err.Error())
	}
	compressor.Write(data)
	compressor.Close()
	//log.Printf("%.2f \n",float64(b.Len())/float64(len(data)))
	return b.Bytes()
}

func Decompress(data []byte)[]byte{
	res:=make([]byte,1024*10)
	Decompressor:=flate.NewReader(bytes.NewReader(data))
	var n int
	n,_=Decompressor.Read(res)
	if n>=1024*10{
		fmt.Println("error")
	}
	Decompressor.Close()
	return res[:n]
}


func Decompress1(data []byte)[]byte{
	var res bytes.Buffer
	Decompressor:=flate.NewReader(bytes.NewReader(data))
	var n int
	n,_=Decompressor.Read(res.Bytes())
	fmt.Println(n)
	
	Decompressor.Close()
	return res.Bytes()
}


