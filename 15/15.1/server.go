package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
)

func checkError(e error) {
	if e != nil {
		panic(fmt.Errorf("出现异常,异常原因: %w", e))
	}
}

var (
	listener net.Listener
	userMu   sync.RWMutex
	//用户当前的连接数量
	users = make(map[string]int)

	shutdownOnce sync.Once
)

func main() {
	as_is()
}

func serve() {
	fmt.Println("Starting the server ...")

	var err error

	listener, err = net.Listen("tcp", "localhost:50000")
	checkError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				fmt.Println("服务器已经关闭")
				return
			}
			fmt.Println("接受连接失败", err)
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	if !scanner.Scan() {
		return
	}
	clientName := strings.TrimSpace(scanner.Text())
	if clientName == "" {
		return
	}
	setClientActive(clientName, true)
	defer setClientActive(clientName, false)
	fmt.Printf("User %s connected\n", clientName)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		switch {
		case strings.EqualFold(text, "WHO"):
			displayClientList()
		case strings.EqualFold(text, "SH"):
			fmt.Println("即将关闭服务器端")
			shutDownServer()
			return
		default:
			fmt.Printf("Received data %s says %s\n", clientName, text)
		}
		// Scanner 遇到正常 EOF 时，Err() 返回 nil。
		if err := scanner.Err(); err != nil {
			fmt.Printf("读取客户端 %s 失败：%v\n", clientName, err)
		}
	}
}

func displayClientList() {
	userMu.RLock()

	names := make([]string, 0, len(users))
	status := make(map[string]int, len(users))

	for name, connCount := range users {
		names = append(names, name)
		if connCount > 0 {
			status[name] = 1
		} else {
			status[name] = 0
		}
	}

	userMu.RUnlock()
	//map遍历顺序不固定,排序输出
	sort.Strings(names)

	fmt.Println("This is the client list: 1:active, 0=inactive")
	for _, name := range names {
		fmt.Printf("User %s is %d\n", name, status[name])
	}

}

func setClientActive(name string, active bool) {
	userMu.Lock()
	defer userMu.Unlock()

	if active {
		users[name]++
		return
	}
	if users[name] > 0 {
		users[name]--
	}
}

func shutDownServer() {
	shutdownOnce.Do(func() {
		_ = listener.Close()
	})
}

func s() string {
	var b bytes.Buffer
	b.WriteString("hello")
	return b.String()
}

//15
//17 18 19     5，6
//20 21     7
