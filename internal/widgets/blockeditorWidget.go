package widgets

import (
	"net/http"
	"project/internal/links"

	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

var _ Widget = (*blockeditorWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewBlockeditotWidget creates a new instance of the blockeditor widget
//
// Parameters:
//   - None
//
// Returns:
//   - *visibleWidget - A pointer to the show widget
func NewBlockeditotWidget() *blockeditorWidget {
	return &blockeditorWidget{}
}

// == WIDGET ================================================================

// blokeditorWidget
//
// Example:
// <x-blockeditor environment="production">content</x-visible>
type blockeditorWidget struct{}

func (w *blockeditorWidget) Alias() string {
	return "x-blockeditor"
}

func (w *blockeditorWidget) Description() string {
	return ""
}

func (w *blockeditorWidget) Render(r *http.Request, content string, params map[string]string) string {
	example := lo.ValueOr(params, "example", "")

	if example == "" {
		example = utils.Req(r, "example", "")
	}

	if example == "example1" {
		return w.Example1(r, content, params)
	}

	if example == "example2" {
		return w.Example2(r, content, params)
	}

	return "Example not found"
}

// == EXAMPLES ==============================================================

func (w *blockeditorWidget) Example1(r *http.Request, content string, params map[string]string) string {
	render := utils.Req(r, "render", "no")

	definitions := []blockeditor.BlockDefinition{
		{
			Icon: hb.I().Class(`bi bi-type-h1`),
			Type: "heading",
			Fields: []form.FieldInterface{
				form.NewField(form.FieldOptions{
					Name:  "content",
					Label: "Content (HTML is supported)",
					Type:  form.FORM_FIELD_TYPE_TEXTAREA,
				}),
			},
			ToTag: func(block ui.BlockInterface) *hb.Tag {
				content := block.Parameter("content")

				if content == "" {
					return hb.Div().
						Text("This block is empty. Open settings to add content.")
				}

				return hb.H1().HTML(content)
			},
		},
		{
			Icon: hb.I().Class(`bi bi-fonts`),
			Type: "paragraph",
			Fields: []form.FieldInterface{
				form.NewField(form.FieldOptions{
					Name:  "content",
					Label: "Content (HTML is supported)",
					Type:  form.FORM_FIELD_TYPE_TEXTAREA,
				}),
			},
			ToTag: func(block ui.BlockInterface) *hb.Tag {
				content := block.Parameter("content")

				if content == "" {
					return hb.Div().
						Text("This block is empty. Open settings to add content.")
				}

				return hb.P().HTML(content)
			},
		},
	}

	if render == "yes" {
		return blockeditor.Handle(nil, r, definitions)
	}

	editor, err := blockeditor.NewEditor(blockeditor.NewEditorOptions{
		ID:    "blockeditor1",
		Name:  "blockeditor1",
		Value: `[{"id":"20241105073231479400417132350866","type":"heading","content":"","parameters":{"content":"Lorem Ipsum"},"children":[]},{"id":"20241105073319278127915335089835","type":"paragraph","content":"","parameters":{"content":"Lectus sagittis elementum nec mi augue dapibus torquent vel orci gravida mus. Morbi consequat pharetra porta id vehicula ullamcorper viverra integer metus volutpat libero lorem. Facilisi platea maecenas et varius himenaeos. Curae; a auctor potenti duis accumsan lacinia justo primis massa rhoncus inceptos congue. Posuere commodo lacus cras massa venenatis lorem urna lacus blandit."},"children":[]},{"id":"20241105073559473087713897308998","type":"paragraph","content":"","parameters":{"content":"Aliquet, proin feugiat dis pharetra ut adipiscing tempus risus! Condimentum elit cum luctus torquent lacinia curae;? Vulputate aenean purus magna penatibus scelerisque tellus eros lobortis ut libero primis. Est id habitant imperdiet nullam habitant sit in potenti rutrum ullamcorper habitant! Eros tortor tellus neque blandit diam venenatis egestas risus. Netus aliquam class nulla elementum augue tellus nostra parturient metus sodales. Nascetur sagittis vivamus, gravida nec condimentum tincidunt sed. Non fusce tellus himenaeos. Ut laoreet lorem facilisis."},"children":[]}]`,
		HandleEndpoint: links.NewWebsiteLinks().Widget(w.Alias(), map[string]string{
			"example": "example1",
			"render":  "yes",
		}),
		BlockDefinitions: definitions,
	})

	if err != nil {
		return "Unable to create editor: " + err.Error()
	}

	return hb.Wrap().
		Child(hb.NewScriptURL(cdn.Htmx_2_0_0())).
		Child(hb.NewScriptURL(cdn.Sweetalert2_11())).
		Child(hb.NewStyleURL(cdn.BootstrapIconsCss_1_11_3())).
		Child(editor).
		ToHTML()
}

func (w *blockeditorWidget) Example2(r *http.Request, content string, params map[string]string) string {
	render := utils.Req(r, "render", "no")

	definitions := []blockeditor.BlockDefinition{
		{
			Icon: hb.I().Class(`bi bi-type-h1`),
			Type: "heading",
			Fields: []form.FieldInterface{
				form.NewField(form.FieldOptions{
					Name:  "content",
					Label: "Content (HTML is supported)",
					Type:  form.FORM_FIELD_TYPE_TEXTAREA,
				}),
			},
			ToTag: func(block ui.BlockInterface) *hb.Tag {
				content := block.Parameter("content")

				if content == "" {
					return hb.Div().
						Text("This block is empty. Open settings to add content.")
				}

				return hb.H1().HTML(content)
			},
		},
		{
			Icon: hb.I().Class(`bi bi-fonts`),
			Type: "paragraph",
			Fields: []form.FieldInterface{
				form.NewField(form.FieldOptions{
					Name:  "content",
					Label: "Content (HTML is supported)",
					Type:  form.FORM_FIELD_TYPE_TEXTAREA,
				}),
			},
			ToTag: func(block ui.BlockInterface) *hb.Tag {
				content := block.Parameter("content")

				if content == "" {
					return hb.Div().
						Text("This block is empty. Open settings to add content.")
				}

				return hb.P().HTML(content)
			},
		},
		{
			Icon:          hb.I().Class(`bi bi-distribute-vertical`),
			Type:          "row",
			AllowChildren: true,
			AllowedChildTypes: []string{
				"column",
			},
			Fields: []form.FieldInterface{},
			ToTag: func(block ui.BlockInterface) *hb.Tag {
				content := block.Parameter("content")

				return hb.Div().Class("row").HTMLIf(content != "", content)
			},
		},
		{
			Icon:          hb.I().Class(`bi bi-columns`),
			Type:          "column",
			AllowChildren: true,
			Fields: []form.FieldInterface{
				form.NewField(form.FieldOptions{
					Name:  "width",
					Label: "Width (1-12)",
					Type:  form.FORM_FIELD_TYPE_SELECT,
					Options: []form.FieldOption{
						{
							Value: "1",
							Key:   "1",
						},
						{
							Value: "2",
							Key:   "2",
						},
						{
							Value: "3",
							Key:   "3",
						},
						{
							Value: "4",
							Key:   "4",
						},
						{
							Value: "5",
							Key:   "5",
						},
						{
							Value: "6",
							Key:   "6",
						},
						{
							Value: "7",
							Key:   "7",
						},
						{
							Value: "8",
							Key:   "8",
						},
						{
							Value: "9",
							Key:   "9",
						},
						{
							Value: "10",
							Key:   "10",
						},
						{
							Value: "11",
							Key:   "11",
						},
						{
							Value: "12",
							Key:   "12",
						},
					},
				}),
			},
			Wrapper: func(block ui.BlockInterface) *hb.Tag {
				width := block.Parameter("width")

				if width == "" {
					return hb.Div().
						Text("This block is empty. Open settings to add content.")
				}

				return hb.Div().Class("col-" + width)
			},

			ToTag: func(block ui.BlockInterface) *hb.Tag {
				return hb.Wrap()
			},
		},
	}

	if render == "yes" {
		return blockeditor.Handle(nil, r, definitions)
	}

	editor, err := blockeditor.NewEditor(blockeditor.NewEditorOptions{
		ID:    "blockeditor2",
		Name:  "blockeditor2",
		Value: `[{"id":"20241105073231479400417132350866","type":"heading","content":"","parameters":{"content":"Lorem Ipsum"},"children":[]},{"id":"20241105123906331465414942482563","type":"row","content":"","parameters":{},"children":[{"id":"20241105123918395155816725940094","type":"column","content":"","parameters":{"width":"5"},"children":[{"id":"20241105124004593311415539635155","type":"paragraph","content":"","parameters":{"content":"Lectus sagittis elementum nec mi augue dapibus torquent vel orci gravida mus. Morbi consequat pharetra porta id vehicula ullamcorper viverra integer metus volutpat libero lorem. Facilisi platea maecenas et varius himenaeos. Curae; a auctor potenti duis accumsan lacinia justo primis massa rhoncus inceptos congue. Posuere commodo lacus cras massa venenatis lorem urna lacus blandit."},"children":[]}]},{"id":"20241105124028669903014932979879","type":"column","content":"","parameters":{"width":"7"},"children":[{"id":"20241105124059517446715517406622","type":"paragraph","content":"","parameters":{"content":"Aliquet, proin feugiat dis pharetra ut adipiscing tempus risus! Condimentum elit cum luctus torquent lacinia curae;? Vulputate aenean purus magna penatibus scelerisque tellus eros lobortis ut libero primis. Est id habitant imperdiet nullam habitant sit in potenti rutrum ullamcorper habitant! Eros tortor tellus neque blandit diam venenatis egestas risus. Netus aliquam class nulla elementum augue tellus nostra parturient metus sodales. Nascetur sagittis vivamus, gravida nec condimentum tincidunt sed. Non fusce tellus himenaeos. Ut laoreet lorem facilisis."},"children":[]}]}]}]`,
		HandleEndpoint: links.NewWebsiteLinks().Widget(w.Alias(), map[string]string{
			"example": "example2",
			"render":  "yes",
		}),
		BlockDefinitions: definitions,
	})

	if err != nil {
		return "Unable to create editor: " + err.Error()
	}

	return hb.Wrap().
		Child(hb.NewScriptURL(cdn.Htmx_2_0_0())).
		Child(hb.NewScriptURL(cdn.Sweetalert2_11())).
		Child(hb.NewStyleURL(cdn.BootstrapIconsCss_1_11_3())).
		Child(editor).
		ToHTML()
}
