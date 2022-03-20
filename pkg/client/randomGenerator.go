package client

import (
	"math/rand"
	"time"
)

//generator of random numbers
type RandomGenerator struct{}

//constructor of rundom numbers generator
func NewRandomGenerator() RandomGenerator {
	rand.Seed(time.Now().UnixNano())
	return RandomGenerator{}
}

//function that generates random integer number less then 100
func (rg RandomGenerator) Generate() int {
	return rand.Intn(100)
}
