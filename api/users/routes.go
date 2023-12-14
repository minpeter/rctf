package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minpeter/rctf-backend/utils"
)

func Routes(userRoutes *gin.RouterGroup) {

	userRoutes.GET("/:id", getUserHandler)

	me := userRoutes.Group("/me")
	{
		me.GET("", getMeHandler)
		me.PATCH("", updateMeHandler)

		auth := me.Group("/auth")
		{
			auth.DELETE("/ctftime", deleteCtftimeAuthHandler)
			auth.PUT("/ctftime", putCtftimeAuthHandler)
			auth.DELETE("/email", deleteEmailAuthHandler)
			auth.PUT("/email", putEmailAuthHandler)
		}

		members := me.Group("/members")
		{
			members.DELETE("/:id", deleteMemberHandler)
			members.GET("", listMembersHandler)
			members.POST("", newMemberHandler)
		}
	}
}

func getUserHandler(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func getMeHandler(c *gin.Context) {
	utils.SendResponse(c, "goodUserData", gin.H{
		"name":             "admin",
		"ctftimeId":        nil,
		"division":         "open",
		"score":            20000,
		"globalPlace":      nil,
		"divisionPlace":    nil,
		"solves":           []string{},
		"teamToken":        "testToken",
		"allowedDivisions": []string{"open"},
		"id":               "5f925ecc-89e3-4e2d-9b5d-1219e9abc8d1",
		"email":            "admin@admin.com",
	})
}

func updateMeHandler(c *gin.Context) {
	c.Status(http.StatusNoContent)
}