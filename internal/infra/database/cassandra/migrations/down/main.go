package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/entity"
	"github.com/rafaelsouzaribeiro/whatsapp-clone-in-go/internal/infra/database/cassandra/migrations"
)

var cql = ""

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("do you want to delete? Yes(y) or No(n): ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == "n" {
		os.Exit(0)
	}

	if text == "y" {
		con, err := migrations.SetVariables()

		if err != nil {
			panic(err)
		}

		defer con.Close()

		setCommands()

		err = con.Query(cql).Exec()

		if err != nil {
			panic(err)
		}
	}
}

func setCommands() {
	cql = fmt.Sprintf(`DROP KEYSPACE IF EXISTS %s;`, entity.KeySpace)

}
