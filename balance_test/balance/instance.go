package balance

import (
	"strconv"
)

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (self *Instance) GetHost() string {
	return self.host
}

func (self *Instance) GetPort() int {
	return self.port
}

func (self *Instance) String() string {
	return self.host + ":" + strconv.Itoa(self.port)
}
