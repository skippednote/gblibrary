package main

type book struct {
	ISBN      string `yaml:"isbn" json:"isbn"`
	Rating    int    `yaml:"rating" json:"rating"`
	Status    string `yaml:"read" json:"read"`
	Title     string `yaml:"title" json:"title"`
	Image     string `yaml:"image" json:"image"`
	Author    string `yaml:"author" json:"author"`
	Slug      string `yaml:"slug" json:"slug"`
	ImagePath string `yaml:"imagepath" json:"imagepath"`
}

type books []book
