package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	s := covertToBin(63)
	fmt.Println(s)

	ss := strings.Split("a b c d e", " ")
	fmt.Println(len(ss))
	//n := make(map[string]string)
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		randomS := genRandomNumber()
	//		if s, ok := n[randomS]; ok {
	//			fmt.Println("存在：", s)
	//		} else {
	//			fmt.Println(randomS)
	//			n[randomS] = randomS
	//		}
	//	}()
	//}
	//time.Sleep(5*time.Second)
}

var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getegid()))
}

func genRandomNumber() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1926817 + 1893122683
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]

}

func m() {
	r := make([]int, 0, 0)
	fmt.Printf("%p", r)
	fmt.Println()
	fmt.Println(cap(r))
	r = append(r, 1)
	fmt.Printf("%p", r)
	fmt.Println()
	fmt.Println(cap(r))
	fmt.Println(len(r))
}

func mm() {
	a := map[string]string{
		"a": "it's a",
		"b": "it's b",
		"c": "it's c",
		"d": "it's d",
	}
	res := ""
	for _, v := range a {
		res += v + "_"
	}
	fmt.Println(res)
}

func tt() {
	ct := time.Now()
	st := time.Date(2022, time.Month(12)+1, 28, 19, 59, 0, 0, ct.Location())
	fmt.Println(st.Date())
	//sub := int(st.Weekday()) - int(ct.Weekday())
	//	fmt.Println(sub)
}

func monthtt() {
	ct := time.Now()
	t := time.Date(2022, 3, 31, 0, 0, 0, 0, ct.Location())
	fmt.Println(t.AddDate(0, 1, 0).Unix())
	fmt.Println(t.Unix())

}

func ZellerFunction2Week(year, month, day uint16) int {
	var y, m, c uint16
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}

	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	return int(week)

}

func covertToBin(n int) []int {
	res := make([]int, 7)
	i := 1
	for n != 0 {
		res[i] = n & 1
		n = n >> 1
		i++
	}
	return res
}
