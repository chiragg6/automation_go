package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fix_automater/server"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

// var app = &cli.App{
// 	Name:  "boom",
// 	Usage: "make an explosive entrance",
// 	Action: func(c *cli.Context) error {
// 		fmt.Println("boom! I say!")
// 		return nil
// 	},
// }

// var author []string{"Chirag"}

func info() {
	app.Name = "Automatic healing system"
	app.Usage = "fix fault/failure in system"
	// app.Authors = []author
	app.Version = "beta"
}

func commands() {
	app.Commands = []cli.Commands{
		{
			Name:    "runserver",
			Aliases: []string{"s"},
			Usage:   "Starts the server",
			Action: func(c *cli.Context) {
				server.StartServer()
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Database migration",
			Action: func(c *cli.Context) {
				storage.Migration()
				fmt.Println("Migrating...")
			},
		},
		{
			Name:    "downonestep",
			Aliases: []string{"dw1"},
			Usage:   "Database roll back",
			Action: func(c *cli.Context) {
				storage.DownOneStep()
				fmt.Println("Version Rolled back by 1 step...")
			},
		},
		{
			Name:    "startsurgeon",
			Aliases: []string{"surgeon"},
			Usage:   "Surgeon starts long polling",
			Action: func(c *cli.Context) {
				fmt.Println("Started Surgeon...")
				surgeon.LongPolling()
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
