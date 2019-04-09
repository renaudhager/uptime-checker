package main

import (
	"os"
	"uptime-checker/check"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var version string

// main function
func main() {
	app := cli.NewApp()
	app.Name = "uptime-checker"
	app.Usage = "Check uptime, if its above the limit, creates a file."
	app.HelpName = "CLI tool to check if uptime is above a specific threshold and touch a file if uptime is above the threshold."
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "l, log-level",
			Usage:  "Setting log level.",
			EnvVar: "UC_LOGLEVEL",
			Value:  "INFO",
		},
		cli.StringFlag{
			Name:   "lf, log-format",
			Usage:  "Setting log format. Value [text|json].",
			EnvVar: "UC_LOGLEVEL",
			Value:  "json",
		},
	}

	app.Commands = []cli.Command{
		check.CheckCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
