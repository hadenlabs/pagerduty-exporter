package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Flags names
const (
	logLevelFlagName = "logLevel"
)

// NewCLI returns a new CLI
// it contains basic info such as version, authors, etc..
// but also wraps subcommands like bw and op
func NewCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "Pagerduty Exporter"
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  logLevelFlagName,
			Value: log.InfoLevel.String(),
			Usage: "Set log level (debug|info|warn|error)",
		},
	}

	app.Before = func(c *cli.Context) (err error) {
		logLevel, err := log.ParseLevel(c.String(logLevelFlagName))
		if err != nil {
			return fmt.Errorf("pagerduty-exporter error: log level %s is unacceptable", c.String(logLevelFlagName))
		}
		log.SetLevel(logLevel)
		return err
	}

	app.Commands = []cli.Command{}

	return app
}
