package main

//模板嵌套

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request){
	//定义一个函数kua
	//要么只有一个返回值，如果有两个返回值的话，第二个返回值必须是error类型的
	kua:=func(name string)(string,error){
		return name + "年轻又帅气！",nil
	}
	//定义模板
	//解析模板
	t:=template.New("f1.tmpl") //创建一个名字为f1.tmpl 的模板对象,创建的模板对象名一定要和解析的模板文件对应上
	//告诉模板引擎，我现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua":kua,
	})
	_,err:=t.ParseFiles("./f1.tmpl")
	if err!=nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	name:="嘎嘎"
	//渲染模板
	t.Execute(w,name)
}

func tmplDemo(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err!=nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	name:="发发"
	//渲染模板
	t.Execute(w,name)
}

func main()  {
	http.HandleFunc("/f1",f1)
	http.HandleFunc("/tmplDemo",tmplDemo)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}
