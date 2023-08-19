package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	TYPE = "tcp"
)

var debug = false

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
				},
				Action: func(c *cli.Context) error {
					Send(host, port, debug)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
