package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, _ := os.Open("foo.txt");
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
