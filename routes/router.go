package routes

import (
	"hh_blog/api/v1/tool"
	"hh_blog/config"
	"hh_blog/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	gin.SetMode(config.GinMode)
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)
	// 解决跨域问题
	r.Use(utils.Cors())

	r.HTMLRender = createMyRender()

	r.Static("/js", "./web/front/dist/js")
	r.Static("/css", "./web/front/dist/css")
	r.Static("/fonts", "./web/front/dist/fonts")
	r.StaticFile("/favicon.ico", "./web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "frontHtml", nil)
	})

	v1Group := r.Group("/api/v1/")
	v1Group.POST("dt_deal", tool.TimeUnixSecDeal)
	r.Run(config.GinListenPort)
}

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("frontHtml", "./web/front/dist/index.html")
	return p
}
