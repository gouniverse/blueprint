package controllers

import (
	"net/http"
	"project/internal/helpers"
	"project/internal/links"
	"project/internal/payment"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

type paypalCancelController struct{}
type paypalCancelControllerData struct {
	invoiceID string
}

func NewPaypalCancelController() *paypalCancelController {
	return &paypalCancelController{}
}

func (controller *paypalCancelController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareDataAndValidate(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewGuestLinks().Home(), 10)
	}

	return hb.NewWebpage().
		SetTitle("View Invoice").
		StyleURL(cdn.BootstrapCss_5_3_0()).
		Child(controller.page(data)).
		ToHTML()
}

func (controller *paypalCancelController) page(data paypalCancelControllerData) hb.TagInterface {
	title := hb.NewHeading1().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		HTML(`PayPal Invoice Payment Canceled`)

	invoiceUrl := links.NewCustomerLinks().InvoiceView(data.invoiceID, map[string]string{})
	invoiceLink := hb.NewHyperlink().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		Href(invoiceUrl).
		HTML(`CLICK HERE`)

	alert := hb.NewDiv().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		Child(hb.NewDiv().
			Class(`alert alert-info`).
			HTML(`You cancelled payment of Invoice Ref. ` + data.invoiceID + ` thorough PayPal.`).
			HTML(`To go back to invoice ` + invoiceLink.ToHTML() + ` ...`))

	return hb.NewDiv().
		Class(`container`).
		Child(title).
		Child(alert)
}

func (controller *paypalCancelController) prepareDataAndValidate(r *http.Request) (data paypalCancelControllerData, errorMessage string) {
	code := utils.Req(r, "code", "")

	if code == "" {
		return data, `Missing payment code`
	}

	codeData, err := payment.GetCancelPaymentData(code)

	if err != nil {
		return data, `Invalid payment code`
	}

	data.invoiceID = codeData.InvoiceID

	return data, ""
}
