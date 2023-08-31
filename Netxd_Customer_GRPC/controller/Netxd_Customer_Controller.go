package controller

import (
	pro "Netxd_Project/Netxd_Customer/customer"
	"Netxd_Project/Netxd_DAL/interfaces"
	"Netxd_Project/Netxd_DAL/models"
	"context"
)

type RPCserver struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCserver) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &models.Customer{
		CustomerID: req.CustomerId}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerID,
		}
		return responseCustomer, nil
	}
}
