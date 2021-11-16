package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go SleepWorker(i)
	}
	wg.Wait()
}

func SleepWorker(idx int) {
	defer wg.Done()
	fmt.Printf("%d worker sleep\n", idx)
	time.Sleep(time.Second * 1)
	fmt.Printf("%d worker stop sleep\n", idx)
}
