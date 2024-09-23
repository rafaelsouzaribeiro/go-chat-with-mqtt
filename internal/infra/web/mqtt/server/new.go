package server

type Broker struct {
	host     string
	port     int
	username string
	password string
}

func NewBroker(host, username, password string, port int) *Broker {
	return &Broker{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}
