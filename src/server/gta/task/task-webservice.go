package task

import(
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini"
)
//GetPath implements webservice.GetPath
func (t *Task) GetPath() string {
	//Associate this service with http://host:port/guestbook
	return "/task"
}
//WebGet implements webservice.WebGet
func (t *Task) WebGet(params martini.Params)(int,string){
	if len(params) == 0 {
		encodedTasks,err:=json.Marshal(GetAll())
		if err!=nil{
			return http.StatusInternalServerError,"error on get all tasks from db"
		}
		return http.StatusOK,string(encodedTasks)
	}
	//TODO: get task by id
	return http.StatusNotFound,"data is empty"
}

func (t *Task) WebPost(
	params martini.Params,
	req *http.Request) (int,string){
	return http.StatusOK,"new task created"
}
