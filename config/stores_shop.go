package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gouniverse/shopstore"
)

func init() {
	if ShopStoreUsed {
		addDatabaseInit(ShopStoreInitialize)
		addDatabaseMigration(ShopStoreAutoMigrate)
	}
}

func ShopStoreInitialize(db *sql.DB) error {
	if !ShopStoreUsed {
		return nil
	}

	shopStoreInstance, err := shopstore.NewStore(shopstore.NewStoreOptions{
		DB:                     db,
		CategoryTableName:      "snv_shop_category",
		DiscountTableName:      "snv_shop_discount",
		MediaTableName:         "snv_shop_media",
		OrderTableName:         "snv_shop_order",
		OrderLineItemTableName: "snv_shop_order_line_item",
		ProductTableName:       "snv_shop_product",
	})

	if err != nil {
		return errors.Join(errors.New("shopstore.NewStore"), err)
	}

	if shopStoreInstance == nil {
		return errors.Join(errors.New("shopStoreInstance is nil"))
	}

	ShopStore = shopStoreInstance

	return nil
}

func ShopStoreAutoMigrate(_ context.Context) error {
	if !ShopStoreUsed {
		return nil
	}

	err := ShopStore.AutoMigrate()

	if err != nil {
		return errors.Join(errors.New("shopstore.AutoMigrate"), err)
	}

	return nil
}
