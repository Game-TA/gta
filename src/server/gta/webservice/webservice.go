package webservice

import(
	"net/http"
	"github.com/codegangsta/martini"
)

type WebService interface{
	GetPath() string
	WebGet(params martini.Params) (int,string)
	WebPost(params martini.Params,req *http.Request) (int,string)
}

func RegisterWebService(
	webService WebService,
	classicMartini *martini.ClassicMartini) {
	path:=webService.GetPath()
	classicMartini.Get(path,webService.WebGet)
	classicMartini.Get(path+"/:id",webService.WebGet)
	classicMartini.Post(path,webService.WebPost)
	classicMartini.Post(path+"/:id",webService.WebPost)
}
