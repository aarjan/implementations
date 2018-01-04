// A alternative implementation of github.com/aarjan/aaml by reading configuration files line-by-line

package main

import (
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
	bufReader := bytes.NewBufferString(config)
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
