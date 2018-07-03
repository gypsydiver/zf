package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	file1 = flag.String("file1", "", "first file (leading)")
	file2 = flag.String("file2", "", "second file (trailing)")
)

func main() {
	flag.Parse()

	f1, err := os.Open(*file1)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open(*file2)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	scanner1.Buffer([]byte{}, 1024*1024)

	scanner2 := bufio.NewScanner(f2)
	scanner2.Buffer([]byte{}, 1024*1024)

	for scanner1.Scan() {
		fmt.Println(scanner1.Text())
		if scanner2.Scan() {
			fmt.Println(scanner2.Text())
		}
	}
}
