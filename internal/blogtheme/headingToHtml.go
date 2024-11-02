package blogtheme

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
)

func (t *theme) headingToHtml(block ui.BlockInterface) *hb.Tag {
	level := block.Parameter("level")
	if level == "" {
		level = "1"
	}

	text := block.Parameter("content")

	levelInt, _ := utils.ToInt(level)
	levelStr := utils.ToString(levelInt)

	return hb.NewTag(`h` + levelStr).
		Style("margin-bottom:20px;margin-top:20px;").
		HTML(text)
}
