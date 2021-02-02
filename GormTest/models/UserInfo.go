package models

import (
	"GormTest/dao"
	"fmt"
)

type UserInfo struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Num int `gorm:"column:学号"`
	Name string `gorm:"column:姓名"`
	Sex string `gorm:"column:性别"`
	Score int `gorm:"column:分数"`
}

type UserInfoModel interface {
	dao.DBmodel
}

func (user UserInfo)Add() (error){
	fmt.Println("请输入信息:")
	fmt.Scanf("%d %d %s %s %d",&user.Id,&user.Num,&user.Name,&user.Sex,&user.Score)
	return dao.DB.Create(&user).Error
}

func (user UserInfo)Get() (error){
	var xuehao string
	fmt.Println("请输入查询学号")
	fmt.Scanf("%s",&xuehao)
	dao.DB.Find(&user,"学号=?",xuehao)
	fmt.Println(user)
	return dao.DB.Find(&user,"学号=?",xuehao).Error
}

func (user UserInfo) Delete() (error){
	var xuehao string
	fmt.Println("请输入删除学号")
	fmt.Scanf("%s",&xuehao)
	dao.DB.Find(&user,"学号=?",xuehao)
	return dao.DB.Delete(&user).Error
}

func (user UserInfo)Updata() (error){
	var xuehao string
	fmt.Println("请输入要修改的学号")
	fmt.Scanf("%s",&xuehao)
	dao.DB.Find(&user,"学号=?",xuehao)
	return dao.DB.Model(&user).Update("分数",200).Error
}