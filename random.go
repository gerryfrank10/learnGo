package main

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())
	for i := 0; i < 5; i++ {
		println(rand.Intn(100))
	}
	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
