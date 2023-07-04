package models

import (
	"bookmanagement/pkg/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectToDB()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {

	DB.Create(&b)
	return b
}

func GetAllBook() []Book {
	var books []Book
	DB.Find(&books)
	return books
}

func GetBookbyID(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	DB = DB.Where("ID=?", Id).Find(&getBook)
	return &getBook, DB

}

func DeleteBook(Id int64) Book {
	var book Book
	DB.Where("id=?", Id).Delete(book)
	return book
}
