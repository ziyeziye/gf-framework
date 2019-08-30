package model

import (
	"database/sql"
)

var TagModel = Model(Tag{})

type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Hot        int    `json:"hot"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t Tag) TableName() string {
	return GetPrefix() + "tags"
}

func GetTag(id int) (tag *Tag, err error) {
	err = TagModel.Where("id", id).Struct(&tag)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetTags(maps map[string]interface{}) (tags []*Tag, err error) {
	query, err := ModelSearch(Tag{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(&tags)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetTagTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = TagModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteTag(tag *Tag) (err error) {
	_, err = TagModel.Where(tag).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateTag(tag *Tag, data map[string]interface{}) (err error) {
	_, err = TagModel.Where(tag).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddTag(tag *Tag) (err error) {
	_, err = TagModel.Data(tag).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
