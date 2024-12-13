package testutils

import (
	"context"
	"project/config"

	"github.com/gouniverse/shopstore"
)

func SeedOrder(orderID string, customerID string) (shopstore.OrderInterface, error) {
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

func SeedProduct(productID string, price float64) (shopstore.ProductInterface, error) {
	product, err := config.ShopStore.ProductFindByID(context.Background(), productID)

	if err != nil {
		return nil, err
	}

	if product != nil {
		return product, nil
	}

	product = shopstore.NewProduct()
	product.SetID(productID)
	product.SetTitle("Test Product")
	product.SetStatus(shopstore.PRODUCT_STATUS_ACTIVE)
	product.SetPriceFloat(price)

	if err := config.ShopStore.ProductCreate(context.Background(), product); err != nil {
		return nil, err
	}

	return product, nil
}
