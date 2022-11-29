package models

type Book struct {
	ID      int
	Title   string
	Author  string
	Release int
}

var Book1 = Book{
	Title:   "Math",
	Author:  "Einsten",
	Release: 1900,
}

var Book2 = Book{
	Title:   "Manga",
	Author:  "Kishimoto",
	Release: 2001,
}
