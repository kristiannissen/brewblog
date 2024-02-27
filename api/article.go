package handler

import (
	"log"
	"os"
	"strings"
)

func GetArticle() (string, error) {
	md, err := os.ReadFile("../content/sample.md")
	if err != nil {
		panic(err)
	}

	log.Printf("%q", strings.Split(string(md), "\n"))

	return string(md), err
}
