package dataservice

import "fmt"

type Rest interface {
	Get()
	Post()
}

type Connection struct {
	connection string
}

func (con Connection) Get() {
	fmt.Println("GET : ", con.connection)
}
func (con Connection) Post() {
	fmt.Println("POST : ", con.connection)
}

func (con Connection) SetNameConnection(name string) {
	con.connection = name
}
