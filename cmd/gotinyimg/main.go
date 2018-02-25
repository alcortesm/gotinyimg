package main

import (
	"fmt"
	"log"

	"github.com/alcortesm/gotinyimg"

	"golang.org/x/text/language"
)

func main() {
	msg, err := gotinyimg.Hello(language.English)
	if err != nil {
		log.Fatalf("helloing: %v", err)
	}
	fmt.Println(msg)
}
