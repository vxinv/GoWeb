package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("./quanbudu.go")
	defer f.Close()
	buff := make([]byte, 1024)
	for {
		n, err := f.Read(buff)
		if n == 1024 {
			fmt.Println(string(buff))
		} else {
			fmt.Println(string(buff[0:n]))
		}
		if err == io.EOF {
			break
		}
	}
}
