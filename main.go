package main

import (
	"math/rand"
	"time"
)

func main() {
	InitDB()
	rand.Seed(time.Now().UnixNano())
}
