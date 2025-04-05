package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func NewConfig() *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "config/local.yaml", "config path, eg: -conf config/local.yaml")
		flag.Parse()
	}
	if envConf == "" {
		envConf = "config/local.yaml"
	}
	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)

}
func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
