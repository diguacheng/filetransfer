package compress

import (
	"bytes"
	"crypto/rand"
	"ffile/util"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	 data:=make([]byte, 1024*12)
	 rand.Read(data)

	fmt.Println(len(data))
	d:=Compress(data) 
	fmt.Println(len(d))
	res:=Decompress1(d)
	fmt.Println(len(res))
	fmt.Println(bytes.Compare(data,res))
}

func Test1Compress(t *testing.T) {
	
	f,_:=os.Open("../123.txt")
	defer f.Close()

	data:=make([]byte,1024*1024)
	off:=int64(0)
	sum1:=0
	sum2:=0
	for {
		n,err:=f.ReadAt(data,off)
		//fmt.Println(Compress(data[:n]))
		ll:=len(Compress(data[:n]))
		fmt.Printf("%.2f%% \n",float64(ll)/float64(n)*100)
		sum1+=n
		off+=int64(n)
		sum2+=ll
		//fmt.Println(n,err)
		if err!=nil{
			if err==io.EOF{
				//fmt.Println("done")
				break
			}
		}

	}
	fmt.Println("before",util.GetSize(float64(sum1)))
	fmt.Println("after",util.GetSize(float64(sum2)))
	//fmt.Println(sum2)
	fmt.Printf("Taotal  %.2f%% \n",float64(sum2)/float64(sum1)*100)
}

