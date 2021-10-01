package customer

import (
	"fmt"
	"strings"
)

type product struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewProduct(_name string) *product {
	return &product{
		name: _name,
	}
}

func (p *product) NotifyAll() {
	for _, Observer := range p.observerList {
		Observer.Update(p.name)
	}
}

func (p *product) UpdateAvalability() {
	fmt.Printf("Product name %s \n", p.name)
	fmt.Printf(strings.Repeat("-", 100))
	fmt.Printf("\n")
	p.inStock = true
	p.NotifyAll()
}

func (p *product) Register(o Observer) {
	p.observerList = append(p.observerList, o)
}

func (p *product) RegisterList(obList []Observer) {
	for _, o := range obList {
		p.observerList = append(p.observerList, o)
	}
}

func RemoveFromList(observerList []Observer, removeObserver Observer) []Observer {
	for index, Observer := range observerList {
		if removeObserver.GetID() == Observer.GetID() {
			return append(observerList[:index], observerList[index+1:]...)
		}
	}
	return observerList
}

func (p *product) Unregister(o Observer) []Observer {
	p.observerList = RemoveFromList(p.observerList, o)
	return p.observerList
}
