package services

import (
	"Netxd_Project/Netxd_DAL/interfaces"
	"Netxd_Project/Netxd_DAL/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c *CustomerService) CreateCustomer(customer *models.Customer) (*models.DBResponse, error) {

	_, err := c.CustomerCollection.InsertOne(c.ctx, &customer)
	if err != nil {
		return nil, err
	}

	response := &models.DBResponse{
		CustomerID: customer.CustomerID,
		CreatedAt:  customer.CreatedAt,
	}

	return response, nil
}

// customer.CustomerID = 4545
// customer.FirstName = "Rohith"
// customer.LastName = "S"
// customer.BankID = 21345
// customer.CreatedAt = "29.08.2023"
// customer.UpdateAt = "30.08.2023"
// customer.Balance = 500.00
// customer.IsActive = true
