package model

import (
	"fmt"
	orm "gogin/app/database"
)

type Region struct {
	Id   uint   `json:"id"`
	Pid  uint   `json:"pid"`
	Name string `json:"name"`
}

type RegionTree struct {
	Region
	Children []RegionTree
}

// setup table_name
// if not, the default table name will be model_Name+s
func (region *Region) TableName() string {
	return "region"
}

// insert
func (region *Region) Insert() (id uint, err error) {
	if _, err := region.FetchOne(); err == nil {
		return 0, fmt.Errorf("id %d is already exists", region.Id)
	}
	result := orm.Mysql.Create(&region)
	id = region.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// fetch all regions
func (region *Region) FetchAll() (returnData []Region, err error) {
	err = orm.Mysql.Find(&returnData).Error
	if err != nil {
		return
	}
	return
}

// fetch regions by Pid
func (region *Region) FetchTree(pid uint) (returnData []RegionTree, err error) {
	data, err := region.FetchByPid(pid)
	if err != nil {
		return
	}
	for _, row := range data {
		children, _ := region.FetchTree(row.Id)
		node := RegionTree{}
		node.Id = row.Id
		node.Pid = row.Pid
		node.Name = row.Name
		node.Children = children
		returnData = append(returnData, node)
	}
	return
}

// fetch regions by Pid
func (region *Region) FetchByPid(pid uint) (returnData []Region, err error) {
	err = orm.Mysql.Where("pid = ?", pid).Find(&returnData).Error
	if err != nil {
		return
	}
	return
}

// fetch one
func (region *Region) FetchOne() (returnData Region, err error) {
	if err = orm.Mysql.First(&returnData, region.Id).Error; err != nil {
		return
	}
	return
}

// find and modify
func (region *Region) Update(id uint) (returnData Region, err error) {
	if err = orm.Mysql.Select([]string{"id", "pid", "name"}).First(&returnData, id).Error; err != nil {
		return
	}

	if err = orm.Mysql.Model(&returnData).Updates(&region).Error; err != nil {
		return
	}
	return
}

// delete
func (region *Region) Delete(id uint) (returnData Region, err error) {

	if err = orm.Mysql.Select([]string{"id"}).First(&returnData, id).Error; err != nil {
		return
	}

	if err = orm.Mysql.Delete(&region).Error; err != nil {
		return
	}
	return
}
