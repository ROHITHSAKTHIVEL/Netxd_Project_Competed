package interfaces

import "Netxd_Project/Netxd_DAL/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
