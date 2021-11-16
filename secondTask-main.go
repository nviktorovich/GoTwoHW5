package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

//Реализуйте функцию для разблокировки мьютекса с помощью defer

var mx = sync.Mutex{}
var wg2 = sync.WaitGroup{}
var cnt int

func incrCnt () {
	cnt++
}

func makeTmpFile(s []byte) (f *os.File, err error) {
	defer f.Close()
	f, err = os.Create("tmp.txt")
	if err != nil {
		err = errors.New("ошибка создания файла")
		return nil, err
	}
	_, err = f.Write(s)
	if err != nil {
		err = errors.New("ошибка записи в файл")
	}
	return f, nil
}
// в этой функции добавил в конце обнуление счетчика. Если убрать мьютексы, то
// параллельно работающие горутины могут выдавать разные значения счетчика, но с
// мьютексами, оно всегда одинаково. Потому что разблокировка происходит уже
// после обнуления.
func logToFile(f *os.File) error {
	defer mx.Unlock()
	defer wg2.Done()
	//time.Sleep(time.Millisecond * 300)
	now := time.Now().Format("150405")
	mx.Lock()

	incrCnt()

	_, err := f.WriteString(now + " " + strconv.Itoa(cnt)+ "\n")
	if err != nil {
		err = errors.New("ошибка записи лога в файл")
		return err
	}
	cnt = 0
	return nil
}

func main() {
	now := time.Now()
	f, err := makeTmpFile([]byte("hello\n"))
	defer f.Close()
	if err != nil {
		log.Printf("ошибка файла, %v", err)
		os.Exit(1)
	}
	wg2.Add(10)
	for i := 0; i < 10; i++ {
		go logToFile(f)

	}
	wg2.Wait()
	fmt.Println("success", time.Since(now))
}
