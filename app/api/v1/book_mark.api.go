package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type BookMarkApi struct {
	gmvc.Controller
}

func ConfigBookMarksRouter(router *ghttp.RouterGroup) {
	controller := BookMarkApi{}
	router.GET("/bookmarks", controller.GetAllBookMarks)
	router.POST("/bookmarks", controller.AddBookMark)
	router.GET("/bookmarks/:id", controller.GetBookMark)
	router.PUT("/bookmarks/:id", controller.UpdateBookMark)
	router.DELETE("/bookmarks/:id", controller.DeleteBookMark)
}

func (c *BookMarkApi) GetAllBookMarks(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetBookMarkTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if bookmarks, err := model.GetBookMarks(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = bookmarks
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *BookMarkApi) GetBookMark(r *ghttp.Request) {
	id := r.GetParam("id")
	bookmark, err := model.GetBookMark(id.Int())

	respJson := response.Json(r)
	if err == nil && bookmark.ID > 0 {
		respJson.SetData(bookmark)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *BookMarkApi) AddBookMark(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	bookmark := model.BookMark{}

	respJson := response.Json(r)
	if err := model.AddBookMark(&bookmark); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, bookmark)
	} else {
		respJson.SetData(bookmark)
	}
	respJson.Return()
}

func (c *BookMarkApi) UpdateBookMark(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	bookmark, err := model.GetBookMark(id.Int())

	respJson := response.Json(r)
	if err == nil && bookmark.ID > 0 {
		if err := model.UpdateBookMark(bookmark, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, bookmark)
		} else {
			respJson.SetData(bookmark)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *BookMarkApi) DeleteBookMark(r *ghttp.Request) {
	id := r.GetParam("id")
	bookmark, err := model.GetBookMark(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteBookMark(bookmark); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
