package file

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewCircleByteBuffer(t *testing.T) {
	fmt.Println("hello world")
	b := NewCircleByteBuffer(1024)
	go func() {
		for i := 0; i < 100000; i++ {
			bts := []byte("么个么个")
			bts = append(bts, []byte(strconv.Itoa(i))...)
			b.Write(bts)
			//time.Sleep(time.Millisecond)
		}
		b.Write(nil)
	}()

	res := make([]byte, 0)
	buf := make([]byte, 10)
	for {
		//time.Sleep(time.Millisecond)
		n, err := b.Read(buf)
		if n > 0 {
			res = append(res, buf[0:n]...)
			fmt.Println("Reads:", string(buf[0:n]))
		}
		if err != nil {
			break
		}
	}
	fmt.Println("end!")
	fmt.Println(string(res))
	forever := make(chan int)
	<-forever
}
