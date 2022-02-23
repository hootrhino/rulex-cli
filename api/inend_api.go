package api

import (
	"encoding/json"
	"fmt"

	"github.com/ngaut/log"
	"github.com/urfave/cli/v2"
)

func AllInends(c *cli.Context) error {
	host := c.String("host")
	result := get(host, "inends")
	fmt.Println(result)
	return nil
}

//
//
//
func CreateInend(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		_, result := post(maps, host, "inends")
		fmt.Println(result)
	}
	return nil
}

//
//
//
func UpdateInEnd(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		result := put(host, "inends", maps)
		fmt.Println(result)
	}
	return nil
}

//
//
//
func RemoveInEnd(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, "inends", "uuid="+uuid)
	fmt.Println(result)
	return nil
}
