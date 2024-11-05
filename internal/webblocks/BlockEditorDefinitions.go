package webblocks

import (
	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
)

func BlockEditorDefinitions() []blockeditor.BlockDefinition {
	definitions := []blockeditor.BlockDefinition{
		blockSectionDefinition(),
		blockContainerDefinition(),
		blockRowDefinition(),
		blockColumnDefinition(),
		blockHeadingDefinition(),
		blockParagraphDefinition(),
		blockImageDefinition(),
		blockHyperlinkDefinition(),
		blockUnorderedListDefinition(),
		blockOrderedListDefinition(),
		blockListItemDefinition(),
		blockCodeDefinition(),
		blockBreadcrumbsDefinition(),
		blockRawHTMLDefinition(),
	}

	for index, definition := range definitions {
		definition.Fields = append([]form.Field{{
			Name:  "status",
			Label: "Status",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "Published",
					Key:   "published",
				},
				{
					Value: "Unpublished",
					Key:   "unpublished",
				},
			},
		}}, definition.Fields...)

		definition.Fields = append([]form.Field{{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Block Settings</legend>`,
		}}, definition.Fields...)

		definition.Fields = append(definition.Fields, form.Field{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		})

		definition.Fields = append(definition.Fields, blockeditor.FieldsHTML()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsAlign()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsText()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsBackground()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsBorder()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsFont()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsMargin()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsPadding()...)

		definitions[index] = definition
	}

	return definitions
}

func blockBreadcrumbsDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-three-dots"),
		Type: "breadcrumbs",
		Fields: []form.Field{
			{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: hb.H2().Text("Breadcrumb 1").ToHTML(),
			},
			{
				Name:  "breadcrumb1_url",
				Label: "Breadcrumb 1 URL",
				Type:  form.FORM_FIELD_TYPE_STRING,
			},
			{
				Name:  "breadcrumb1_text",
				Label: "Breadcrumb 1 Text",
				Type:  form.FORM_FIELD_TYPE_STRING,
			},
			{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: hb.H2().Text("Breadcrumb 2").ToHTML(),
			},
			{
				Name:  "breadcrumb2_url",
				Label: "Breadcrumb 2 URL",
				Type:  form.FORM_FIELD_TYPE_STRING,
			},
			{
				Name:  "breadcrumb2_text",
				Label: "Breadcrumb 2 Text",
				Type:  form.FORM_FIELD_TYPE_STRING,
			},
			{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: hb.H2().Text("Breadcrumb 3").ToHTML(),
			},
			{
				Name:  "breadcrumb3_url",
				Label: "Breadcrumb 3 URL",
				Type:  form.FORM_FIELD_TYPE_STRING,
			},
			{
				Name:  "breadcrumb3_text",
				Label: "Breadcrumb 3 Text",
				Type:  form.FORM_FIELD_TYPE_STRING,
			},
		},
	}
}

func blockOrderedListDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-list-ol"),
		Type:          "ordered_list",
		Fields:        []form.Field{},
		AllowChildren: true,
		AllowedChildTypes: []string{
			"list_item",
		},
	}
}

func blockHeadingDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-type-h1"),
		Type:          "heading",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "level",
				Label: "Level",
				Type:  form.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Value: "Heading Level 1",
						Key:   "1",
					},
					{
						Value: "Heading Level 2",
						Key:   "2",
					},
					{
						Value: "Heading Level 3",
						Key:   "3",
					},
					{
						Value: "Heading Level 4",
						Key:   "4",
					},
					{
						Value: "Heading Level 5",
						Key:   "5",
					},
					{
						Value: "Heading Level 6",
						Key:   "6",
					},
				},
			},
			{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			level := block.Parameter("level")
			content := block.Parameter("content")

			if level == "" {
				level = "1"
			}

			if content == "" {
				content = "Add heading text"
			}

			return hb.NewTag("h"+level).
				HTMLIf(content != "", content)
		},
	}
}

func blockCodeDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class(`bi bi-code`),
		Type: "code",
		Fields: []form.Field{
			{
				Name:  "language",
				Label: "Language",
				Type:  form.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Value: "bash",
						Key:   "bash",
					},
					{
						Value: "css",
						Key:   "css",
					},
					{
						Value: "golang",
						Key:   "golang",
					},
					{
						Value: "html",
						Key:   "html",
					},
					{
						Value: "javascript",
						Key:   "javascript",
					},
					{
						Value: "json",
						Key:   "json",
					},
					{
						Value: "markdown",
						Key:   "markdown",
					},
					{
						Value: "php",
						Key:   "php",
					},
					{
						Value: "python",
						Key:   "python",
					},
					{
						Value: "sql",
						Key:   "sql",
					},
				},
			},
			{
				Name:  "content",
				Label: "Code",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
	}
}

func blockColumnDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-columns`),
		Type:          "column",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "width",
				Label: "Width",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
			},
		},
	}
}

func blockContainerDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-layout-text-window`),
		Type:          "container",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Text",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
	}
}

func blockHyperlinkDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-link-45deg"),
		Type:          "hyperlink",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "url",
				Label: "URL",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
			{
				Name:  "content",
				Label: "Text (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
			{
				Name:  "target",
				Label: "Target",
				Type:  form.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Value: "_blank",
						Key:   "_blank",
					},
					{
						Value: "_self",
						Key:   "_self",
					},
					{
						Value: "_parent",
						Key:   "_parent",
					},
					{
						Value: "_top",
						Key:   "_top",
					},
				},
			},
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")
			url := block.Parameter("url")
			target := block.Parameter("target")

			if url == "" {
				url = "#"
			}

			if target == "" {
				target = "_self"
			}

			return hb.A().
				Href(url).
				HTMLIf(content != "", content).
				Target(target)
		},
	}
}

func blockImageDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-image"),
		Type: "image",
		Fields: []form.Field{
			{
				Name:  "image_url",
				Label: "Image URL (may be base64)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			imageUrl := block.Parameter("image_url")

			if imageUrl == "" {
				imageUrl = utils.PicsumURL(100, 100, utils.PicsumURLOptions{
					Grayscale: true,
					Seed:      "no image",
				})
			}

			return hb.Img(imageUrl)
		},
	}
}

func blockUnorderedListDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-list-ul"),
		Type:          "unordered_list",
		AllowChildren: true,
		Fields:        []form.Field{},
		AllowedChildTypes: []string{
			"list_item",
		},
	}
}

func blockListItemDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-list"),
		Type:          "list_item",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")

			if content == "" {
				content = "Add list item text"
			}

			return hb.Raw(content)
		},
	}
}

func blockParagraphDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-paragraph"),
		Type:          "paragraph",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")

			if content == "" {
				content = "Add paragraph text"
			}

			return hb.P().
				HTML(content)
		},
	}
}

func blockRawHTMLDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class(`bi bi-code-slash`),
		Type: "raw_html",
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "HTML",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
	}
}

func blockRowDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-distribute-vertical`),
		Type:          "row",
		AllowChildren: true,
		Fields:        []form.Field{},
	}
}

func blockSectionDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-distribute-vertical`),
		Type:          "section",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Text",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
	}
}

func blockTextDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-fonts`),
		Type:          "text",
		AllowChildren: true,
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Text",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
	}
}
