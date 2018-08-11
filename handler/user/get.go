package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/yufenghui/apiserver/handler"
	"github.com/yufenghui/apiserver/model"
	"github.com/yufenghui/apiserver/pkg/errno"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")

	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, user)
}
