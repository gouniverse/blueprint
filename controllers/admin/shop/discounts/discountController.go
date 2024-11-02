package admin

import (
	"errors"
	"net/http"
	"project/config"
	"project/controllers/admin/shop/shared"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/cdn"
	crud "github.com/gouniverse/crud/v2"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/shopstore"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type discountController struct {
}

func NewDiscountController() *discountController {
	return &discountController{}
}

func (discountController *discountController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	discountsCrud, err := crud.New(crud.Config{
		EntityNameSingular: "Discount",
		EntityNamePlural:   "Discounts",
		Endpoint:           shared.NewLinks().Discounts(map[string]string{}),
		ColumnNames: []string{
			"Title",
			"Status",
			"Type",
			"Amount",
			"Period Valid",
			"Discount Code",
			"Created",
		},
		CreateFields: []form.Field{
			{
				Label: "Title",
				Name:  "title",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
		},
		ReadFields: []form.Field{
			{
				Label: "Title",
				Name:  "title",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Status",
				Name:  "status",
				Type:  crud.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Key:   "",
						Value: "",
					},
					{
						Key:   shopstore.DISCOUNT_STATUS_DRAFT,
						Value: shopstore.DISCOUNT_STATUS_DRAFT,
					},
					{
						Key:   shopstore.DISCOUNT_STATUS_INACTIVE,
						Value: shopstore.DISCOUNT_STATUS_INACTIVE,
					},
					{
						Key:   shopstore.DISCOUNT_STATUS_ACTIVE,
						Value: shopstore.DISCOUNT_STATUS_ACTIVE,
					},
				},
			},
			{
				Label: "Type",
				Name:  "type",
				Type:  crud.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Key:   "",
						Value: "",
					},
					{
						Key:   shopstore.DISCOUNT_TYPE_AMOUNT,
						Value: shopstore.DISCOUNT_TYPE_AMOUNT,
					},
					{
						Key:   shopstore.DISCOUNT_TYPE_PERCENT,
						Value: shopstore.DISCOUNT_TYPE_PERCENT,
					},
				},
			},
			{
				Label: "Code",
				Name:  "code",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Starts",
				Name:  "starts_at",
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
			},
			{
				Label: "Ends",
				Name:  "ends_at",
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
			},
			{
				Label: "Created",
				Name:  "created_at",
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
			},
			{
				Label: "Updated",
				Name:  "updated_at",
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
			},
			{
				Label: "Description",
				Name:  "description",
				Type:  crud.FORM_FIELD_TYPE_HTMLAREA,
			},
		},
		UpdateFields: []form.Field{
			{
				Label: "Status",
				Name:  "status",
				Type:  crud.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Key:   "",
						Value: "",
					},
					{
						Key:   shopstore.DISCOUNT_STATUS_DRAFT,
						Value: shopstore.DISCOUNT_STATUS_DRAFT,
					},
					{
						Key:   shopstore.DISCOUNT_STATUS_INACTIVE,
						Value: shopstore.DISCOUNT_STATUS_INACTIVE,
					},
					{
						Key:   shopstore.DISCOUNT_STATUS_ACTIVE,
						Value: shopstore.DISCOUNT_STATUS_ACTIVE,
					},
				},
			},
			{
				Label: "Title",
				Name:  "title",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Type",
				Name:  "type",
				Type:  crud.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Key:   "",
						Value: "",
					},
					{
						Key:   shopstore.DISCOUNT_TYPE_AMOUNT,
						Value: shopstore.DISCOUNT_TYPE_AMOUNT,
					},
					{
						Key:   shopstore.DISCOUNT_TYPE_PERCENT,
						Value: shopstore.DISCOUNT_TYPE_PERCENT,
					},
				},
			},
			{
				Label: "Amount",
				Name:  "amount",
				Type:  crud.FORM_FIELD_TYPE_NUMBER,
			},
			{
				Label: "Discount Code",
				Name:  "code",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Starts",
				Name:  "starts_at",
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
			},
			{
				Label: "Ends",
				Name:  "ends_at",
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
			},
			{
				Label: "Description",
				Name:  "description",
				Type:  crud.FORM_FIELD_TYPE_HTMLAREA,
			},
		},
		FuncRows:            discountController.FuncRows,
		FuncCreate:          discountController.FuncCreate,
		FuncFetchReadData:   discountController.FuncFetchReadData,
		FuncFetchUpdateData: discountController.FuncFetchUpdateData,
		FuncTrash:           discountController.FuncTrash,
		FuncUpdate:          discountController.FuncUpdate,
		FuncLayout:          discountController.FuncLayout,
		HomeURL:             links.NewAdminLinks().Home(map[string]string{}),
	})

	if err != nil {
		return "Error: " + err.Error()
	}

	discountsCrud.Handler(w, r)
	return ""
}

