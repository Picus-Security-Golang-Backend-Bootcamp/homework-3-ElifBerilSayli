package bookLib

import (
	"fmt"
	"math/rand"
	"strconv"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName string
	AuthorInfo
	StockCode   string `gorm:"type:varchar(100);column:StockCode"`
	ISBNno      int
	PageNumber  int
	Price       int
	StockNumber int
	IsDeleted   bool
}

type AuthorInfo struct {
	Name  string
	Birth string
}

// Book struct constructor
func NewBook(bookName string, authName string) Book {
	p := new(Book)

	p.BookName = bookName
	p.ISBNno = rand.Intn(100)
	p.PageNumber = rand.Intn(100)
	p.Price = rand.Intn(100)
	p.Name = authName
	p.Birth = "1980"
	p.StockNumber = rand.Intn(100)
	p.StockCode = "book" + strconv.Itoa(p.ISBNno)
	p.IsDeleted = false
	return *p
}

func (Book) TableName() string {
	return "Book"
}

//Printing operations
func (c *Book) ToString() string {
	return fmt.Sprintf(" BookName : %s, AuthName : %s , StockCode : %s, ISBNno : %d ,StockNumber : %d", c.BookName, c.Name, c.StockCode, c.ISBNno, c.StockNumber)
}
