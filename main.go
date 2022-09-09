package main

import (
	"encoding/json"
	"fmt"
	"msp-go/config"
	"msp-go/db"
)

func main() {

	conf := config.New(false)
	dbc := db.New(*conf)
	containers := dbc.GetContainers()
	for _, container := range containers {
		output, err := json.MarshalIndent(container, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}
