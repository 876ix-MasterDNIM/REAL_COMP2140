package servehtml
import (
	"path/filepath"
	"mime"
	"io/ioutil"
	"net/http"
)

func ServeHtml (path string, response http.ResponseWriter) {
	page, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	} else {
		response.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
		response.Write(page)
	}
}
