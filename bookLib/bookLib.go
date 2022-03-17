package bookLib

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Book repository database related operations
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) FindAll() []Book {
	var books []Book
	r.db.Find(&books)

	return books
}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Book{})
}

func (r *BookRepository) GetById(id int) Book {
	var books Book
	result := r.db.First(&books, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Book not found with id : %d", id)
		return Book{}
	}
	return books
}

func (r *BookRepository) Create(c Book) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) InsertSampleData() {
	books := []Book{
		{BookName: "name1", StockCode: "string", ISBNno: 111,
			PageNumber: 1112, Price: 123, StockNumber: 123, IsDeleted: true},
	}
	fmt.Println(books)
	for _, c := range books {
		r.db.Create(&c)
	}
}

func (r *BookRepository) InsertData(bookStructSlice []Book) {
	for _, c := range bookStructSlice {
		r.db.Create(&c)

	}
}

func (r *BookRepository) Delete(c Book) error {
	result := r.db.Delete(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *BookRepository) DeleteById(id int) error {
	result := r.db.Delete(&Book{}, id)
	if result.Error != nil {
		return result.Error
	} else {

	}

	return nil
}
func (r *BookRepository) FindByName(name string) []Book {
	var books []Book
	r.db.Where("BookName LIKE ?", "%"+strings.ToLower(name)+"%").Find(&books)

	return books
}

func (r *BookRepository) FindByBookOrAuthorName(name string) []Book {
	var books []Book
	r.db.Where("BookName = ?", name).Or("Name = ?", name).Find(&books)
	return books
}

func (r *BookRepository) FindByNameWithRawSQL(name string) []Book {
	var books []Book
	r.db.Raw("SELECT * FROM Book WHERE BookName LIKE ?", "%"+name+"%").Scan(&books)

	return books
}

func (r *BookRepository) UpdateStockNumber(numberOfBooksToBuy int, c Book) error {
	result := r.db.Save(c)
	r.db.Model(&c).Update("StockNumber", (c.StockNumber - numberOfBooksToBuy))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func List(bookStructSlice []Book) {
	for _, c := range bookStructSlice {
		fmt.Println(c.ToString())
		fmt.Println()
	}
}
