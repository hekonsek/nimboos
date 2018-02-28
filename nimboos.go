package main

import "fmt"

func main() {
	data, err := Asset("data/hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
