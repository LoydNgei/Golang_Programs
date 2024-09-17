// Panic is a mechanism that allows you to halt normal execution of a program when an unexpected situation occurs

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	SecretKey string `json:"secret_key"`
}

func main() {
	// Read the config file
	configData, err : ioutil.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("Error reading configuration file: %v", err))
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing configuration data: %v", err))
	}

	if config.SecretKey == "" {
		panic("Secret key is missing from configuration")
	}

	fmt.Println("Applicaton started successfully")
}



// Recover is used to regain control of a panicking goroutine

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	SecretKey string `json:"secret_key"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Prinf("An error occured: %v\n", r)
			fmt.Println("Application terminated gracefully")
		} else {
			fmt.Println("Application executed successfully")
		}
	}
	// Read the config file
	configData, err : ioutil.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("Error reading configuration file: %v", err))
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing configuration data: %v", err))
	}

	if config.SecretKey == "" {
		panic("Secret key is missing from configuration")
	}

	fmt.Println("Applicaton started successfully")
}