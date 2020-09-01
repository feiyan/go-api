package region

import (
	"encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gogin/app/middleware"
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

type repSub struct {
	repDetail
	Children []*repSub
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
	region := model.Region{
	    Id: req.Id,
	    Pid: req.Pid,
	    Name: req.Name,
    }
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

	middleware.Message("fetch region detail " + strconv.Itoa(int(req.Id)))

	// redis cache
	var data repDetail
	r := util.RedisClient
	cacheKey := "region_detail_" + strconv.Itoa(int(req.Id))
	jsonStr, err := r.Get(cacheKey).Result()
	fmt.Println(jsonStr)
	if err != nil {
		// fetch data
		region := model.Region{
            Id: req.Id,
        }
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
		data = repDetail{
            Id: returnData.Id,
            Pid: returnData.Pid,
            Name: returnData.Name,
        }
		jsonStr, _ := json.Marshal(data)

		// 需要时间的地方使用time库
		r.Set(cacheKey, jsonStr, time.Hour)
	} else {
		// convert json to struct
		_ = json.Unmarshal([]byte(jsonStr), &data)
	}

	// return data
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
		"data":     data,
	})
}

// 递归方式获取无限级regions
func walk(pid uint, nodes map[uint][]*repSub) (returnData []*repSub) {
	if subNodes, ok := nodes[pid]; ok {
		for _, node := range subNodes {
			node.Children = walk(node.Id, nodes)
			returnData = append(returnData, node)
		}
	}
	return
}

// get tree with one query
func Sub(c *gin.Context) {
	// fetch data
	region := model.Region{}
	data, err := region.FetchAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 1000,
			"err_msg":  "Fetch Data Failure: " + err.Error(),
		})
		return
	}

	// build tree
	dataMap := make(map[uint][]*repSub)
	for _, row := range data {
	    pid := row.Pid
		node := &repSub{
		    repDetail{
		        Id: row.Id,
		        Pid: pid,
		        Name: row.Name,
            },
            nil,
        }
		dataMap[pid] = append(dataMap[pid], node)
	}
	retData := walk(0, dataMap)

	// return data
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "Success",
		"data":     retData,
	})
}

// return a list of regions
func Tree(c *gin.Context) {
	// check request params
	var req reqSub
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// fetch data
	var data interface{}
	var err interface{ Error() string }
	if req.Pid > 0 {
		region := model.Region{}
		data, err = region.FetchByPid(req.Pid)
	} else {
		region := model.RegionTree{}
		data, err = region.FetchTree(0)
	}

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
