package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

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

func intArray(n []string) []int {
	var m []int
	for i, _ := range n {
		v, _ := strconv.Atoi(n[i])
		m = append(m, v)
	}
	return m
}

func uint64Array(n []string) []uint64 {
	var m []uint64
	for i, _ := range n {
		v, _ := strconv.ParseUint(n[i], 10, 64)
		m = append(m, v)
	}
	return m
}

func float64Array(n []string) []float64 {
	var m []float64
	for i, _ := range n {
		v, _ := strconv.ParseFloat(n[i], 10)
		m = append(m, v)
	}
	return m
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

func main() {
	_ = int(nextInt())
	L := float64Array(strings.Split(nextLine(), " "))

	for {
		if len(L) == 1 {
			break
		}

		sort.SliceStable(L, func(a, b int) bool { return L[a] < L[b] })
		L = append(L[2:], (L[0]+L[1])/2)
	}

	fmt.Println(L[0])
}
