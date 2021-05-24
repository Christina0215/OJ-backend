package route

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/http"
	"qkcode/controller/auth"
	"qkcode/controller/contest"
	"qkcode/controller/file"
	"qkcode/controller/problem"
	"qkcode/controller/record"
	"qkcode/middleware"
	"qkcode/controller/rank"
	"runtime"
)

func AddRoute() {
	global := http.Router.Group("/api")
	global.Use(middleware.FilterOptions())
	{

		global.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"version": runtime.Version(),
			})
		})
		_auth := global.Group("/auth")
		{
			_auth.POST("/register", auth.Register)
			_auth.POST("/login", auth.Login)
			_auth.POST("/code", auth.GetCode)
			_auth.Use(middleware.AuthServiceProvider()).GET("", auth.GetAuth)
		}

		_problem := global.Group("/problem")
		{
			_problem.Use(middleware.AuthServiceProvider()).POST("", problem.Create)
			_problem.GET("", problem.GetList)
			_problem.GET("/:problemId", problem.GetDetail)
			_problem.Use(middleware.AuthServiceProvider()).PUT("/:problemId", problem.Modify)
			_problem.Use(middleware.AuthServiceProvider()).DELETE("/:problemId", problem.Delete)

			_record := _problem.Group("/:problemId").Use(middleware.AuthServiceProvider())
			{
				_record.POST("/record", record.Create)
				_record.GET("/record", record.GetList)
				_record.GET("/record/:recordId", record.GetDetail)

			}
		}

		_contest := global.Group("/contest")
		{
			_contest.Use(middleware.AuthServiceProvider()).POST("", contest.Create)
			_contest.GET("", contest.GetList)
			_contest.Use(middleware.AuthServiceProvider()).PUT("/:contestId", contest.Modify)
			_contest.GET("/:contestId", contest.GetDetail)
			//_contest.Use(Middleware.Admin()).GET("/contest/:contestId/rank", Contest.GetRank)
			//_contest.Use(middleware.AuthServiceProvider()).PUT("/contestId/signUp", contest.SignUp)
		}

		_rank := global.Group("/rank")
		{
			_rank.GET("", rank.GetList)
			//_contest.GET("/:rankId", rank.GetDetail)
		}

		upload := global.Group("/upload")
		{
			upload.POST("", file.Upload)
		}
	}

}
