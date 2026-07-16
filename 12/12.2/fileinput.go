package main

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
go语言使用*os.File类型来表示一个文件 文件句柄
*/
func fileinput() {
	//程序运行时的当前工作目录
	inputFile, err := os.Open("input.dat")
	if err != nil {
		fmt.Printf("发生了一个错误 \n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 1024)

	for {
		// inputString, readerr := inputReader.ReadString('\n')
		// fmt.Printf("The input was %s", inputString)
		// if readerr == io.EOF {
		// 	return
		// }
		n, err := inputReader.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[0:n]), "字符数", n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("文件读取失败%v", err)
		}
	}
}

func read_write_file1() {
	input := "input.dat"
	output := "output.dat"
	buf, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取文件失败:%v", err)
		return
	}

	fmt.Printf("%s\n", string(buf))
	// 0 6 4 4
	//   110 100 100
	//   写 读 执行
	//文件所有者 同组用户 其他用户
	err = os.WriteFile(output, buf, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败:%v", err)
		return
	}
	fmt.Println("文件复制完成")
}

func read_file2() {
	file, err := os.Open("products.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanf(file, "%q;%q;%q\n", &v1, &v2, &v3)
		//_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func path() {
	path := "D:/Code/A_WorkSpace/vs/hello/products.txt"
	fmt.Println("完整路径", path)
	fmt.Println("所在目录", filepath.Dir(path))
	fmt.Println("文件名字", filepath.Base(path))
	fmt.Println("扩展名", filepath.Ext(path))
	fmt.Println("不含扩展名的文件名", strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
}

type product struct {
	title string
	price float64
	num   int
}

func read_csv() {
	res := make([]product, 0)
	inputfile, err := os.Open("products.txt")
	if err != nil {
		return
	}
	defer inputfile.Close()
	inputReader := bufio.NewReader(inputfile)
	for {
		cur := product{}
		_, err = fmt.Fscanf(inputReader, "%q;%f;%d\n", &cur.title, &cur.price, &cur.num)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取失败", err)
			break
		}
		res = append(res, cur)
	}

	for _, v := range res {
		fmt.Println(v)
	}

}

func read_csv1() {
	res := make([]product, 0)
	inputfile, err := os.Open("products.txt")
	if err != nil {
		return
	}
	defer inputfile.Close()
	inputReader := csv.NewReader(inputfile)
	inputReader.Comma = ';'
	for {
		record, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取失败", err)
			break
		}
		p, _ := strconv.ParseFloat(record[1], 64)
		n, _ := strconv.Atoi(record[2])
		cur := product{record[0], p, n}
		res = append(res, cur)
	}

	for _, v := range res {
		fmt.Println(v)
	}

}

// compress提供读取压缩的能力 支持的格式有 bzip2、flate、gzip、lzw 和 zlib
func gzipped() {
	filename := "MyFile.gz"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建文件失败", err)
		return
	}
	defer file.Close()
	/*
		gzip文件包括 gzip头部 压缩数据 gzip尾部
		gzip尾部由 gzipWriter.Close()进行写入
	*/
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()
	content := `第一行:Hello Go
第二行:这是一个gzip文件
第三行:Windows 可以正常读写`
	_, err = gzipWriter.Write([]byte(content))
	if err != nil {
		fmt.Println("写入压缩数据失败", err)
	}
}

func readGzip() {
	filename := "MyFile.gz"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("", err)
		return
	}
	defer gzipReader.Close()

	input := bufio.NewReader(gzipReader)

	for {
		s, err := input.ReadString('\n')
		if len(s) > 0 {
			fmt.Print(s)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			break
		}

	}
}

func fileoutput() {
	outfile, err := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("打开或创建出现问题", err)
		return
	}
	defer outfile.Close()
	//os.WriteFile(filename,data []bytes,perm)
	//outfile.Write(buf) 创建缓冲区写

	outputstring := "hello go \n"
	//直接写
	outfile.WriteString(outputstring)
	//outfile.Write()
	//缓冲写
	//bufio.Writer 默认缓冲区通常是 4096 字节
	outputwriter := bufio.NewWriter(outfile)
	outputwriter.WriteString(outputstring)

	outputwriter.Flush()
}
func filewrite() {
	os.Stdout.WriteString("hello word\n")
	file, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	file.WriteString("hello world in a file\n")
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return os.WriteFile(p.Title, p.Body, 0o66)
}
func load(tit string) (*Page, error) {
	//打开该文件，并立即把原有内容清空，文件长度变为 0
	//os.Create()
	res, err := os.ReadFile(tit)
	if err != nil {
		fmt.Print("读取失败")
		return nil, fmt.Errorf("读取文件失败%w", err)
	}
	fmt.Println(string(res))
	return &Page{tit, res}, nil
}

func main() {
	page := &Page{"test1.txt", []byte("这是内容\n hello go")}
	page.save()
	load("test1.txt")
}
