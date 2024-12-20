package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/migrations"
)

var cql = make([]string, 6)

func main() {

	con, err := migrations.SetVariables()

	if err != nil {
		panic(err)
	}

	defer con.Close()

	setCommands()

	for _, v := range cql {
		err = con.Query(v).Exec()

		if err != nil {
			panic(err)
		}
	}

}

func setCommands() {
	cql[0] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.messages (
                        id TEXT,
                        userid TEXT,
                        receive TEXT,
            			message TEXT,
           				pages INT,
                        username TEXT,
                        types TEXT,
            			times TIMESTAMP,
                PRIMARY KEY ((userid, receive), pages, times)
        );`, entity.KeySpace)

	cql[1] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_users (
		id TEXT PRIMARY KEY,
		page INT,
		total INT);`, entity.KeySpace)

	cql[2] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_messages (
    id TEXT PRIMARY KEY,
    page INT,
    total INT
	);`, entity.KeySpace)

	cql[3] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users (
                                id TEXT,
                                pages INT,
                                username TEXT,
								password TEXT,
                                times TIMESTAMP,
                                photo TEXT,
                                PRIMARY KEY((pages), times,id)
                        ) WITH CLUSTERING ORDER BY (times DESC);`, entity.KeySpace)

	cql[4] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users_login (
			id TEXT,
			username TEXT,
			password TEXT,
			times TIMESTAMP,
			photo TEXT,
			PRIMARY KEY(username,password)
	);`, entity.KeySpace)

	cql[5] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users_status (
                        id TEXT,
                        times TIMESTAMP,
                        status TEXT,
						photo TEXT,
						username TEXT,
                        PRIMARY KEY(id)
        );`, entity.KeySpace)
}
