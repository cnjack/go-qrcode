package qrcode

import (
	"github.com/cnjack/go-qrcode/Api"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/throttle"
	"time"
)

func Run() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(throttle.Policy(&throttle.Quota{
		Limit:  1,
		Within: time.Second,
	}))

	m.Get("/", Api.IndexHandler)
	m.Get("/api", Api.ApiHandler)
	m.NotFound(Api.NotFoundHandler)

	m.RunOnAddr(":8080")
}
