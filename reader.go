// Reader implementation using different stdlibs

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const str = "Hey\tUniverse\nYou're\tgonna\tget\truined\n"

func main() {
	scanner()
}

func scanner() {
	// Default delimiter is newline
	// Need to define a split function for custom delimiter-> https://golang.org/pkg/bufio/#Scanner
	sc := bufio.NewScanner(strings.NewReader(str))
	for sc.Scan() {
		fmt.Println(string(sc.Bytes()))
	}

}

func stringsPkg() {
	bufReader := bufio.NewReader(strings.NewReader(str))
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println("strings+bufio -> ", line)
	}
}
func bufioPkg() {
	byt := bytes.NewBuffer([]byte(str))
	bufReader := bufio.NewReader(byt)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println("buffio -> ", line)
	}
}
func bytesPkg() {
	buf := bytes.NewBufferString(str)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println("bytes -> ", line)
	}
}
