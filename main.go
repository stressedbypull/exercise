package main

import "dev/exercise/customer"

func main() {
	ps5 := customer.NewProduct("Sony Playstation 5")
	ps4 := customer.NewProduct("Sony Playstation 4")
	osservatore2 := &customer.Customer{Id: "martin@martinfowler.com "}
	osservatore1 := &customer.Customer{Id: "bora@borakasmer.com"}
	ps5.RegisterList([]customer.Observer{osservatore2, osservatore1})
	ps4.RegisterList([]customer.Observer{osservatore2})
	ps5.UpdateAvalability()
	ps4.UpdateAvalability()

}
