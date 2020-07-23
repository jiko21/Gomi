package main

import (
	"jiko21/gomi/git"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GOMI"
	app.Usage = "Branch delete tool made by Golang"
	app.Version = "0.2.1"
	app.Action = func(c *cli.Context) error {
		gitInst := git.ConstructGit(".gomiignore")
		gitInst.Delete()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
