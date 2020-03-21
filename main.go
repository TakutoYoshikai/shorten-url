package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	InitDB()
	r := InitServer()
	r.Run()
}
