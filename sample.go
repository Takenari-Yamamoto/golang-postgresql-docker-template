package main

import "os"

func main() {
	// os.Exit(1)
	// fmt.Println("Start")

	// defer func () {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(123, err)
	// }

		// Args
		// fmt.Println(os.Args[0])
		// fmt.Println(os.Args[1])
		// fmt.Println(os.Args[2])
		// fmt.Println(os.Args[3])

		// // os.Argsの要素数を表示
		// fmt.Printf("length=%d\n", len(os.Args))
		// for i, v := range os.Args {
		// 	fmt.Println(i, v)
		// }

		// f, err := os.Open("text.txt")
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// defer f.Close()

		f, _ := os.Create("foo.txt")
		f.Write([]byte("Hello\n"))
		f.WriteAt([]byte("Golang"), 6)
		f.Seek(0, os.SEEK_END)
		f.WriteString("Yeah")

}
