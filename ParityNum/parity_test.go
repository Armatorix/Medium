package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestParityFuncs(test *testing.T) {
	tt := []struct {
		arg int
		res bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{-1, false},
		{-2, true},
		{2137, false},
		{998, true},
	}

	tf := []struct {
		name     string
		function func(int) bool
	}{
		{"Modular check", isEvenMod},
		{"Shift check", isEvenShift},
		{"Bitwise and check", isEvenAnd},
	}

	var wg sync.WaitGroup
	wg.Add(len(tt) * len(tf))
	for i := range tf {
		for j := range tt {
			go func(i, j int) {
				defer wg.Done()
				f := tf[i]
				t := tt[j]
				r := f.function(t.arg)
				if r != t.res {
					test.Errorf("%s f(%d)\nShould be:\n%v\nWas:\n%v\n", f.name, t.arg, t.res, r)
				}
			}(i, j)
		}
	}
	wg.Wait()
}

func TestParityMutualRandom(test *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	testNo := 1000000
	var wg sync.WaitGroup
	wg.Add(testNo)
	for i := 0; i < testNo; i++ {
		n := r.Int()
		go func(n int) {
			defer wg.Done()

			resMod, resShift, resAnd := isEvenMod(n), isEvenShift(n), isEvenAnd(n)
			if resMod != resShift || resShift != resAnd {
				test.Errorf("failed for %d\n", n)
			}
		}(n)
	}
}
