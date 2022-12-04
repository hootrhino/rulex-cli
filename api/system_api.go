package api

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const SYSTEM_API_NAME = "system"

func SystemInfo(c *cli.Context) error {
	host := c.String("host")
	result := get(host, SYSTEM_API_NAME)
	fmt.Println(result)
	return nil
}
