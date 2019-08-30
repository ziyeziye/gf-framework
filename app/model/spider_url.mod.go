package model

import (
	"database/sql"
)

var SpiderUrlModel = Model(SpiderUrl{})

type SpiderUrl struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Host       string `json:"host"`
	Type       string `json:"type"`
	LinkID     int    `json:"link_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (s SpiderUrl) TableName() string {
	return GetPrefix() + "spider_urls"
}

func GetSpiderUrl(id int) (spiderurl *SpiderUrl, err error) {
	err = SpiderUrlModel.Where("id", id).Struct(spiderurl)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetSpiderUrls(maps map[string]interface{}) (spiderurls []*SpiderUrl, err error) {
	query, err := ModelSearch(SpiderUrl{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(spiderurls)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetSpiderUrlTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = SpiderUrlModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteSpiderUrl(spiderurl *SpiderUrl) (err error) {
	_, err = SpiderUrlModel.Where(spiderurl).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateSpiderUrl(spiderurl *SpiderUrl, data map[string]interface{}) (err error) {
	_, err = SpiderUrlModel.Where(spiderurl).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddSpiderUrl(spiderurl *SpiderUrl) (err error) {
	_, err = SpiderUrlModel.Data(spiderurl).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
