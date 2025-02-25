package models

type ToDoGroup struct {
	Id        string     `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	GroupName string     `json:"group_name"`
	Status    bool       `json:"status"`
	ToDoList  []ToDoList `gorm:"foreignKey:GroupId;references:id" json:"to_do_list"`
	OwnerId   string     `json:"owner_id"`
	IsDeleted bool
}

func (e *ToDoGroup) TableName() string {
	return "to_do_group"
}
