package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
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

	h1 := sha256.Sum256([]byte(words[0]))
	h2 := sha256.Sum256([]byte(words[1]))
	h3 := sha256.Sum256([]byte(words[2]))
	h4 := sha256.Sum256([]byte(words[3]))

	fmt.Printf("h1: %x %v", h1, words[0])
	fmt.Printf("h2: %x %v", h2, words[1])
	fmt.Printf("h3: %x %v", h3, words[2])
	fmt.Printf("h4: %x %v", h4, words[3])

	stringByte := fmt.Sprintf("%x", h1[:])
	stringByte2 := fmt.Sprintf("%x", h2[:])
	fmt.Printf("stringnbyte: %v\n", stringByte)

	var b bytes.Buffer
	b.Write([]byte(stringByte))
	b.Write([]byte(stringByte2))

	fmt.Printf("buf: %v\n", b.String())

	h12 := sha256.Sum256([]byte(b.String()))

	fmt.Printf("h12: %x\n", h12)

	//node12 = append(node12, h2...)

	//h1234 := sha256.New()
}

//func CreateTree(args []string) {}

// PrintSlice Print a given string slice
func PrintSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	for _, word := range s {
		fmt.Printf("Word: %v", word)
	}
}
