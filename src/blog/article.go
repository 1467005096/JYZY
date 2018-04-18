package main 

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Iarticle interface{
	GetAll()//获取所有文章
	GetOne()//根据URL传递的Id值查询相应的一篇文章
	Add()//添加文章
	Update()//更新文章
	Del()//根据URL传递的ID值删除相应的一篇文章
}

type Article struct {
	Id                int     `json:"id"`
	Author            string  `json:"author"`
	Title             string  `json:"title"`
	Content           string  `json:"content"`
	Time_publish      string  `json:"time_publish"`
	Time_last_update  string  `json:"time_last_update"`
}

//获取所有文章
func (article Article) GetAll(c *gin.Context){
	//连接数据库并设置相应参数
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ginblog?charset=utf8")
	if err != nil {
	    log.Fatal(err)
	}
	if err != nil{
	  log.Fatalln(err)
	 }
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil{
	  log.Fatalln(err)
	}
	
	//查询数据库
	rows, err := db.Query("SELECT id,author,title,content,time_publish,time_last_update FROM article")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	articles := make([]Article, 0)//用来保存将查询到的数据
	
	//遍历查询结果
	for rows.Next() {
	    if err := rows.Scan(&article.Id, &article.Author, &article.Title, &article.Content, &article.Time_publish, &article.Time_last_update); err != nil {
	        log.Fatal(err)
	        c.JSON(http.StatusOK, gin.H{
			"articles":   nil,
			"msg":    "查询失败！",
			"status": http.StatusOK,
		})
	    }
	    //将当前article添加到articles切片中
	    articles = append(articles, article)
	}
	fmt.Println(articles)
	
	c.JSON(http.StatusOK, gin.H{
			"article":   articles,
			"msg":    "成功获取所有文章！",
			"status": http.StatusOK,
		})//渲染JSON
}

//根据ID获取文章
func (article Article) GetOne(c *gin.Context){
	fmt.Println("根据ID获取文章")
	//连接数据库并设置相应参数
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ginblog?charset=utf8")
	if err != nil {
	    log.Fatal(err)
	}
	if err != nil{
	  log.Fatalln(err)
	 }
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil{
	  log.Fatalln(err)
	}
	
	//获取请求数据
	sid :=c.Param("id")
	id, err := strconv.Atoi(sid)//将参数的类型转换为int
	if err != nil{
	  log.Fatalln(err)
	 }
	//var article Article////用来保存将查询到的数据
	
	//查询数据库
	row := db.QueryRow("SELECT id,author,title,content,time_publish,time_last_update FROM article WHERE id=?",id)
	err1 :=row.Scan(&article.Id, &article.Author, &article.Title, &article.Content, &article.Time_publish, &article.Time_last_update)
	if err1 != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"article":   nil,
			"msg":    "未找到此id的文章！",
			"status": http.StatusOK,
		})//渲染JSON
		return
	}
	fmt.Println(article)
	
	c.JSON(http.StatusOK, gin.H{
			"article":   article,
			"msg":    "成功获取此id的文章！",
			"status": http.StatusOK,
		})//渲染JSON
}

//添加文章
func (article Article) Add(c *gin.Context){
	fmt.Println("添加文章")
	//连接数据库并设置相应参数
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ginblog?charset=utf8")
	if err != nil {
	    log.Fatal(err)
	}
	if err != nil{
	  log.Fatalln(err)
	 }
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil{
	  log.Fatalln(err)
	}
	
	//获取请求数据
	author:=c.Request.FormValue("author")
	title:=c.Request.FormValue("title")
	content:=c.Request.FormValue("content")
	fmt.Println("author=",author)
	fmt.Println("title=",title)
	fmt.Println("content",content)
	
	//获取当前时间
	t := time.Now()
	time_publish := t.Format("2006-01-02 15:04:05")
	time_last_update := time_publish
	
	//查询数据库获取最大Id，此次新文章的Id就是maxId+1
	var maxId int;
	row := db.QueryRow("SELECT MAX(id) FROM article")
	err1 :=row.Scan(&maxId)
	if err1 != nil {
		log.Fatal(err)
		maxId = 0//如果查询结果为空，说明表里还没有数据，maxId=0
	}
	fmt.Println("maxid:",maxId)
	
	//执行插入语句
	rs, err := db.Exec("INSERT INTO article(id,author,title,content,time_publish,time_last_update) VALUES (?,?,?,?,?,?)", maxId+1,author,title,content,time_publish,time_last_update)
	if err != nil {
        log.Fatalln(err)
	}
	iid, err := rs.LastInsertId()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("insert article Id {}", iid)
    msg := fmt.Sprintf("insert successful %d", iid)
    
    c.JSON(http.StatusOK, gin.H{
            "msg": msg,
            "status":  http.StatusOK,
    })//渲染JSON
}

//更新文章
func (article Article) Update(c *gin.Context){
	fmt.Println("更新文章")
	//连接数据库并设置相应参数
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ginblog?charset=utf8")
	if err != nil {
	    log.Fatal(err)
	}
	if err != nil{
	  log.Fatalln(err)
	 }
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil{
	  log.Fatalln(err)
	  c.JSON(http.StatusOK, gin.H{
			"msg":     "文章信息有误！",
			"data":    nil,
			"status":  http.StatusOK,
		})//渲染JSON
	}
	
	//获取请求数据
	sid:=c.Request.FormValue("id")
	id, err := strconv.Atoi(sid)//将参数的类型转换为int
	title:=c.Request.FormValue("title")
	content:=c.Request.FormValue("content")
	fmt.Println("id=",id)
	fmt.Println("title=",title)
	fmt.Println("content",content)
	
	//获取当前时间
	t := time.Now()
	time_last_update := t.Format("2006-01-02 15:04:05")
	
	article = Article{Id: id}
	err = c.Bind(&article)
	if err != nil {
		log.Println(err)
	}
	
	//执行更新语句
	stmt, err := db.Prepare("UPDATE article SET title=?,content=?,time_last_update=? WHERE id=?")
	if err != nil {
        log.Fatalln(err)
	}
	defer stmt.Close()
	rs, err := stmt.Exec(title,content,time_last_update,id)
	if err != nil {
		log.Println(err)
	}
	ra, err := rs.RowsAffected()
    if err != nil {
        log.Fatalln(err)
    }
    msg := fmt.Sprintf("Update article %d successful %d", article.Id, ra)
    c.JSON(http.StatusOK, gin.H{
	        "msg": msg,
	        "status":  http.StatusOK,
    })//渲染JSON
}

//删除文章
func (article Article) Del(c *gin.Context){
	fmt.Println("删除文章")
	//连接数据库并设置相应参数
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ginblog?charset=utf8")
	if err != nil {
	    log.Fatal(err)
	}
	if err != nil{
	  log.Fatalln(err)
	 }
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil{
	  log.Fatalln(err)
	}
	//获取请求数据
	sid :=c.Param("id")
	id, err := strconv.Atoi(sid)//将参数的类型转换为int
	
	//执行删除语句
	rs, err := db.Exec("DELETE FROM article WHERE id=?", id)
	if err != nil {
        log.Fatalln(err)
	}
	ra, err := rs.RowsAffected()
	if err != nil {
	    log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete article %d successful %d", id, ra)
	
	c.JSON(http.StatusOK, gin.H{
	    "msg": msg,
	    "status": http.StatusOK,
	})//渲染JSON
}




