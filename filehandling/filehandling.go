package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr2 := flag.String("fpath", "learn.txt", "file path to read from")
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