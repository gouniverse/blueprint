package testutils

import (
	"project/config"

	"github.com/gouniverse/shopstore"
	"github.com/gouniverse/userstore"
)

func SeedOrder(orderID string, customerID string) (shopstore.OrderInterface, error) {
	order, err := config.ShopStore.OrderFindByID(orderID)

	if err != nil {
		return nil, err
	}

	if order != nil {
		return order, nil
	}

	order = shopstore.NewOrder()
	order.SetID(orderID)
	order.SetCustomerID(customerID)

	if err := config.ShopStore.OrderCreate(order); err != nil {
		return nil, err
	}

	return order, nil
}

func SeedProduct(productID string, price float64) (shopstore.ProductInterface, error) {
	product, err := config.ShopStore.ProductFindByID(productID)

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

	if err := config.ShopStore.ProductCreate(product); err != nil {
		return nil, err
	}

	return product, nil
}

// SeedUser find existing or generates a new user with the given ID
func SeedUser(userID string) (userstore.UserInterface, error) {
	user, err := config.UserStore.UserFindByID(userID)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	user = userstore.NewUser().
		SetID(userID).
		SetStatus(userstore.USER_STATUS_ACTIVE)

	if userID == USER_01 {
		user.SetRole(userstore.USER_ROLE_USER)
	}

	if userID == ADMIN_01 {
		user.SetRole(userstore.USER_ROLE_ADMINISTRATOR)
	}

	err = config.UserStore.UserCreate(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
