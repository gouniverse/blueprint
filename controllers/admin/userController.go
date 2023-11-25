package admin

import (
	"net/http"
	"project/internal/links"
	"project/models"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/crud"
	"github.com/samber/lo"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (userController *userController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	usersCrud, err := crud.NewCrud(crud.CrudConfig{
		EntityNameSingular: "User",
		EntityNamePlural:   "Users",
		Endpoint:           links.NewAdminLinks().Users(),
		ColumnNames: []string{
			"Name",
			"BusinessName",
			"Email",
			"Created",
		},
		CreateFields: []crud.FormField{
			{
				Label: "First Name",
				Name:  "first_name",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Last Name",
				Name:  "last_name",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Email",
				Name:  "email",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
		},
		UpdateFields: []crud.FormField{
			{
				Label: "Status",
				Name:  "status",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "First Name",
				Name:  "first_name",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Last Name",
				Name:  "last_name",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			{
				Label: "Email",
				Name:  "email",
				Type:  crud.FORM_FIELD_TYPE_STRING,
			},
			// {
			// 	Label: "Phone",
			// 	Name:  "phone",
			// 	Type:  crud.FORM_FIELD_TYPE_STRING,
			// },
			// {
			// 	Label: "Business Name",
			// 	Name:  "business_name",
			// 	Type:  crud.FORM_FIELD_TYPE_STRING,
			// },
			{
				Label: "Role",
				Name:  "role",
				Type:  crud.FORM_FIELD_TYPE_SELECT,
				Options: []crud.FormFieldOption{
					{
						Key:   models.USER_ROLE_USER,
						Value: models.USER_ROLE_USER,
					},
					{
						Key:   models.USER_ROLE_MANAGER,
						Value: models.USER_ROLE_MANAGER,
					},
					{
						Key:   models.USER_ROLE_ADMINISTRATOR,
						Value: models.USER_ROLE_ADMINISTRATOR,
					},
					{
						Key:   models.USER_ROLE_SUPERUSER,
						Value: models.USER_ROLE_SUPERUSER,
					},
				},
			},
		},
		FuncRows:            userController.FuncRows,
		FuncCreate:          userController.FuncCreate,
		FuncTrash:           userController.FuncTrash,
		FuncUpdate:          userController.FuncUpdate,
		FuncFetchUpdateData: userController.FuncFetchUpdateData,
		FuncLayout:          userController.FuncLayout,
	})

	if err != nil {
		return "Error: " + err.Error()
	}

	usersCrud.Handler(w, r)
	return ""
}

func (userController *userController) FuncLayout(w http.ResponseWriter, r *http.Request, title string, content string, styleURLs []string, style string, scriptURLs []string, script string) string {
	scriptURLs = append([]string{
		cdn.Jquery_3_6_4()}, scriptURLs...,
	)

	return layout(r, layoutOptions{
		Title:      title + " | Admin",
		Content:    content,
		StyleURLs:  styleURLs,
		ScriptURLs: scriptURLs,
		Scripts:    []string{script},
		Styles: []string{
			`nav#Toolbar {border-bottom: 4px solid red;}`,
			style,
		},
	}).ToHTML()
}

// func (userController *userController) FuncCreate(entityID string, data map[string]string) error {

// }

func (userController *userController) FuncRows() ([]crud.Row, error) {
	users, err := models.NewUserService().UserList(models.UserQueryOptions{})

	if err != nil {
		return nil, err
	}

	rows := lo.Map(users, func(user models.User, _ int) crud.Row {
		return crud.Row{
			ID: user.ID(),
			Data: []string{
				user.FirstName() + " " + user.LastName(),
				user.BusinessName(),
				user.Email(),
				user.CreatedAtCarbon().Format("d M Y"),
			},
		}
	})

	return rows, nil
}

func (userController *userController) FuncUpdate(entityID string, data map[string]string) error {
	user, err := models.NewUserService().UserFindByID(entityID)

	if err != nil {
		return err
	}

	user.SetEmail(data["email"])
	user.SetFirstName(data["first_name"])
	user.SetLastName(data["last_name"])
	user.SetRole(data["role"])
	user.SetStatus(data["status"])

	err = models.NewUserService().UserUpdate(user)

	if err != nil {
		return err
	}

	return nil
}

func (userController *userController) FuncFetchUpdateData(userID string) (map[string]string, error) {
	user, err := models.NewUserService().UserFindByID(userID)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"first_name":    user.FirstName(),
		"last_name":     user.LastName(),
		"email":         user.Email(),
		"business_name": user.BusinessName(),
		"phone":         user.Phone(),
		"role":          user.Role(),
		"status":        user.Status(),
	}, nil
}

func (userController *userController) FuncCreate(data map[string]string) (userID string, err error) {
	user := models.NewUser()
	user.SetFirstName(data["first_name"])
	user.SetLastName(data["last_name"])
	user.SetEmail(data["email"])
	err = models.NewUserService().UserCreate(user)

	if err != nil {
		return "", err
	}

	return user.ID(), nil
}

func (userController *userController) FuncTrash(userID string) error {
	err := models.NewUserService().UserDeleteByID(userID)
	return err
}
