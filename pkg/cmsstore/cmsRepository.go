package cmsstore

import (
	"errors"
	"project/config"

	"github.com/gouniverse/cms"
	"github.com/gouniverse/entitystore"
	"github.com/samber/lo"
)

type CmsTemplateQueryOptions struct {
	ID     string
	Handle string
	// Status   string
	// StatusIn []string
	// PerPage    int
	// PageNumber int
	Offset    int
	Limit     int
	SortOrder string
	SortBy    string
	CountOnly bool
}

type CmsPageQueryOptions struct {
	ID     string
	Handle string
	// Status   string
	// StatusIn []string
	// PerPage    int
	// PageNumber int
	Offset    int
	Limit     int
	SortOrder string
	SortBy    string
	CountOnly bool
}

// const coursesCourseTableName = "snv_courses_course"
// const coursesCourseCategoryTableName = "snv_courses_category"
// const coursesCourseCategoryMapTableName = "snv_courses_course_category_map"

func NewCmsRepository() *CmsRepository {
	return &CmsRepository{}
}

type CmsRepository struct {
	Debug bool
}

func (r CmsRepository) PageFindByID(pageID string) (*CmsPage, error) {
	if pageID == "" {
		return nil, errors.New("page id is empty")
	}

	list, err := r.PageList(CmsPageQueryOptions{
		ID:    pageID,
		Limit: 1,
	})

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, nil
}

func (r CmsRepository) PageList(options CmsPageQueryOptions) ([]CmsPage, error) {
	entityList, errEntityList := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
		EntityType:   cms.ENTITY_TYPE_PAGE,
		ID:           options.ID,
		EntityHandle: options.Handle,
		Limit:        uint64(options.Limit),
		Offset:       uint64(options.Offset),
		CountOnly:    options.CountOnly,
	})

	if errEntityList != nil {
		return []CmsPage{}, errEntityList
	}

	list := []CmsPage{}
	var errMap error = nil

	lo.ForEach(entityList, func(entity entitystore.Entity, index int) {
		attrs, err := entity.GetAttributes()
		if err != nil {
			errMap = err
		}
		cmsPageMap := map[string]string{
			"id":         entity.ID(),
			"handle":     entity.Handle(),
			"created_at": entity.CreatedAt().String(),
			"updated_at": entity.UpdatedAt().String(),
		}

		lo.ForEach(attrs, func(attr entitystore.Attribute, index int) {
			cmsPageMap[attr.AttributeKey()] = attr.AttributeValue()
		})

		cmsPage := NewCmsPageFromExistingData(cmsPageMap)
		list = append(list, *cmsPage)
	})

	if errMap != nil {
		return []CmsPage{}, errEntityList
	}

	return list, nil
}

func (r CmsRepository) TemplateList(options CmsTemplateQueryOptions) ([]CmsTemplate, error) {
	entityList, errEntityList := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
		EntityType:   cms.ENTITY_TYPE_TEMPLATE,
		ID:           options.ID,
		EntityHandle: options.Handle,
		Limit:        uint64(options.Limit),
		Offset:       uint64(options.Offset),
		CountOnly:    options.CountOnly,
	})

	if errEntityList != nil {
		return []CmsTemplate{}, errEntityList
	}

	list := []CmsTemplate{}
	var errMap error = nil

	lo.ForEach(entityList, func(entity entitystore.Entity, index int) {
		attrs, err := entity.GetAttributes()
		if err != nil {
			errMap = err
		}
		cmsTemplateMap := map[string]string{
			"id":         entity.ID(),
			"handle":     entity.Handle(),
			"created_at": entity.CreatedAt().String(),
			"updated_at": entity.UpdatedAt().String(),
		}

		lo.ForEach(attrs, func(attr entitystore.Attribute, index int) {
			cmsTemplateMap[attr.AttributeKey()] = attr.AttributeValue()
		})

		cmsTemplate := NewCmsTemplateFromExistingData(cmsTemplateMap)
		list = append(list, *cmsTemplate)
	})

	if errMap != nil {
		return []CmsTemplate{}, errEntityList
	}

	return list, nil
}

func (r CmsRepository) TemplateFindByID(templateID string) (*CmsTemplate, error) {
	if templateID == "" {
		return nil, errors.New("template id is empty")
	}

	list, err := r.TemplateList(CmsTemplateQueryOptions{
		ID:    templateID,
		Limit: 1,
	})

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, nil
}
