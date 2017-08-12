package main

type book struct {
	ISBN      string `yaml:"isbn"`
	Rating    int    `yaml:"rating"`
	Status    string `yaml:"status"`
	Title     string `yaml:"title"`
	Image     string `yaml:"image"`
	Author    string `yaml:"author"`
	Slug      string `yaml:"slug"`
	ImagePath string `yaml:"imagepath"`
	Date      string `yaml:"date"`
}

type books []book
