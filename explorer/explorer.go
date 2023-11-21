package explorer

import (
	"fmt"
	"github.com/LSapee/potatocoin/blockchain"
	"log"
	"net/http"
	"text/template"
)

const (
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, request *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)
}
func add(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		req.ParseForm()
		data := req.Form.Get("blockData")
		blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, req, "/", http.StatusPermanentRedirect)
	}

}

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
