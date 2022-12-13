package main

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("login", "views/users/base.html", "views/users/login.html")
	r.AddFromFilesFuncs(
		"index",
		template.FuncMap{},
		"views/main/base.html",
		"views/main/index.html",
	)
	r.AddFromFilesFuncs(
		"detail",
		template.FuncMap{},
		"views/main/base.html",
		"views/main/detail.html",
	)
	return r
}
