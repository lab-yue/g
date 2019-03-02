# g

shortcuts for git commands

This package just provides extra shortcut options, commands not mentioned would be the same as origin git commands.

e.g. `g diff`

## To install

`$ go get github.com/rainy-me/g`

## Shortcuts map

| command | origin |
| - | - |
| cm | commit -m |
| rao | remote add origin |
| po | push origin |
| pom | push origin master |
| a | add |
| b | branch |
| c | checkout |
| om | origin master |
| m | merge |
| cb | checkout -b |
| db | branch -d |

## To build readme:  

`go build -o dist/doc doc/doc.go`

 `./dist/doc`