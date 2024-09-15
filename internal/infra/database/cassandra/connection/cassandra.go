package connection

import "github.com/gocql/gocql"

func NewCassandraConnect(hosts []string, username, password string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return session, nil
}
