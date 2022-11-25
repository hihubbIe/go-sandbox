package main

import (
	"fmt"
	"pipeline"
	"time"
)

func main() {
	p1 := pipeline.From(func(i float32) int {
		return int(i)
	})
	p2 := pipeline.From(func(i int) float32 {
		return float32(i)
	})
	p3 := pipeline.From(func(i float32) string {
		return fmt.Sprintf("%f", i)
	})
	pipeline.Connect(p1, p2)
	pipeline.Connect(p2, p3)
	chout1 := p1.GetChan()
	chout3 := p3.GetChan()

	p1.Emit(1.2, 3.54, 8.21)

	defer close(chout3)
	defer close(chout1)

	run := true
	lastrcv := time.Now().UnixNano()
	for run {
		select {
		case v := <-chout1:
			println(v)
			lastrcv = time.Now().UnixNano()
		case v := <-chout3:
			println(v)
			lastrcv = time.Now().UnixNano()
		default:
			now := time.Now()
			if (now.UnixNano() - lastrcv) > int64(time.Second) {
				run = false
			}
		}
	}
}
