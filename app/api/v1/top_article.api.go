package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type TopArticleApi struct {
	gmvc.Controller
}

func ConfigTopArticlesRouter(router *ghttp.RouterGroup) {
	controller := TopArticleApi{}
	router.GET("/toparticles", controller.GetAllTopArticles)
	router.POST("/toparticles", controller.AddTopArticle)
	router.GET("/toparticles/:id", controller.GetTopArticle)
	router.PUT("/toparticles/:id", controller.UpdateTopArticle)
	router.DELETE("/toparticles/:id", controller.DeleteTopArticle)
}

func (c *TopArticleApi) GetAllTopArticles(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetTopArticleTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if toparticles, err := model.GetTopArticles(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = toparticles
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *TopArticleApi) GetTopArticle(r *ghttp.Request) {
	id := r.GetParam("id")
	toparticle, err := model.GetTopArticle(id.Int())

	respJson := response.Json(r)
	if err == nil && toparticle.ID > 0 {
		respJson.SetData(toparticle)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *TopArticleApi) AddTopArticle(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	toparticle := model.TopArticle{}

	respJson := response.Json(r)
	if err := model.AddTopArticle(&toparticle); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, toparticle)
	} else {
		respJson.SetData(toparticle)
	}
	respJson.Return()
}

func (c *TopArticleApi) UpdateTopArticle(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	toparticle, err := model.GetTopArticle(id.Int())

	respJson := response.Json(r)
	if err == nil && toparticle.ID > 0 {
		if err := model.UpdateTopArticle(toparticle, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, toparticle)
		} else {
			respJson.SetData(toparticle)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *TopArticleApi) DeleteTopArticle(r *ghttp.Request) {
	id := r.GetParam("id")
	toparticle, err := model.GetTopArticle(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteTopArticle(toparticle); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
