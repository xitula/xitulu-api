// Package router 路由包，路由的配置项与公共函数等
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"xitulu/router/handler"
)

var db = make(map[string]string)

// SetupRouter 配置路由规则
func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	// 配置cors
	r.Use(cors.Default())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 用户相关
	// 获取所有用户
	r.GET("/users", handler.UserGetAll)
	// 新增用户
	r.POST("/users", handler.UserAdd)
	// 更新用户数据
	r.PUT("/users", handler.UserUpdate)
	// 删除用户
	r.DELETE("/users/:id", handler.UserDelete)
	// 用户登录
	r.POST("/users/login", handler.UserLogin)
	// 用户登出
	r.GET("/users/logout", handler.UserLogout)

	// 待办相关
	// 返回所有数据
	r.GET("/todo", handler.TodoGet)
	// 返回指定ID的数据
	r.GET("/todo/:id", handler.TodoGetOne)
	// 新增
	r.POST("/todo", handler.TodoAdd)
	// 更新
	r.PUT("/todo", handler.TodoUpdate)
	// 删除
	r.DELETE("/todo/:id", handler.TodoDelete)

	// 随感相关
	// 查询所有随感数据
	r.GET("/causeries", handler.CauseriesAll)
	// 新增随感
	r.POST("/causeries", handler.CauseriesAdd)
	// 更新随感
	r.PUT("/causeries", handler.CauseriesUpdate)
	// 依据ID删除随感
	r.DELETE("/causeries/:id", handler.CauseriesDelete)

	// 文章相关
	// 新增文章
	r.POST("/articles", handler.ArticleAdd)
	//r.GET("/articles/:id", getArticle)
	// 获取所有文章
	r.GET("/articles", handler.Articles)
	// 更新文章
	r.PUT("/articles", handler.ArticleUpdate)
	// 删除文章
	r.DELETE("/articles/:id", handler.ArticleDelete)

	// 设置静态上传的文件目录
	r.Static("/uploaded", "../uploaded")
	r.MaxMultipartMemory = 1 << 20 // 1MB
	r.POST("/upload/avatar", handler.Upload)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	// 	"foo":  "bar",
	// 	"manu": "123",
	// }))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
