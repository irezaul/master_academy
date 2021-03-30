package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)
	createFile("reza.txt")
}

func createFile(fileName string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer posf.Close()

	_, err = posf.Write([]byte("this is write file"))
	//fmt.Println(n, err)
	if err != nil {
		return false
	}

	return true
}
