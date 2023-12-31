package controller

import (
	"context"

	pro "github.com/ROHITHSAKTHIVEL/Netxd_Customer_Proto/proto"

	"github.com/ROHITHSAKTHIVEL/Netxd_DAL/interfaces"
	"github.com/ROHITHSAKTHIVEL/Netxd_DAL/models"
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
