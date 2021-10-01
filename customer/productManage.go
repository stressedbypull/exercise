package customer

type ProductManage interface {
	Register(Observer Observer)
	RegisterList(Observer []Observer)
	UnRegister(Observer Observer)
	NotifyAll()
}
