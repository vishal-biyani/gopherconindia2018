package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "golang-kube"
	app.Usage = "Using client-go effectively with Kubernetes api"
	app.Version = "0.1"

	exampleSubCommands := []cli.Command{
		{Name: "one", Usage: "Run example one.go", Action: funcOne},
		{Name: "two", Usage: "Run example two.go", Action: funcTwo},
		{Name: "three", Usage: "Run example three.go", Action: funcThree},
		{Name: "four", Usage: "Run example three.go", Action: funcFour},
	}

	app.Commands = []cli.Command{
		{Name: "run", Usage: "Run examples", Subcommands: exampleSubCommands},
	}

	app.Run(os.Args)
}
