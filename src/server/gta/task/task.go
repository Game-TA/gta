package task

import (
	"github.com/Game-TA/gta/src/server/gta/db"
	"github.com/golang/go/src/pkg/encoding/json"
)

type Task struct {
	Id       int
	ParentId int
	Name     string
	Content  string
}

func Add(parentId int, name, content string) int {
	//TODO:add task object into mysql db
	result, err := db.Mysql().Exec(
		"insert into t_task(parent_id,name,content) values(?,?,?)",
		parentId, name, content)
	if err !=nil{
		println(err)
	}
	println("print out the added task  result",result)
	if result!=nil{
		id, err := result.LastInsertId()
		if err!=nil{
			println(err)
		}
		return int(id)
	}

	return  0
}

func GetAll() string {
	//TODO:query all tasks from mysql db,convent to json format
	rows, _ := db.Mysql().Query("select * from t_task limit 0,10")
	var tasks = []Task{}
	var task = Task{}
	for rows.Next() {
		rows.Scan(
			&task.Id,
			&task.ParentId,
			&task.Name,
			&task.Content)
		tasks = append(tasks, task)
	}
	encoded,_:=json.Marshal(tasks)
	return string(encoded)
}
