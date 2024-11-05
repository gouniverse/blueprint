package webtheme

import (
	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/samber/lo"
)

type theme struct {
	blocks          []ui.BlockInterface
	tree            *blockeditor.FlatTree
	availableBlocks []struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}
}

func New(blocks []ui.BlockInterface) *theme {
	tree := blockeditor.NewFlatTree(blocks)

	t := &theme{
		blocks: blocks,
		tree:   tree,
	}

	t.addBlockRenderer("breadcrumbs", t.breadcrumbsToHtml)
	t.addBlockRenderer("heading", t.headingToHtml)
	t.addBlockRenderer("hyperlink", t.hyperlinkToHtml)
	t.addBlockRenderer("image", t.imageToHtml)
	t.addBlockRenderer("paragraph", t.paragraphToHtml)
	t.addBlockRenderer("raw_html", t.rawToHtml)
	t.addBlockRenderer("unordered_list", t.ulToHtml)
	t.addBlockRenderer("list_item", t.liToHtml)
	t.addBlockRenderer("ordered_list", t.olToHtml)
	t.addBlockRenderer("section", t.sectionToHtml)
	t.addBlockRenderer("container", t.containerToHtml)

	return t
}

// func New(blocksJSON string) (*theme, error) {
// 	blocks, err := ui.BlocksFromJson(blocksJSON)

// 	if err != nil {
// 		return nil, err
// 	}

// 	tree := blockeditor.NewFlatTree(blocks)

// 	t := &theme{
// 		blocks: blocks,
// 		tree:   tree,
// 	}

// 	t.addBlockRenderer("heading", t.headingToHtml)
// 	t.addBlockRenderer("hyperlink", t.hyperlinkToHtml)
// 	t.addBlockRenderer("image", t.imageToHtml)
// 	t.addBlockRenderer("paragraph", t.paragraphToHtml)
// 	t.addBlockRenderer("raw", t.rawToHtml)
// 	t.addBlockRenderer("unordered_list", t.ulToHtml)
// 	t.addBlockRenderer("list_item", t.liToHtml)
// 	t.addBlockRenderer("ordered_list", t.olToHtml)

// 	return t, nil
// }

func (t *theme) Style() string {
	style := `
.BlogTitle {
	font-family: Roboto, sans-serif;
}
.BlogContent {
	font-family: Roboto, sans-serif;
}
h1 { 
	margin-bottom: 20px;
	font-size: 48px;
}
h2 { 
	margin-bottom: 20px;
	font-size: 36px;
}
h3 {
	margin-bottom: 20px;
	font-size: 24px;
}
h4 {
	margin-bottom: 20px;
	font-size: 18px;
}
h5 {
	margin-bottom: 20px;
	font-size: 16px;
}
h6 {
	margin-bottom: 20px;
	font-size: 14px;
}
	`
	return style
}

func (t *theme) addBlockRenderer(blockType string, toTag func(block ui.BlockInterface) *hb.Tag) {
	t.availableBlocks = append(t.availableBlocks, struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}{
		Type:  blockType,
		ToTag: toTag,
	})
}

func (t *theme) renderBlock(block ui.BlockInterface) *hb.Tag {
	status := block.Parameter("status")

	if status != "published" {
		return nil
	}

	childrenTags := lo.Map(block.Children(), func(block ui.BlockInterface, _ int) hb.TagInterface {
		return t.renderBlock(block)
	})

	blockTag := t.blockToTag(block).Children(childrenTags)

	t.applyCommonParameters(block, blockTag)
	t.applyMarginParameters(block, blockTag)
	t.applyPaddingParameters(block, blockTag)
	t.applyBackgroundParameters(block, blockTag)
	t.applyAlignmentParameters(block, blockTag)
	t.applyFontParameters(block, blockTag)

	return blockTag
}

