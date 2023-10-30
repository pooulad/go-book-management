package models

import (
	"github.com/pooulad/go-book-management/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `gorm:"" json:"publication"`
}

func init() {
	config.Connect()

	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func (b *Book) DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
