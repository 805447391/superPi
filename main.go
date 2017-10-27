package main

import (
	"flag"
	"fmt"
)

func main() {
	var L int
	flag.IntVar(&L, "n", 1000, "Number of Pi")
	flag.Parse()

	N := (L)/4 + 1

	s := make([]int, N+3)
	w := make([]int, N+3)
	v := make([]int, N+3)
	q := make([]int, N+3)
	n := (int)(float32(L)/1.39793 + 1)

	w[0] = 16 * 5
	v[0] = 4 * 239

	for k := 1; k <= n; k++ {
		div(w, 25, w, N)
		div(v, 57121, v, N)
		sub(w, v, q, N)
		div(q, 2*k-1, q, N)
		if k%2 != 0 {
			add(s, q, s, N)
		} else {
			sub(s, q, s, N)
		}
	}
	fmt.Printf("%d.", s[0])
	for k := 1; k < N; k++ {
		fmt.Printf("%04d", s[k])
	}
	return
}

func add(a []int, b []int, c []int, N int) {
	i, carry := 0, 0
	for i = N + 1; i >= 0; i-- {
		c[i] = a[i] + b[i] + carry
		if c[i] < 10000 {
			carry = 0
		} else {
			c[i] = c[i] - 10000
			carry = 1
		}
	}
}

func sub(a []int, b []int, c []int, N int) {
	i, borrow := 0, 0
	for i = N + 1; i >= 0; i-- {
		c[i] = a[i] - b[i] - borrow
		if c[i] >= 0 {
			borrow = 0
		} else {
			c[i] = c[i] + 10000
			borrow = 1
		}
	}
}

func div(a []int, b int, c []int, N int) {
	i, tmp, remain := 0, 0, 0
	for i = 0; i <= N+1; i++ {
		tmp = a[i] + remain
		c[i] = tmp / b
		remain = (tmp % b) * 10000
	}
}
