package admin

import (
	"errors"
	"net/http"
	"project/config"
	"project/internal/links"

	"github.com/gouniverse/api"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/crud"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"

	// "github.com/gouniverse/crud"
	"github.com/samber/lo"
)

func NewBlogController() *blogController {
	return &blogController{}
}

type blogController struct{}

// Index is the handler function for the blogController
//
// Parameters:
//   - w http.ResponseWriter
//   - r *http.Request
//
// Returns:
//   - string
func (c blogController) Index(w http.ResponseWriter, r *http.Request) string {
	entityID := utils.Req(r, "entity_id", "")
	contentType := ""
	contentScript := ""
	if entityID != "" {
		blogpost, err := config.BlogStore.PostFindByID(entityID)
		if err != nil {
			config.LogStore.ErrorWithContext("At blogController > AnyIndex", err.Error())
		}
		if blogpost != nil {
			editor := blogpost.Meta("editor")
			contentType = lo.Ternary(editor == blogstore.POST_EDITOR_HTMLAREA, crud.FORM_FIELD_TYPE_HTMLAREA, crud.FORM_FIELD_TYPE_TEXTAREA)
			contentScript = lo.Ternary(editor == blogstore.POST_EDITOR_BLOCKAREA, hb.NewScript(`setTimeout(() => {
		const blockArea = new BlockArea('content');
		blockArea.setParentId('`+entityID+`');
		blockArea.registerBlock(BlockAreaHeading);
		blockArea.registerBlock(BlockAreaText);
		blockArea.registerBlock(BlockAreaImage);
		blockArea.registerBlock(BlockAreaCode);
		blockArea.registerBlock(BlockAreaRawHtml);
		blockArea.init();
	}, 2000)`).
				ToHTML(), "")
		}
	}

	crudInstance, err := crud.NewCrud(crud.CrudConfig{
		HomeURL:  links.NewAdminLinks().Home(),
		Endpoint: links.NewAdminLinks().Blog(),
		// FileManagerURL:      links.NewAdminLinks().FileManager(map[string]string{}),
		EntityNamePlural:    "Blogs",
		EntityNameSingular:  "Blog",
		ColumnNames:         []string{"Title", "Status", "Created", "Modified"},
		FuncLayout:          adminCrudLayout,
		FuncCreate:          c.postCreate,
		FuncRows:            c.postRows,
		FuncFetchUpdateData: c.postFetchUpdateData,
		FuncTrash:           c.postTrash,
		FuncUpdate:          c.postUpdate,
		CreateFields: []crud.FormField{
			{
				Type:  crud.FORM_FIELD_TYPE_STRING,
				Label: "Title",
				Name:  "title",
			},
		},
		UpdateFields: []crud.FormField{
			{
				Type:  "select",
				Label: "Status",
				Name:  "status",
				Options: []crud.FormFieldOption{
					{
						Key:   "",
						Value: "- not selected - ",
					},
					{
						Key:   blogstore.POST_STATUS_DRAFT,
						Value: "Draft",
					},
					{
						Key:   blogstore.POST_STATUS_PUBLISHED,
						Value: "Published",
					},
					{
						Key:   blogstore.POST_STATUS_UNPUBLISHED,
						Value: "Unpublished",
					},
				},
			},
			{
				Type:  crud.FORM_FIELD_TYPE_STRING,
				Label: "Title",
				Name:  "title",
				Help:  "The title of this blog as will be seen everywhere",
			},
			{
				Type:  crud.FORM_FIELD_TYPE_TEXTAREA,
				Label: "Summary",
				Name:  "summary",
				Help:  "A short summary of this blog post to display on the list page.",
			},
			{
				Type:  crud.FORM_FIELD_TYPE_SELECT,
				Label: "Editor",
				Name:  "editor",
				Help:  "The editor that will be used while editing this blogpost.",
				Options: []crud.FormFieldOption{
					{
						Key:   "",
						Value: "- not selected - ",
					},
					{
						Key:   blogstore.POST_EDITOR_BLOCKAREA,
						Value: "BlockArea",
					},
					{
						Key:   blogstore.POST_EDITOR_MARKDOWN,
						Value: "Markdown",
					},
					{
						Key:   blogstore.POST_EDITOR_HTMLAREA,
						Value: "HTML Area",
					},
					{
						Key:   blogstore.POST_EDITOR_TEXTAREA,
						Value: "Text Area",
					},
				},
			},
			{
				Type:  crud.FORM_FIELD_TYPE_RAW,
				Value: contentScript,
			},
			{
				Type:  contentType,
				ID:    "content",
				Label: "Content",
				Name:  "content",
				Help:  "The content of this blog post to display on the details page.",
			},
			{
				Type:  crud.FORM_FIELD_TYPE_IMAGE_INLINE,
				Label: "Image URL",
				Name:  "image_url",
				Help:  "The image URL for the blogpost.",
			},
			{
				Type:  crud.FORM_FIELD_TYPE_DATETIME,
				Label: "Date Published",
				Name:  "published_at",
				Help:  "The date this blog post was published.",
			},
			{
				Type: crud.FORM_FIELD_TYPE_SELECT,
				Options: []crud.FormFieldOption{
					{
						Key:   "",
						Value: "- not selected - ",
					},
					{
						Key:   blogstore.NO,
						Value: "No",
					},
					{
						Key:   blogstore.YES,
						Value: "Yes",
					},
				},
				Label: "Is Featured?",
				Name:  "featured",
				Help:  "Should this blogpost be displayed as featured (on the homepage).",
			},
			{
				Type:  crud.FORM_FIELD_TYPE_STRING,
				Label: "SEO. Meta Keywords",
				Name:  "meta_keywords",
				Help:  "The meta keywords for the blogpost. Separate multiple keywords with a comma.",
			},
			{
				Type:  crud.FORM_FIELD_TYPE_STRING,
				Label: "SEO. Meta Description",
				Name:  "meta_description",
				Help:  "The meta description for the blogpost.",
			},
			{
				Type: "select",
				Options: []crud.FormFieldOption{
					{
						Key:   "",
						Value: "- not selected - ",
					},
					{
						Key:   "",
						Value: "INDEX, FOLLOW",
					},
					{
						Key:   "",
						Value: "NOINDEX, FOLLOW",
					},
					{
						Key:   "",
						Value: "INDEX, NOFOLLOW",
					},
					{
						Key:   "",
						Value: "NOINDEX, NOFOLLOW",
					},
				},
				Label: "SEO. Meta Robots Follow",
				Name:  "meta_robots",
				Help:  "The meta robots follow for the blogpost.",
			},
			{
				Type:  "string",
				Label: "SEO. Cannonical URL",
				Name:  "canonical_url",
				Help:  "The canonical URL that this page should point to. Leave empty to default to default page URL. ",
			},
		},
	})

	if err != nil {
		api.Respond(w, r, api.Error("Error: "+err.Error()))
		return ""
	}

	crudInstance.Handler(w, r)
	return ""
}

