package main

import (
	"GormTest/dao"
	"GormTest/models"
	"fmt"
)

var order int
var user models.UserInfo

func main(){
	err := dao.Startsql()
	if err != nil{
		fmt.Println("数据库启动错误%v",err)
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&models.UserInfo{})
	fmt.Printf("1.插入数据\n")
	fmt.Printf("2.查询数据\n")
	fmt.Printf("3.修改数据\n")
	fmt.Printf("4.删除数据\n")
	fmt.Scanf("%d",&order)
	switch order {
	case 1:if err = user.Add();err!=nil{
		fmt.Println("掺入数据出错，%v",err)
	}
	case 2:if err = user.Get();err!=nil{
		fmt.Println("查询数据出错，%v",err)
	}
	case 3:if err = user.Updata();err!=nil{
		fmt.Println("修改数据出错，%v",err)
	}
	case 4:if err=user.Delete();err!=nil{
		fmt.Println("删除数据出错，%v",err)
	}
	}
}