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

type paypalSuccessController struct{}
type paypalSuccessControllerData struct {
	invoiceID string
}

func NewPaypalSuccessController() *paypalSuccessController {
	return &paypalSuccessController{}
}

func (controller *paypalSuccessController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
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

func (controller *paypalSuccessController) page(data paypalSuccessControllerData) hb.TagInterface {
	title := hb.NewHeading1().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		HTML(`PayPal Invoice Payment Successful`)

	invoiceUrl := links.NewCustomerLinks().InvoiceView(data.invoiceID, map[string]string{})
	invoiceLink := hb.NewHyperlink().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		Href(invoiceUrl).
		HTML(`CLICK HERE`)

	alert := hb.NewDiv().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		Child(hb.NewDiv().
			Class(`alert alert-success`).
			HTML(`You successfully paid for Invoice Ref. ` + data.invoiceID + ` thorough PayPal.`).
			HTML(`To go back to invoice ` + invoiceLink.ToHTML() + ` ...`))

	return hb.NewDiv().
		Class(`container`).
		Child(title).
		Child(alert)
}

func (controller *paypalSuccessController) prepareDataAndValidate(r *http.Request) (data paypalSuccessControllerData, errorMessage string) {
	code := utils.Req(r, "code", "")

	if code == "" {
		return data, `Missing payment code`
	}

	// cfmt.Warningln("HERE 1", code)

	codeData, err := payment.GetSuccessPaymentData(code)

	// cfmt.Warningln("HERE 2")

	if err != nil {
		return data, `Invalid payment code`
	}

	// cfmt.Warningln("HERE 3")

	data.invoiceID = codeData.InvoiceID

	return data, ""
}
