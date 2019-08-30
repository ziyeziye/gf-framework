package model

import (
	"database/sql"
)

var CommentModel = Model(Comment{})

type Comment struct {
	ID         int    `json:"id"`
	IssueID    int    `json:"issue_id"`
	UserID     int    `json:"user_id"`
	ReplyMsg   string `json:"reply_msg"`
	Like       int    `json:"like"`
	Dislike    int    `json:"dislike"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (c Comment) TableName() string {
	return GetPrefix() + "comment"
}

func GetComment(id int) (comment *Comment, err error) {
	err = CommentModel.Where("id", id).Struct(comment)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetComments(maps map[string]interface{}) (comments []*Comment, err error) {
	query, err := ModelSearch(Comment{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(comments)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetCommentTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = CommentModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteComment(comment *Comment) (err error) {
	_, err = CommentModel.Where(comment).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateComment(comment *Comment, data map[string]interface{}) (err error) {
	_, err = CommentModel.Where(comment).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddComment(comment *Comment) (err error) {
	_, err = CommentModel.Data(comment).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
