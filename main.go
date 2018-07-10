package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TODO: Custom separator
func main() {
	scanners := make(map[*bufio.Scanner]bool)
	for _, v := range os.Args[1:] {
		f, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanners[NewCustomScanner(f)] = false
	}

	for {
		var allDone = true
		for k := range scanners {
			var ok bool
			if ok = k.Scan(); ok {
				fmt.Println(k.Text())
			}
			allDone = allDone && !ok
		}
		if allDone {
			break
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
