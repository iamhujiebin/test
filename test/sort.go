package test

import (
	"fmt"
	"math/rand"
	"time"
	// "os"
	// "os/signal"
)

const (
	num      = 10
	rangeNum = 1000
)

type People struct {
	Name string
	Age  int
}

func (p *People) test() {
	fmt.Println("test")
}

func main() {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	//buf = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println(buf)
	t := time.Now()
	//冒泡排序
	// maopao(buf)
	// 选择排序
	// xuanze(buf)
	// 插入排序
	// charu(buf)
	//希尔排序
	// xier(buf)
	shellshort(buf)
	//快速排序
	//kuaisu(buf)
	//kuaisu2(buf)
	// 归并排序
	// guibing(buf)
	// 堆排序
	//duipai(buf)
	//duipai2(buf)
	fmt.Println(buf)

	// fmt.Println(buf)
	fmt.Println(time.Since(t))

	//等待退出
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, os.Kill)
	// <-c
	// fmt.Println("Receive ctrl-c")
}

// 冒泡排序
func maopao(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {
			if buf[j-1] > buf[j] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
}

// 选择排序
func xuanze(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			tmp := buf[i]
			buf[i] = buf[min]
			buf[min] = tmp
		}
	}
	fmt.Println("xuanze times: ", times)
}

// 插入排序
func charu(buf []int) {
	times := 0
	for i := 1; i < len(buf); i++ {
		for j := i; j > 0; j-- {
			if buf[j] < buf[j-1] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
			} else {
				break
			}
		}
	}
	fmt.Println("charu times: ", times)
}

// 希尔排序
func xier(buf []int) {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length
	for {
		incre /= 2
		for k := 0; k < incre; k++ { //根据增量分为若干子序列
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
			}
		}
		if incre == 1 {
			break
		}
	}
	fmt.Println("xier times: ", times)
}

func shellshort(buf []int) {
	n := len(buf)
	for gap := n / 2; gap > 0; gap /= 2 { //分成各个gap
		for i := gap; i < n; i++ { //i:=gap,就是把i赋值为分组内的第二个偏移值的数。当gap=1时，i=1
			for j := i - gap; j >= 0 && buf[j] > buf[j+gap]; j -= gap { //对gap内的各个分组，进行插入排序
				buf[j], buf[j+gap] = buf[j+gap], buf[j]
			}
		}
	}
}

// 快速排序
func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i == j
	a[i] = key
	kuai(a, l, i-1)
	kuai(a, i+1, r)
}

func kuaisu2(buf []int) {
	kuai2(buf, 0, len(buf)-1)
}

func kuai2(a []int, l, r int) {
	if l > r {
		return
	}
	i, j, key := l, r, a[l]
	for i < j {
		for i < j && a[j] >= key {
			j--
		}
		for i < j && a[i] <= key {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[l], a[i] = a[i], a[l]
	kuai2(a, l, i-1)
	kuai2(a, i+1, r)
}

//归并排序
func guibing(buf []int) {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
}

func merge_sort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)       //左半部分排好序
		merge_sort(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
}

// 堆排序
func duipai(buf []int) {
	temp, n := 0, len(buf)

	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(buf, i, n)
	}

	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MinHeapFixdown(buf, 0, i)
	}
}

func MinHeapFixdown(a []int, i, n int) {
	j, temp := 2*i+1, 0
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

		i = j
		j = 2*i + 1
	}
}

func duipai2(buf []int) {
	n := len(buf)
	for i := n/2 - 1; i >= 0; i-- { //最后一个叶子节点开始建堆
		adjustHeadFromDown(buf, i, n)
	}

	for i := n - 1; i >= 0; i-- { //从头开始重新建堆，并且交换首尾
		buf[0], buf[i] = buf[i], buf[0]
		adjustHeadFromDown(buf, 0, i)
	}
}

func adjustHeadFromDown(a []int, i, n int) {
	j := 2*i + 1
	for j < n {
		if j+1 < n && a[j+1] > a[j] {
			j++
		}
		if a[i] >= a[j] {
			break
		}
		a[i], a[j] = a[j], a[i]
		i = j
		j = 2*j + 1
	}
}
