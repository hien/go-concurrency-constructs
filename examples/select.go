package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes generates random strings of length n
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomString(messages chan string) {
	messages <- RandStringRunes(rand.Intn(10) + 1)
}

func randomInt(numbers chan int) {
	numbers <- rand.Intn(100)
}

func main() {
	messages := make(chan string)
	numbers := make(chan int)

	for i := 1; i <= 5; i++ {
		go randomString(messages)
		go randomInt(numbers)
	}

	select {
	case msg := <-messages:
		fmt.Println("Message: ", msg)
	case num := <-numbers:
		fmt.Println("Number: ", num)
	default:
		fmt.Println("None")
	}
}
