package main

import (
	"fmt"
	"strings"
	// "strings"
)

//判断是否为旋转词
func isRotation(a, b string) bool {
	if a == "" || b == "" { //如果俩个输入的字符串为空的话 直接return false
		return false
	}
	b2 := b + b                    //把b相加起来拼成长 串
	return getIndexOf(b2, a) != -1 //如果getIndexOf 为-1 则返回false
}

func getIndexOf(s, m string) int { //如果匹配返回相应的index  不匹配返回-1
	if len(s) < len(m) { //如果待匹配的串比匹配的串短  则一定不匹配  返回-1
		return -1
	}
	ss := []byte(s) //先将输入的 俩个string 转化成 【】byte型  方便遍历每一个字符
	ms := []byte(m)
	si := 0 //分别是ms ss 的遍历时的索引下标 初始化为0
	mi := 0
	next := getNextArray(ms)           //KMP算法   生成next数组
	for si < len(ss) && mi < len(ms) { //遍历俩个byte数组
		if ss[si] == ms[mi] { //如果遇到俩个数组有相等的字符
			si++ //将他们的索引游标同时后移
			mi++
		} else if next[mi] == -1 { //如果匹配到某个位置不相等的话   查看他的next数组值 如果为-1，则把si 后移
			si++
		} else { //如果mi当时的next数组值不是-1
			mi = next[mi] //  则把next相应的值和 ss中si（即刚刚不匹配的地方 进行比较，，转回第一个if条件）
		}
	}
	if mi == len(ms) { //如果最后匹配成功的话mi 应该是ms的长度
		return si - mi //则返回si-mi 即刚刚匹配的开始index
	} else {
		return -1 //匹配到最后还没有将mi移到ms的末尾  说明不匹配
	}

}

func getNextArray(ms []byte) []int { //得到next数组
	if len(ms) == -1 { //最特殊的情况  ms的长度为-1  直接返回 -1
		return []int{-1}
	}
	length := len(ms)
	next := make([]int, length) //定义了一个切片  比较方便go语言
	next[0] = -1                //第一个字符 前面没有东西  所以为-1
	next[1] = 0                 //第二个字符虽然前面有一个字符但是 它的前面只有一个所以没有最长公共前缀后缀  即为0
	pos := 2                    //定义的是字符串比较靠后的 位置
	cn := 0                     //定义前面的位置   为了求最长公共前缀后缀的长度
	for pos < len(next) {       //遍历整个ms
		if ms[pos-1] == ms[cn] { //比较不匹配的前一个字符与cn是否相等
			cn = cn + 1
			next[pos] = cn //因为要比较最后不匹配的位置和前面最长的前缀的下一个字符 所以要加一
			pos++          //递增  以遍历
		} else if cn > 0 { //如果 cn  不再是0的话  意味着前面的next数组值已经算出来了 所以后面可以不用管
			cn = next[cn] //因为刚刚直接比较cn和 pos-1 不想等  找到cn的最长公共前缀  比较
		} else { //如果以上几步都不行  则说明 他没有  给它0吧
			next[pos] = 0
			pos++
		}
	}
	fmt.Println(next)
	return next
}
func main() {
	str1 := "abcdabd"
	str2 := "abdabcd"
	//str1, str2 = "ab", "ab"
	fmt.Println(isRotation(str1, str2))
	str3 := str1 + str1
	fmt.Println(strings.Contains(str3, str2))
}
