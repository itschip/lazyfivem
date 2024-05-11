package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func getConfigValues() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	user_config_dir, _ := os.UserConfigDir()

	// Should work on any OS
	viper.AddConfigPath(os.ExpandEnv("$LAZYFIVEM_CONFIG_HOME"))
	viper.AddConfigPath(user_config_dir + "/lazyfivem/")

	err := viper.ReadInConfig()

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
