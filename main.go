package main

import "fmt"

func main() {
	fmt.Println("hello")
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 24 {
		return "long"
	}
	return "short"
}
