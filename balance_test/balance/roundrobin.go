package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

func (self *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("No Instances")
		return
	}
	if self.curIndex >= lens {
		self.curIndex = 0
	}
	inst = insts[self.curIndex]
	self.curIndex++
	//self.curIndex = inst[self.curIndex+1] % lens
	return
}
