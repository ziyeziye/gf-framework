package model

import (
	"database/sql"
)

var ChapterModel = Model(Chapter{})

type Chapter struct {
	ID          int    `json:"id"`
	BookID      int    `json:"book_id"`
	ChapterName string `json:"chapter_name"`
	Content     string `json:"content"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	Sort        int    `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (c Chapter) TableName() string {
	return GetPrefix() + "chapters"
}

func GetChapter(id int) (chapter *Chapter, err error) {
	err = ChapterModel.Where("id", id).Struct(chapter)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetChapters(maps map[string]interface{}) (chapters []*Chapter, err error) {
	query, err := ModelSearch(Chapter{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(chapters)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetChapterTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = ChapterModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteChapter(chapter *Chapter) (err error) {
	_, err = ChapterModel.Where(chapter).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateChapter(chapter *Chapter, data map[string]interface{}) (err error) {
	_, err = ChapterModel.Where(chapter).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddChapter(chapter *Chapter) (err error) {
	_, err = ChapterModel.Data(chapter).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
