# g

### This this created as a way to learn go, and had no reason to implement in go at all. Maybe you want to checkout this [shell script](https://github.com/rainy-me/dotfiles/blob/master/scripts/g)

---

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