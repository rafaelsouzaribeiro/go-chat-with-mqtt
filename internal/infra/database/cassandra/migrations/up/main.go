package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/migrations"
)

var cql = make([]string, 4)

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
			userid TEXT,
            message TEXT,
            pages INT,
			username TEXT,
            times TIMESTAMP,
            PRIMARY KEY (userid,pages,times )
        ) WITH CLUSTERING ORDER BY (pages ASC,times DESC);`, entity.KeySpace)

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
                                times TIMESTAMP,
                                photo TEXT,
                                PRIMARY KEY ((pages), id, times)
                        );`, entity.KeySpace)
}
