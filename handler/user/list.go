package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/yufenghui/apiserver/handler"
	"github.com/yufenghui/apiserver/model"
)

func List(c *gin.Context) {

	users, err := model.List()
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, err, users)
}
