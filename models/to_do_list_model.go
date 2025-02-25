package models

type ToDoList struct {
	Id        string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	ToDoTask  string `json:"to_do_task"`
	Status    bool   `json:"status"`
	GroupId   string `json:"group_id"`
	OwnerId   string `json:"owner_id"`
	IsDeleted bool
}

func (e *ToDoList) TableName() string {
	return "to_do_list"
}
