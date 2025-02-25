package dto

type ToDoListPostRequest struct {
	ToDoTask string `json:"to_do_task"`
	Status   bool   `json:"status"`
	GroupId  string `json:"group_id"`
	OwnerId  string `json:"owner_id"`
}
type ToDoGroupPostRequest struct {
	GroupName string `json:"group_name"`
	Status    bool   `json:"status"`
	GroupId   string `json:"group_id"`
	OwnerId   string `json:"owner_id"`
}

type ToDoListPutRequest struct {
	Id       string `json:"id"`
	ToDoTask string `json:"to_do_task"`
	GroupId  string `json:"group_id"`
}

type ToDoListCheck struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}
