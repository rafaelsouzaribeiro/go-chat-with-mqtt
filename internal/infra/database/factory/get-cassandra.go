package factory

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/connection"
	"github.com/spf13/viper"
)

func (f *Factory) GetConCassandra() (*gocql.Session, error) {
	viper.AutomaticEnv()

	hostsDocker := strings.Split(viper.GetString("HOST_CASSANDRA_DOCKER"), ",")
	hosts := strings.Split(viper.GetString("HOST_CASSANDRA"), ",")

	if hostsDocker[0] != "" {
		hosts = hostsDocker
	}

	cassandra, err := connection.NewCassandraConnect(hosts, viper.GetString("USER_CASSANDRA"), viper.GetString("PASSWORD_CASSANDRA"))

	if err != nil {
		return nil, err
	}

	return cassandra, nil
}