func (discountController *discountController) FuncLayout(w http.ResponseWriter, r *http.Request, title string, content string, styleURLs []string, style string, scriptURLs []string, script string) string {
	scriptURLs = append([]string{
		cdn.Jquery_3_6_4(),
	}, scriptURLs...)

	return layouts.NewAdminLayout(r, layouts.Options{
		Request:    r,
		Title:      title + " | Admin",
		Content:    hb.Wrap().HTML(content),
		StyleURLs:  styleURLs,
		ScriptURLs: scriptURLs,
		Scripts:    []string{script},
		Styles: []string{
			`nav#Toolbar {border-bottom: 4px solid red;}`,
			style,
		},
	}).ToHTML()
}

func (discountController *discountController) FuncRows() ([]crud.Row, error) {
	discounts, err := config.ShopStore.DiscountList(shopstore.DiscountQueryOptions{})

	if err != nil {
		return nil, err
	}

	rows := lo.Map(discounts, func(discount shopstore.DiscountInterface, _ int) crud.Row {
		return crud.Row{
			ID: discount.ID(),
			Data: []string{
				discount.Title(),
				discount.Status(),
				discount.Type(),
				utils.ToString(discount.Amount()),
				discount.StartsAtCarbon().Format("d M Y") + " - " + discount.EndsAtCarbon().Format("d M Y"),
				discount.Code(),
				discount.CreatedAtCarbon().Format("d M Y"),
			},
		}
	})

	return rows, nil
}

func (discountController *discountController) FuncUpdate(entityID string, data map[string]string) error {
	discount, err := config.ShopStore.DiscountFindByID(entityID)

	if err != nil {
		return err
	}

	amountStr := data["amount"]
	startsAt := data["starts_at"]
	endsAt := data["ends_at"]
	title := data["title"]
	code := data["code"]
	status := data["status"]
	discountType := data["type"]

	if title == "" {
		return errors.New("title is required")
	}

	if status == "" {
		return errors.New("status is required")
	}

	if code == "" {
		return errors.New("code is required")
	}

	if discountType == "" {
		return errors.New("discount type is required")
	}

	if startsAt == "" {
		return errors.New("starts_at is required")
	}

	if endsAt == "" {
		return errors.New("ends_at is required")
	}

	if amountStr == "" {
		amountStr = "0"
	}

	amount, _ := utils.ToFloat(amountStr)
	startsAt = carbon.Parse(startsAt).ToDateTimeString(carbon.UTC)
	endsAt = carbon.Parse(endsAt).ToDateTimeString(carbon.UTC)

	discount.SetTitle(title)
	discount.SetDescription(data["description"])
	discount.SetStatus(status)
	discount.SetAmount(amount)
	discount.SetType(discountType)
	discount.SetCode(code)
	discount.SetStartsAt(startsAt)
	discount.SetEndsAt(endsAt)

	err = config.ShopStore.DiscountUpdate(discount)

	if err != nil {
		return err
	}

	return nil
}

func (discountController *discountController) FuncFetchReadData(discountID string) ([][2]string, error) {
	discount, err := config.ShopStore.DiscountFindByID(discountID)

	if err != nil {
		return nil, err
	}

	data := [][2]string{
		{"Title", discount.Title()},
		{"Status", discount.Status()},
		{"Description", discount.Description()},
		{"Type", discount.Type()},
		{"Amount", utils.ToString(discount.Amount())},
		{"Starts At", discount.StartsAtCarbon().Format("d M Y")},
		{"Ends At", discount.EndsAtCarbon().Format("d M Y")},
		{"Created", discount.CreatedAtCarbon().Format("d M Y")},
		{"Updated", discount.UpdatedAtCarbon().Format("d M Y")},
	}

	return data, nil
}

func (discountController *discountController) FuncFetchUpdateData(discountID string) (map[string]string, error) {
	discount, err := config.ShopStore.DiscountFindByID(discountID)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"title":       discount.Title(),
		"status":      discount.Status(),
		"amount":      utils.ToString(discount.Amount()),
		"description": discount.Description(),
		"type":        discount.Type(),
		"code":        discount.Code(),
		"starts_at":   discount.StartsAtCarbon().ToDateTimeString(),
		"ends_at":     discount.EndsAtCarbon().ToDateTimeString(),
		"created_at":  discount.CreatedAtCarbon().ToDateTimeString(),
		"updated_at":  discount.CreatedAtCarbon().ToDateTimeString(),
	}, nil
}

func (discountController *discountController) FuncCreate(data map[string]string) (discountID string, err error) {
	discount := shopstore.NewDiscount()
	discount.SetTitle(data["title"])
	discount.SetStatus(shopstore.DISCOUNT_STATUS_DRAFT)
	discount.SetAmount(0.00)
	err = config.ShopStore.DiscountCreate(discount)

	if err != nil {
		return "", err
	}

	return discount.ID(), nil
}

func (discountController *discountController) FuncTrash(discountID string) error {
	err := config.ShopStore.DiscountSoftDeleteByID(discountID)
	return err
}