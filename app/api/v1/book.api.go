package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type BookApi struct {
	gmvc.Controller
}

func ConfigBooksRouter(router *ghttp.RouterGroup) {
	controller := BookApi{}
	router.GET("/books", controller.GetAllBooks)
	router.POST("/books", controller.AddBook)
	router.GET("/books/:id", controller.GetBook)
	router.PUT("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)
}

func (c *BookApi) GetAllBooks(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetBookTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if books, err := model.GetBooks(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = books
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *BookApi) GetBook(r *ghttp.Request) {
	id := r.GetParam("id")
	book, err := model.GetBook(id.Int())

	respJson := response.Json(r)
	if err == nil && book.ID > 0 {
		respJson.SetData(book)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *BookApi) AddBook(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	book := model.Book{}

	respJson := response.Json(r)
	if err := model.AddBook(&book); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, book)
	} else {
		respJson.SetData(book)
	}
	respJson.Return()
}

func (c *BookApi) UpdateBook(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	book, err := model.GetBook(id.Int())

	respJson := response.Json(r)
	if err == nil && book.ID > 0 {
		if err := model.UpdateBook(book, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, book)
		} else {
			respJson.SetData(book)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *BookApi) DeleteBook(r *ghttp.Request) {
	id := r.GetParam("id")
	book, err := model.GetBook(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteBook(book); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
