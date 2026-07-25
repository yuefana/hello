package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
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
		<input id="input" type="text" name="in" value="{{.Input}}">
		<button type="submit">提交</button>
	</form>

	{{if .Submitted}}
		<p>
			NUMBER:
			{{range .Result.Number}}
				  {{.}}
			{{end}}
			{{range $index, $number := .Result.Number}}
					{{if $index}}, {{end}}
					{{$number}}
			{{end}}
		</p>
		<p>COUNT:{{.Result.Count}}</p>
		<p>MEAN:{{.Result.Mean}}</p>
		<p>MEDIAN:{{.Result.Median}}</p>
		
	{{end}}
</body>
</html>
`

type result struct {
	Number []int
	Count  int
	Mean   float64
	Median float64
}

type formData struct {
	Submitted bool
	Result    result
	Input     string
}

var formTemplate = template.Must(
	template.New("form").Parse(formPage),
)

func compute(str string) result {
	var res result
	numstr := strings.Fields(str)
	res.Count = len(numstr)
	if res.Count <= 0 {
		return res
	}
	res.Number = make([]int, res.Count)

	sum := 0.0
	for i, v := range numstr {
		n, _ := strconv.Atoi(v)
		sum += float64(n)
		res.Number[i] = n
		//append要创建len为0的slice
		//res.Number = append(res.Number, n)
	}
	res.Mean = sum / float64(res.Count)
	sorted := append([]int(nil), res.Number...)
	sort.Ints(sorted)
	index := res.Count / 2
	if res.Count&1 == 0 {
		//偶数
		res.Median = (float64(sorted[index-1]) + float64(sorted[index])) / 2.0
	} else {
		res.Median = float64(sorted[index])
	}
	return res
}

func simpleServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)
	if _, err := io.WriteString(
		w,
		"<!doctype html><h1>hello world</h1>",
	); err != nil {
		log.Printf("写入响应失败:%v", err)
	}
}

func showForm(w http.ResponseWriter, r *http.Request) {
	renderForm(w, formData{})
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(
			w,
			"表单格式错误",
			http.StatusBadRequest,
		)
		return
	}
	input := r.PostForm.Get("in")
	renderForm(w, formData{
		Submitted: true,
		Input:     input,
		Result:    compute(input),
	})
}

func renderForm(w http.ResponseWriter, data formData) {
	var buf bytes.Buffer
	if err := formTemplate.Execute(&buf, data); err != nil {
		log.Printf("执行模板失败:%v", err)
		http.Error(
			w,
			"生成页面失败",
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set(
		"Content-Type",
		"text/html; charset=utf-8",
	)

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("写入响应失败:%v", err)
	}
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /test1", simpleServer)
	mux.HandleFunc("GET /test2", showForm)
	mux.HandleFunc("POST /test2", submitForm)

	log.Println("服务器正在监听:http://localhost:8088")

	log.Fatal(http.ListenAndServe(":8088", mux))
}
