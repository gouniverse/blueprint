package controllers

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"project/internal/payment"
	"project/pkg/accountingstore"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

type paypalNotifyController struct{}

type paypalNotifyControllerData struct {
	invoiceID string
}

func NewPaypalNotifyController() *paypalNotifyController {
	return &paypalNotifyController{}
}

func (controller *paypalNotifyController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
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

func (controller *paypalNotifyController) page(data paypalNotifyControllerData) hb.TagInterface {
	title := hb.NewHeading1().
		Style(`margin-top: 10px; margin-bottom: 10px`).
		HTML(`PayPal Invoice Payment Notified`)

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

func (controller *paypalNotifyController) prepareDataAndValidate(r *http.Request) (data paypalNotifyControllerData, errorMessage string) {
	code := utils.Req(r, "code", "")

	if code == "" {
		return data, `Missing payment code`
	}

	codeData, err := payment.GetNotifyPaymentData(code)

	if err != nil {
		return data, `Invalid payment code`
	}

	data.invoiceID = codeData.InvoiceID

	if data.invoiceID == "" {
		return data, "Invoice ID is missing"
	}

	invoice, err := config.AccountingStore.InvoiceFindByID(data.invoiceID)

	if err != nil {
		config.LogStore.ErrorWithContext("At invoiceViewController > prepareDataAndValidate", err.Error())
		return data, "Invoice not found"
	}

	if invoice == nil {
		return data, "Invoice not found"
	}

	invoice.SetStatus(accountingstore.INVOICE_STATUS_PAID)

	err = config.AccountingStore.InvoiceUpdate(invoice)

	if err != nil {
		config.LogStore.ErrorWithContext("At invoiceViewController > prepareDataAndValidate", err.Error())
		return data, "Error updating invoice. Please contact an administrator."
	}

	return data, ""
}
