package main

import (
	"log"
	"strings"
)

func main() {

	key := "abc=123&q=123&ccc=123123"
	i := strings.IndexAny(key,"&;")
	log.Println(key[:i],key[i+1:])

}
