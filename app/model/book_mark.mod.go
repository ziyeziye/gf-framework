package model

import (
	"database/sql"
)

var BookMarkModel = Model(BookMark{})

type BookMark struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	BookID       int    `json:"book_id"`
	ChapterID    int    `json:"chapter_id"`
	MarkType     string `json:"mark_type"`
	ChapterName  string `json:"chapter_name"`
	ChapterValue string `json:"chapter_value"`
	CommText     string `json:"comm_text"`
	Range        string `json:"range"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b BookMark) TableName() string {
	return GetPrefix() + "book_marks"
}

func GetBookMark(id int) (bookmark *BookMark, err error) {
	err = BookMarkModel.Where("id", id).Struct(bookmark)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetBookMarks(maps map[string]interface{}) (bookmarks []*BookMark, err error) {
	query, err := ModelSearch(BookMark{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(bookmarks)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetBookMarkTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = BookMarkModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteBookMark(bookmark *BookMark) (err error) {
	_, err = BookMarkModel.Where(bookmark).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateBookMark(bookmark *BookMark, data map[string]interface{}) (err error) {
	_, err = BookMarkModel.Where(bookmark).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddBookMark(bookmark *BookMark) (err error) {
	_, err = BookMarkModel.Data(bookmark).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
