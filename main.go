package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"

	"io"
	"log"
	"os"

	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ElifBerilSayli/bookLib"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ElifBerilSayli/database"
)

// initialize of global variables
var bookStructSlice = []bookLib.Book{}
var bookSlice = []string{}
var authSlice = []string{}

var (
	bookRepository *bookLib.BookRepository
)

// Error handling
var ErrInArgument = errors.New("ERROR: Arguments are invalid")
var ErrInId = errors.New("ERROR: Book ıd or another argument have problem. (Need integer value)")
var ErrInvalidIdNumber = errors.New("ERROR: Book ıd not found")

// Book related and database initializations operations
func init() {
	db := database.NewMySQLDB("root:Password123!@tcp(127.0.0.1:3306)/location?parseTime=True&loc=Local")
	bookRepository = bookLib.NewBookRepository(db)
	bookRepository.Migration()
	// Reading csv file to obtain author and book information
	bookSlice, authSlice = readCsvOfBooks()
	// Create new Book Structs
	for i := range bookSlice {
		n := bookLib.NewBook(strings.ToLower(bookSlice[i]), authSlice[i])
		bookStructSlice = append(bookStructSlice, n)
	}
	bookRepository.InsertData(bookStructSlice)

}

// Reading operations of csv and filling book and author arrays
func readCsvOfBooks() ([]string, []string) {

	f, err := os.Open("books.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cnt := 0
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		bookSlice = append(bookSlice, rec[0])
		authSlice = append(authSlice, rec[1])
		//fmt.Printf("%s\n", bookSlice[cnt])
		cnt = cnt + 1
	}
	return bookSlice, authSlice
}

func main() {

	args := os.Args

	// Arguments and operations to list search buy and delete books
	if args[1] == "list" {
		bookLib.List(bookStructSlice)
		return
	}

	err := checkCommandSize(args)
	if err != nil {
		fmt.Printf("error running program: %s \n", err.Error())
		return
	}
	// Checks whether user search book and operate searching process
	if args[1] == "search" {
		err := checkCommandSize(args)
		if err != nil {
			fmt.Printf("error running program: %s \n", err.Error())
		} else {
			var bookNameSlice []string
			for i := 2; i < len(args); i++ {
				bookNameSlice = append(bookNameSlice, args[i])
			}
			bookName := strings.Join(bookNameSlice, " ")
			books := bookRepository.FindByBookOrAuthorName(strings.ToLower(bookName))
			bookLib.List(books)
		}
	}
	// Checks whether user buy book and operate bought process
	if args[1] == "buy" {
		err := checkCommandSize(args)
		if err != nil {
			fmt.Printf("error running program: %s \n", err.Error())
		} else {
			ıd, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("error running program: %s \n", err.Error())
			} else {
				err = checkIdValidError(ıd)
				if err != nil {
					fmt.Printf("error running program: %s \n", err.Error())
				} else {
					numberOfBooksToBuy, err := strconv.Atoi(args[3])
					if err != nil {
						fmt.Printf("error running program: %s \n", err.Error())
					} else {
						books := bookRepository.UpdateStockNumber(numberOfBooksToBuy, bookRepository.GetById(ıd))
						fmt.Println(books)
					}
				}
			}
		}
	}
	// Checks whether user delete book and operate deletion process
	if args[1] == "delete" {
		deletionId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("error running program: %s \n", err.Error())
		} else {
			err = checkIdValidError(deletionId)
			if err != nil {
				fmt.Printf("error running program: %s \n", err.Error())
			} else {
				books := bookRepository.DeleteById(deletionId)
				if books == nil {
					fmt.Printf("Successful deletion !! \n")
				} else {
					fmt.Printf("Error in deletion !! \n")
				}
			}
		}

	}
}

//func checkCommandSize checks arguments size for error handling.
func checkCommandSize(args []string) error {
	if len(args) <= 2 {
		return ErrInArgument
	}
	return nil
}

//func checkIdValidError checks whether id is valid or not.
func checkIdValidError(id int) error {
	for i := range bookStructSlice {
		if int(bookStructSlice[i].ID) == id {
			return nil
		}
	}
	return ErrInvalidIdNumber
}
