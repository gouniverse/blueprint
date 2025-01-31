package helpers

import (
	"regexp"
)

// escapedBracketOpening := strings.ReplaceAll(sh.bracketOpening, "[", "\\[")
// 	escapedBracketOpening = strings.ReplaceAll(escapedBracketOpening, "(", "\\(")
// 	escapedBracketClosing := strings.ReplaceAll(sh.bracketClosing, "]", "\\]")
// 	escapedBracketClosing = strings.ReplaceAll(escapedBracketClosing, ")", "\\)")
// 	attr := `(\s+[^` + escapedBracketClosing + `]+)?`
// 	start := escapedBracketOpening + shortcode + attr + escapedBracketClosing
// 	end := escapedBracketOpening + `/` + shortcode + escapedBracketClosing
// 	//content := `([^` + escapedBracketClosing + `]*)`
// 	content := `([\S\s]+?.*?|\s?)`

// Define BBCode to HTML replacements as a package-level variable
var content = `([\S\s]+?.*?|\s?)`
var replacements = map[string]string{
	// Paragraph
	`\[paragraph\]` + content + `\[/paragraph\]`: `<p>$1</p>`,
	`\[p\]` + content + `\[/p\]`:                 `<p>$1</p>`,

	// Headings
	`\[heading1\]` + content + `\[/heading1\]`: `<h1>$1</h1>`,
	`\[heading2\]` + content + `\[/heading2\]`: `<h2>$1</h2>`,
	`\[heading3\]` + content + `\[/heading3\]`: `<h3>$1</h3>`,
	`\[heading4\]` + content + `\[/heading4\]`: `<h4>$1</h4>`,
	`\[heading5\]` + content + `\[/heading5\]`: `<h5>$1</h5>`,
	`\[heading6\]` + content + `\[/heading6\]`: `<h6>$1</h6>`,
	`\[h1\]` + content + `\[/h1\]`:             `<h1>$1</h1>`,
	`\[h2\]` + content + `\[/h2\]`:             `<h2>$1</h2>`,
	`\[h3\]` + content + `\[/h3\]`:             `<h3>$1</h3>`,
	`\[h4\]` + content + `\[/h4\]`:             `<h4>$1</h4>`,
	`\[h5\]` + content + `\[/h5\]`:             `<h5>$1</h5>`,
	`\[h6\]` + content + `\[/h6\]`:             `<h6>$1</h6>`,

	// Preformatted, quote and code blocks
	`\[pre\]` + content + `\[/pre\]`:         `<pre>$1</pre>`,
	`\[quote\]` + content + `\[/quote\]`:     `<blockquote>$1</blockquote>`,
	`\[code\]` + content + `\[/code\]`:       `<code>$1</code>`,
	`\[code=(.*?)\]` + content + `\[/code\]`: `<code lang="$1">$2</code>`,

	// Lists
	`\[list\]` + content + `\[/list\]`:   `<ul>$1</ul>`,
	`\[list=1\]` + content + `\[/list\]`: `<ol>$1</ol>`,
	`\[item\]` + content + `\[/item\]`:   `<li>$1</li>`,
	`\\*` + content + `\\n`:              `<li>$1</li>`,

	// Formatting
	`\[bold\]` + content + `\[/bold\]`:           `<b>$1</b>`,
	`\[italic\]` + content + `\[/italic\]`:       `<i>$1</i>`,
	`\[underline\]` + content + `\[/underline\]`: `<u>$1</u>`,
	`\[strike\]` + content + `\[/strike\]`:       `<s>$1</s>`,
	`\[color=(.*?)\]` + content + `\[/color\]`:   `<span style="color:$1">$2</span>`,
	`\[size=(.*?)\]` + content + `\[/size\]`:     `<span style="font-size:$1">$2</span>`,

	// Links
	`\[url\]` + content + `\[/url\]`:       `<a href="$1">$1</a>`,
	`\[url=(.*?)\]` + content + `\[/url\]`: `<a href="$1">$2</a>`,

	// Images
	`\[img\]` + content + `\[/img\]`:       `<img src="$1" />`,
	`\[img=(.*?)\]` + content + `\[/img\]`: `<img src="$1" alt="$2" />`,

	// Section
	`\[section\]` + content + `\[/section\]`: `<section>$1</section>`,
	`\[div\]` + content + `\[/div\]`:         `<div>$1</div>`,

	// Divider
	`\[divider\]` + content + `\[/divider\]`: `<hr />`,
	`\[rule\]`:                               `<hr />`,
	`\[rule\]` + content + `\[/rule\]`:       `<hr />`,
	`\[hr\]` + content + `\[/hr\]`:           `<hr />`,
	`\[hr]`:                                  `<hr />`,
	`\[hr=(.*?)\]`:                           `<hr />`,

	// Misc
	`\[break\]`:                          `<br />`,
	`\[break\]` + content + `\[/break\]`: `<br />`,
	`\[br\]` + content + `\[/br\]`:       `<br />`,
	`\[br\]`:                             `<br />`,
	`\[email\]` + content + `\[/email\]`: `<a href="mailto:$1">$1</a>`,
}

func BbcodeToHtml(input string) string {
	// Perform replacements
	for bbcode, html := range replacements {
		re := regexp.MustCompile(bbcode)
		input = re.ReplaceAllString(input, html)
	}

	return input
}
