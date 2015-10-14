package qrcode

import (
	"fmt"
	"git.oschina.net/cnjack/qrcode/Api"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Run() {
	fmt.Println("The App Is Running...")
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", Api.IndexHandler)

	m.NotFound(Api.NotFoundHandler)
	m.RunOnAddr(":8080")
}
