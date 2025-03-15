package webtheme

import (
	"github.com/gouniverse/blockeditor"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/samber/lo"
)

// == TYPE ====================================================================
type theme struct {
	blocks          []ui.BlockInterface
	tree            *blockeditor.FlatTree
	availableBlocks []struct {
		Type  string
		ToTag func(block ui.BlockInterface) *hb.Tag
	}
}

// == CONSTRUCTOR =============================================================

func New(blocks []ui.BlockInterface) *theme {
	tree := blockeditor.NewFlatTree(blocks)

	t := &theme{
		blocks: blocks,
		tree:   tree,
	}

	t.addBlockRenderer(TYPE_BREADCRUMBS, t.breadcrumbsToHtml)
	t.addBlockRenderer(TYPE_COLUMN, t.columnToHtml)
	t.addBlockRenderer(TYPE_CONTAINER, t.containerToHtml)
	t.addBlockRenderer(TYPE_DIV, t.divToHtml)
	t.addBlockRenderer(TYPE_HEADING, t.headingToHtml)
	t.addBlockRenderer(TYPE_HYPERLINK, t.hyperlinkToHtml)
	t.addBlockRenderer(TYPE_IMAGE, t.imageToHtml)
	t.addBlockRenderer(TYPE_ICON, t.iconToHtml)
	t.addBlockRenderer(TYPE_LIST_ITEM, t.liToHtml)
	t.addBlockRenderer(TYPE_PARAGRAPH, t.paragraphToHtml)
	t.addBlockRenderer(TYPE_RAW_HTML, t.rawToHtml)
	t.addBlockRenderer(TYPE_ROW, t.rowToHtml)
	t.addBlockRenderer(TYPE_ORDERED_LIST, t.olToHtml)
	t.addBlockRenderer(TYPE_TEXT, t.textToHtml)
	t.addBlockRenderer(TYPE_SECTION, t.sectionToHtml)
	t.addBlockRenderer(TYPE_UNORDERED_LIST, t.ulToHtml)

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
	// 	style := `
	// .BlogTitle {
	// 	font-family: Roboto, sans-serif;
	// }
	// .BlogContent {
	// 	font-family: Roboto, sans-serif;
	// }
	// h1 {
	// 	margin-bottom: 20px;
	// 	font-size: 48px;
	// }
	// h2 {
	// 	margin-bottom: 20px;
	// 	font-size: 36px;
	// }
	// h3 {
	// 	margin-bottom: 20px;
	// 	font-size: 24px;
	// }
	// h4 {
	// 	margin-bottom: 20px;
	// 	font-size: 18px;
	// }
	// h5 {
	// 	margin-bottom: 20px;
	// 	font-size: 16px;
	// }
	// h6 {
	// 	margin-bottom: 20px;
	// 	font-size: 14px;
	// }
	// 	`
	return ``
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

	return blockTag
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
