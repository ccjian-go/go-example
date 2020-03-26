package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
)

/**
	一. 将整个文件读取到内存
		使用 ioutil 包中的 ReadFile 函数
	二. 使用绝对文件路径
	三. 使用命令行标记来传递文件路径
		通过 String 函数，创建了一个字符串标记，名称是 fpath，默认值是 test.txt，
		描述为 file path to read from。这个函数返回存储 flag 值的字符串变量的地址。
		在程序访问 flag 之前，必须先调用 flag.Parse()
		命令行
			go install filehandling
			wrkspacepath/bin/filehandling -fpath=/path-of-file/test.txt
	四. 将文件绑定在二进制文件中
		安装 packr 包
		packr 会把静态文件（例如 .txt 文件）转换为 .go 文件
	五. 分块读取文件
	六. 逐行读取文件
 */

func main() {
	//data, err := ioutil.ReadFile("test.txt")
	data, err := ioutil.ReadFile("D:\\项目\\go-dir\\go-example\\test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
		}
	fmt.Println("Contents of file:", string(data))

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)

	//box := packr.NewBox("../filehandling")
	box := packr.NewBox("./")
	data2 := box.String("test.txt")
	fmt.Println("Contents of file:", data2)

}