package customer

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) Update(productName string) {
	fmt.Printf("customer id %s  | product name %s \n", c.Id, productName)
}
func (c *Customer) GetID() string {
	return c.Id
}
