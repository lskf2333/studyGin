package main

//模板继承

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter,r *http.Request)  {
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./index.tmpl")
	if err!=nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	msg:="哒咩"
	//渲染模板
	t.Execute(w,msg)
}

func home(w http.ResponseWriter,r *http.Request)  {
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./home.tmpl")
	if err!=nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	msg:="哒咩"
	//渲染模板
	t.Execute(w,msg)
}

func index2(w http.ResponseWriter,r *http.Request)  {
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/index2.tmpl")
	if err!=nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	msg:="哒咩index2"
	//渲染模板
	t.ExecuteTemplate(w,"index2.tmpl",msg)
}

func home2(w http.ResponseWriter,r *http.Request)  {
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/home2.tmpl")
	if err!=nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	msg:="哒咩home2"
	//渲染模板
	t.ExecuteTemplate(w,"home2.tmpl",msg)
}

func main()  {
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	http.HandleFunc("/index2",index2)
	http.HandleFunc("/home2",home2)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}