package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/yufenghui/apiserver/handler"
	"github.com/yufenghui/apiserver/model"
)

// @Summary List the users in the database
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body model.UserModel true "List users"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /user [get]
func List(c *gin.Context) {

	users, err := model.List()
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, err, users)
}
