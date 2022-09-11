package repository

import (
	pgsql "assignment-2/helpers/database"
	"assignment-2/models"
)

func GetOrders() (orders []models.Order, err error) {
	db := pgsql.GetDB()

	err = db.Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, err
}

func CreateOrders(payload models.Order) (orders *models.Order, err error) {
	db := pgsql.GetDB().Begin()

	err = db.Create(&payload).Error
	if err != nil {
		db.Rollback()
		return &payload, err
	}
	db.Commit()
	return &payload, err
}

func DeleteOrders(payload uint64) (orders models.Order, err error) {
	db := pgsql.GetDB().Begin()

	err = db.Preload("Items").Where("Order_ID = ?", payload).First(&orders).Error
	if err != nil {
		db.Rollback()
		return orders, err
	}

	for _, v := range orders.Items {
		err = db.Delete(&v).Error
		if err != nil {
			db.Rollback()
			return orders, err
		}
	}
	err = db.Delete(&orders).Error
	if err != nil {
		db.Rollback()
		return orders, err
	}
	db.Commit()
	return orders, err
}

func UpdateOrders(payload models.Order) (updated models.Order, err error) {
	db := pgsql.GetDB().Begin()

	for _, newItem := range payload.Items {
		item := models.Item{}
		err = db.Where("Order_ID = ? AND Item_ID = ?", &payload.Order_ID, &newItem.Item_ID).First(&item).Error
		if err != nil {
			db.Rollback()
			return payload, err
		} else {
			err = db.Updates(&newItem).Error
			if err != nil {
				db.Rollback()
				return payload, err
			}
		}
	}
	err = db.Updates(&payload).Error
	if err != nil {
		db.Rollback()
		return payload, err
	}
	db.Commit()
	return payload, err
}
