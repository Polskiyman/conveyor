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

var delim = byte('\r') // for Windows

func init() {
	if runtime.GOOS != "windows" {
		delim = '\n' // for UNIX
	}
}

func main() {
	wg := &sync.WaitGroup{}
	input, squares := make(chan int), make(chan int)

	go readNumbers(input)

	go sqr(input, squares)

	wg.Add(1)
	go double(wg, squares)

	wg.Wait()
	fmt.Println("конец")
}

func readNumbers(input chan int) {
	in := bufio.NewReader(os.Stdin)
	for {
		v, err := in.ReadString(delim)
		v = strings.Trim(v, "\r\n")
		if err == io.EOF || v == "стоп" {
			close(input)
			break
		}
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("не удалось превратить в число: %v", err)
			continue
		}

		fmt.Printf("Ввод: %v\n", number)
		input <- number
	}
}

func sqr(input, squares chan int) {
	for {
		v, ok := <-input
		if !ok {
			close(squares)
			break
		}
		squares <- v * v
		fmt.Printf("Квадрат равен: %v\n", v*v)
	}
}

func double(wg *sync.WaitGroup, squares chan int) {
	for {
		v, ok := <-squares
		if !ok {
			break
		}
		fmt.Printf("Произведение: %v\n", 2*v)
	}
	wg.Done()
}
