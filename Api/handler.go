package Api

import (
	"github.com/martini-contrib/render"
)

func IndexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}
func NotFoundHandler(r render.Render) {
	r.HTML(400, "404", nil)
}
