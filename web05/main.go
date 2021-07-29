package main

//go 模板引擎

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	//结构体属性的首字母一定要大写，不然渲染不了那个属性
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter,r *http.Request)  {
	// 2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil {
		fmt.Printf("Parse temp failed,err:%v",err)
		return
	}
	// 3.渲染模板
	u1:=User{
		Name: "拉拉",
		Gender: "男",
		Age: 19,
	}
	m1:=map[string]interface{}{
		"name": "哈哈",
		"gender": "女",
		"age": 18,
	}
	hobbyList:=[]string{
		"篮球",
		"足球",
		"双色球",
	}
	err=t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
	if err!=nil {
		fmt.Printf("Execute template failed,err:%v",err)
		return
	}
}

func main()  {
	http.HandleFunc("/hello",sayHello)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}
