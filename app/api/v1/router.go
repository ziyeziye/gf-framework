package v1

import "github.com/gogf/gf/g/net/ghttp"

func ConfigRouter(router *ghttp.RouterGroup) {
	ConfigTagsRouter(router)

}
