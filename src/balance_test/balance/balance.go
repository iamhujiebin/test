package balance

type Balancer interface {
	//多个参数，自定义hash算法
	DoBalance([]*Instance, ...string) (*Instance, error)
}
