package service

import (
	"github.com/gin-gonic/gin"
	"mazhj.com/relaxoj/define"
)

func GetProblemList(c *gin.Context) {
	_ = c.DefaultQuery("page", define.DefaultPage)
	_ = c.DefaultQuery("size", define.DefaultSize)
	_ = c.Query("keyword")

}
