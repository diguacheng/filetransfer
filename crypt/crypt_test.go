package crypt

import (
	"encoding/hex"
	"ffile/compress"
	"fmt"
	"log"
	"testing"
)

func TestGetKey(t *testing.T) {
	key := GetKey([]byte("helxvcvbxvrld"))
	fmt.Println(hex.EncodeToString(key), len(key))
}

func TestEnAndDe(t *testing.T) {
	key := GetKey([]byte("123345"))
	data := []byte("hello worldfsdfsdfsdfsdfsfsdfd")
	fmt.Println(hex.EncodeToString(data))
	encrypted, err := Encrypt(key, data)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(hex.EncodeToString(encrypted))
	d, err := Decrypt(key, encrypted)
	fmt.Println(hex.EncodeToString(d))

}

func TestEnAndDe1(t *testing.T) {
	key := GetKey([]byte("123345"))
	data := []byte(`"Biden's win in Arizona was not for a lack of trying on Trump's part. The President held seven events in the state in 2020. Biden held one event after the Democratic National Convention over the summer, a bus tour around Maricopa in October.
	To Slugocki, those visits did little to break through the voters' focus on education, health care and the economy.
	"Clearly, voters wanted something new from Arizona. Voters were energized and enthused to vote. Maricopa County's elections are safe, secure and transparent," said the county party chair. "A bright future is ahead for Maricopa County and I couldn't be prouder."`)
	fmt.Println(len(data))
	data=compress.Compress(data)
	encrypted, err := Encrypt(key, data)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(len(encrypted))

	d, err := Decrypt(key, encrypted)
	d=compress.Decompress(d)
	fmt.Println(len(d))

}