func (c blogController) postCreate(data map[string]string) (serviceID string, err error) {
	title := data["title"]

	if title == "" {
		return "", errors.New("title is required field")
	}

	post := blogstore.NewPost().SetTitle(title).SetStatus(blogstore.POST_STATUS_UNPUBLISHED)
	err = config.BlogStore.PostCreate(post)

	if err != nil {
		config.LogStore.ErrorWithContext("At blogController > postCreate", err.Error())
		return "", errors.New("post failed to be created")
	}

	return post.ID(), nil
}

func (c blogController) postFetchUpdateData(entityID string) (map[string]string, error) {
	post, err := config.BlogStore.PostFindByID(entityID)

	if err != nil {
		config.LogStore.ErrorWithContext("At blogController > postFetchUpdateData", err.Error())
		return nil, errors.New("user fetch failed")
	}

	data := post.Data()
	data["editor"] = post.Meta("editor")

	publishedAt := lo.Substring(lo.ValueOr(data, "published_at", sb.NULL_DATETIME), 0, 19)
	data["published_at"] = publishedAt

	return data, nil
}

func (c blogController) postTrash(entityID string) error {
	post, err := config.BlogStore.PostFindByID(entityID)
	if err != nil {
		config.LogStore.ErrorWithContext("At blogController > postTrash", err.Error())
		return err
	}
	return config.BlogStore.PostTrash(post)
}

func (c blogController) postUpdate(entityID string, data map[string]string) error {
	post, err := config.BlogStore.PostFindByID(entityID)

	if err != nil {
		config.LogStore.ErrorWithContext("At blogController > postUpdate", err.Error())
		return errors.New("user find failed")
	}

	title := data["title"]
	status := data["status"]
	editor := data["editor"]
	delete(data, "editor")

	if title == "" {
		return errors.New("title is required field")
	}

	if status == "" {
		return errors.New("status is required field")
	}

	if editor == "" {
		return errors.New("editor is required field")
	}

	publishedAt := lo.Substring(lo.ValueOr(data, "published_at", sb.NULL_DATETIME), 0, 19)
	data["published_at"] = publishedAt

	post.SetData(data)
	post.AddMetas(map[string]string{
		"editor": editor,
	})

	errUpdate := config.BlogStore.PostUpdate(post)

	if errUpdate != nil {
		config.LogStore.ErrorWithContext("At blogController > postUpdate", errUpdate.Error())
		return errUpdate
	}

	return nil
}

// postRows returns a list of posts
func (c blogController) postRows() ([]crud.Row, error) {
	pageNumber := 0
	perPage := 1000
	posts, err := config.BlogStore.
		PostList(blogstore.PostQueryOptions{
			Limit:  perPage,
			Offset: pageNumber * perPage,
			StatusIn: []string{
				blogstore.POST_STATUS_DRAFT,
				blogstore.POST_STATUS_PUBLISHED,
				blogstore.POST_STATUS_UNPUBLISHED,
			},
		})

	if err != nil {
		config.LogStore.ErrorWithContext("At blogController > postRows", err.Error())
		return []crud.Row{}, err
	}

	rows := []crud.Row{}
	for _, post := range posts {
		// created := post.CreatedAtCarbon().Format("d M Y")
		updated := post.UpdatedAtCarbon().Format("d M Y")
		published := post.PublishedAtCarbon().Format("d M Y")

		row := crud.Row{}
		row.ID = post.ID()
		row.Data = append(row.Data, post.Title())
		row.Data = append(row.Data, post.Status())
		row.Data = append(row.Data, published)
		row.Data = append(row.Data, updated)
		rows = append(rows, row)
	}

	return rows, nil
}
