package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	const BlockSize = 64

	words := []string{}
	reader := bufio.NewReader(os.Stdin)

	w1 := "stall"
	w2 := "sphere"
	w3 := "carbon"
	w4 := "news"

	for i := 1; i < 5; i++ {
		fmt.Println("Enter word number", i, ": ")

		text, _ := reader.ReadString('\n')
		fmt.Printf(text)

		words = append(words, text)

	}

	fmt.Println("Printing Slice:")
	printSlice(words)
	fmt.Println("Henlo, world!")

	h1 := sha256.New()
	h1.Write([]byte(w1))

	h2 := sha256.New()
	h2.Write([]byte(w2))

	h3 := sha256.New()
	h3.Write([]byte(w3))

	h4 := sha256.New()
	h4.Write([]byte(w4))
}

func CreateTree(args []string) {

}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	for _, word := range s {
		fmt.Printf("Word: %v", word)
	}
}
