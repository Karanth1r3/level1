package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start", time.Now())
	sleep1(time.Second * 1)
	fmt.Println("mid", time.Now())
	sleep2(2)
	fmt.Println("end", time.Now())
}

func sleep1(dur time.Duration) {

	ticker := time.NewTicker(dur)
	<-ticker.C
	ticker.Stop()
}

func sleep2(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}
