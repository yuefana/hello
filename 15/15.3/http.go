package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", helloHandler)
	mux.HandleFunc("GET /shouthello/{name}", shouthelloHandler)
	server := &http.Server{
		Addr:              "localhost:9999",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	//http.ListenAndServeTLS()
	//http.ListenAndServe("localhost:8080", mux)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Http 服务器停止", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	//r.FormValue()
	// if err := r.ParseForm(); err != nil {
	// 	http.Error(w, "表单格式错误", http.StatusBadRequest)
	// 	return
	// }
	// values, found := r.Form["var1"]
	//可能得到 []string{"go", "http"}
	// if !found || len(values) == 0 {
	// 	http.Error(w, "缺少字段", http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println(w, values[0])
	// fmt.Println(name)
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	if _, err := fmt.Fprintf(w, "hello,%s!\n", name); err != nil {
		log.Printf("写入响应失败:%v", err)
	}
}
func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	if _, err := fmt.Fprintf(w, "hello,%s!\n", name); err != nil {
		log.Printf("写入响应失败:%v", err)
	}
}

type hello struct {
}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func hello_Server() {
	http.ListenAndServe(":8080", &hello{})
}

var urls = []string{
	"https://example.com/",
}

func main() {
	//xmltest()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for {
		fmt.Println("请输入URL")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.EqualFold(input, "Q") {
			fmt.Println("退出")
			return
		}
		if strings.EqualFold(input, "") {
			continue
		}
		if err := fetch(client, input); err != nil {
			log.Fatal(err)
		}
	}
	// if len(os.Args) != 2 {
	// 	log.Fatalf("用法%s <xml-url>", os.Args[0])
	// }

	// user, err := fetchUser(client, os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("状态%s\n", user.Status.Text)
}

const maxBodySize = 1 << 20 //1MB

func fetch(client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("请求%s失败:%w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("服务器返回异常状态:%s", resp.Status)
	}
	// 多读取一个字节，用来判断正文是否超出限制。
	//io.Copy(os.Stdout, resp.Body)
	data, err := io.ReadAll(io.LimitReader(resp.Body, maxBodySize+1))
	if err != nil {
		return fmt.Errorf("读取响应正文失败:%w", err)
	}
	if len(data) > maxBodySize {
		return fmt.Errorf("响应正文超过%d字节", maxBodySize)
	}
	fmt.Printf("%s\n", data)
	return nil
}

func pollURL(client *http.Client, url string) {
	resp, err := client.Head(url)
	if err != nil {
		log.Printf("访问%s失败:%v", url, err)
		return
	}
	defer resp.Body.Close()
	// decoder:=xml.NewDecoder(resp.Body)
	fmt.Printf("%s:%s\n", url, resp.Status)
}

type Status struct {
	Text string `xml:"text"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Status  Status   `xml:"status"`
	ID      int      `xml:"id,attr"`
}

func fetchUser(client *http.Client, url string) (*User, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求XML失败:%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("服务器返回异常状态:%s", resp.Status)
	}
	var user User
	decoder := xml.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		return nil, fmt.Errorf("解析XML失败:%w", err)
	}
	return &user, nil
}

func xmltest() {
	data := []byte(`
	<user id="1001">
			<status>
				<text>Hello Go</text>
			</status>
		</user>
	`)
	var user User

	decode := xml.NewDecoder(bytes.NewReader(data))
	decode.Decode(&user)

	// if err := xml.Unmarshal(data, &user); err != nil {
	// 	fmt.Println("解析失败", err)
	// }
	fmt.Printf("XMLName:%v\n", user.XMLName)
	fmt.Printf("ID:%d\n", user.ID)
	fmt.Printf("状态%s\n", user.Status.Text)
}
