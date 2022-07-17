package services

import (
	"github.com/rajajamal/resto-backend/models"
	"github.com/rajajamal/resto-backend/repositories"
)

type Customer struct {
	Repository repositories.Customer
}

func (s Customer) Get(id int) (models.Customer, error) {
	return s.Repository.Find(id)
}

// func (s Customer) Reservation(customer models.Customer, tableNumber int) (models.Order, error) {
// 	order := models.Order{
// 		CustomerID:  customer.ID,
// 		TableNumber: tableNumber,
// 		Status:      types.ORDER_PENDING,
// 		Detail:      []*models.OrderDetail{},
// 	}

// 	return order, s.Order.Saves(&order)
// }

func (s Customer) Save(customer models.Customer) (models.Customer, error) {
	exist, _ := s.Repository.FindByPhoneNumber(customer.PhoneNumber)
	if exist.PhoneNumber == customer.PhoneNumber {
		customer.ID = exist.ID
	}

	return customer, s.Repository.Saves(&customer)
}
