package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// TODO: N-way file zip
// TODO: Custom separator
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

	scanner1 := NewCustomScanner(f1)
	scanner2 := NewCustomScanner(f2)

	// assumes file 1 is bigger
	// TODO: flush contents of file two
	for scanner1.Scan() {
		fmt.Println(scanner1.Text())
		if scanner2.Scan() {
			fmt.Println(scanner2.Text())
		}
	}
}

// NewCustomScanner sets maxTokenSize to 1024 * 1024
// default is 64 * 1024
func NewCustomScanner(r io.Reader) (scn *bufio.Scanner) {
	scn = bufio.NewScanner(r)
	scn.Buffer([]byte{}, 1024*1024)
	return
}
