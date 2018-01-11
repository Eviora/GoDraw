package main

import (
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateColor() string {
	return colorful.Hsv(rand.Float64()*360.0, 0.8, 0.8).Hex()
}
