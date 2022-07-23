package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("hello docker container!")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
