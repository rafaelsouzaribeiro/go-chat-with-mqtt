package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	HostName       string `mapstructure:"HOST_NAME_CONF"`
	WsEndPoint     string `mapstructure:"WS_ENDPOINT_CONF"`
	Port           string `mapstructure:"PORT_CONF"`
	HostCassaandra string `mapstructure:"HOST_CASSANDRA"`
	UserCassaandra string `mapstructure:"USER_CASSANDRA"`
	PassCassaandra string `mapstructure:"PASSWORD_CASSANDRA"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("go_cassandra")
	viper.SetConfigType("env")
	viper.SetConfigFile(path + ".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, err
}
