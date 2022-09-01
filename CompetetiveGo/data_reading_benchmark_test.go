package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func BenchmarkFmtFscanWithBufio(b *testing.B) {
	cleanup, err := mockStdin(b.N)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < b.N; i++ {
		x := readInt64Fscanf(r)
		if x == 0 {
			panic("unexpected value")
		}
	}
	cleanup()
}

func BenchmarkFmtFscanWithOsStdin(b *testing.B) {
	cleanup, err := mockStdin(b.N)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := readInt64Fscanf(os.Stdin)
		if x == 0 {
			panic("unexpected value")
		}
	}
	cleanup()
}

func BenchmarkFmtScan(b *testing.B) {
	cleanup, err := mockStdin(b.N)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := readInt64Scan()
		if x == 0 {
			panic("unexpected value")
		}
	}
	cleanup()
}

// mockStdin is a helper function that lets the test pretend dummyInput as os.Stdin.
// It will return a function for `defer` to clean up after the test.
// KUDOS: https://gist.github.com/KEINOS/76857bc6339515d7144e00f17adb1090
func mockStdin(n int) (funcDefer func(), err error) {
	oldOsStdin := os.Stdin
	tmpfile, err := os.CreateTemp(os.TempDir(), "*")

	if err != nil {
		return nil, err
	}

	for i := 0; i < n; i++ {
		if _, err := tmpfile.Write([]byte(fmt.Sprintf("%d ", rand.Intn(2137)+1))); err != nil {
			return nil, err
		}
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	os.Stdin = tmpfile

	return func() {
		os.Stdin = oldOsStdin
		os.Remove(tmpfile.Name())
	}, nil
}
