package cmsstore

import (
	"github.com/gouniverse/dataobject"
)

// var _ dataobject.DataObjectFluentInterface = (*CmsPage)(nil) // verify it extends the data object interface

const CMSPAGE_STATUS_ACTIVE = "active"
const CMSPAGE_STATUS_INACTIVE = "inactive"
const CMSPAGE_STATUS_DELETED = "deleted"

// Unused
// func NewCmsPage() *CmsPage {
// 	o := &CmsPage{}
// 	o.SetID(uid.HumanUid()).
// 		SetContent("").
// 		SetName("")

// 	return o
// }

func NewCmsPageFromExistingData(data map[string]string) *CmsPage {
	o := &CmsPage{}
	o.Hydrate(data)
	return o
}

type CmsPage struct {
	dataobject.DataObject
}

func (o *CmsPage) Alias() string {
	return o.Get("alias")
}

func (o *CmsPage) CreatedAt() string {
	return o.Get("created_at")
}

func (o *CmsPage) CanonicalUrl() string {
	return o.Get("canonical_url")
}

func (o *CmsPage) Content() string {
	return o.Get("content")
}

func (o *CmsPage) ID() string {
	return o.Get("id")
}

func (o *CmsPage) Handle() string {
	return o.Get("handle")
}

func (o *CmsPage) MetaDescription() string {
	return o.Get("meta_description")
}

func (o *CmsPage) MetaKeywords() string {
	return o.Get("meta_keywords")
}

func (o *CmsPage) MetaRobots() string {
	return o.Get("meta_robots")
}

func (o *CmsPage) Name() string {
	return o.Get("name")
}

func (o *CmsPage) Status() string {
	return o.Get("status")
}

func (o *CmsPage) Title() string {
	return o.Get("title")
}

func (o *CmsPage) TemplateID() string {
	return o.Get("template_id")
}

func (o *CmsPage) UpdatedAt() string {
	return o.Get("updated_at")
}

// func (o *CmsPage) SetCreatedAt(createdAt string) *CmsPage {
// 	o.Set("created_at", createdAt)
// 	return o
// }

// func (o *CmsPage) SetContent(content string) *CmsPage {
// 	o.Set("content", content)
// 	return o
// }

// func (o *CmsPage) SetHandle(handle string) *CmsPage {
// 	o.Set("handle", handle)
// 	return o
// }

// func (o *CmsPage) SetName(name string) *CmsPage {
// 	o.Set("name", name)
// 	return o
// }

// func (o *CmsPage) SetID(id string) *CmsPage {
// 	o.Set("id", id)
// 	return o
// }

// func (o *CmsPage) SetStatus(status string) *CmsPage {
// 	o.Set("status", status)
// 	return o
// }

// func (o *CmsPage) SetUpdatedAt(updatedAt string) *CmsPage {
// 	o.Set("updated_at", updatedAt)
// 	return o
// }
