package response

import (
	"github.com/gogf/gf/g/net/ghttp"
)

type jsonType struct {
	r     *ghttp.Request `json:"-"`
	Code  int            `json:"code"`
	Msg   string         `json:"msg"`
	State bool           `json:"state"`
	Data  interface{}    `json:"data"`
}

func Json(r *ghttp.Request) (respJson *jsonType) {
	return &jsonType{r: r, Code: SUCCESS, State: true}
}

func (json *jsonType) Set(code int, msg string, state bool, data interface{}) {
	json.SetCode(code).SetMsg(msg).SetState(state).SetData(data)
}

func (json *jsonType) SetCode(code int) *jsonType {
	json.Code = code
	return json
}

func (json *jsonType) SetMsg(msg string) *jsonType {
	json.Msg = msg
	return json
}

func (json *jsonType) SetState(state bool) *jsonType {
	json.State = state
	return json
}

func (json *jsonType) SetData(data interface{}) *jsonType {
	json.Data = data
	return json
}

func (json *jsonType) Return() {
	if json.Msg == "" {
		json.Msg = GetMsg(json.Code) //ok
	}
	json.r.Response.WriteJson(json)
	json.r.Exit()
}
