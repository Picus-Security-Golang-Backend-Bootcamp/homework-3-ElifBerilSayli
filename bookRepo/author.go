package bookRepo

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name     string
	Birth    string
	BookName string
}

func NewAuthor(aName string, birth string, bookName string) Author {
	p := new(Author)

	p.BookName = bookName
	p.Birth = birth
	p.Name = aName

	return *p
}

func (Author) TableName() string {
	return "Author"
}

//Printing operations
func (c *Author) ToString() string {
	return fmt.Sprintf(" BookName : %s AuthName : %s AuthBirth : %s", c.BookName, c.Name, c.Birth)
}
