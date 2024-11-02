package blocks

import (
	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func BlockEditorDefinitions() []blockeditor.BlockDefinition {
	return []blockeditor.BlockDefinition{
		blockHeadingDefinition(),
		blockParagraphDefinition(),
		blockImageDefinition(),
		blockHyperlinkDefinition(),
		blockUnorderedListDefinition(),
		blockOrderedListDefinition(),
		blockListItemDefinition(),
	}
}

func blockUnorderedListDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:   hb.I().Class("bi bi-list-ul"),
		Type:   "unordered_list",
		Fields: []form.Field{},
		AllowedChildren: []string{
			"list_item",
		},
	}
}

func blockOrderedListDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:   hb.I().Class("bi bi-list-ol"),
		Type:   "ordered_list",
		Fields: []form.Field{},
		AllowedChildren: []string{
			"list_item",
		},
	}
}

func blockListItemDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-list"),
		Type: "list_item",
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
		ToHTML: func(block ui.BlockInterface) string {
			content := block.Parameter("content")

			if content == "" {
				return "Add list item text"
			}

			return hb.Raw(content).ToHTML()
		},
	}
}

func blockHeadingDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-type-h1"),
		Type: "heading",
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
		ToHTML: func(block ui.BlockInterface) string {
			level := block.Parameter("level")
			content := block.Parameter("content")

			if content == "" {
				return "Add heading text"
			}

			if level == "" {
				level = "1"
			}

			return hb.NewTag("h" + level).
				HTML(content).
				ToHTML()
		},
	}
}

func blockParagraphDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-paragraph"),
		Type: "paragraph",
		Fields: []form.Field{
			{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			},
		},
		ToHTML: func(block ui.BlockInterface) string {
			content := block.Parameter("content")

			if content == "" {
				return "Add paragraph text"
			}

			return hb.P().
				HTML(content).
				ToHTML()
		},
	}
}

func blockHyperlinkDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-link-45deg"),
		Type: "hyperlink",
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
		ToHTML: func(block ui.BlockInterface) string {
			content := block.Parameter("content")

			if content == "" {
				return "Add link text"
			}

			return hb.A().
				Href(block.Parameter("url")).
				HTML(block.Parameter("content")).
				Target(block.Parameter("target")).
				ToHTML()
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
		ToHTML: func(block ui.BlockInterface) string {
			imageUrl := block.Parameter("image_url")

			if imageUrl == "" {
				return "Add image URL"
			}

			return hb.Img(imageUrl).ToHTML()
		},
	}
}
