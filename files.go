package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"

	"github.com/spf13/viper"
)

func getConfigValues() {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	home_dir := usr.HomeDir

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Win
	viper.AddConfigPath(home_dir + "/.lazyfivem/")
	// viper.AddConfigPath(os.ExpandEnv("~/.lazyfivem/"))

	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Failed to load lazyfivem config: %w", err))
	}
}

func getAllServers() {
	keys := viper.AllKeys()

	for _, k := range keys {
		serverName := strings.ToUpper(k)
		serverPath := viper.Get(k)

		Servers[serverName] = serverPath.(string)
	}
}
