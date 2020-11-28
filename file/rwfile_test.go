package file

import (
	"fmt"
	"log"
	"testing"
	"ffile/util"
)

func TestGetfileinfo(t *testing.T) {
	filename:="../xxx.rar"
	info,err:=Getfileinfo(filename,false,false)
	if err!=nil{
		log.Println(err.Error())
	}
	fmt.Printf(" %+v \n", info)

	fmt.Println(util.Get(float64(info.Size),1.0))
}
