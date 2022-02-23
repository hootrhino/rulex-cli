package api

import (
	"encoding/json"
	"fmt"

	"github.com/ngaut/log"
	"github.com/urfave/cli/v2"
)

//
//
//
func CreateRule(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		_, result := post(maps, host, "rules")
		fmt.Println(result)
	}
	return nil
}

//
//
//
func AllRules(c *cli.Context) error {
	host := c.String("host")
	result := get(host, "rules")
	fmt.Println(result)
	return nil
}

//
//
//
func UpdateRule(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		result := put(host, "rules", maps)
		fmt.Println(result)
	}
	return nil
}

//
//
//
func RemoveRule(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, "rules", "uuid="+uuid)
	fmt.Println(result)
	return nil
}
