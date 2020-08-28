package region

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gogin/app/model"
	"net/http"
)

type reqAdd struct {
	Id   uint   `form:"id" json:"id" binding:"required"`
	Pid  uint   `form:"pid" json:"pid" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type reqDetail struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

type reqSub struct {
	Pid uint `validate:"required"`
}

type DetailData struct {
	Id   uint   `json:"id"`
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

var validate *validator.Validate

// add new region
func Add(c *gin.Context) {
	var req reqAdd
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	region := model.Region{}
	region.Id = req.Id
	region.Pid = req.Pid
	region.Name = req.Name
	id, err := region.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 1000,
			"err_msg":  "Insert Data Failure",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
		"data":     id,
	})
}

// return one region information
func Detail(c *gin.Context) {
	// check request params
	var req reqDetail
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// fetch data
	region := model.Region{}
	region.Id = req.Id
	data, err := region.FetchOne()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 1000,
			"err_msg":  "Fetch Data Failure: " + err.Error(),
		})
		return
	}

	// return data
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
		"data":     data,
	})
}

// return a list of regions
func Sub(c *gin.Context) {
	// check request params
	var req reqSub
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// fetch data
	region := model.Region{}
	region.Pid = req.Pid
	data, err := region.List()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 1000,
			"err_msg":  "Fetch Data Failure: " + err.Error(),
		})
		return
	}

	// return data
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
		"data":     data,
	})
}
