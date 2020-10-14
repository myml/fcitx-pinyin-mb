package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("gbkpy.org")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(b, []byte{'\n'})
	m := make(map[string]string)
	for i := range lines {
		fields := bytes.Split(lines[i], []byte{' '})
		if len(fields) != 2 {
			continue
		}
		m[string(fields[1])] = string(fields[0])
	}
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		vals := strings.Split(scan.Text(), "")
		var keys []string
		for i := range vals {
			if _, ok := m[vals[i]]; !ok {
				keys = []string{}
				break
			}
			keys = append(keys, m[vals[i]])
		}
		if len(keys) > 0 {
			fmt.Println(strings.Join(keys, "'"), strings.Join(vals, ""))
		}
	}
}
