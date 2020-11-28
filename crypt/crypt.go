package crypt

import (
	"crypto/aes"
	"crypto/cipher"

	"crypto/rand"
	"crypto/sha256"
	"errors"
	"log"
	"golang.org/x/crypto/pbkdf2"
)


var Key []byte


var Code []byte=[]byte("hello world")

func init() {
	Key=GetKey(Code)
}

var salt []byte=[]byte("asdfghjk")

func GetKey(password []byte) []byte {
	key:=pbkdf2.Key(password,salt,512,16,sha256.New)
	//key:=KKey(password,salt,512,16,sha256.New)
	return key
}



func Encrypt(key,data []byte)(encrypted []byte, err error){
	block,err:=aes.NewCipher(key)
	if err!=nil{
		log.Println(err.Error())
		return nil,err
	}
	aesgcm,err:=cipher.NewGCM(block)
	if err!=nil{
		log.Println(err.Error())
		return nil,err
	}
	iv:=make([]byte,aesgcm.NonceSize())
	rand.Read(iv)
	encrypted=aesgcm.Seal(iv,iv,data,nil)
	//encrypted = append(iv,encrypted... )
	return encrypted,nil
}


func Decrypt(key,encrypted []byte)(data []byte,err error){
	block,err:=aes.NewCipher(key)
	if err != nil{
		log.Println(err.Error())
		return nil,err
	}
	aesgcm,err:=cipher.NewGCM(block)
	if err != nil{
		log.Println(err.Error())
		return nil,err
	}
	if len(encrypted)<=aesgcm.NonceSize(){
		err=errors.New("incorrect encrypted data")
		return nil,err
	}
	iv:=encrypted[:aesgcm.NonceSize()]
	data,err=aesgcm.Open(nil,iv,encrypted[aesgcm.NonceSize():],nil)
	if err != nil{
		return nil,err
	}
	return data,nil
	

}


// func KKey(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
// 	prf := hmac.New(h, password)
// 	hashLen := prf.Size()
// 	fmt.Println(hashLen)
// 	numBlocks := (keyLen + hashLen - 1) / hashLen

// 	var buf [4]byte
// 	dk := make([]byte, 0, numBlocks*hashLen)
// 	U := make([]byte, hashLen)
// 	for block := 1; block <= numBlocks; block++ {
// 		// N.B.: || means concatenation, ^ means XOR
// 		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
// 		// U_1 = PRF(password, salt || uint(i))
// 		prf.Reset()
// 		prf.Write(salt)
// 		buf[0] = byte(block >> 24)
// 		buf[1] = byte(block >> 16)
// 		buf[2] = byte(block >> 8)
// 		buf[3] = byte(block)
// 		prf.Write(buf[:4])
// 		dk = prf.Sum(dk)
// 		T := dk[len(dk)-hashLen:]
// 		copy(U, T)

// 		// U_n = PRF(password, U_(n-1))
// 		for n := 2; n <= iter; n++ {
// 			prf.Reset()
// 			prf.Write(U)
// 			U = U[:0]
// 			U = prf.Sum(U)
// 			for x := range U {
// 				T[x] ^= U[x]
// 			}
// 		}
// 	}
// 	return dk[:keyLen]
// }
