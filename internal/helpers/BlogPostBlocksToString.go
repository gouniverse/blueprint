package helpers

import (
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type Block struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Sequence   int            `json:"sequence"`
	ParentID   string         `json:"ParentId"`
	Text       string         `json:"content"`
	Attributes map[string]any `json:"attributes"`
}

func BlogPostBlocksToString(blocksString string) string {
	blocksAny, err := utils.FromJSON(blocksString, []map[string]any{})

	if err != nil {
		return "Error parsing content. Please try again later."
	}

	blocksMap := maputils.MapAnyToArrayMapStringAny(blocksAny)

	html := ""
	for _, blockMap := range blocksMap {
		blockType := blockMap["Type"].(string)
		blockID := blockMap["Id"].(string)
		parentID := blockMap["ParentId"].(string)
		attributes := blockMap["Attributes"].(map[string]any)
		sequence := blockMap["Sequence"].(float64)
		sequenceInt, _ := utils.ToInt(sequence)

		block := Block{
			ID:         blockID,
			Type:       blockType,
			Sequence:   int(sequenceInt),
			ParentID:   parentID,
			Attributes: attributes,
		}

		html += processBlock(block)

	}

	return html
}

func processBlock(block Block) string {
	if block.Type == "code" || block.Type == "Code" {
		return blockEditorBlockCodeToHtml(block)
	} else if block.Type == "heading" || block.Type == "Heading" {
		return blockEditorBlockHeadingToHtml(block)
	} else if block.Type == "image" || block.Type == "Image" {
		return blockEditorBlockImageToHtml(block)
	} else if block.Type == "text" || block.Type == "Text" {
		return blockEditorBlockTextToHtml(block)
	} else if block.Type == "raw-html" || block.Type == "RawHtml" {
		return blockEditorBlockRawHtmlToHtml(block)
	}

	return "Block " + block.Type + " renderer does not exist"
}

func blockEditorBlockCodeToHtml(block Block) string {
	code := lo.ValueOr(block.Attributes, "Code", "").(string)
	language := lo.ValueOr(block.Attributes, "Language", "").(string)

	html := ``
	html += `<div class="card" style="margin-bottom:20px;">`
	html += `  <div class="card-header">Language: ` + language + `</div>`
	html += `  <div class="card-body"><pre><code>` + code + `</code></pre></div>`
	html += `</div>`
	return html
}

func blockEditorBlockHeadingToHtml(block Block) string {
	level := lo.ValueOr(block.Attributes, "Level", "1").(string)
	text := lo.ValueOr(block.Attributes, "Text", "").(string)
	levelInt, _ := utils.ToInt(level)
	levelStr := utils.ToString(levelInt)

	return `<h` + levelStr + ` style="margin-bottom:20px;margin-top:20px;">` + text + `</h` + levelStr + `>`
}

func blockEditorBlockImageToHtml(block Block) string {
	url := lo.ValueOr(block.Attributes, "Url", "").(string)
	return `<img src="` + url + `" class="img img-responsive img-thumbnail" />`
}

func blockEditorBlockTextToHtml(block Block) string {
	text := lo.ValueOr(block.Attributes, "Text", "").(string)
	return `<p>` + text + `</p>`
}

func blockEditorBlockRawHtmlToHtml(block Block) string {
	text := lo.ValueOr(block.Attributes, "Text", "").(string)
	return text
}
