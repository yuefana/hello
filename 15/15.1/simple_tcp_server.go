package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
)

const maxRead = 25

func s_t_s() {
	flag.Parse()
	if flag.NArg() != 2 {
		host := flag.Arg(0)
		port := flag.Arg(1)
		//hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
		address := net.JoinHostPort(host, port)
		listener, err := net.Listen("tcp", address)
		if err != nil {
			fmt.Println("建立连接失败")
		}
		fmt.Printf("服务器正在监听%s\n", listener.Addr())

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		go func() {
			<-ctx.Done()
			fmt.Println("收到退出信号,停止接受连接")
			if err := listener.Close(); err != nil {
				fmt.Println("关门监听器失败", err.Error())
			}
		}()
		for {
			conn, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					break
				}
				fmt.Println("接受连接失败", err.Error())
				continue
			}
			go connectionHandler(conn)
		}
		fmt.Println("服务器已经停止")
	}
}
func connectionHandler(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("客户端已连接", remoteAddr)
	defer conn.Close()
	if _, err := io.WriteString(conn, "let's go\n"); err != nil {
		fmt.Println("写入数据失败")
		return
	}

	buf := make([]byte, maxRead)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			fmt.Println(buf[:n])
		}
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("读取数据失败")
			return
		}
	}
}

var ErrNotFound = errors.New("数据不存在")

type QueryError struct {
	Key string
	Err error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("查询 %q 失败：%v", e.Key, e.Err)
}

func (e *QueryError) Unwrap() error {
	return e.Err
}
func findData() error {
	return &QueryError{Key: "user-1001", Err: ErrNotFound}
}

func service() error {
	err := findData()
	if err != nil {
		return fmt.Errorf("业务处理失败:%w", err)
	}
	return nil
}

func as_is() {
	err := service()
	fmt.Println("完整错误", err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Is:错误原因数据暴怒存在")
	}
	var queryErr *QueryError
	if errors.As(err, &queryErr) {
		fmt.Println("as find queryerr")
		fmt.Println("查询建", queryErr.Key)
		fmt.Println("内部错误：", queryErr.Err)
	}

}
