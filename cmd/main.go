package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.io/MXuDong/example/config"
	"github.io/MXuDong/example/internal/server"
	"os"
	"strings"
)

// ====================
// @author: MXuDong
// project for:
//   The project for mock the application in product environment, and
//   use it to provide some func or handle to mock the action of
//   application.
// Create at: 2021-04-30
// ====================

var ConfigPath = ""

// main function, the application run here
func main() {

	configPath := flag.String("c", "./config/conf.yaml", "配置文件路径")
	flag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	err := viper.Unmarshal(&config.Config, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	})
	if err != nil {
		logrus.Error(err)
	}

	c := config.Config
	_ = c

	preHandler()

	go func() {
		err := server.TcpServerStart()
		if err != nil {
			logrus.Error(err)
		}
	}()
	server.Run()
}

// deal with the config, the pre handler of the program
func preHandler() {
	_ = config.InitInnerValue()
}

func init() {
	envs := os.Environ()
	for item := range envs {
		fmt.Println(item)
	}
}
