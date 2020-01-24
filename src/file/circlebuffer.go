package file

import (
	"errors"
	"fmt"
	"io"
	"time"
)

//参考https://blog.csdn.net/mgr9525/article/details/86765303 环形缓冲区
/*
说明：
start指针是读的位置
end指针是写的位置
当写入空数据，isEnd为true
当主动close，isClose为true
读数据会根据isClose/isEnd做一定的返回数据操作

双线程安全读写。一个流而已
*/

type CircleByteBuffer struct {
	io.Reader
	io.Writer
	io.Closer
	datas []byte

	start   int
	end     int
	size    int
	isClose bool
	isEnd   bool
}

func NewCircleByteBuffer(len int) *CircleByteBuffer {
	var e = new(CircleByteBuffer)
	e.datas = make([]byte, len)
	e.start = 0
	e.end = 0
	e.size = len
	e.isClose = false
	e.isEnd = false
	return e
}

func (e *CircleByteBuffer) getLen() int {
	if e.start == e.end {
		return 0
	} else if e.start < e.end {
		return e.end - e.start
	} else {
		return e.start - e.end
	}
}

func (e *CircleByteBuffer) putByte(b byte) error {
	if e.isClose {
		return io.EOF
	}
	e.datas[e.end] = b
	var pos = e.end + 1
	if pos == e.size {
		pos = 0
	}

	for pos == e.start { //一直阻塞，如果没有在读的话
		if e.isClose {
			return io.EOF
		}
		time.Sleep(time.Nanosecond)
		//fmt.Println("write block")
	}
	e.end = pos
	return nil
}

func (e *CircleByteBuffer) getByte() (byte, error) {
	if e.isClose {
		return 0, io.EOF
	}
	if e.isEnd && e.getLen() <= 0 {
		return 0, io.EOF
	}
	if e.getLen() <= 0 {
		//fmt.Println("get block")
		return 0, errors.New("no datas")
	}
	var ret = e.datas[e.start]
	e.start++
	if e.start == e.size {
		e.start = 0
	}
	return ret, nil
}

func (e *CircleByteBuffer) Close() error {
	e.isClose = true
	return nil
}

func (e *CircleByteBuffer) Read(bts []byte) (int, error) {
	if e.isClose {
		return 0, io.EOF
	}
	if bts == nil {
		return 0, errors.New("bts is nil")
	}
	var ret = 0
	for i := 0; i < len(bts); i++ {
		b, err := e.getByte()
		if err != nil {
			if err == io.EOF {
				return ret, err
			}
			return ret, nil
		}
		bts[i] = b
		ret++
	}
	if e.isClose {
		return ret, io.EOF
	}
	return ret, nil
}

func (e *CircleByteBuffer) Write(bts []byte) (int, error) {
	if e.isClose {
		return 0, io.EOF
	}
	if bts == nil {
		e.isEnd = true
		return 0, io.EOF
	}
	var ret = 0
	for i := 0; i < len(bts); i++ {
		err := e.putByte(bts[i])
		if err != nil {
			fmt.Println("Write bts err:", err)
			return ret, err
		}
		ret++
	}
	if e.isClose {
		return ret, io.EOF
	}
	return ret, nil
}
