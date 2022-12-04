package api

import (
	"encoding/json"
	"fmt"

	"github.com/ngaut/log"
	"github.com/urfave/cli/v2"
)

const RULE_API_NAME = "rules"

func CreateRule(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		_, result := post(maps, host, OUTEND_API_NAME)
		fmt.Println(result)
	}
	return nil
}

func AllRules(c *cli.Context) error {
	host := c.String("host")
	result := get(host, OUTEND_API_NAME)
	fmt.Println(result)
	return nil
}

func UpdateRule(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		result := put(host, OUTEND_API_NAME, maps)
		fmt.Println(result)
	}
	return nil
}

func RemoveRule(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, OUTEND_API_NAME, "uuid="+uuid)
	fmt.Println(result)
	return nil
}
