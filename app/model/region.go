package model

import (
	orm "gogin/app/database"
)

type Region struct {
	Id   uint   `json:"id"`
	Pid  uint   `json:"pid"`
	Name string `json:"name"`
}

// setup table_name
// if not, the default table name will be model_Name+s
func (region *Region) TableName() string {
	return "region"
}

// insert
func (region *Region) Insert() (id uint, err error) {
	result := orm.Mysql.Create(&region)
	id = region.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// select
func (region *Region) List() (returnData []Region, err error) {
	if err = orm.Mysql.Find(&returnData).Error; err != nil {
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
