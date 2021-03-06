package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() uint64 {
	sc.Scan()
	i, e := strconv.ParseUint(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func int64Array(n []string) []int64 {
	var m []int64
	for i, _ := range n {
		v, _ := strconv.ParseInt(n[i], 10, 64)
		m = append(m, v)
	}
	return m
}

func uint64Array(arr []string) (newArr []uint64) {
	for i, _ := range arr {
		n, _ := strconv.ParseUint(arr[i], 10, 64)
		newArr = append(newArr, n)
	}
	return
}

func intExistIn(v int, arr []int) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}

	return false
}

func stringExistIn(v string, arr []string) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}

	return false
}

func gcd_and_lcm(m, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(m)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func uintArrayToStr(a []uint64) string {
	var S string

	for _, v := range a {
		S += strconv.Itoa(int(v))
	}
	return S
}

func remove(x []uint64, y int) []uint64 {
	var newArray []uint64
	for i, _ := range x {
		if i != y {
			newArray = append(newArray, x[i])
		}
	}
	return newArray
}

func ascendSort(a []uint64) []uint64 {
	for i := 0; i < len(a); i++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				tmp := a[i+1]
				a[i+1] = a[i]
				a[i] = tmp
			}
		}
	}
	return a
}

func descendSort(a []uint64) []uint64 {
	for i := 0; i < len(a); i++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] < a[i+1] {
				tmp := a[i+1]
				a[i+1] = a[i]
				a[i] = tmp
			}
		}
	}
	return a
}

func maxAndIndex(a []uint64) (uint64, int) {
	max := uint64(0)
	index := 0
	for i, v := range a {
		if v > max {
			max = v
			index = i
		}
	}
	return max, index
}

func sum(a []uint64) uint64 {
	sum := uint64(0)
	for _, v := range a {
		sum += v
	}
	return sum
}

func defeat(count *int64, hero *int64, monster *int64) {
	if *hero >= *monster {
		*count, *hero, *monster = *count+*monster, *hero-*monster, int64(0)
		return
	}
	*count, *hero, *monster = *count+*hero, int64(0), *monster-*hero
}

func main() {
	// sc.Split(bufio.ScanWords)
	N, _ := strconv.Atoi(readLine())
	monsters := int64Array(strings.Split(readLine(), " "))
	heroes := int64Array(strings.Split(readLine(), " "))

	var count int64

	for i := 0; i < N; i++ {
		defeat(&count, &heroes[i], &monsters[i])
		defeat(&count, &heroes[i], &monsters[i+1])
	}

	fmt.Println(count)
}
