package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
	}
	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
	letters := [5]string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		fmt.Println(letter)
	}
	wg.Done()
}
