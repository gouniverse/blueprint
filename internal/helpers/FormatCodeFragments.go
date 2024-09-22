package helpers

import "strings"

// formatCodeFragments formats code fragments in the given text.
//
// It replaces all occurrences of "[code]" with "<pre><code>" and
// all occurrences of "[/code]" with "</code></pre>" in the text.
// The modified text is then returned.
//
// Parameters:
//   - text: the text that contains code fragments
//
// Return:
//   - the modified text with formatted code fragments
func FormatCodeFragments(text string) string {
	text = strings.ReplaceAll(text, "[code]", `<pre class="code"><code>`)
	text = strings.ReplaceAll(text, "[/code]", `</code></pre>`)
	return text
}
