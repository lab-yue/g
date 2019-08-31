package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rainy-me/g/shortcut"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "g"
	app.Usage = "git command shortcut"

	app.Action = func(c *cli.Context) error {

		cmd := exec.Command("git")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		args := []string{}

		for _, arg := range c.Args() {
			if val, ok := shortcut.Map[arg]; ok {
				args = append(args, val...)
			} else {
				args = append(args, arg)
			}
		}

		cmd.Args = append(cmd.Args, args...)

		fmt.Println(cmd.Args)
		cmd.Run()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
