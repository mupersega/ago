package helper

import (
	"net/http"
	"strconv"
)

type twoDArray [][]int

func (a twoDArray) Get(x, y int) int {
	return a[x][y]
}

func (a twoDArray) Set(x, y, value int) {
	a[x][y] = value
}

func (a twoDArray) display() {
	// print in columns and rows
	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a[0]); x++ {
			print(a.Get(x, y), " ")
		}
		println()
	}
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func GetIpFromRequest(r *http.Request) string {
	return r.RemoteAddr
}
