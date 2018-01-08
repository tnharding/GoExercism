package main

import (
	"fmt"
	"time"

	"github.com/tnharding/GoExercism/Problem1/gigasecond"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	t = gigasecond.AddGigasecond(t)
	fmt.Println(t)
}
