package dto

import "github.com/kahfi/to-do-list-app/models"

type ToDoListResponse struct {
	Id        string            `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	GroupName string            `json:"group_name"`
	Details   []models.ToDoList `json:"details"`
	Status    bool              `json:"status"`
	OwnerId   string            `json:"owner_id"`
}

type ToDoListDetail struct {
	Id       string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	ToDoTask string `json:"to_do_task"`
	Status   bool   `json:"status"`
	GroupId  string `json:"group_id"`
	OwnerId  string `json:"owner_id"`
}
