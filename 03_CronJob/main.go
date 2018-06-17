package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("1 * * * * *", helloWorld)
	fmt.Println("Before Start")
	c.Run()
	fmt.Println("After Start")

	// c.Stop()
}

func helloWorld() {
	fmt.Println(time.Now())
}
