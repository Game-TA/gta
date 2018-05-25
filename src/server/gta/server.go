package main

import (
	"github.com/codegangsta/martini"
	"github.com/Game-TA/gta/src/server/gta/task"
	"github.com/Game-TA/gta/src/server/gta/webservice"
)

func main(){
	martiniClassic:= martini.Classic()
	t:=&task.Task{}
	webservice.RegisterWebService(t,martiniClassic)
	martiniClassic.Run()
}
