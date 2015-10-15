package Api

import (
	"github.com/martini-contrib/render"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
	"strconv"
)

func IndexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var png []byte
	var size int = 256
	//level && size
	r.ParseForm()
	qr_string := r.FormValue("qr_string")
	if qr_string == "" {
	}
	size_str := r.FormValue("size")
	if size_str != "" {
		size, err = strconv.Atoi(size_str)
		if err != nil {
			size = 256
		}
	}
	png, err = qrcode.Encode(qr_string, qrcode.High, size)
	if err == nil {
		w.Header().Set("content-type", "image/png")
		w.Write(png)
	}
}
func NotFoundHandler(r render.Render) {
	r.HTML(400, "404", nil)
}
