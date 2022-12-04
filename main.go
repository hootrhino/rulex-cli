package main

//
// Cli client of rulex
//
import (
	"fmt"
	"os"
	"rulexc/api"

	"github.com/ngaut/log"
	"github.com/urfave/cli/v2"
)

// main
func main() {
	app := &cli.App{
		Name: "rulec: command line tools of rulex framework!",
		Usage: `------------------------------------------------------------------------------
|*| 查看系统参数: rulexc system-info -host 127.0.0.1
|*| 查看输入资源列表: rulexc inend-list -host 127.0.0.1
|*| 查看输出资源列表: rulexc outend-list -host 127.0.0.1
|*| 查看设备列表: rulexc device-list -host 127.0.0.1
|*| 查看规则列表: rulexc rules-list -host 127.0.0.1
|*| 创建输入资源: rulexc inend-create --config '[config]' -host 127.0.0.1
|*| 创建输出资源: rulexc outend-create --config  '[config]' -host 127.0.0.1
|*| 创建规则脚本: rulexc rules-create --config  '[config]' -host 127.0.0.1
|*| 创建设备: rulexc device-create --config  '[config]' -host 127.0.0.1
------------------------------------------------------------------------------
		`,
		Action: func(c *cli.Context) error {
			fmt.Printf("Unknown command:%v, those below are valid supported command:\n", c.Args().Slice())
			fmt.Println(c.App.Usage)
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
			// List all devices
			{
				Name:  "device-list",
				Usage: "device-list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Host of rulex",
						Value: "127.0.0.1",
					},
				},
				Action: api.AllDevices,
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
