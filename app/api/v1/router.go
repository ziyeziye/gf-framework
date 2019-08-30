package v1

import "github.com/gogf/gf/g/net/ghttp"

func ConfigRouter(router *ghttp.RouterGroup) {
	ConfigBookMarksRouter(router)
	ConfigBooksRouter(router)
	ConfigBookshelvesRouter(router)
	ConfigChaptersRouter(router)
	ConfigCommentsRouter(router)
	ConfigRepliesRouter(router)
	ConfigSpiderUrlsRouter(router)
	ConfigTagsRouter(router)
	ConfigTopArticlesRouter(router)
	ConfigUsersRouter(router)

}
