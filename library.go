package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gosimple/slug"
	"github.com/sillyotter/gbsearch"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var gb = gbsearch.Options{}
var tmpl *template.Template

func getBooksList() (books, error) {
	f, err := os.Open("books.yml")
	if err != nil {
		return nil, err
	}

	fc, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var bs []book
	err = yaml.Unmarshal(fc, &bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func getBook(b book, c chan book) {
	r, err := gbsearch.ISBNSearch(b.ISBN, &gb)
	if err != nil {
		log.Fatal(err)
	}

	b.Author = r.Items[0].VolumeInfo.Authors[0]
	b.Title = r.Items[0].VolumeInfo.Title
	b.Image = r.Items[0].VolumeInfo.ImageLinks.Thumbnail
	b.Slug = slug.Make(r.Items[0].VolumeInfo.Title)
	b.ImagePath = "/imgs/library/" + b.Slug + ".jpg"

	c <- b
}

func saveImage(b book) {
	f, err := os.Create("static" + b.ImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := http.Get(b.Image)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	io.Copy(f, res.Body)
}

func saveBook(b book) {
	t, err := time.Parse("2006-01-02", b.Date)
	if err != nil {
		log.Fatal(err)
	}
	date := generateDate(t)
	bookPath := "content/library/" + date + "-" + b.Slug + ".md"

	f, err := os.Create(bookPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tmpl = template.Must(template.ParseFiles("book.tmpl.html"))
	tmpl.Execute(f, b)
}

func save(b book) {
	logrus.WithFields(logrus.Fields{
		"Title":  b.Title,
		"Author": b.Author,
	}).Info("Saving book to Library.")
	saveImage(b)
	saveBook(b)
}
