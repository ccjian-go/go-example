package main

import (
	"fmt"
	"unicode/utf8"
)

/**
	一. Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。
	二. 由于字符串是一个字节切片，所以我们可以获取字符串的每一个字节
		打印出来的字符是 "Hello World" 以 Unicode UTF-8 编码的结果
		来理解一下什么是 Unicode 和 UTF-8
		https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/
	三. %c 格式限定符用于打印字符串的字符
		%x 格式限定符用于打印字符串的 Unicode UTF-8编码
	四. 获取字符串的每一个字符，虽然看起来是合法的，但却有一个严重的 bug
	五.	分割 Señor 就出现错误，这是因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。
		它的 UTF-8 编码占用了 c3 和 b1 两个字节。它的 UTF-8 编码占用了两个字节 c3 和 b1。
		而我们打印字符时，却假定每个字符的编码只会占用一个字节，这是错误的。
		在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。
		使用rune解决
	六. rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。
		代码点无论占用多少个字节，都可以用一个 rune 来表示
	七. %d 格式限定符打印数字类型
	八. 字节切片构造字符串，[]byte{...} +  UTF-8 编码后的 16 进制字节或10进制 均可运行
	九. 用 rune 切片构造字符串，[]rune{...} +  UTF-8 编码后的 16 进制字节或10进制 均可运行
	十. 使用 utf8.RuneCountInString(s) 获取字符串长度
		utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度。
		这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。
	十一. 为了修改字符串，可以把字符串转化为一个 rune 切片。
		然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。

*/
func main() {
	name := "Hello World"
	fmt.Println(name)
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

	name2 := "Señor"
	printBytes(name2)
	fmt.Printf("\n")
	printChars(name2)
	fmt.Printf("\n")
	printChars2(name2)
	fmt.Printf("\n")
	printCharsAndBytes(name2)

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	byteSlice2 := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str2 := string(byteSlice2)
	fmt.Println(str2)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str3 := string(runeSlice)
	fmt.Println(str3)

	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	h := "hello"
	fmt.Println(mutate([]rune(h)))
}

func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%c ",s[i])
	}
}

func printChars2(s string) {
	runes := []rune(s)
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ",runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, r := range s {
		fmt.Printf("%c starts at byte %d\n", r, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune)string {
	s[0] = 'a'//any valid unicode character within single quote is a rune
	return string(s)
}
