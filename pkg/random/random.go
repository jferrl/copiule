package random

import "math/rand/v2"

type Random struct {
}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) Intn(n int) int {
	return rand.IntN(n)
}
