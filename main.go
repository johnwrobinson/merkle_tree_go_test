package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	words := []string{"stall\n", "sphere\n", "carbon\n", "news\n"}

	//reader := bufio.NewReader(os.Stdin)
	/*for i := 1; i < 5; i++ {
		fmt.Println("Enter word number", i, ": ")
		text, _ := reader.ReadString('\n')
		fmt.Printf(text)
		words = append(words, text)
	}*/

	fmt.Println("Printing Slice:")
	PrintSlice(words)

	h1 := GetHash(words[0])
	h2 := GetHash(words[1])
	h3 := GetHash(words[2])
	h4 := GetHash(words[3])

	fmt.Printf("h1: %v %v", h1, words[0])
	fmt.Printf("h2: %v %v", h2, words[1])
	fmt.Printf("h3: %v %v", h3, words[2])
	fmt.Printf("h4: %v %v", h4, words[3])

	b := Concat(h1, h2)
	h12 := GetHash(b.String())

	b = Concat(h3, h4)
	h34 := GetHash(b.String())

	fmt.Printf("h12: %v\n", h12)
	fmt.Printf("h34: %v\n", h34)

	b = Concat(h12, h34)
	h1234 := GetHash(b.String())

	fmt.Printf("h1234: %v\n", h1234)
}

// GetHash takes a string and returns its sha256 hash as a string literal
func GetHash(b string) (hash string) {
	// Hash input b and capture/format its output as a string literal
	hash = fmt.Sprintf("%X", sha256.Sum256([]byte(b)[:]))
	return
}

// Concat takes two strings, concatonates them, and returns a buffer containing the result
func Concat(x string, y string) (z bytes.Buffer) {
	z.Write([]byte(x))
	z.Write([]byte(y))
	return
}

// Future function
//func CreateTree(args []string) {}

// PrintSlice Print a given string slice
func PrintSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	for _, word := range s {
		fmt.Printf("Word: %v", word)
	}
}
