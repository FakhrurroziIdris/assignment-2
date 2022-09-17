package service

import (
	"assignment-2/models"
	brandRepo "assignment-2/repository"
)

func GetOrders(c chan models.Response) {
	go brandRepo.GetOrders(c)
	response := <-c
	if len(response.Data) > 0 {
		c <- response
	} else {
		c <- response
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
