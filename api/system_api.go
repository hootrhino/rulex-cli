package api

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

//
//
//
func SystemInfo(c *cli.Context) error {
	host := c.String("host")
	result := get(host, "system")
	fmt.Println(result)
	return nil
}
