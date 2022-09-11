package repository

import (
	pgsql "assignment-2/helpers/database"
	"assignment-2/models"
	"fmt"
)

func CreateBrand(payload models.Brand) (models.Brand, error) {
	db := pgsql.GetDB()

	err := db.Create(&payload).Error
	if err != nil {
		fmt.Println("Error creating brand data:", err)
	}
	fmt.Println("New Brand data:", payload)
	return payload, err
}

func GetBrand(brandID uint64) (brand models.Brand, err error) {
	db := pgsql.GetDB()

	err = db.First(&brand, "id=?", brandID).Error
	if err != nil {
		fmt.Print("Brand not found with ID:", brandID)
	}
	return brand, err
}

func UpdateBrand(payload models.Brand) (models.Brand, int64, error) {
	db := pgsql.GetDB()

	update := db.Model(&payload).Where("id=?", payload.ID).Updates(payload)
	if update.Error != nil {
		fmt.Print("Updating Brand failed with ID:", update.Error)
	}
	fmt.Print("Updating Brand success with ID:", payload.ID)
	return payload, update.RowsAffected, update.Error
}

func DeleteBrand(payload models.Brand) {
	db := pgsql.GetDB()

	err1 := db.Model(&payload).Where("id=?", payload.ID).Delete(payload).Error
	if err1 != nil {
		fmt.Print("Deleting Brand failed with ID:", payload.ID)
	}
	fmt.Print("Deleting Brand success with ID:", payload.ID)
}
