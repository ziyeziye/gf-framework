package v1

import (
	"framework/library/request"
	"framework/library/response"
	"net/http"

	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/net/ghttp"

	"framework/app/model"
)

type UserApi struct {
	gmvc.Controller
}

func ConfigUsersRouter(router *ghttp.RouterGroup) {
	controller := UserApi{}
	router.GET("/users", controller.GetAllUsers)
	router.POST("/users", controller.AddUser)
	router.GET("/users/:id", controller.GetUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}

func (c *UserApi) GetAllUsers(r *ghttp.Request) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var total int
	total, _ = model.GetUserTotal(maps)
	data["total"] = total

	maps = request.GetPage(r, maps, false)
	respJson := response.Json(r)
	if users, err := model.GetUsers(maps); err != nil {
		respJson.SetState(false).SetMsg("error")
	} else {
		data["list"] = users
		respJson.SetData(data)
	}

	respJson.Return()
}

func (c *UserApi) GetUser(r *ghttp.Request) {
	id := r.GetParam("id")
	user, err := model.GetUser(id.Int())

	respJson := response.Json(r)
	if err == nil && user.ID > 0 {
		respJson.SetData(user)
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *UserApi) AddUser(r *ghttp.Request) {
	//maps := make(map[string]interface{})

	user := model.User{}

	respJson := response.Json(r)
	if err := model.AddUser(&user); err != nil {
		respJson.Set(http.StatusInternalServerError, "新增失败", false, user)
	} else {
		respJson.SetData(user)
	}
	respJson.Return()
}

func (c *UserApi) UpdateUser(r *ghttp.Request) {
	id := r.GetParam("id")
	maps := make(map[string]interface{})

	user, err := model.GetUser(id.Int())

	respJson := response.Json(r)
	if err == nil && user.ID > 0 {
		if err := model.UpdateUser(user, maps); err != nil {
			respJson.Set(http.StatusInternalServerError, "修改失败", false, user)
		} else {
			respJson.SetData(user)
		}
	} else {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}
	respJson.Return()
}

func (c *UserApi) DeleteUser(r *ghttp.Request) {
	id := r.GetParam("id")
	user, err := model.GetUser(id.Int())

	respJson := response.Json(r)
	if err != nil {
		respJson.SetState(false).SetCode(response.ERROR_NOT_EXIST)
	}

	if err := model.DeleteUser(user); err != nil {
		respJson.Set(http.StatusInternalServerError, "删除失败", false, nil)
	}
	respJson.Return()
}
