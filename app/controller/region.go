package region

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DetailParams struct {
	Id uint8 `validate:"required"`
}

type SubParams struct {
	Pid uint8 `validate:"required"`
}

type DetailData struct {
	Id   uint8  `json:"id"`
	Name string `json:"name"`
}

var SubData = []*DetailData{
	{
		Id:   1,
		Name: "Beijing",
	},
	{
		Id:   2,
		Name: "Shanghai",
	},
}

// return one region information
func Detail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
	})
}

// return a list of regions
func Sub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
		"results":  &SubData,
	})
}
