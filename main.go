package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	TYPE = "tcp"
)

var debug = false
var dryRun = false

func main() {
	var host string
	var port int

	app := &cli.App{
		Name:           "xcopy",
		Description:    "Clipboard remote access",
		DefaultCommand: "client",
		Commands: []*cli.Command{
			{
				Name: "server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "host",
						Aliases:     []string{"l"},
						Value:       "localhost",
						Destination: &host,
					},
					&cli.IntFlag{
						Name:        "port",
						Aliases:     []string{"p"},
						Value:       9001,
						Destination: &port,
					},
				},
				Action: func(c *cli.Context) error {
					err := ServerInit()
					if err != nil {
						return err
					}
					ServerStart(host, port)
					return nil
				},
			},
			{
				Name: "client",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "host",
						Aliases:     []string{"s"},
						Value:       "localhost",
						Destination: &host,
					},
					&cli.IntFlag{
						Name:        "port",
						Aliases:     []string{"p"},
						Value:       9001,
						Destination: &port,
					},
					&cli.BoolFlag{
						Name:        "verbose",
						Aliases:     []string{"v"},
						Value:       false,
						Destination: &debug,
					},
					&cli.BoolFlag{
						Name:        "dry-run",
						Aliases:     []string{"n"},
						Value:       false,
						Destination: &dryRun,
					},
				},
				Action: func(c *cli.Context) error {
					Send(host, port, debug, dryRun)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
