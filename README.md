# g

shortcuts for git commands

This package just provides extra shortcut options, commands not mentioned would be the same as origin git commands.

e.g. `g diff`

## To install

`$ go get github.com/rainy-me/g`

## Shortcuts map

| command | origin |
| - | - |
| c | checkout |
| cm | commit -m |
| rso | remote set-url origin |
| a | add |
| m | merge |
| om | origin master |
| rao | remote add origin |
| b | branch |
| po | push origin |
| pom | push origin master |
| db | branch -d |
| s | status |
| rgo | remote get-url origin |
| pox | push origin $(git branch --show-current) |
| cb | checkout -b |

## To build readme:  

`go build -o dist/doc doc/doc.go`

 `./dist/doc`