package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"context"
	"github.com/Celitech/CelitechSDKGo"
	"github.com/Celitech/CelitechSDKGo/destinations"
)

func main() {
	loadEnv()

	config := celitech.NewConfig()
	config.SetClientID("CLIENT_ID")
	config.SetClientSecret("CLIENT_SECRET")
	client := celitech.NewCelitech(config)

	params := destinations.ListDestinationsRequestParams{
		Accept: celitech.Nullable[string]("application/json"),
	}

	response, err := client.Destinations.ListDestinations(context.Background(), params)
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
