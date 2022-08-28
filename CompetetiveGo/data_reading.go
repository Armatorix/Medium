package main

import (
	"fmt"
	"io"
)

func readInt64Scan() int64 {
	var n int64
	fmt.Scan(&n)
	return n
}

func readInt64Fscanf(r io.Reader) int64 {
	var n int64
	fmt.Fscan(r, &n)
	return n
}
