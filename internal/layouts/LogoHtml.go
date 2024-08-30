package layouts

import "strings"

func LogoHTML() string {
	logo := `
<div style="display: inline-block; justify-content: space-between; align-items: center; width: fit-content; padding: 0px; border: 3px solid orange; background:orange; color: white; font-family: sans-serif; font-size: 20px; letter-spacing: 2px;">
  <span style="color: white; font-family: sans-serif; font-size: 20px; letter-spacing: 2px;">
    BLUE
  </span>
  <span style="background-color: white; color: orange; padding: 5px;">
    PRINT
  </span>
</div>
`

	return strings.TrimSpace(logo)
}
