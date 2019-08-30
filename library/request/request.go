package request

import (
	"framework/config"

	"github.com/gogf/gf/g/net/ghttp"
)

func GetPage(r *ghttp.Request, maps map[string]interface{}, must bool) map[string]interface{} {
	page := r.GetInt("page", 0)
	if page > 0 || must {
		size := r.GetInt("size")

		if page < 1 {
			page = 1
		}

		maxSize := config.GetCfg().GetInt("max_page_size", 50)
		if size < 1 {
			size = config.GetCfg().GetInt("page_size", 10)
		} else if size > maxSize {
			size = maxSize
		}

		maps["page"] = (page - 1) * size
		maps["size"] = size
	}
	return maps
}
