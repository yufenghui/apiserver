package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/yufenghui/apiserver/handler"
	"github.com/yufenghui/apiserver/model"
	"github.com/yufenghui/apiserver/pkg/errno"
	"strconv"
)

// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserModel true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [put]
func Update(c *gin.Context) {

	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// validate data
	if userId == 0 {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// update the record based on the user id.
	u.Id = uint64(userId)

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// update
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
