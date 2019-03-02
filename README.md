# g

shortcuts for git commands

This package just provides extra shortcut options, commands not mentioned would be the same as origin git commands.

e.g. `g diff`

## To install

`$ go get github.com/rainy-me/g`

## Shortcuts map

| command | origin |
| - | - |
| a | add |
| om | origin master |
| pom | push origin master |
| cb | checkout -b |
| b | branch |
| m | merge |
| cm | commit -m |
| rao | remote add origin |
| po | push origin |

## To build readme:  

`go build -o dist/doc doc/doc.go`

 `./dist/doc`