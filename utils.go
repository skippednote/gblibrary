package main

import (
	"log"
	"os"
)

func makeDir() {
	err := os.MkdirAll("content/library", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("static/imgs/library", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	os.Remove("content/library")
	os.Remove("static/imgs/library")
	makeDir()
}
