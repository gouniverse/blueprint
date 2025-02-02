package testutils

import (
	"context"
	"errors"
	"project/config"

	"github.com/gouniverse/shopstore"
)

func SeedOrder(orderID string, customerID string) (shopstore.OrderInterface, error) {
	if !config.ShopStoreUsed {
		return nil, errors.New("shopstore is not used")
	}

	if config.ShopStore == nil {
		return nil, errors.New("shopstore is nil")
	}

	order, err := config.ShopStore.OrderFindByID(context.Background(), orderID)

	if err != nil {
		return nil, err
	}

	if order != nil {
		return order, nil
	}

	order = shopstore.NewOrder()
	order.SetID(orderID)
	order.SetCustomerID(customerID)

	if err := config.ShopStore.OrderCreate(context.Background(), order); err != nil {
		return nil, err
	}

	return order, nil
}
