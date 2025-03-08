package testutils

import (
	"context"
	"project/config"

	"github.com/gouniverse/cmsstore"
)

func SeedTemplate(siteID, templateID string) (err error) {
	templateContent := `
	<html>
	    <head>
			<title>[[PageTitle]]</title>
		</head>
		<body>
			[[PageContent]]
		</body>
	</html>
	`

	template := cmsstore.NewTemplate().
		SetID(templateID).
		SetSiteID(siteID).
		SetStatus(cmsstore.TEMPLATE_STATUS_ACTIVE).
		SetName(templateID).
		SetContent(templateContent)

	return config.CmsStore.TemplateCreate(context.Background(), template)
}
