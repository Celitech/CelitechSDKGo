package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Celitech/CelitechSDKGo/pkg/celitech"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
)

func main() {
	loadEnv()

	config := celitechconfig.NewConfig()
	config.SetClientId("CLIENT_ID")
	config.SetClientSecret("CLIENT_SECRET")
	client := celitech.NewCelitech(config)

	response, err := client.Destinations.ListDestinations(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
