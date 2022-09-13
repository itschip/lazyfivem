package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func getConfigValues() {
	viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  // Win
  viper.AddConfigPath("$HOME/.lazyfivem/")
  // Linux
  viper.AddConfigPath("etc/lazyfivem/")

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

