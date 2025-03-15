package webtheme

import (
	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

func BlockEditorDefinitions() []blockeditor.BlockDefinition {
	definitions := []blockeditor.BlockDefinition{
		blockSectionDefinition(),
		blockContainerDefinition(),
		blockRowDefinition(),
		blockColumnDefinition(),
		blockDivDefinition(),
		blockParagraphDefinition(),
		blockHeadingDefinition(),
		blockTextDefinition(),
		blockIconDefinition(),
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
		definition.Fields = append([]form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "status",
				Label: "Status",
				Type:  form.FORM_FIELD_TYPE_SELECT,
				Options: []form.FieldOption{
					{
						Value: "",
						Key:   "",
					},
					{
						Value: "Published",
						Key:   "published",
					},
					{
						Value: "Unpublished",
						Key:   "unpublished",
					},
				},
			})}, definition.Fields...)

		definition.Fields = append([]form.FieldInterface{
			blockeditor.FieldGroupStart("block_settings", "Block Settings", false),
		}, definition.Fields...)

		definition.Fields = append(definition.Fields, blockeditor.FieldGroupEnd())

		definition.Fields = append(definition.Fields, blockeditor.FieldsHTML()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsDisplay()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsPositioning()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsAlign()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsText()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsBackground()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsBorder()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsFont()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsSize()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsPadding()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsMargin()...)
		definition.Fields = append(definition.Fields, blockeditor.FieldsFlexBox()...)
		definitions[index] = definition
	}

	return definitions
}

func applyAllParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	blockeditor.ApplyAlignmentParameters(block, blockTag)
	blockeditor.ApplyAnimationParameters(block, blockTag)
	blockeditor.ApplyBackgroundParameters(block, blockTag)
	blockeditor.ApplyBorderParameters(block, blockTag)
	blockeditor.ApplyDisplayParameters(block, blockTag)
	blockeditor.ApplyFontParameters(block, blockTag)
	blockeditor.ApplyFlexBoxParameters(block, blockTag)
	blockeditor.ApplyHTMLParameters(block, blockTag)
	blockeditor.ApplyMarginParameters(block, blockTag)
	blockeditor.ApplyPaddingParameters(block, blockTag)
	blockeditor.ApplyPositionParameters(block, blockTag)
	blockeditor.ApplySizeParameters(block, blockTag)
	blockeditor.ApplyTextParameters(block, blockTag)
	blockeditor.ApplyTransitionParameters(block, blockTag)
}

func blockBreadcrumbsDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-three-dots"),
		Type: "breadcrumbs",
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: hb.H2().Text("Breadcrumb 1").ToHTML(),
			}),
			form.NewField(form.FieldOptions{
				Name:  "breadcrumb1_url",
				Label: "Breadcrumb 1 URL",
				Type:  form.FORM_FIELD_TYPE_STRING,
			}),
			form.NewField(form.FieldOptions{
				Name:  "breadcrumb1_text",
				Label: "Breadcrumb 1 Text",
				Type:  form.FORM_FIELD_TYPE_STRING,
			}),
			form.NewField(form.FieldOptions{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: hb.H2().Text("Breadcrumb 2").ToHTML(),
			}),
			form.NewField(form.FieldOptions{
				Name:  "breadcrumb2_url",
				Label: "Breadcrumb 2 URL",
				Type:  form.FORM_FIELD_TYPE_STRING,
			}),
			form.NewField(form.FieldOptions{
				Name:  "breadcrumb2_text",
				Label: "Breadcrumb 2 Text",
				Type:  form.FORM_FIELD_TYPE_STRING,
			}),
			form.NewField(form.FieldOptions{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: hb.H2().Text("Breadcrumb 3").ToHTML(),
			}),
			form.NewField(form.FieldOptions{
				Name:  "breadcrumb3_url",
				Label: "Breadcrumb 3 URL",
				Type:  form.FORM_FIELD_TYPE_STRING,
			}),
			form.NewField(form.FieldOptions{
				Name:  "breadcrumb3_text",
				Label: "Breadcrumb 3 Text",
				Type:  form.FORM_FIELD_TYPE_STRING,
			}),
		},
	}
}

func blockCodeDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class(`bi bi-code`),
		Type: "code",
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
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
			}),
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Code",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
	}
}

func blockColumnDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-columns`),
		Type:          "column",
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "width_xs",
				Label: "Width XS",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
				Help:  `The width of the column on extra small devices (<576px, mobile phones). `,
			}),
			form.NewField(form.FieldOptions{
				Name:  "width_sm",
				Label: "Width SM",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
				Help:  `The width of the column on small devices (≥576px, portrait phones, tablets).`,
			}),
			form.NewField(form.FieldOptions{
				Name:  "width_md",
				Label: "Width MD",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
				Help:  `The width of the column on medium devices (≥768px, desktops).`,
			}),
			form.NewField(form.FieldOptions{
				Name:  "width_lg",
				Label: "Width LG",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
				Help:  `The width of the column on large devices (≥992px, large desktops).`,
			}),
			form.NewField(form.FieldOptions{
				Name:  "width_xl",
				Label: "Width XL",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
				Help:  `The width of the column on extra large devices (≥1200px, extra large desktops).`,
			}),
			form.NewField(form.FieldOptions{
				Name:  "width_xxl",
				Label: "Width XXL",
				Type:  form.FORM_FIELD_TYPE_NUMBER,
				Help:  `The width of the column on extra extra large devices (≥1400px, extra extra large desktops).`,
			}),
		},
		Wrapper: func(block ui.BlockInterface) *hb.Tag {
			widthXs := block.Parameter("width_xs")
			widthSm := block.Parameter("width_sm")
			widthMd := block.Parameter("width_md")
			widthLg := block.Parameter("width_lg")
			widthXl := block.Parameter("width_xl")
			widthXxl := block.Parameter("width_xxl")

			if widthXs == "" && widthSm == "" && widthMd == "" && widthLg == "" && widthXl == "" && widthXxl == "" {
				widthSm = "12"
			}

			return hb.Div().ClassIf(widthXs != "", "col-xs-"+widthXs).
				ClassIf(widthSm != "", "col-sm-"+widthSm).
				ClassIf(widthMd != "", "col-md-"+widthMd).
				ClassIf(widthLg != "", "col-lg-"+widthLg).
				ClassIf(widthXl != "", "col-xl-"+widthXl).
				ClassIf(widthXxl != "", "col-xxl-"+widthXxl)
		},
	}
}

func blockContainerDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-layout-text-window`),
		Type:          "container",
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Text",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
	}
}

func blockDivDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-layout-text-window`),
		Type:          TYPE_DIV,
		AllowChildren: true,
		Fields:        []form.FieldInterface{},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			tag := hb.Div()
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockHeadingDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-type-h1"),
		Type:          "heading",
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
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
			}),
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			level := block.Parameter("level")
			content := block.Parameter("content")

			if level == "" {
				level = "1"
			}

			tag := hb.NewTag("h"+level).
				HTMLIf(content != "", content)

			applyAllParameters(block, tag)

			return tag
		},
	}
}

func blockHyperlinkDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-link-45deg"),
		Type:          "hyperlink",
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "url",
				Label: "URL",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Text (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
			form.NewField(form.FieldOptions{
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
			}),
		},
		Wrapper: func(block ui.BlockInterface) *hb.Tag {
			return hb.Div().Style(`display: inline-block`)
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

			tag := hb.A().
				Href(url).
				HTMLIf(content != "", content).
				Target(target)

			applyAllParameters(block, tag)

			return tag
		},
	}
}

func blockIconDefinition() blockeditor.BlockDefinition {
	icons := bootstrapIconList()

	rowSearch := hb.Div().
		Class("row").
		Child(hb.Div().
			Class("col-12 mb-3").
			Child(hb.Input().
				Class("form-control").
				Type("text").
				Placeholder("Search...").
				OnKeyUp("iconSearch(this.value)")))
	rowIcons := hb.Div().
		Class("row g-3").
		Children(lo.Map(icons, func(icon struct {
			Name string
			Icon string
		}, _ int) hb.TagInterface {
			return hb.Div().
				Class("col-xs-6 col-sm-4 col-md-3 col-lg-2").
				Child(hb.Button().
					Type("button").
					Title(icon.Name).
					Class("btn btn-light w-100 h-100 p-3 ButtonIcon").
					OnClick("iconSelected('" + icon.Icon + "')").
					Child(hb.I().
						Class(icon.Icon).
						Style(`font-size: 36px;`)).
					Child(hb.Div().Text(icon.Name).
						Style("font-size: 10px;")))
		}))
	sectionIcons := hb.Section().
		Child(hb.Div().
			Child(hb.Span().
				Text("Browse Icons").
				Style("cursor: pointer;").
				Style("text-decoration: underline;").
				OnClick(`document.getElementById('ListIcons').style.display = 'block'`))).
		Child(hb.Div().
			ID("ListIcons").
			Style(`display: none;`).
			// Style("max-height: 300px; overflow-y: scroll;").
			Style("background: aquablue;border: 1px solid lightblue;padding: 10px;").
			Child(rowSearch).
			Child(rowIcons))

	script := hb.NewScript(`
