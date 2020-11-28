package trans

import (
	"encoding/json"
	"ffile/file"
	"fmt"
	"testing"
)

func TestNewConnection(t *testing.T) {
	info,_:=file.Getfileinfo("../ttt.pdf",false,false)
	data,_:=json.Marshal(info)
	fmt.Println(string(data))
	var f file.Fileinfo
	_=json.Unmarshal(data,&f)
	fmt.Printf("%+v \n",f)

}
