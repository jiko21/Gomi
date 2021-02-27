package main

import (
	"jiko21/gomi/git"
	"jiko21/gomi/initializer"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GOMI"
	app.Usage = "Branch delete tool made by Golang"
	app.Version = "0.2.6"
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 1 && c.Args().Get(0) == "init" {
			initObj, err := initializer.New()
			if err != nil {
				return err
			}
			return initObj.Exec()
		} else if c.NArg() == 0 {
			gitInst := git.ConstructGit(".gomiignore")
			return gitInst.Delete()
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
