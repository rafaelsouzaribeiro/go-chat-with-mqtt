package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	HostMqtt          string `mapstructure:"HOST_MQTT"`
	PortMqtt          string `mapstructure:"PORT_MQTT"`
	UserNameMqtt      string `mapstructure:"USERNAME_MQTT"`
	PasswordMqtt      string `mapstructure:"PASSWORD_MQTT"`
	HostWebsocketMqtt string `mapstructure:"HOST_MQTT_WEBSOCKET"`
	PortWebsocketMqtt string `mapstructure:"PORT_MQTT_WEBSOCKET"`
	TopicMqtt         string `mapstructure:"TOPIC_MQTT"`
	HostCassaandra    string `mapstructure:"HOST_CASSANDRA"`
	UserCassaandra    string `mapstructure:"USER_CASSANDRA"`
	PassCassaandra    string `mapstructure:"PASSWORD_CASSANDRA"`
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
