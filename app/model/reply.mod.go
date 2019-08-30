package model

import (
	"database/sql"
)

var ReplyModel = Model(Reply{})

type Reply struct {
	ID         int    `json:"id"`
	CommentID  int    `json:"comment_id"`
	FromUserID int    `json:"from_user_id"`
	ToUserID   int    `json:"to_user_id"`
	ReplyMsg   string `json:"reply_msg"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Like       int    `json:"like"`
	Dislike    int    `json:"dislike"`
}

// TableName sets the insert table name for this struct type
func (r Reply) TableName() string {
	return GetPrefix() + "reply"
}

func GetReply(id int) (reply *Reply, err error) {
	err = ReplyModel.Where("id", id).Struct(reply)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func GetReplies(maps map[string]interface{}) (replies []*Reply, err error) {
	query, err := ModelSearch(Reply{}, maps)
	if err != nil {
		return
	}
	err = query.Structs(replies)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return
}

func GetReplyTotal(maps map[string]interface{}) (count int, err error) {
	cond, values, err := whereBuild(maps)
	if err != nil {
		return 0, err
	}
	count, err = ReplyModel.Filter().Where(cond, values...).Count()

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return
}

func DeleteReply(reply *Reply) (err error) {
	_, err = ReplyModel.Where(reply).Delete()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func UpdateReply(reply *Reply, data map[string]interface{}) (err error) {
	_, err = ReplyModel.Where(reply).Data(data).Update()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func AddReply(reply *Reply) (err error) {
	_, err = ReplyModel.Data(reply).Insert()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}
