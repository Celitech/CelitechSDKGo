package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Celitech/CelitechSDKGo/pkg/celitech"
	"github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
	"github.com/Celitech/CelitechSDKGo/pkg/oauth"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

func main() {
	loadEnv()

	config := celitechconfig.NewConfig()
	client := celitech.NewCelitech(config)

	request := oauth.GetAccessTokenRequest{}

	response, err := client.OAuth.GetAccessToken(context.Background(), request)
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
