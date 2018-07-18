package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	const BlockSize = 64
	words := []string{}

	w1 := "stall\n"
	w2 := "sphere\n"
	w3 := "carbon\n"
	w4 := "news\n"

	//reader := bufio.NewReader(os.Stdin)
	/*for i := 1; i < 5; i++ {
		fmt.Println("Enter word number", i, ": ")
		text, _ := reader.ReadString('\n')
		fmt.Printf(text)
		words = append(words, text)
	}*/

	words = append(words, w1, w2, w3, w4)

	fmt.Println("Printing Slice:")
	PrintSlice(words)
	fmt.Println("Henlo, world!")

	h1 := sha256.New()
	h2 := sha256.New()
	h3 := sha256.New()
	h4 := sha256.New()

	h1.Write([]byte(words[0]))
	h2.Write([]byte(words[1]))
	h3.Write([]byte(words[2]))
	h4.Write([]byte(words[3]))

	fmt.Printf("Sum h1: %x %v", h1.Sum(nil), words[0])
	fmt.Printf("Sum h2: %x %v", h2.Sum(nil), words[1])
	PrintSlice(h1)
	node12 := h1.Sum(nil) + h2.Sum(nil)

	b := bytes.Join(h1.Sum(nil), h2.Sum(nil))

	h12 := sha256.New()
	h34 := sha256.New()

	h12.Write(h1.Sum(nil).Join(h2.Sum(nil)))

	//h1234 := sha256.New()
}

func CreateTree(args []string) {

}

func PrintSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	for _, word := range s {
		fmt.Printf("Word: %v", word)
	}
}
