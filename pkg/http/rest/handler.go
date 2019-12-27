package rest

import (
	"github.com/RainJoe/go-web-project-template/pkg/adding"
	"github.com/RainJoe/go-web-project-template/pkg/listing"
	"github.com/gin-gonic/gin"
)

//Handler provides services
func Handler(a adding.Service, l listing.Service) *gin.Engine {
	router := gin.Default()

	router.POST("/v1/register", addUser(a))
	router.GET("/v1/users", getAllUsers(l))

	return router
}

func addUser(s adding.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user adding.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "params error",
			})
			return
		}

		if err := s.AddUser(user); err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "adding user to db error",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	}
}

func getAllUsers(l listing.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := l.GetAllUsers()
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "adding user to db error",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": users,
		})
	}
}
