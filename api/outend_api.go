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
func AllOutEnds(c *cli.Context) error {
	host := c.String("host")
	result := get(host, "outends")
	fmt.Println(result)
	return nil
}

//
//
//
func CreateOutends(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		_, result := post(maps, host, "outends")
		fmt.Println(result)
	}
	return nil
}

//
//
//
func UpdateOutEnd(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		result := put(host, "outends", maps)
		fmt.Println(result)
	}
	return nil
}

//
//
//
func RemoveOutEnd(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, "outends", "uuid="+uuid)
	fmt.Println(result)
	return nil
}
