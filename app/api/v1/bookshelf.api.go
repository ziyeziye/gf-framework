package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type BookshelfApi struct {
	gmvc.Controller
}

func ConfigBookshelvesRouter(router *ghttp.RouterGroup) {
	controller := BookshelfApi{}
	router.GET("/bookshelves", controller.GetAllBookshelves)
	router.POST("/bookshelves", controller.AddBookshelf)
	router.GET("/bookshelves/:id", controller.GetBookshelf)
	router.PUT("/bookshelves/:id", controller.UpdateBookshelf)
	router.DELETE("/bookshelves/:id", controller.DeleteBookshelf)
}

func (c *BookshelfApi) GetAllBookshelves(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetBookshelfTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if bookshelves, err := model.GetBookshelves(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = bookshelves
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *BookshelfApi) GetBookshelf(r *ghttp.Request) {
	id := r.GetParam("id")
	bookshelf, err := model.GetBookshelf(id.Int())

	respJson := response.Json(r)
	if err == nil && bookshelf.ID > 0 {
		respJson.SetData(bookshelf)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *BookshelfApi) AddBookshelf(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	bookshelf := model.Bookshelf{}

	respJson := response.Json(r)
	if err := model.AddBookshelf(&bookshelf); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, bookshelf)
	} else {
		respJson.SetData(bookshelf)
	}
	respJson.Return()
}

func (c *BookshelfApi) UpdateBookshelf(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	bookshelf, err := model.GetBookshelf(id.Int())

	respJson := response.Json(r)
	if err == nil && bookshelf.ID > 0 {
		if err := model.UpdateBookshelf(bookshelf, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, bookshelf)
		} else {
			respJson.SetData(bookshelf)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *BookshelfApi) DeleteBookshelf(r *ghttp.Request) {
	id := r.GetParam("id")
	bookshelf, err := model.GetBookshelf(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteBookshelf(bookshelf); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
