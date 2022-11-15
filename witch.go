package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scan_dan() (a int, b int, c int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := strings.Split(scanner.Text(), " ")
	a, _ = strconv.Atoi(x[0])
	b, _ = strconv.Atoi(x[1])
	c, _ = strconv.Atoi(x[2])
	if a < -1000000000 || b < -1000000000 || c < -1000000000 || a > 1000000000 || b > 1000000000 || c > 1000000000 {
		fmt.Print("Error ")
		os.Exit(449)
	}
	return a, b, c
}

func sred(a int, b int, c int) int {
	if (a >= b && a <= c) || (a <= b && a >= c) {
		return a
	} else if (b >= a && b <= c) || (b <= a && b >= c) {
		return b
	} else {
		return c
	}
}

func result(a int, b int, c int, x int) {
	if (a >= x) && ((x != b) && (x != c)) {
		fmt.Println(a-x, 2)
	} else if (a >= x) && (x == b) && (x == c) {
		fmt.Println(a-x, 0)
	} else if a >= x {
		fmt.Println(a-x, 1)
	} else if (a < x) && ((x != b) && (x != c)) {
		fmt.Println(x-a, 2)
	} else if (a < x) && (x == b) && (x == c) {
		fmt.Println(x-a, 0)
	} else if a < x {
		fmt.Println(x-a, 1)
	}
	if (b >= x) && ((x != a) && (x != c)) {
		fmt.Println(b-x, 2)
	} else if (b >= x) && (x == a) && (x == c) {
		fmt.Println(b-x, 0)
	} else if b >= x {
		fmt.Println(b-x, 1)
	} else if (b < x) && ((x != a) && (x != c)) {
		fmt.Println(x-b, 2)
	} else if (b < x) && ((x == a) && (x == c)) {
		fmt.Println(x-b, 0)
	} else if b < x {
		fmt.Println(x-b, 1)
	}
	if (c >= x) && ((x != b) && (x != a)) {
		fmt.Println(c-x, 2)
	} else if (c >= x) && (x == b) && (x == a) {
		fmt.Println(c-x, 0)
	} else if c >= x {
		fmt.Println(c-x, 1)
	} else if (c < x) && ((x != b) && (x != a)) {
		fmt.Println(x-c, 2)
	} else if (c < x) && (x == b) && (x == a) {
		fmt.Println(x-c, 0)
	} else if c < x {
		fmt.Println(x-c, 1)
	}
}

func main() {
	a, b, c := scan_dan()
	x := sred(a, b, c)
	result(a, b, c, x)
}
