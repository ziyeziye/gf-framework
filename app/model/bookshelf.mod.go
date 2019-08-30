package model

import (
	"database/sql"
)

var BookshelfModel = Model(Bookshelf{})

type Bookshelf struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	BookID     int    `json:"book_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b Bookshelf) TableName() string {
	return GetPrefix() + "bookshelf"
}

func GetBookshelf(id int) (bookshelf *Bookshelf, err error) {
	err = BookshelfModel.Where("id", id).Struct(bookshelf)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetBookshelves(maps map[string]interface{}) (bookshelves []*Bookshelf, err error) {
	query, err := ModelSearch(Bookshelf{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(bookshelves)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetBookshelfTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = BookshelfModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteBookshelf(bookshelf *Bookshelf) (err error) {
	_, err = BookshelfModel.Where(bookshelf).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateBookshelf(bookshelf *Bookshelf, data map[string]interface{}) (err error) {
	_, err = BookshelfModel.Where(bookshelf).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddBookshelf(bookshelf *Bookshelf) (err error) {
	_, err = BookshelfModel.Data(bookshelf).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
