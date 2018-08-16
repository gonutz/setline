package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		panic(`expecting two arguments:
  the line index (starting at 0)
  the new line to set`)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("first argument must be the line index (starting at 0)")
	}
	newLine := []byte(os.Args[2])
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("unable to read stdin: " + err.Error())
	}
	lines := bytes.Split(data, []byte{'\n'})
	if 0 <= i && i < len(lines) {
		line := lines[i]
		if bytes.HasSuffix(line, []byte{'\r'}) {
			newLine = append(newLine, '\r')
		}
		lines[i] = newLine
	}
	output := bytes.Join(lines, []byte{'\n'})
	fmt.Printf("%s", output)
}
