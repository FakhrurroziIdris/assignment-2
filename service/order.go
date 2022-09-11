package service

import (
	"assignment-2/models"
	brandRepo "assignment-2/repository"
	"errors"
)

func GetOrders() (orders []models.Order, err error) {
	orders, err = brandRepo.GetOrders()
	if len(orders) > 0 {
		return
	} else {
		return orders, errors.New("Orders is empty.")
	}
}

func CreateOrders(payload models.Order) (created *models.Order, err error) {
	created, err = brandRepo.CreateOrders(payload)
	return created, err
}

func DeleteOrders(payload uint64) (deleted models.Order, err error) {
	deleted, err = brandRepo.DeleteOrders(payload)
	return deleted, err
}

func UpdateOrders(payload models.Order) (updated models.Order, err error) {
	updated, err = brandRepo.UpdateOrders(payload)
	return updated, err
}
