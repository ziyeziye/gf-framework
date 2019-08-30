package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type SpiderUrlApi struct {
	gmvc.Controller
}

func ConfigSpiderUrlsRouter(router *ghttp.RouterGroup) {
	controller := SpiderUrlApi{}
	router.GET("/spiderurls", controller.GetAllSpiderUrls)
	router.POST("/spiderurls", controller.AddSpiderUrl)
	router.GET("/spiderurls/:id", controller.GetSpiderUrl)
	router.PUT("/spiderurls/:id", controller.UpdateSpiderUrl)
	router.DELETE("/spiderurls/:id", controller.DeleteSpiderUrl)
}

func (c *SpiderUrlApi) GetAllSpiderUrls(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetSpiderUrlTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if spiderurls, err := model.GetSpiderUrls(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = spiderurls
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *SpiderUrlApi) GetSpiderUrl(r *ghttp.Request) {
	id := r.GetParam("id")
	spiderurl, err := model.GetSpiderUrl(id.Int())

	respJson := response.Json(r)
	if err == nil && spiderurl.ID > 0 {
		respJson.SetData(spiderurl)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *SpiderUrlApi) AddSpiderUrl(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	spiderurl := model.SpiderUrl{}

	respJson := response.Json(r)
	if err := model.AddSpiderUrl(&spiderurl); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, spiderurl)
	} else {
		respJson.SetData(spiderurl)
	}
	respJson.Return()
}

func (c *SpiderUrlApi) UpdateSpiderUrl(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	spiderurl, err := model.GetSpiderUrl(id.Int())

	respJson := response.Json(r)
	if err == nil && spiderurl.ID > 0 {
		if err := model.UpdateSpiderUrl(spiderurl, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, spiderurl)
		} else {
			respJson.SetData(spiderurl)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *SpiderUrlApi) DeleteSpiderUrl(r *ghttp.Request) {
	id := r.GetParam("id")
	spiderurl, err := model.GetSpiderUrl(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteSpiderUrl(spiderurl); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
