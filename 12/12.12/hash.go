package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

// sha 和md5是hash算法
func hash() {
	md5er := md5.New()
	hasher := sha1.New()
	//计算 SHA1 校验值
	//通过 io.WriteString 或 hasher.Write 将给定的 []byte 附加到当前的 hash.Hash 对象中
	io.WriteString(hasher, "test")
	md5er.Write([]byte("test"))
	//hasher.Write([]byte("test"))
	b := []byte{}
	//
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	hasher.Write(data)
}

// AES-GCM加密解密
func encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("创建AES失败:%w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("创建GCM失败:%w", err)
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err := rand.Read(nonce); err != nil {
		//if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("生成nonce失败%w", err)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
func decrypt(encodedCipher string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCipher)
	if err != nil {
		return "", fmt.Errorf("Base64 解码失败 %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("创建AES 失败 %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("创建AES失败 %w", err)
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("密文长度不正确")
	}
	nonce := ciphertext[:nonceSize]
	encryptedData := ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", fmt.Errorf("解密失败或密文被修改: %w", err)
	}
	return string(plaintext), nil
}
func main() {
	key := []byte("12345678901234567890123456789012")
	original := "we shall overcome"
	encrypted, err := encrypt(original, key)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("原始数据：", original)
	fmt.Println("加密结果：", encrypted)

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("解密结果：", decrypted)
}