func (t *theme) applyFontParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	fontFamily := block.Parameter("font_family")

	if fontFamily != "" {
		blockTag.Style("font-family:" + fontFamily)
	}

	fontSize := block.Parameter("font_size")

	if fontSize != "" {
		blockTag.Style("font-size:" + fontSize)
	}

	fontWeight := block.Parameter("font_weight")

	if fontWeight != "" {
		blockTag.Style("font-weight:" + fontWeight)
	}
}

func (t *theme) applyAlignmentParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	textAlign := block.Parameter("text_align")

	if textAlign != "" {
		blockTag.Style("text-align:" + textAlign)
	}

	verticalAlign := block.Parameter("vertical_align")

	if verticalAlign != "" {
		blockTag.Style("vertical-align:" + verticalAlign)
	}
}

func (t *theme) applyBackgroundParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	backgroundColor := block.Parameter("background_color")

	if backgroundColor != "" {
		blockTag.Style("background-color:" + backgroundColor)
	}

	backgroundImage := block.Parameter("background_image")

	if backgroundImage != "" {
		blockTag.Style("background-image:url(" + backgroundImage + ")")
	}

	backgroundPosition := block.Parameter("background_position")

	if backgroundPosition != "" {
		blockTag.Style("background-position:" + backgroundPosition)
	}

	backgroundRepeat := block.Parameter("background_repeat")

	if backgroundRepeat != "" {
		blockTag.Style("background-repeat:" + backgroundRepeat)
	}

	backgroundSize := block.Parameter("background_size")

	if backgroundSize != "" {
		blockTag.Style("background-size:" + backgroundSize)
	}

	backgroundAttachment := block.Parameter("background_attachment")

	if backgroundAttachment != "" {
		blockTag.Style("background-attachment:" + backgroundAttachment)
	}
}

func (t *theme) applyCommonParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	id := block.Parameter("html_id")

	if id != "" {
		blockTag.ID(id)
	}

	class := block.Parameter("html_class")

	if class != "" {
		blockTag.Class(class)
	}

	style := block.Parameter("html_style")

	if style != "" {
		blockTag.Style(style)
	}

	color := block.Parameter("text_color")

	if color != "" {
		blockTag.Style("color:" + color)
	}
}

func (t *theme) applyMarginParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	marginTop := block.Parameter("margin_top")

	if marginTop != "" {
		blockTag.Style("margin-top:" + marginTop)
	}

	marginBottom := block.Parameter("margin_bottom")

	if marginBottom != "" {
		blockTag.Style("margin-bottom:" + marginBottom)
	}

	marginLeft := block.Parameter("margin_left")

	if marginLeft != "" {
		blockTag.Style("margin-left:" + marginLeft)
	}

	marginRight := block.Parameter("margin_right")

	if marginRight != "" {
		blockTag.Style("margin-right:" + marginRight)
	}
}

func (t *theme) applyPaddingParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	paddingTop := block.Parameter("padding_top")

	if paddingTop != "" {
		blockTag.Style("padding-top:" + paddingTop)
	}

	paddingBottom := block.Parameter("padding_bottom")

	if paddingBottom != "" {
		blockTag.Style("padding-bottom:" + paddingBottom)
	}

	paddingLeft := block.Parameter("padding_left")

	if paddingLeft != "" {
		blockTag.Style("padding-left:" + paddingLeft)
	}

	paddingRight := block.Parameter("padding_right")

	if paddingRight != "" {
		blockTag.Style("padding-right:" + paddingRight)
	}
}

func (t *theme) ToHtml() string {
	wrap := hb.Wrap()

	for _, block := range t.blocks {
		blockTag := t.renderBlock(block)
		wrap.Child(blockTag)
	}

	return wrap.ToHTML()
}

func (t *theme) isSupportedBlock(block ui.BlockInterface) bool {
	for _, availableBlock := range t.availableBlocks {
		if block.Type() == availableBlock.Type {
			return true
		}
	}

	return false
}

