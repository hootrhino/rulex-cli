package api

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const USER_API_NAME = "users"

func AllUsers(c *cli.Context) error {
	host := c.String("host")
	result := get(host, USER_API_NAME)
	fmt.Println(result)
	return nil
}

func CreateUser(c *cli.Context) error {
	host := c.String("host")
	username := c.String("u")
	password := c.String("p")
	maps := map[string]interface{}{}
	maps["username"] = username
	maps["password"] = password
	maps["role"] = "admin"
	_, result := post(maps, host, USER_API_NAME)
	fmt.Println(result)
	return nil
}
