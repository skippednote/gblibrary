package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func makeDir() {
	logrus.Info("Creating new directories for Library.")
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
	logrus.Warn("Deleting existing library files and images.")
	os.RemoveAll("content/library")
	os.RemoveAll("static/imgs/library")
}

func setup() {
	cleanup()
	makeDir()
}

func generateDate(t time.Time) string {
	day := t.Day()
	month := int(t.Month())
	year := t.Year()
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(year), strconv.Itoa(month), strconv.Itoa(day))
}
