package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type TagApi struct {
	gmvc.Controller
}

func ConfigTagsRouter(router *ghttp.RouterGroup) {
	controller := TagApi{}
	router.GET("/tags", controller.GetAllTags)
	router.POST("/tags", controller.AddTag)
	router.GET("/tags/:id", controller.GetTag)
	router.PUT("/tags/:id", controller.UpdateTag)
	router.DELETE("/tags/:id", controller.DeleteTag)
}

func (c *TagApi) GetAllTags(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetTagTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if tags, err := model.GetTags(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = tags
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *TagApi) GetTag(r *ghttp.Request) {
	id := r.GetParam("id")
	tag, err := model.GetTag(id.Int())

	respJson := response.Json(r)
	if err == nil && tag.ID > 0 {
		respJson.SetData(tag)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *TagApi) AddTag(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	tag := model.Tag{}

	respJson := response.Json(r)
	if err := model.AddTag(&tag); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, tag)
	} else {
		respJson.SetData(tag)
	}
	respJson.Return()
}

func (c *TagApi) UpdateTag(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	tag, err := model.GetTag(id.Int())

	respJson := response.Json(r)
	if err == nil && tag.ID > 0 {
		if err := model.UpdateTag(tag, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, tag)
		} else {
			respJson.SetData(tag)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *TagApi) DeleteTag(r *ghttp.Request) {
	id := r.GetParam("id")
	tag, err := model.GetTag(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteTag(tag); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
