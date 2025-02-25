package repository

import (
	"fmt"

	"github.com/kahfi/to-do-list-app/database"
	"gorm.io/gorm/clause"
)

func GetData(model interface{}, ownerid string) error {
	db := database.DB.Debug()
	if err := db.Preload(clause.Associations).Where("ownerid = ?", ownerid).Find(model).Error; err != nil {
		return err
	}
	return nil
}
func GetDataByPredicate(model interface{}, field string, value string) error {
	db := database.DB.Debug()
	if err := db.Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(model).Error; err != nil {
		return err
	}
	return nil
}

func Save(model interface{}) error {
	db := database.DB.Debug()
	err := db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateData(model interface{}, id any) error {
	db := database.DB.Debug()
	err := db.Model(model).Where("id = ?", id).Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckUnchekData(model interface{}, id any) error {
	db := database.DB.Debug()
	err := db.Model(model).Where("id = ?", id).Select("status").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
func SoftDeleteById(model interface{}, idField string, idValue any) error {
	// Ensure that the model is a pointer to a struct
	if err := database.DB.Model(model).Where(fmt.Sprintf("%s = ?", idField), idValue).Update("isdeleted", true).Error; err != nil {
		return err
	}
	return nil
}
