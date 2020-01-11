package main

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c *Cat) Quack() {
	println(c.Name + " meow")
}

func main() {
	var c Duck = &Cat{Name: "grooming"}
	c.Quack()        //动态的
	c.(*Cat).Quack() //这个是指定的，会快一些。运行时编译器的指令会少一些
	//动态派发生成的指令会带来 ~18% 左右的额外性能开销。
	//开启默认的编译器优化之后，动态派发的额外开销会降低至 ~5% 左右，对应用性能的整体影响就更小了。
}
