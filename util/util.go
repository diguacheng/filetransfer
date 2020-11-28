package util

import "fmt"

func Get(size, time float64) string {
	x := size / time
	unit:=[]string{"B/s","KB/s","MB/s","GB/s"}
	c:=0
	for x>=1024&&c<3{
		x=x/1024
		c++
	}
	if c==4{
		c=3
	}
	return fmt.Sprintf("%.2f"+unit[c],x)
}

func GetSize(size float64) string {
	x := size 
	unit:=[]string{"B","KB","MB","GB"}
	c:=0
	for x>=1024&&c<3{
		x=x/1024
		c++
	}
	if c==4{
		c=3
	}
	return fmt.Sprintf("%.2f"+unit[c],x)
}


