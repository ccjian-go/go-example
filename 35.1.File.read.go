package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"log"
	"os"
)

/**
	一. 将整个文件读取到内存
		使用 ioutil 包中的 ReadFile 函数
	二. 使用绝对文件路径
	三. 使用命令行标记来传递文件路径
		通过 String 函数，创建了一个字符串标记，名称是 fpath，默认值是 learn.txt，
		描述为 file path to read from。这个函数返回存储 flag 值的字符串变量的地址。
		在程序访问 flag 之前，必须先调用 flag.Parse()
		命令行
			go install filehandling
			wrkspacepath/bin/filehandling -fpath=/path-of-file/learn.txt
	四. 将文件绑定在二进制文件中
		安装 packr 包
		packr 会把静态文件（例如 .txt 文件）转换为 .go 文件
	五. 分块读取文件
		我们学习了如何把整个文件读取到内存。当文件非常大时，尤其在 RAM 存储量不足的情况下，把整个文件都读入内存是没有意义的。更好的方法是分块读取文件。这可以使用 bufio 包来完成。
		在上述程序的第 15 行，我们使用命令行标记传递的路径，打开文件。
		在第 19 行，我们延迟了文件的关闭操作。
		在上面程序的第 24 行，我们新建了一个缓冲读取器（buffered reader）。在下一行，我们创建了长度和容量为 3 的字节切片，程序会把文件的字节读取到切片中。
		第 27 行的 Read 方法会读取 len(b) 个字节（达到 3 字节），并返回所读取的字节数。当到达文件最后时，它会返回一个 EOF 错误
	六. 逐行读取文件
		使用 Go 逐行读取文件。这可以使用 bufio 来实现。
		逐行读取文件涉及到以下步骤。
			打开文件；
			在文件上新建一个 scanner；
			扫描文件并且逐行读取。
 */

func main() {
	//data, err := ioutil.ReadFile("learn.txt")
	data, err := ioutil.ReadFile("D:\\项目\\go-dir\\go-example\\learn.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
		}
	fmt.Println("Contents of file:", string(data))

	fptr := flag.String("fpath", "learn.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)

	//box := packr.NewBox("../filehandling")
	box := packr.NewBox("./")
	data2 := box.String("learn.txt")
	fmt.Println("Contents of file:", data2)


    fptr2 := flag.String("fpath2", "learn.txt", "file path to read from")
    flag.Parse()
    f, err := os.Open(*fptr2)
    if err != nil {
    	log.Fatal(err)
    }
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
	    _, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			    break
			}
		fmt.Println(string(b))
	}


	fptr3 := flag.String("fpath3", "learn.txt", "file path to read from")
	flag.Parse()
	f2, err := os.Open(*fptr3)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f2.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f2)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}