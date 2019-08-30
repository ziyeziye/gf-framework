package model

import (
	"database/sql"
)

var TopArticleModel = Model(TopArticle{})

type TopArticle struct {
	ID       int    `json:"id"`
	Str      string `json:"str"`
	DataType string `json:"dataType"`
	Title    string `json:"title"`
	IsShow   int    `json:"isShow"`
	Rss      string `json:"rss"`
}

// TableName sets the insert table name for this struct type
func (t TopArticle) TableName() string {
	return GetPrefix() + "top_articles"
}

func GetTopArticle(id int) (toparticle *TopArticle, err error) {
	err = TopArticleModel.Where("id", id).Struct(toparticle)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetTopArticles(maps map[string]interface{}) (toparticles []*TopArticle, err error) {
	query, err := ModelSearch(TopArticle{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(toparticles)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetTopArticleTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = TopArticleModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteTopArticle(toparticle *TopArticle) (err error) {
	_, err = TopArticleModel.Where(toparticle).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateTopArticle(toparticle *TopArticle, data map[string]interface{}) (err error) {
	_, err = TopArticleModel.Where(toparticle).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddTopArticle(toparticle *TopArticle) (err error) {
	_, err = TopArticleModel.Data(toparticle).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
