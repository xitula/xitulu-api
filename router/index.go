/*
@Description 路由包，路由的配置项与公共函数等
*/
package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	t "xitulu/types"
)

var db = make(map[string]string)

/*
@Description 默认接口返回
*/
func response(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusOK, t.Res{Code: 1, Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, t.Res{Code: 0, Message: "ok"})
	}
}

/*
@Description 带数据的默认接口返回
*/
func responseData(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusOK, t.Res{Code: 1, Message: err.Error(), Data: err})
	} else {
		c.JSON(http.StatusOK, t.Res{Code: 0, Message: "ok", Data: data})
	}
}

/*
@Description 设置Gin路由
*/
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

	registerTodo(r)
	registerUsers(r)
	registerCauseries(r)
	registerArticles(r)

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
