package main

import (
	"log"
	"os"

	commands "github.com/AdhityaRamadhanus/checkupd/commands/checklist"
	"github.com/AdhityaRamadhanus/checkupd/config"
	"github.com/urfave/cli"
)

func cmdNotFound(c *cli.Context, command string) {
	log.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}

func main() {
	// Init Config
	config.SetDefaultConfig()
	app := cli.NewApp()
	app.Name = "checklist"
	app.Author = "Adhitya Ramadhanus"
	app.Email = "adhitya.ramadhanus@gmail.com"

	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound // Inspired by docker machine
	app.Usage = "Checkup server cli "
	app.Version = "1.0.0"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
