package main

//
// Cli client of rulex
//
import (
	"os"
	"rulexc/api"

	"github.com/ngaut/log"
	"github.com/urfave/cli/v2"
)

// main
func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			// SystemInfo
			{
				Name:  "system-info",
				Usage: "system-info",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
				},
				Action: api.SystemInfo,
			},
			// List all inends
			{
				Name:  "inend-list",
				Usage: "inend-list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
				},
				Action: api.AllInends,
			},
			// List all outends
			{
				Name:  "outend-list",
				Usage: "outend-list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
				},
				Action: api.AllOutEnds,
			},
			// List all rules
			{
				Name:  "rules-list",
				Usage: "rules-list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
				},
				Action: api.AllRules,
			},
			// Create InEnd
			{
				Name:  "inend-create",
				Usage: "inend-create",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
					&cli.StringFlag{
						Name:     "config",
						Usage:    "Config of rulex",
						Value:    "",
						Required: true,
					},
				},
				Action: api.CreateInend,
			},
			// create outend
			{
				Name:  "outend-create",
				Usage: "outend-create",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
					&cli.StringFlag{
						Name:     "config",
						Usage:    "Config of rulex",
						Value:    "",
						Required: true,
					},
				},
				Action: api.CreateOutends,
			},
			// create rule
			{
				Name:  "rule-create",
				Usage: "rule-create",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
					&cli.StringFlag{
						Name:     "config",
						Usage:    "Config of rulex",
						Value:    "",
						Required: true,
					},
				},
				Action: api.CreateRule,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
