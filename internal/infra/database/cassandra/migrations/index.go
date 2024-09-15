package migrations

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/configs"
	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/infra/database/cassandra/connection"
	"github.com/spf13/viper"
)

func SetVariables() (*gocql.Session, error) {

	viper.AutomaticEnv()
	var hosts []string
	hostsDocker := strings.Split(viper.GetString("HOST_CASSANDRA_DOCKER"), ",")

	Conf, err := configs.LoadConfig("./cmd/")

	if err != nil {
		panic(err)
	}

	user := Conf.UserCassaandra
	password := Conf.PassCassaandra
	hosts = strings.Split(Conf.HostCassaandra, ",")

	if hostsDocker[0] != "" {
		hosts = hostsDocker
	}

	con, err := connection.NewCassandraConnect(hosts, user, password)

	if err != nil {
		return nil, err
	}

	return con, nil

}
