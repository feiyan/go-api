package region

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gogin/app/model"
	"gogin/app/util"
	"net/http"
	"strconv"
	"time"
)

type reqAdd struct {
	Id   uint   `form:"id" json:"id" binding:"required"`
	Pid  uint   `form:"pid" json:"pid" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type reqDetail struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

// binding: required, 必填
// binding: omitempty, 空时忽略
type reqSub struct {
	Pid uint `form:"pid" json:"pid" binding:"omitempty"`
}

type repDetail struct {
	Id   uint   `json:"id"`
	Pid  uint   `json:"pid"`
	Name string `json:"name"`
}

var validate *validator.Validate

// add new region
// test with CURL: curl -d "id=110200&pid=110000&name=Chaoyang" "http://localhost:8080/region/add"
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
			"err_msg":  "Insert Data Failure: " + err.Error(),
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

	// redis cache
	var data repDetail
	r := util.RedisClient
	cacheKey := "region_detail_" + strconv.Itoa(int(req.Id))
	jsonStr, err := r.Get(cacheKey).Result()
	if err != nil {
		// fetch data
		region := model.Region{}
		region.Id = req.Id
		returnData, err := region.FetchOne()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err_code": 1000,
				"err_msg":  "Fetch Data Failure: " + err.Error(),
			})
			return
		}

		// convert struct to json
		// set json val to cacheKey
		data.Id = returnData.Id
		data.Pid = returnData.Pid
		data.Name = returnData.Name
		jsonStr, _ := json.Marshal(data)

		// 需要时间的地方使用time库
		r.Set(cacheKey, jsonStr, time.Hour)
	} else {
		// convert json to struct
		json.Unmarshal([]byte(jsonStr), &data)
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
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// fetch data
	region := model.Region{}
	data, err := region.FetchByPid(req.Pid)

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
