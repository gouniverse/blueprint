package admin

// import (
// 	"errors"
// 	"net/http"
// 	"project/config"
// 	"project/controllers/admin/partials"
// 	"project/app/links"
// 	"project/models"

// 	"github.com/gouniverse/api"
// 	"github.com/gouniverse/crud"
// )

// func NewFaqController() *faqController {
// 	return &faqController{}
// }

// type faqController struct{}

// func (c faqController) Handler(w http.ResponseWriter, r *http.Request) {
// 	crudInstance, err := crud.NewCrud(crud.CrudConfig{
// 		HomeURL:             links.NewAdminLinks().Home(map[string]string{}),
// 		Endpoint:            links.NewAdminLinks().Faq(map[string]string{}),
// 		FileManagerURL:      links.NewAdminLinks().FileManager(map[string]string{}),
// 		EntityNamePlural:    "FAQ",
// 		EntityNameSingular:  "FAQ",
// 		ColumnNames:         []string{"Question", "Status", "Category"},
// 		FuncLayout:          partials.AdminDashboardCrudLayout,
// 		FuncCreate:          c.faqCreate,
// 		FuncRows:            c.faqRows,
// 		FuncFetchUpdateData: c.faqFetchUpdateData,
// 		FuncTrash:           c.faqTrash,
// 		FuncUpdate:          c.faqUpdate,
// 		CreateFields: []crud.FormField{
// 			{
// 				Type:  crud.FORM_FIELD_TYPE_STRING,
// 				Label: "Question",
// 				Name:  "question",
// 			},
// 		},
// 		UpdateFields: []crud.FormField{
// 			{
// 				Type:  "select",
// 				Label: "Status",
// 				Name:  "status",
// 				Options: []crud.FormFieldOption{
// 					{
// 						Key:   "",
// 						Value: "- not selected - ",
// 					},
// 					{
// 						Key:   models.FAQ_STATUS_DRAFT,
// 						Value: "Draft",
// 					},
// 					{
// 						Key:   models.FAQ_STATUS_ACTIVE,
// 						Value: "Active",
// 					},
// 					{
// 						Key:   models.FAQ_STATUS_INACTIVE,
// 						Value: "Inactive",
// 					},
// 				},
// 			},
// 			{
// 				Type:  crud.FORM_FIELD_TYPE_STRING,
// 				Label: "Question",
// 				Name:  "question",
// 				Help:  "The question of this FAQ as will be seen everywhere",
// 			},
// 			{
// 				Type:  crud.FORM_FIELD_TYPE_HTMLAREA,
// 				Label: "Answer",
// 				Name:  "answer",
// 				Help:  "The content of this blog faq to display on the details page.",
// 			},
// 			{
// 				Type: crud.FORM_FIELD_TYPE_SELECT,
// 				Options: []crud.FormFieldOption{
// 					{
// 						Key:   "",
// 						Value: "- not selected - ",
// 					},
// 					{
// 						Key:   models.FAQ_CATEGORY_GENERAL_GUIDE,
// 						Value: "General Guide",
// 					},
// 					{
// 						Key:   models.FAQ_CATEGORY_SCHOOL_GUIDE,
// 						Value: "School Guide",
// 					},
// 					{
// 						Key:   models.FAQ_CATEGORY_STUDENT_GUIDE,
// 						Value: "Student Guide",
// 					},
// 					{
// 						Key:   models.FAQ_CATEGORY_TEACHER_GUIDE,
// 						Value: "Teacher Guide",
// 					},
// 				},
// 				Label: "Category",
// 				Name:  "category",
// 				Help:  "In which category should this FAQ be listed.",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		api.Respond(w, r, api.Error("Error: "+err.Error()))
// 		return
// 	}

// 	crudInstance.Handler(w, r)
// }

// func (c faqController) faqCreate(data map[string]string) (serviceID string, err error) {
// 	question := data["question"]

// 	if question == "" {
// 		return "", errors.New("question is required field")
// 	}

// 	faq := models.NewFAQ().SetQuestion(question).SetStatus(models.FAQ_STATUS_DRAFT)
// 	err = models.NewFaqRepository().FaqCreate(faq)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At faqController > faqCreate", err.Error())
// 		return "", errors.New("faq failed to be created")
// 	}

// 	return faq.ID(), nil
// }

// func (c faqController) faqFetchUpdateData(entityID string) (map[string]string, error) {
// 	faq, err := models.NewFaqRepository().FaqFindByID(entityID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At faqController > faqFetchUpdateData", err.Error())
// 		return nil, errors.New("user fetch failed")
// 	}

// 	data := faq.Data()

// 	// publishedAt := lo.Substring(lo.ValueOr(data, "published_at", models.NULL_DATETIME), 0, 19)
// 	// data["published_at"] = publishedAt

// 	return data, nil
// }

// func (c faqController) faqTrash(entityID string) error {
// 	faq, err := models.NewFaqRepository().FaqFindByID(entityID)
// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At faqController > faqTrash", err.Error())
// 		return err
// 	}
// 	return models.NewFaqRepository().FaqTrash(faq)
// }

// func (c faqController) faqUpdate(entityID string, data map[string]string) error {
// 	faq, err := models.NewFaqRepository().FaqFindByID(entityID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At faqController > faqUpdate", err.Error())
// 		return errors.New("faq find failed")
// 	}

// 	question := data["question"]
// 	status := data["status"]
// 	answer := data["answer"]
// 	category := data["category"]

// 	if question == "" {
// 		return errors.New("question is required field")
// 	}

// 	if status == "" {
// 		return errors.New("status is required field")
// 	}

// 	if answer == "" {
// 		return errors.New("answer is required field")
// 	}

// 	if category == "" {
// 		return errors.New("category is required field")
// 	}

// 	faq.SetData(data)

// 	errUpdate := models.NewFaqRepository().FaqUpdate(faq)

// 	if errUpdate != nil {
// 		config.Cms.LogStore.ErrorWithContext("At faqController > faqUpdate", errUpdate.Error())
// 		return errUpdate
// 	}

// 	return nil
// }

// // faqRows returns a list of faqs
// func (c faqController) faqRows() ([]crud.Row, error) {
// 	pageNumber := 0
// 	perPage := 1000
// 	faqs, err := models.NewFaqRepository().FaqList(models.FaqQueryOptions{
// 		Limit:  perPage,
// 		Offset: pageNumber * perPage,
// 		// StatusIn: []string{
// 		// models.BLOGPOST_STATUS_DRAFT,
// 		// models.BLOGPOST_STATUS_PUBLISHED,
// 		// models.BLOGPOST_STATUS_UNPUBLISHED,
// 		// },
// 	})

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At faqController > faqRows", err.Error())
// 		return []crud.Row{}, err
// 	}

// 	rows := []crud.Row{}
// 	for _, faq := range faqs {
// 		row := crud.Row{}
// 		row.ID = faq.ID()
// 		row.Data = append(row.Data, faq.Question())
// 		row.Data = append(row.Data, faq.Status())
// 		row.Data = append(row.Data, faq.Category())
// 		rows = append(rows, row)
// 	}

// 	return rows, nil
// }
