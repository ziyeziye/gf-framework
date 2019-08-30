package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type ReplyApi struct {
	gmvc.Controller
}

func ConfigRepliesRouter(router *ghttp.RouterGroup) {
	controller := ReplyApi{}
	router.GET("/replies", controller.GetAllReplies)
	router.POST("/replies", controller.AddReply)
	router.GET("/replies/:id", controller.GetReply)
	router.PUT("/replies/:id", controller.UpdateReply)
	router.DELETE("/replies/:id", controller.DeleteReply)
}

func (c *ReplyApi) GetAllReplies(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetReplyTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if replies, err := model.GetReplies(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = replies
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *ReplyApi) GetReply(r *ghttp.Request) {
	id := r.GetParam("id")
	reply, err := model.GetReply(id.Int())

	respJson := response.Json(r)
	if err == nil && reply.ID > 0 {
		respJson.SetData(reply)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *ReplyApi) AddReply(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	reply := model.Reply{}

	respJson := response.Json(r)
	if err := model.AddReply(&reply); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, reply)
	} else {
		respJson.SetData(reply)
	}
	respJson.Return()
}

func (c *ReplyApi) UpdateReply(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	reply, err := model.GetReply(id.Int())

	respJson := response.Json(r)
	if err == nil && reply.ID > 0 {
		if err := model.UpdateReply(reply, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, reply)
		} else {
			respJson.SetData(reply)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *ReplyApi) DeleteReply(r *ghttp.Request) {
	id := r.GetParam("id")
	reply, err := model.GetReply(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteReply(reply); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
