package util

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	fmt.Println(Get(1023, 1.0))
	fmt.Println(Get(10256, 1.0))
}

func TestGetSize(t *testing.T) {

	fmt.Println(GetSize(1024))
	fmt.Println(GetSize(1024 * 1087))
}
