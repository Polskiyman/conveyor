package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var delim = '\r' // for Windows

func init() {
	if runtime.GOOS != "windows" {
		delim = '\n'
	}
}

func main() {
	var wg sync.WaitGroup
	respChan := make(chan int)
	in := bufio.NewReader(os.Stdin)

	for {
		v, err := in.ReadString(byte(delim))
		v = strings.Trim(v, "\r")
		v = strings.Trim(v, "\n")
		if err == io.EOF || v == "стоп" {
			break
		}
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("не удалось превратить в число %v", err)
		}
		wg.Add(2)

		go func(wg *sync.WaitGroup, number int) {
			fmt.Printf("Ввод: %v, \n", number)
			res := number * number
			respChan <- res
			defer wg.Done()
			fmt.Printf("Квадрат равен: %v \n", res)
		}(&wg, number)

		go func(wg *sync.WaitGroup) {
			v := <-respChan
			res := v * 2
			defer wg.Done()
			fmt.Printf("Произведение: %v \n", res)
		}(&wg)
		wg.Wait()
	}
}
