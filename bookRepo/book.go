package bookRepo

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName string
	Author
	StockCode   string `gorm:"type:varchar(100);column:StockCode"`
	ISBNno      int
	PageNumber  int
	Price       int
	StockNumber int
	IsDeleted   bool
}

// Book struct constructor
func NewBook(bookName string, authName string, stockCodeSlice string, ISBSlice int, pageNumber int, price int, stockNumber int) Book {
	p := new(Book)

	p.BookName = bookName
	p.ISBNno = ISBSlice
	p.PageNumber = pageNumber
	p.Price = price
	p.StockNumber = stockNumber
	p.StockCode = stockCodeSlice
	p.IsDeleted = false
	return *p
}

func (Book) TableName() string {
	return "Book"
}

//Printing operations
func (c *Book) ToString() string {
	return fmt.Sprintf(" BookName : %s,  StockCode : %s, ISBNno : %d ,StockNumber : %d", c.BookName, c.StockCode, c.ISBNno, c.StockNumber)
}
