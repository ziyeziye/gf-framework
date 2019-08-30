package model

import (
	"database/sql"
)

var BookModel = Model(Book{})

type Book struct {
	ID         int    `json:"id"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"book_author"`
	BookImg    string `json:"book_img"`
	BookDesc   string `json:"book_desc"`
	BookTag    string `json:"book_tag"`
	State      int    `json:"state"`
	Hot        int    `json:"hot"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b Book) TableName() string {
	return GetPrefix() + "books"
}

func GetBook(id int) (book *Book, err error) {
	err = BookModel.Where("id", id).Struct(book)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetBooks(maps map[string]interface{}) (books []*Book, err error) {
	query, err := ModelSearch(Book{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(books)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetBookTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = BookModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteBook(book *Book) (err error) {
	_, err = BookModel.Where(book).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateBook(book *Book, data map[string]interface{}) (err error) {
	_, err = BookModel.Where(book).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddBook(book *Book) (err error) {
	_, err = BookModel.Data(book).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
