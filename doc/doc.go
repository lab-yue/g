package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/rainy-me/g/shortcut"
)

func main() {

	out := `# g

## git commands shortcuts

This package just provides extra shortcut options, commands not mentioned would be the same as origin git command.

e.g. ` + "`g diff`" + `

### shortcuts map

| command | origin |
| - | - |
`

	for k, v := range shortcut.Map {
		out += fmt.Sprintf("| %s | %s |\r\n", k, strings.Join(v, " "))
	}

	out += "\r\n### To build readme:  \r\n\r\n`go build -o dist/doc doc/doc.go`\r\n\r\n `./dist/doc`"

	fmt.Println(out)

	ioutil.WriteFile("README.md", []byte(out), 0644)

}
