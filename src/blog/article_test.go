package main 

import (
	"testing"
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
)

//测试GetAll
func TestGetAll(t *testing.T){
	//发送GET请求
	resp, err := http.Get("http://127.0.0.1:8010/GetAll")
	if err != nil {
		t.Error("请求失败！")
	}
	if resp.StatusCode != 200 {
		t.Fatalf("http错误码为%d", resp.StatusCode)
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("请求未返回数据！")
    }
    //声明Value为空接口的map，用来转换获取的数据
    var datas map[string]interface{}
    json.Unmarshal(body, &datas)
    fmt.Println("datas=",datas)
    status := datas["status"]
    if status != float64(200) {
		t.Error("获取数据失败！")
	}
}

//测试GetOne
func TestGetOne(t *testing.T){
	//发送GET请求
	resp, err := http.Get("http://127.0.0.1:8010/GetOne/1")
	if err != nil {
		t.Error("请求失败！")
	}
	if resp.StatusCode != 200 {
		t.Fatalf("http错误码为%d", resp.StatusCode)
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("请求未返回！")
    }
    //声明Value为空接口的map，用来转换获取的数据
    var datas map[string]interface{}
    json.Unmarshal(body, &datas)
    fmt.Println("datas=",datas)
    status := datas["status"]
    if status != float64(200) {
		t.Error("获取数据失败！")
	}
}

//测试Add
func TestAdd(t *testing.T){
	//以Http.PostFrom方法创建模拟数据并发送POST请求
	resp, err := http.PostForm("http://127.0.0.1:8010/Add", url.Values{"author": {"迹翼之羽"}, "title": {"文章标题"}, "content": {"文章内容"}})
	if err != nil {
		t.Error("请求失败！")
	}
	if resp.StatusCode != 200 {
		t.Fatalf("http错误码为%d", resp.StatusCode)
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("请求未返回！")
    }
    //声明Value为空接口的map，用来转换获取的数据
    var datas map[string]interface{}
    json.Unmarshal(body, &datas)
    fmt.Println("datas=",datas)
    status := datas["status"]
    if status != float64(200) {
		t.Error("添加文章失败！")
	}
}

//测试Update
func TestUpdate(t *testing.T){
	//以Http.PostFrom方法创建模拟数据并发送POST请求
	resp, err := http.PostForm("http://127.0.0.1:8010/Update", url.Values{"id": {"1"}, "title": {"新的标题"}, "content": {"新的内容"}})
	if err != nil {
		t.Error("请求失败！")
	}
	if resp.StatusCode != 200 {
		t.Fatalf("http错误码为%d", resp.StatusCode)
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("请求未返回！")
    }
    //声明Value为空接口的map，用来转换获取的数据
    var datas map[string]interface{}
    json.Unmarshal(body, &datas)
    fmt.Println("datas=",datas)
    status := datas["status"]
    if status != float64(200) {
		t.Error("更新文章失败！")
	}
}

//测试Delete
func TestDel(t *testing.T){
	//发送GET请求
	resp, err := http.Get("http://127.0.0.1:8010/Del/2")
	if err != nil {
		t.Error("请求失败！")
	}
	if resp.StatusCode != 200 {
		t.Fatalf("http错误码为%d", resp.StatusCode)
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error("请求未返回！")
    }
    //声明Value为空接口的map，用来转换获取的数据
    var datas map[string]interface{}
    json.Unmarshal(body, &datas)
    fmt.Println("datas=",datas)
    status := datas["status"]
    if status != float64(200) {
		t.Error("删除文章失败！")
	}
}

////获取数据类型
//func typeof(v interface{}) string {
//    return fmt.Sprintf("%T", v)
//}

