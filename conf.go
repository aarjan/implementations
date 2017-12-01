package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

var config = `
	url = example.com
	user = manager
	pass = 1234
	root_user = false
 
`

func main() {
	byt := bytes.NewBuffer([]byte(config))
	bufReader := bufio.NewReader(byt)

	for {
		l, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("err", err)
		}
		l = strings.TrimSpace(l)
		if l == "\n" || l == "\t" || l == "" {
			continue
		}

		i := strings.IndexAny(l, "=:")
		str := l[0:i]
		val := l[i+1:]
		fmt.Printf("str: %s, val: %s\n", str, val)
	}

}
