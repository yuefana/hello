package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

const (
	typeInt    byte = 1
	typeString byte = 2
	typeBool   byte = 3
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

func EncodeStruct(w io.Writer, data any) error {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return errors.New("不能编码 nil指针")
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("只能编码结构体，收到的是 %s", value.Kind())
	}
	valueType := value.Type()

	fieldCount := uint32(0)

	for i := 0; i < value.NumField(); i++ {
		structField := valueType.Field(i)
		if structField.PkgPath != "" {
			continue
		}
		switch value.Field(i).Kind() {
		case reflect.Int, reflect.String, reflect.Bool:
			fieldCount++
		}
	}

	if err := binary.Write(w, binary.BigEndian, fieldCount); err != nil {
		return err
	}

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		structField := valueType.Field(i)
		if structField.PkgPath != "" {
			continue
		}
		switch fieldValue.Kind() {
		case reflect.Int:
			if err := writeString(w, structField.Name); err != nil {
				return err
			}
			if err := binary.Write(w, binary.BigEndian, typeInt); err != nil {
				return err
			}
			if err := binary.Write(w, binary.BigEndian, fieldValue.Int()); err != nil {
				return err
			}
		case reflect.String:
			if err := writeString(w, structField.Name); err != nil {
				return err
			}
			if err := binary.Write(w, binary.BigEndian, typeString); err != nil {
				return err
			}
			if err := writeString(w, fieldValue.String()); err != nil {
				return err
			}
		case reflect.Bool:
			if err := writeString(w, structField.Name); err != nil {
				return err
			}
			if err := binary.Write(w, binary.BigEndian, typeBool); err != nil {
				return err
			}
			if err := binary.Write(w, binary.BigEndian, fieldValue.Bool()); err != nil {
				return err
			}

		}
	}
	return nil
}

func DecodeStruct(r io.Reader, target any) error {
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return errors.New("解码目标必须是指针")
	}
	value = value.Elem()

	if value.Kind() != reflect.Struct {
		return errors.New("解码目标必须指向结构体")
	}

	var fieldCount uint32
	if err := binary.Read(r, binary.BigEndian, &fieldCount); err != nil {
		return err
	}

	for i := uint32(0); i < fieldCount; i++ {
		fieldName, err := readString(r)
		if err != nil {
			return err
		}
		var encodedType byte
		if err := binary.Read(r, binary.BigEndian, &encodedType); err != nil {
			return err
		}
		targetField := value.FieldByName(fieldName)

		switch encodedType {
		case typeInt:
			var number int64
			if err := binary.Read(r, binary.BigEndian, &number); err != nil {
				return err
			}
			if !targetField.IsValid() || !targetField.CanSet() {
				continue
			}
			switch targetField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:

				if targetField.OverflowInt(number) {
					return fmt.Errorf("字段 %s 的值 %d 超出 %s 范围", fieldName, number, targetField.Type())
				}

				targetField.SetInt(number)
			}
		case typeString:
			text, err := readString(r)
			if err != nil {
				return err
			}
			if !targetField.IsValid() || !targetField.CanSet() {
				continue
			}
			if targetField.Kind() == reflect.String {
				targetField.SetString(text)
			}
		case typeBool:
			var encodedBool byte
			if err := binary.Read(r, binary.BigEndian, &encodedBool); err != nil {
				return err
			}
			if targetField.Kind() == reflect.Bool {
				targetField.SetBool(encodedBool != 0)
			}
		}
	}
	return nil
}

func writeString(w io.Writer, test string) error {
	data := []byte(test)
	length := uint32(len(data))
	if err := binary.Write(w, binary.BigEndian, length); err != nil {
		return err
	}
	_, err := w.Write(data)
	return err
}

func readString(r io.Reader) (string, error) {
	var length uint32
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return "", err
	}
	data := make([]byte, length)

	if _, err := io.ReadFull(r, data); err != nil {
		return "", err
	}
	return string(data), nil
}

type Sender struct {
	Name    string
	Age     int
	Enabled bool
	Address string
	private string
}

type Receiver struct {
	Age     int64
	Name    string
	Enabled bool
}

func main() {
	source := Sender{
		Name:    "Laura",
		Age:     20,
		Enabled: true,
		Address: "Beijing",
		private: "不会被编码",
	}

	var buffer bytes.Buffer

	if err := EncodeStruct(&buffer, source); err != nil {
		fmt.Println("编码失败：", err)
		return
	}

	fmt.Println("编码后的字节数：", buffer.Len())

	var target Receiver

	if err := DecodeStruct(&buffer, &target); err != nil {
		fmt.Println("解码失败：", err)
		return
	}

	fmt.Printf("解码结果：%+v\n", target)
}
