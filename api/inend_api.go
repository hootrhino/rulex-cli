package api

import (
	"encoding/json"
	"fmt"

	"log"

	"github.com/urfave/cli/v2"
)

const INEND_API_NAME = "inends"

func AllInends(c *cli.Context) error {
	host := c.String("host")
	result := get(host, INEND_API_NAME)
	fmt.Println(result)
	return nil
}

func CreateInend(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Println(config, err)
	} else {
		_, result := post(maps, host, INEND_API_NAME)
		fmt.Println(result)
	}
	return nil
}

func UpdateInEnd(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Println(config, err)
	} else {
		result := put(host, INEND_API_NAME, maps)
		fmt.Println(result)
	}
	return nil
}

func RemoveInEnd(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, INEND_API_NAME, "uuid="+uuid)
	fmt.Println(result)
	return nil
}
