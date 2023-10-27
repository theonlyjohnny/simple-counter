package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getWaitDuration() time.Duration {
	fmt.Println("Checking WAIT_DURATION env var for wait duration")

	var seconds int

	if envValue, err := strconv.Atoi(os.Getenv("WAIT_DURATION")); err == nil {
		fmt.Println("Detected WAIT_DURATION env var")
		seconds = envValue
	}

	if seconds == 0 {
		seconds = rand.Intn(10)
	}

	return time.Duration(seconds) * time.Second
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	fmt.Printf("Starting counter at %s \n", time.Now().String())

	waitDuration := getWaitDuration()
	fmt.Printf("Waiting %s seconds before exiting\n", waitDuration)

	time.Sleep(waitDuration)
}