func (t *theme) blockToTag(block ui.BlockInterface) *hb.Tag {
	if !t.isSupportedBlock(block) {
		return hb.Div().
			Class("alert alert-warning").
			Text("Block " + block.Type() + " renderer does not exist")
	}

	renderer, found := lo.Find(t.availableBlocks, func(availableBlock struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}) bool {
		return availableBlock.Type == block.Type()
	})

	if !found {
		return hb.Div().
			Class("alert alert-warning").
			Text("Block " + block.Type() + " renderer does not exist")
	}

	return renderer.ToTag(block)
}

// type Block struct {
// 	ID         string         `json:"id"`
// 	Type       string         `json:"type"`
// 	Sequence   int            `json:"sequence"`
// 	ParentID   string         `json:"ParentId"`
// 	Text       string         `json:"content"`
// 	Attributes map[string]any `json:"attributes"`
// }

// func BlogPostBlocksToString(blocksString string) string {
// 	blocksAny, err := utils.FromJSON(blocksString, []map[string]any{})

// 	if err != nil {
// 		return "Error parsing content. Please try again later."
// 	}

// 	blocksMap := maputils.MapAnyToArrayMapStringAny(blocksAny)

// 	html := ""
// 	for _, blockMap := range blocksMap {
// 		blockType := blockMap["Type"].(string)
// 		blockID := blockMap["Id"].(string)
// 		parentID := blockMap["ParentId"].(string)
// 		attributes := blockMap["Attributes"].(map[string]any)
// 		sequence := blockMap["Sequence"].(float64)
// 		sequenceInt, _ := utils.ToInt(sequence)

// 		block := Block{
// 			ID:         blockID,
// 			Type:       blockType,
// 			Sequence:   int(sequenceInt),
// 			ParentID:   parentID,
// 			Attributes: attributes,
// 		}

// 		html += processBlock(block)

// 	}

// 	return html
// }

// func processBlock(block Block) string {
// 	if block.Type == "code" || block.Type == "Code" {
// 		return blockEditorBlockCodeToHtml(block)
// 	} else if block.Type == "heading" || block.Type == "Heading" {
// 		return blockEditorBlockHeadingToHtml(block)
// 	} else if block.Type == "image" || block.Type == "Image" {
// 		return blockEditorBlockImageToHtml(block)
// 	} else if block.Type == "text" || block.Type == "Text" {
// 		return blockEditorBlockTextToHtml(block)
// 	} else if block.Type == "raw-html" || block.Type == "RawHtml" {
// 		return blockEditorBlockRawHtmlToHtml(block)
// 	}

// 	return "Block " + block.Type + " renderer does not exist"
// }

// func blockEditorBlockCodeToHtml(block Block) string {
// 	code := lo.ValueOr(block.Attributes, "Code", "").(string)
// 	language := lo.ValueOr(block.Attributes, "Language", "").(string)

// 	html := ``
// 	html += `<div class="card" style="margin-bottom:20px;">`
// 	html += `  <div class="card-header">Language: ` + language + `</div>`
// 	html += `  <div class="card-body"><pre><code>` + code + `</code></pre></div>`
// 	html += `</div>`
// 	return html
// }

// func blockEditorBlockHeadingToHtml(block Block) string {
// 	level := lo.ValueOr(block.Attributes, "Level", "1").(string)
// 	text := lo.ValueOr(block.Attributes, "Text", "").(string)
// 	levelInt, _ := utils.ToInt(level)
// 	levelStr := utils.ToString(levelInt)

// 	return `<h` + levelStr + ` style="margin-bottom:20px;margin-top:20px;">` + text + `</h` + levelStr + `>`
// }

// func blockEditorBlockImageToHtml(block Block) string {
// 	url := lo.ValueOr(block.Attributes, "Url", "").(string)
// 	return `<img src="` + url + `" class="img img-responsive img-thumbnail" />`
// }

// func blockEditorBlockTextToHtml(block Block) string {
// 	text := lo.ValueOr(block.Attributes, "Text", "").(string)
// 	return `<p>` + text + `</p>`
// }

// func blockEditorBlockRawHtmlToHtml(block Block) string {
// 	text := lo.ValueOr(block.Attributes, "Text", "").(string)
// 	return text
// }
