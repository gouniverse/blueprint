package blogblocks

import (
	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
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
		Icon:          hb.I().Class("bi bi-list-ul"),
		Type:          "unordered_list",
		Fields:        []form.Field{},
		AllowChildren: true,
		AllowedChildTypes: []string{
			"list_item",
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
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")
			
			if content == "" {
				content = "Add list item text"
			}

			return hb.LI().HTML(content)
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
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			level := block.Parameter("level")
			content := block.Parameter("content")

			if level == "" {
				level = "1"
			}

			return hb.NewTag("h"+level).
				HTMLIf(content != "", `Add heading text`).
				HTML(content)
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
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")

			return hb.P().
				HTMLIf(content == "", `Add paragraph text`).
				HTML(content)
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
