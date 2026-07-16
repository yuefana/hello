package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

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

func t1() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"y", "f", []*Address{pa, wa}, ""}
	//转换为json数据
	//js, _ := json.Marshal(vc)
	//fmt.Printf("json format is %s", js)
	f, _ := os.OpenFile("vcaed.json", os.O_RDONLY|os.O_CREATE, 0666)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	//将对象vc转换为json编码并写入
	if err := enc.Encode(vc); err != nil {
		log.Println("Error in encoding json")
	}

}

func t2() {
	// number -> float64
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f any
	json.Unmarshal(b, &f)
	m := f.(map[string]any)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []any:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			//reflect.ValueOf(&vv).Elem().Set(reflect.ValueOf(12.0))
			fmt.Println(k, "is of a type I don’t know how to handle", reflect.TypeOf(vv), vv)

		}

	}
}
func main() {
	t1()
}
