package cmsstore

import (
	"github.com/gouniverse/dataobject"
)

// var _ dataobject.DataObjectFluentInterface = (*CmsTemplate)(nil) // verify it extends the data object interface

const CMSTEMPLATE_STATUS_ACTIVE = "active"
const CMSTEMPLATE_STATUS_INACTIVE = "inactive"
const CMSTEMPLATE_STATUS_DELETED = "deleted"

// Unused
// func NewCmsTemplate() *CmsTemplate {
// 	o := &CmsTemplate{}
// 	o.SetID(uid.HumanUid()).
// 		SetContent("").
// 		SetName("")

// 	return o
// }

func NewCmsTemplateFromExistingData(data map[string]string) *CmsTemplate {
	o := &CmsTemplate{}
	o.Hydrate(data)
	return o
}

type CmsTemplate struct {
	dataobject.DataObject
}

func (o *CmsTemplate) CreatedAt() string {
	return o.Get("created_at")
}

func (o *CmsTemplate) Content() string {
	return o.Get("content")
}

func (o *CmsTemplate) ID() string {
	return o.Get("id")
}

func (o *CmsTemplate) Handle() string {
	return o.Get("handle")
}

func (o *CmsTemplate) Status() string {
	return o.Get("status")
}

func (o *CmsTemplate) Name() string {
	return o.Get("name")
}

func (o *CmsTemplate) UpdatedAt() string {
	return o.Get("updated_at")
}

// func (o *CmsTemplate) SetCreatedAt(createdAt string) *CmsTemplate {
// 	o.Set("created_at", createdAt)
// 	return o
// }

// func (o *CmsTemplate) SetContent(content string) *CmsTemplate {
// 	o.Set("content", content)
// 	return o
// }

// func (o *CmsTemplate) SetHandle(handle string) *CmsTemplate {
// 	o.Set("handle", handle)
// 	return o
// }

// func (o *CmsTemplate) SetName(name string) *CmsTemplate {
// 	o.Set("name", name)
// 	return o
// }

// func (o *CmsTemplate) SetID(id string) *CmsTemplate {
// 	o.Set("id", id)
// 	return o
// }

// func (o *CmsTemplate) SetStatus(status string) *CmsTemplate {
// 	o.Set("status", status)
// 	return o
// }

// func (o *CmsTemplate) SetUpdatedAt(updatedAt string) *CmsTemplate {
// 	o.Set("updated_at", updatedAt)
// 	return o
// }
