package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkParityFuncs(b *testing.B) {
	funcs := []struct {
		f    func(int) bool
		name string
	}{
		{isEvenAnd, "And"},
		{isEvenShift, "Shift"},
		{isEvenMod, "Mod"},
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for _, tf := range funcs {
		b.Run(fmt.Sprintf("%s", tf.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tf.f(r.Int())
				tf.f(i)
			}
		})
	}
}