function iconSelected(icon) {
	document.getElementById("icon").value = icon;
	document.getElementById('ListIcons').style.display = 'none';
}
function iconSearch(search) {
   const buttonIcons=document.getElementsByClassName("ButtonIcon");
   for (let i = 0; i < buttonIcons.length; i++) {
    	const title = buttonIcons[i].title.toLowerCase();
	  	if(title.includes(search.toLowerCase())) {
	  		buttonIcons[i].parentElement.style.display = "block";
		} else {
		 	buttonIcons[i].parentElement.style.display = "none";
		}
   }
}
	`)

	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-brilliance"),
		Type: TYPE_ICON,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				ID:        "icon",
				Name:      "icon",
				Label:     "Icon",
				Type:      form.FORM_FIELD_TYPE_STRING,
				Invisible: true,
			}),
			form.NewField(form.FieldOptions{
				Type:  form.FORM_FIELD_TYPE_RAW,
				Value: sectionIcons.ToHTML() + script.ToHTML(),
			}),
		},
		Wrapper: func(block ui.BlockInterface) *hb.Tag {
			return hb.Div().Style(`display: inline-block`)
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			icon := block.Parameter("icon")

			if icon == "" {
				icon = "bi bi-brilliance"
			}

			tag := hb.I().Class(icon)
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockImageDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class("bi bi-image"),
		Type: "image",
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "image_url",
				Label: "Image URL (may be base64)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			imageUrl := block.Parameter("image_url")

			if imageUrl == "" {
				imageUrl = utils.PicsumURL(100, 100, utils.PicsumURLOptions{
					Grayscale: true,
					Seed:      "no image",
				})
			}

			tag := hb.Img(imageUrl)
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockListItemDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-list"),
		Type:          "list_item",
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
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

func blockOrderedListDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-list-ol"),
		Type:          "ordered_list",
		Fields:        []form.FieldInterface{},
		AllowChildren: true,
		AllowedChildTypes: []string{
			"list_item",
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			tag := hb.Div()
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockParagraphDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-paragraph"),
		Type:          "paragraph",
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Content (HTML)",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")

			if content == "" {
				content = "Add paragraph text"
			}

			tag := hb.P().
				HTML(content)

			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockRawHTMLDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon: hb.I().Class(`bi bi-code-slash`),
		Type: "raw_html",
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "HTML",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
	}
}

func blockRowDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-distribute-vertical`),
		Type:          TYPE_ROW,
		AllowChildren: true,
		Fields:        []form.FieldInterface{},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			tag := hb.Div().Class("row")
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockSectionDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-distribute-vertical`),
		Type:          TYPE_SECTION,
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Text",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			tag := hb.Div().Class("section")
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockTextDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class(`bi bi-fonts`),
		Type:          TYPE_TEXT,
		AllowChildren: true,
		Fields: []form.FieldInterface{
			form.NewField(form.FieldOptions{
				Name:  "content",
				Label: "Text",
				Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			}),
		},
		Wrapper: func(block ui.BlockInterface) *hb.Tag {
			return hb.Div().Style(`display: inline-block`)
		},
		ToTag: func(block ui.BlockInterface) *hb.Tag {
			content := block.Parameter("content")

			if content == "" {
				content = "Add text..."
			}

			tag := hb.Span().HTMLIf(content != "", content)
			applyAllParameters(block, tag)
			return tag
		},
	}
}

func blockUnorderedListDefinition() blockeditor.BlockDefinition {
	return blockeditor.BlockDefinition{
		Icon:          hb.I().Class("bi bi-list-ul"),
		Type:          "unordered_list",
		AllowChildren: true,
		Fields:        []form.FieldInterface{},
		AllowedChildTypes: []string{
			"list_item",
		},
	}
}
