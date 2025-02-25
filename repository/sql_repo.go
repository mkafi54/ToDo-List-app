package repository

import "github.com/kahfi/to-do-list-app/database"

func GetData(model interface{}, ownerid string) error {
	db := database.DB.Debug()
	if err := db.Where("ownerid = ?", ownerid).Find(model).Error; err != nil {
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
