package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

const formPage = `<!doctype html>
<html lang="zh-CN">
<head>
	<meta charset="utf-8">
	<title>表单示例</title>
</head>
<body>
	<h1>输入表单</h1>

	<form action="/test2" method="post">
		<label for="input">请输入内容：</label>

		<input
			id="input"
			type="text"
			name="in"
			value="{{.Input}}"
		>

		<button type="submit">提交</button>
	</form>

	{{if .Error}}
		<p>错误：{{.Error}}</p>
	{{end}}

	{{if .Submitted}}
		<p>你输入的是：{{.Input}}</p>
	{{end}}
</body>
</html>
`

type formData struct {
	Input     string
	Submitted bool
	Error     string
}

var formTemplate = template.Must(
	template.New("from").Parse(formPage),
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				recovered := recover()
				if recovered == nil {
					return
				}

				if recovered == http.ErrAbortHandler {
					panic(recovered)
				}

				log.Printf(
					"处理请求时发生panic:method=%s path=%s remotr=%s panic=%v\n%s",
					r.Method,
					r.URL.Path,
					r.RemoteAddr,
					recovered,
					debug.Stack(),
				)

				http.Error(w, "服务器内部错误", http.StatusInternalServerError)
			}()

			next.ServeHTTP(w, r)
		},
	)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)
	if _, err := w.Write([]byte("<!doctype html><h1>Hello, world!</h1>")); err != nil {
		log.Printf("写入响应失败:%v", err)
	}
}
func showForm(w http.ResponseWriter, r *http.Request) {
	renderForm(w, formData{}, http.StatusOK)
}
func submitForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		renderForm(w, formData{Error: "表单格式错误"}, http.StatusBadRequest)
		return
	}
	input := strings.TrimSpace(r.PostForm.Get("in"))
	if input == "" {
		renderForm(w, formData{Error: "请输入内容"}, http.StatusBadRequest)
		return
	}
	renderForm(w, formData{Input: input, Submitted: true}, http.StatusOK)
}
func renderForm(w http.ResponseWriter, data formData, statusCode int) {
	var buf bytes.Buffer

	if err := formTemplate.Execute(&buf, data); err != nil {
		log.Printf("执行模板失败:%v", err)
		http.Error(w, "生成页面失败", http.StatusInternalServerError)
		return
	}
	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)
	w.WriteHeader(statusCode)
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("写入响应失败:%v", err)
	}
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("用于测试的panic")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /test1", simpleHandler)
	mux.HandleFunc("GET /test2", showForm)
	mux.HandleFunc("POST /test2", submitForm)
	mux.HandleFunc("GET /panic", panicHandler)
	handler := recoverMiddleware(mux)

	log.Println("服务器正在监听:http://localhost:8088")

	if err := http.ListenAndServe(":8088", handler); err != nil {
		log.Fatal("HTTP服务器停止:", err)
	}
}
