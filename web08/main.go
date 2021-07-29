package main

//gin 渲染模板文件

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
		return
	}
	//渲染模板
	name := "飒飒"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Execute template failed,err:%v", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	//解析模板之前定义一个自定义的函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
		return
	}
	//渲染模板
	msg1 := "<script>alert(\"123\");</script>"
	msg2 := "<a href='http://www.baidu.com'>百度</a>"
	err = t.Execute(w, map[string]string{
		"msg1": msg1,
		"msg2": msg2,
	})
	if err != nil {
		fmt.Printf("Execute template failed,err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
}
