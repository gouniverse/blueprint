package cmsstore

import (
	"project/config"
	"testing"

	"github.com/gouniverse/cms"
)

func TestCmsRepositoryPageFindByID(t *testing.T) {
	repo := NewCmsRepository()
	repo.Debug = true

	config.TestsConfigureAndInitialize()
	// Initialize()

	page, errPageCreate := config.Cms.EntityStore.EntityCreateWithTypeAndAttributes(cms.ENTITY_TYPE_PAGE, map[string]string{
		"name":             "PageName",
		"status":           CMSPAGE_STATUS_ACTIVE,
		"content":          "PageContent",
		"content_editor":   "PageContentEditor",
		"meta_description": "PageMetaDescription",
		"meta_keywords":    "PageMetaKeywords",
		"meta_robots":      "PageMetaRobots",
		"canonical_url":    "PageCanonicalUrl",
		"title":            "PageTitle",
		"alias":            "PageAlias",
		"template_id":      "PageTemplateId",
	})

	if errPageCreate != nil {
		t.Error("unexpected error:", errPageCreate)
	}

	pageFound, errFind := repo.PageFindByID(page.ID())
	if errFind != nil {
		t.Error("unexpected error:", errFind)
	}
	if pageFound == nil {
		t.Error("Page MUST NOT be nil")
	}

	if pageFound.Alias() != "PageAlias" {
		t.Error("Page alias MUST BE 'PageAlias', found: ", pageFound.Alias())
	}

	if pageFound.Name() != "PageName" {
		t.Error("Page name MUST BE 'PageName', found: ", pageFound.Name())
	}

	if pageFound.Content() != "PageContent" {
		t.Error("Page content MUST BE 'PageContent', found: ", pageFound.Content())
	}

	if pageFound.Status() != CMSTEMPLATE_STATUS_ACTIVE {
		t.Error("Page status MUST BE 'Unpublished', found: ", pageFound.Status())
	}

	if pageFound.CreatedAt() == "" {
		t.Error("Page created MUST NOT BE empty, found: ", pageFound.CreatedAt())
	}

	if pageFound.UpdatedAt() == "" {
		t.Error("Page updated MUST NOT BE empty, found: ", pageFound.UpdatedAt())
	}

	if pageFound.MetaDescription() != "PageMetaDescription" {
		t.Error("Page meta description MUST BE 'PageMetaDescription', found: ", pageFound.MetaDescription())
	}

	if pageFound.MetaKeywords() != "PageMetaKeywords" {
		t.Error("Page meta keywords MUST BE 'PageMetaDescription', found: ", pageFound.MetaKeywords())
	}

	if pageFound.MetaRobots() != "PageMetaRobots" {
		t.Error("Page meta robots MUST BE 'PageMetaRobots', found: ", pageFound.MetaRobots())
	}
	if pageFound.CanonicalUrl() != "PageCanonicalUrl" {
		t.Error("Page canonical URL MUST BE 'PageCanonicalUrl', found: ", pageFound.CanonicalUrl())
	}
	if pageFound.TemplateID() != "PageTemplateId" {
		t.Error("Page template id MUST BE 'PageTemplateId', found: ", pageFound.TemplateID())
	}
	if pageFound.Title() != "PageTitle" {
		t.Error("Page name MUST BE 'PageTitle', found: ", pageFound.Title())
	}
}

func TestCmsRepositoryTemplateFindByID(t *testing.T) {
	repo := NewCmsRepository()
	repo.Debug = true

	config.TestsConfigureAndInitialize()
	// Initialize()

	template, errTemplateCreate := config.Cms.EntityStore.EntityCreateWithTypeAndAttributes(cms.ENTITY_TYPE_TEMPLATE, map[string]string{
		"name":    "TemplateName",
		"status":  CMSTEMPLATE_STATUS_ACTIVE,
		"content": "TemplateContent",
	})

	if errTemplateCreate != nil {
		t.Error("unexpected error:", errTemplateCreate)
	}

	templateFound, errFind := repo.TemplateFindByID(template.ID())
	if errFind != nil {
		t.Error("unexpected error:", errFind)
	}
	if templateFound == nil {
		t.Error("Template MUST NOT be nil")
	}

	if templateFound.Name() != "TemplateName" {
		t.Error("Template name MUST BE 'TemplateName', found: ", templateFound.Name())
	}

	if templateFound.Content() != "TemplateContent" {
		t.Error("Template content MUST BE 'TemplateContent', found: ", templateFound.Content())
	}

	if templateFound.Status() != CMSTEMPLATE_STATUS_ACTIVE {
		t.Error("Template status MUST BE 'Unpublished', found: ", templateFound.Status())
	}

	if templateFound.CreatedAt() == "" {
		t.Error("Template created MUST NOT BE empty, found: ", templateFound.CreatedAt())
	}

	if templateFound.UpdatedAt() == "" {
		t.Error("Template updated MUST NOT BE empty, found: ", templateFound.UpdatedAt())
	}
}
