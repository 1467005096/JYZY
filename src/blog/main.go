package main 

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ia := new(Article)
	//注册路由器
	router := gin.Default()
    router.GET("/GetAll",ia.GetAll)//获取所有文章
    router.GET("/GetOne/:id",ia.GetOne)//根据URL传递的Id值查询相应的一篇文章
    router.POST("/Add",ia.Add)//添加文章
    router.POST("/Update",ia.Update)//更新文章
    router.GET("/Del/:id",ia.Del)//根据URL传递的ID值删除相应的一篇文章
    router.Run(":8010")//设置端口，完成路由器配置
}

