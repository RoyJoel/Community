package web

import (
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/config"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/web/controller"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/web/interceptor"
	"github.com/gin-gonic/gin"
)

func RunHttp() {
	r := gin.Default()
	//增加拦截器
	r.Use(interceptor.HttpInterceptor())
	//解决跨域
	r.Use(config.CorsConfig())

		//路由组
	commonInfo := r.Group("/")
	{
		commonInfo.POST("auth", controller.NewIHCommunityControllerImpl().Auth)
		commonInfo.POST("get-proposal-info", controller.NewIHCommunityControllerImpl().GetProposalInfo)
	}

	userInfo := r.Group("/user")
	{
		userInfo.POST("/add", controller.NewIHCommunityControllerImpl().AddUser)
	}

	commentInfo := r.Group("/comment")
	{
		commentInfo.POST("/get-list", controller.NewIHCommunityControllerImpl().GetCommentsForUser)
		commentInfo.POST("/update", controller.NewIHCommunityControllerImpl().UpdateComment)
		commentInfo.POST("/delete", controller.NewIHCommunityControllerImpl().DeleteComment)
		commentInfo.POST("/add", controller.NewIHCommunityControllerImpl().AddComment)
	}

	postInfo := r.Group("/post")
	{
		postInfo.POST("/get-list", controller.NewIHCommunityControllerImpl().GetAllPosts)
		postInfo.POST("/update", controller.NewIHCommunityControllerImpl().UpdatePost)
		postInfo.POST("/delete", controller.NewIHCommunityControllerImpl().DeletePost)
		postInfo.POST("/add", controller.NewIHCommunityControllerImpl().AddPost)
		postInfo.POST("/save", controller.NewIHCommunityControllerImpl().SavePost)
	}

	adminInfo := r.Group("/admin")
	{
		adminInfo.POST("/auth", controller.NewIHCommunityControllerImpl().AuthAdmin)
	}

	r.Run(":8080")
}