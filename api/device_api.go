package api

import (
	"encoding/json"
	"fmt"

	"github.com/ngaut/log"
	"github.com/urfave/cli/v2"
)

const DEVICE_API_NAME = "devices"

func AllDevices(c *cli.Context) error {
	host := c.String("host")
	result := get(host, DEVICE_API_NAME)
	fmt.Println(result)
	return nil
}

func CreateDevice(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		_, result := post(maps, host, DEVICE_API_NAME)
		fmt.Println(result)
	}
	return nil
}

func UpdateDevice(c *cli.Context) error {
	host := c.String("host")
	config := c.String("config")
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(config), &maps)
	if err != nil {
		log.Error(config, err)
	} else {
		result := put(host, DEVICE_API_NAME, maps)
		fmt.Println(result)
	}
	return nil
}

func RemoveDevice(c *cli.Context) error {
	host := c.String("host")
	uuid := c.String("uuid")
	result := delete(host, DEVICE_API_NAME, "uuid="+uuid)
	fmt.Println(result)
	return nil
}
