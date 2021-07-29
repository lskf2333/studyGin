package main

//go 模板的初使用

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request)  {
	// 2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil {
		fmt.Printf("Parse temp failed,err:%v",err)
		return
	}
	// 3.渲染模板
	err=t.Execute(w,"么么")
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
