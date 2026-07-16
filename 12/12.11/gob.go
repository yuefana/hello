package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

/*
Gob 通常用于远程方法调用  应用程序和机器之间的数据传输
Gob 特定地用于纯 Go 的环境中  编码和解码用到了go的反射
*/
type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

func gob1() {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	decoder := gob.NewDecoder(&buf)
	err := encoder.Encode(P{3, 4, 5, "P"})
	if err != nil {
		log.Fatal("encode error", err)
	}
	var q Q
	err = decoder.Decode(&q)
	if err != nil {
		log.Fatal("encode error", err)
	}
	fmt.Println(*q.X, *q.Y, q.Name)
}

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func gob2() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	//文件
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(vc)

	file, _ = os.Open("vcard.gob")
	defer file.Close()
	var vcard VCard
	decoder := gob.NewDecoder(file)
	decoder.Decode(&vcard)
	fmt.Print(vcard.FirstName, vcard.LastName, vcard.Addresses, vcard.Remark)
}

func main() {
	gob2()
}
