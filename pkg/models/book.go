package models

type Book struct {
	ID      int
	Title   string
	Author  string
	Release int
}

var Store = make([]Book, 0)

var Books = []Book{
	{
		Title:   "Naruto",
		Author:  "Kishimoto",
		Release: 2002,
	},
	{
		Title:   "Shelock Holmes",
		Author:  "Sir Arthur Conan",
		Release: 1895,
	},
	{
		Title:   "Algorithm",
		Author:  "Noel Kun",
		Release: 2016,
	},
	{
		Title:   "Art",
		Author:  "Thomas Doll El-Kunbar",
		Release: 1574,
	},
	{
		Title:   "The History of Sociolocius",
		Author:  "Kholib Al Kimir Mauk",
		Release: 291,
	},
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

/*
[
    {
		"Title":   "Naruto",
		"Author":  "Kishimoto",
		"Release": 2002
	},
	{
		"Title":   "Shelock Holmes",
		"Author":  "Sir Arthur Conan",
		"Release": 1895
	},
	{
		"Title":   "Algorithm",
		"Author":  "Noel Kun",
		"Release": 2016
	},
	{
		"Title":   "Art",
		"Author":  "Thomas Doll El-Kunbar",
		"Release": 1574
	},
	{
		"Title":   "The History of Sociolocius",
		"Author":  "Kholib Al Kimir Mauk",
		"Release": 291
	}
]
*/
