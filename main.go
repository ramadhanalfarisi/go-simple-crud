package main

import "github.com/ramadhanalfarisi/go-simple-crud/app"

func main(){
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}