package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type ChapterApi struct {
	gmvc.Controller
}

func ConfigChaptersRouter(router *ghttp.RouterGroup) {
	controller := ChapterApi{}
	router.GET("/chapters", controller.GetAllChapters)
	router.POST("/chapters", controller.AddChapter)
	router.GET("/chapters/:id", controller.GetChapter)
	router.PUT("/chapters/:id", controller.UpdateChapter)
	router.DELETE("/chapters/:id", controller.DeleteChapter)
}

func (c *ChapterApi) GetAllChapters(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetChapterTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if chapters, err := model.GetChapters(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = chapters
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *ChapterApi) GetChapter(r *ghttp.Request) {
	id := r.GetParam("id")
	chapter, err := model.GetChapter(id.Int())

	respJson := response.Json(r)
	if err == nil && chapter.ID > 0 {
		respJson.SetData(chapter)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *ChapterApi) AddChapter(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	chapter := model.Chapter{}

	respJson := response.Json(r)
	if err := model.AddChapter(&chapter); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, chapter)
	} else {
		respJson.SetData(chapter)
	}
	respJson.Return()
}

func (c *ChapterApi) UpdateChapter(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	chapter, err := model.GetChapter(id.Int())

	respJson := response.Json(r)
	if err == nil && chapter.ID > 0 {
		if err := model.UpdateChapter(chapter, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, chapter)
		} else {
			respJson.SetData(chapter)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *ChapterApi) DeleteChapter(r *ghttp.Request) {
	id := r.GetParam("id")
	chapter, err := model.GetChapter(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteChapter(chapter); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
