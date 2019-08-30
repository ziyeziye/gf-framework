package router

import (
	v1 "framework/app/api/v1"
	"framework/config"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

// 用于应用初始化。
func AppInit() *ghttp.Server {
	v := g.View()
	s := g.Server()

	// 模板引擎配置
	v.AddPath(config.GetRealPath("setting.view_path"))
	v.SetDelimiters("${", "}")

	// glog配置
	logpath := config.GetRealPath("setting.log_path")
	glog.SetPath(logpath)
	glog.SetStdoutPrint(true)
	//glog.SetDebug(config.GetCfg().GetBool("app.run_debug"))

	// Web Server配置
	s.SetServerRoot(config.GetRealPath("setting.public_path"))
	s.SetLogPath(logpath)
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)

	s.SetPort(config.GetCfg().GetInt("server.http_port", 8081))
	//s.SetReadTimeout(time.Duration(config.GetCfg().GetInt("server.read_timeout")))
	//s.SetWriteTimeout(time.Duration(config.GetCfg().GetInt("server.write_timeout")))

	router(s)
	return s
}

// 统一路由注册.
func router(s *ghttp.Server) {
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	//s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	// 用户模块 路由注册 - 使用执行对象注册方式
	//s.BindObject("/user", new(user.Controller))
	apiv1 := s.Group("/api/v1")
	v1.ConfigTagsRouter(apiv1)

}
