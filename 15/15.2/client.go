package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(fmt.Errorf("出现异常,异常原因: %w", e))
	}
}

func client() {
	//客户端所有操作位于main goroutine 可以捕获panic
	defer func() {
		if v, ok := recover().(error); ok {
			fmt.Println(v)
		}
	}()

	conn, err := net.Dial("tcp", "localhost:50000")
	checkError(err)
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First,what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)
	if clientName == "" {
		fmt.Println("用户名不能为空")
		return
	}
	_, err = fmt.Fprintln(conn, clientName)
	checkError(err)
	for {
		fmt.Println("输入消息;WHO 查看用户;SH 关闭服务器;Q 退出:")
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "Q" {
			fmt.Println("quit success")
			return
		}
		if input == "" {
			continue
		}
		//_, err = conn.Write([]byte(clientName + " say:" + input))
		_, err = fmt.Fprintln(conn, input)
		checkError(err)
		if strings.EqualFold(input, "SH") {
			fmt.Println("服务器端已经下线")
			return
		}
	}
}

func main() {
	client()
}

func dial() {
	checkConnection("TCP IPv4", "tcp4", "127.0.0.1:50000")
	checkConnection("TCP IPv6", "tcp6", "[::1]:50000")
	checkConnection("UDP IPv4", "udp4", "127.0.0.1:50000")
}

func checkConnection(name, network, address string) {
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Printf("%s 连接失败 %v\n", name, err)
		return
	}
	defer conn.Close()
	fmt.Printf("%s 连接成功\n", name)
	fmt.Printf("本地地址：%v\n", conn.LocalAddr())
	fmt.Printf("远程地址：%v\n\n", conn.RemoteAddr())
}

func socket() {
	const host = "www.apache.org"
	//443 是https的默认端口
	con, err := tls.Dial(
		"tcp",
		host+":443",
		&tls.Config{ServerName: host, MinVersion: tls.VersionTLS12},
	)
	if err != nil {
		fmt.Println("TSL连接失败")
		return
	}
	defer con.Close()
	request := "GET / HTTP/1.1\r\n" +
		"Host: " + host + "\r\n" +
		"Connection: close\r\n" +
		"\r\n"
	io.WriteString(con, request)
	// for {
	// 	n, err := con.Read(data)
	// 	if n > 0 {
	// 		fmt.Print(string(data[0:n]))
	// 	}
	// 	if err != nil {
	// 		return
	// 	}
	// }
	io.Copy(os.Stdout, con)
}
func socket1() {
	resp, err := http.Get("http://www.apache.org/")
	if err != nil {
		fmt.Println("请求失败", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("最终状态", resp.Status)
	fmt.Println("最终地址", resp.Request.URL)
	io.Copy(os.Stdout, resp.Body)
}
