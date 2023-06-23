package api

import (
	"encoding/json"
	"fmt"

	"log"

	"github.com/urfave/cli/v2"
)

const OUTEND_API_NAME = "outends"

func AllOutEnds(c *cli.Context) error {
	host := c.String("host")
	result := get(host, OUTEND_API_NAME)
	fmt.Println(result)
	return nil
}

func CreateOutends(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Println(config, err)
	} else {
		_, result := post(maps, host, OUTEND_API_NAME)
		fmt.Println(result)
	}
	return nil
}

func UpdateOutEnd(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Println(config, err)
	} else {
		result := put(host, OUTEND_API_NAME, maps)
		fmt.Println(result)
	}
	return nil
}

func RemoveOutEnd(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, OUTEND_API_NAME, "uuid="+uuid)
	fmt.Println(result)
	return nil
}
