package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type CommentApi struct {
	gmvc.Controller
}

func ConfigCommentsRouter(router *ghttp.RouterGroup) {
	controller := CommentApi{}
	router.GET("/comments", controller.GetAllComments)
	router.POST("/comments", controller.AddComment)
	router.GET("/comments/:id", controller.GetComment)
	router.PUT("/comments/:id", controller.UpdateComment)
	router.DELETE("/comments/:id", controller.DeleteComment)
}

func (c *CommentApi) GetAllComments(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetCommentTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if comments, err := model.GetComments(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = comments
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *CommentApi) GetComment(r *ghttp.Request) {
	id := r.GetParam("id")
	comment, err := model.GetComment(id.Int())

	respJson := response.Json(r)
	if err == nil && comment.ID > 0 {
		respJson.SetData(comment)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *CommentApi) AddComment(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	comment := model.Comment{}

	respJson := response.Json(r)
	if err := model.AddComment(&comment); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, comment)
	} else {
		respJson.SetData(comment)
	}
	respJson.Return()
}

func (c *CommentApi) UpdateComment(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	comment, err := model.GetComment(id.Int())

	respJson := response.Json(r)
	if err == nil && comment.ID > 0 {
		if err := model.UpdateComment(comment, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, comment)
		} else {
			respJson.SetData(comment)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *CommentApi) DeleteComment(r *ghttp.Request) {
	id := r.GetParam("id")
	comment, err := model.GetComment(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteComment(comment); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
