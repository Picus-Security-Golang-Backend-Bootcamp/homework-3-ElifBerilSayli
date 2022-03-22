package bookRepo

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Author repository database related operations
type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) FindAll() []Author {
	var authors []Author
	r.db.Find(&authors)

	return authors
}

func (r *AuthorRepository) Migration() {
	r.db.AutoMigrate(&Author{})
}

func (r *AuthorRepository) GetById(id int) Author {
	var authors Author
	result := r.db.First(&authors, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Author not found with id : %d", id)
		return Author{}
	}
	return authors
}

func (r *AuthorRepository) Create(c Author) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AuthorRepository) InsertData(AuthorStructSlice []Author) {
	for _, c := range AuthorStructSlice {
		r.db.Create(&c)

	}
}

func (r *AuthorRepository) Delete(c Author) error {
	result := r.db.Delete(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *AuthorRepository) DeleteById(id int) error {
	result := r.db.Delete(&Author{}, id)
	if result.Error != nil {
		return result.Error
	} else {

	}

	return nil
}
func (r *AuthorRepository) FindByName(name string) []Author {
	var Authors []Author
	r.db.Where("AuthorName LIKE ?", "%"+strings.ToLower(name)+"%").Find(&Authors)

	return Authors
}

func (r *AuthorRepository) FindByAuthorOrAuthorName(name string) []Author {
	var Authors []Author
	r.db.Where("AuthorName = ?", name).Or("Name = ?", name).Find(&Authors)
	return Authors
}

func (r *AuthorRepository) FindByNameWithRawSQL(name string) []Author {
	var Authors []Author
	r.db.Raw("SELECT * FROM Author WHERE AuthorName LIKE ?", "%"+name+"%").Scan(&Authors)

	return Authors
}

func ListAuthor(AuthorStructSlice []Author) {
	for _, c := range AuthorStructSlice {
		fmt.Println(c.ToString())
		fmt.Println()
	}
}
