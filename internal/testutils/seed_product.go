package testutils

import (
	"context"
	"project/config"

	"github.com/gouniverse/shopstore"
)

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
	product.SetQuantityInt(10)

	if err := config.ShopStore.ProductCreate(context.Background(), product); err != nil {
		return nil, err
	}

	return product, nil
}
