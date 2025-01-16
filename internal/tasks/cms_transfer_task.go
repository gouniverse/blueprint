package tasks

import (
	"context"
	"errors"
	"project/config"
	"strings"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/cms"
	"github.com/gouniverse/cmsstore"
	"github.com/gouniverse/entitystore"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/taskstore"
)

// ============================================================================
// cmsTransferTask
// ============================================================================
// Transfers the old CMS data to the new CMS Store
// ============================================================================
// Example:
// - go run . task CmsTransfer --index=all --truncate=yes
// ============================================================================
type cmsTransferTask struct {
	taskstore.TaskHandlerBase

	// allowed indexes
	allowedIndexes []string

	// index to rebuild
	index string

	// truncate the index table
	truncate bool
}

var _ taskstore.TaskHandlerInterface = (*cmsTransferTask)(nil) // verify it extends the task interface

// == CONSTRUCTOR =============================================================

func NewCmsTransferTask() *cmsTransferTask {
	return &cmsTransferTask{}
}

// == IMPLEMENTATION ==========================================================

func (task *cmsTransferTask) Alias() string {
	return "CmsTransfer"
}

func (task *cmsTransferTask) Title() string {
	return "CMS Transfer (old to new)"
}

func (task *cmsTransferTask) Description() string {
	return "Transfers the old CMS data to the new CMS Store"
}

func (task *cmsTransferTask) Enqueue(index string) (queuedTask taskstore.QueueInterface, err error) {
	if config.TaskStore == nil {
		return nil, errors.New("task store is nil")
	}

	return config.TaskStore.TaskEnqueueByAlias(task.Alias(), map[string]any{})
}

func (task *cmsTransferTask) Handle() bool {
	if task.checkAndEnqueueTask() {
		return true
	}

	if config.CmsStore == nil {
		task.LogError("CMS Store not configured.")
		return false
	}

	if task.cmsTransfer() {
		task.LogSuccess("Task completed successfully.")
	} else {
		task.LogError("Task failed.")
	}

	return true
}

func (task *cmsTransferTask) cmsTransfer() bool {
	task.LogInfo("Transferring CMS data...")

	task.LogInfo(" - Finding main site...")

	if config.CmsStore == nil {
		task.LogError("CMS Store not configured.")
		return false
	}

	sites, err := config.CmsStore.SiteList(context.Background(), cmsstore.SiteQuery().
		SetStatus(cmsstore.SITE_STATUS_ACTIVE).
		SetOrderBy(cmsstore.COLUMN_UPDATED_AT).
		SetSortOrder(sb.DESC).
		SetLimit(1))

	if err != nil {
		task.LogError("Error finding main site: " + err.Error())
		return false
	}

	if len(sites) == 0 {
		task.LogError("No main site found.")
		return false
	}

	site := sites[0]

	domainNames, err := site.DomainNames()

	if err != nil {
		task.LogError("Error getting domain names: " + err.Error())
		return false
	}

	task.LogInfo(" - Site found: " + site.ID() + " (" + strings.Join(domainNames, ", ") + ")")

	if !task.transferTemplates(site) {
		task.LogError("Failed to transfer templates.")
		return false
	}

	if !task.transferPages(site) {
		task.LogError("Failed to transfer pages.")
		return false
	}

	if !task.transferBlocks(site) {
		task.LogError("Failed to transfer blocks.")
		return false
	}

	return true
}

func (task *cmsTransferTask) transferTemplates(site cmsstore.SiteInterface) bool {
	task.LogInfo("Transferring templates...")

	entities, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
		EntityType: cms.ENTITY_TYPE_TEMPLATE,
	})

	if err != nil {
		task.LogError("Error retrieving entities: " + err.Error())
		return false
	}

	for _, entity := range entities {
		if !task.transferTemplate(site, entity) {
			task.LogError("- Failed to transfer entity: " + entity.ID() + ". Aborted.")
			return false
		}
	}

	return true
}

func (task *cmsTransferTask) transferTemplate(site cmsstore.SiteInterface, entity entitystore.Entity) bool {
	task.LogInfo("Transferring entity: " + entity.ID())

	if config.CmsStore == nil {
		task.LogError("CMS Store not configured.")
		return false
	}

	template, err := config.CmsStore.TemplateFindByID(context.Background(), entity.ID())

	if err != nil {
		task.LogError("Error retrieving template: " + err.Error())
		return false
	}

	if template != nil {
		task.LogInfo(" - Template found: " + template.ID())
		return true
	}

	content, err := entity.GetString("content", "")

	if err != nil {
		task.LogError("Error retrieving content: " + err.Error())
		return false
	}

	name, err := entity.GetString("name", "")

	if err != nil {
		task.LogError("Error retrieving title: " + err.Error())
		return false
	}

	status, err := entity.GetString("status", "")

	if err != nil {
		task.LogError("Error retrieving status: " + err.Error())
		return false
	}

	template = cmsstore.NewTemplate().
		SetID(entity.ID()).
		SetName(name).
		SetContent(content).
		SetStatus(status).
		SetSiteID(site.ID()).
		SetCreatedAt(carbon.CreateFromStdTime(entity.CreatedAt(), carbon.UTC).ToDateTimeString(carbon.UTC))

	err = config.CmsStore.TemplateCreate(context.Background(), template)

	if err != nil {
		task.LogError("Error creating template: " + err.Error())
		return false
	}

	return true
}

func (task *cmsTransferTask) transferPages(site cmsstore.SiteInterface) bool {
	task.LogInfo("Transferring pages...")

	entities, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
		EntityType: cms.ENTITY_TYPE_PAGE,
	})

	if err != nil {
		task.LogError("Error retrieving entities: " + err.Error())
		return false
	}

	for _, entity := range entities {
		if !task.transferPage(site, entity) {
			task.LogError("- Failed to transfer entity: " + entity.ID() + ". Aborted.")
			return false
		}
	}

	return true
}

func (task *cmsTransferTask) transferPage(site cmsstore.SiteInterface, entity entitystore.Entity) bool {
	task.LogInfo("Transferring entity: " + entity.ID())

	if config.CmsStore == nil {
		task.LogError("CMS Store not configured.")
		return false
	}

	page, err := config.CmsStore.PageFindByID(context.Background(), entity.ID())

	if err != nil {
		task.LogError("Error retrieving page: " + err.Error())
		return false
	}

	if page != nil {
		task.LogInfo(" - Page found: " + page.ID())
		return true
	}

	alias, err := entity.GetString("alias", "")

	if err != nil {
		task.LogError("Error retrieving alias: " + err.Error())
		return false
	}

	canonicalURL, err := entity.GetString("canonical_url", "")

	if err != nil {
		task.LogError("Error retrieving canonical_url: " + err.Error())
		return false
	}

	content, err := entity.GetString("content", "")

	if err != nil {
		task.LogError("Error retrieving content: " + err.Error())
		return false
	}

	handle, err := entity.GetString("handle", "")

	if err != nil {
		task.LogError("Error retrieving handle: " + err.Error())
		return false
	}

	metaDescription, err := entity.GetString("meta_description", "")

	if err != nil {
		task.LogError("Error retrieving meta_description: " + err.Error())
		return false
	}

	metaKeywords, err := entity.GetString("meta_keywords", "")

	if err != nil {
		task.LogError("Error retrieving meta_keywords: " + err.Error())
		return false
	}

	metaRobots, err := entity.GetString("meta_robots", "")

	if err != nil {
		task.LogError("Error retrieving meta_robots: " + err.Error())
		return false
	}

	name, err := entity.GetString("name", "")

	if err != nil {
		task.LogError("Error retrieving title: " + err.Error())
		return false
	}

	status, err := entity.GetString("status", "")

	if err != nil {
		task.LogError("Error retrieving status: " + err.Error())
		return false
	}

	templateID, err := entity.GetString("template_id", "")

	if err != nil {
		task.LogError("Error retrieving template_id: " + err.Error())
		return false
	}

	title, err := entity.GetString("title", "")

	if err != nil {
		task.LogError("Error retrieving title: " + err.Error())
		return false
	}

	page = cmsstore.NewPage().
		SetAlias(alias).
		SetCanonicalUrl(canonicalURL).
		SetMetaDescription(metaDescription).
		SetMetaKeywords(metaKeywords).
		SetMetaRobots(metaRobots).
		SetHandle(handle).
		SetID(entity.ID()).
		SetName(name).
		SetContent(content).
		SetStatus(status).
		SetSiteID(site.ID()).
		SetTemplateID(templateID).
		SetTitle(title).
		SetCreatedAt(carbon.CreateFromStdTime(entity.CreatedAt()).ToDateTimeString(carbon.UTC))

	err = config.CmsStore.PageCreate(context.Background(), page)

	if err != nil {
		task.LogError("Error creating template: " + err.Error())
		return false
	}

	return true
}

func (task *cmsTransferTask) transferBlocks(site cmsstore.SiteInterface) bool {
	task.LogInfo("Transferring blocks...")

	entities, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
		EntityType: cms.ENTITY_TYPE_BLOCK,
	})

	if err != nil {
		task.LogError("Error retrieving entities: " + err.Error())
		return false
	}

	for _, entity := range entities {
		if !task.transferBlock(site, entity) {
			task.LogError("- Failed to transfer entity: " + entity.ID() + ". Aborted.")
			return false
		}
	}

	return true
}

func (task *cmsTransferTask) transferBlock(site cmsstore.SiteInterface, entity entitystore.Entity) bool {
	task.LogInfo("Transferring entity: " + entity.ID())

	if config.CmsStore == nil {
		task.LogError("CMS Store not configured.")
		return false
	}

	block, err := config.CmsStore.BlockFindByID(context.Background(), entity.ID())

	if err != nil {
		task.LogError("Error retrieving block: " + err.Error())
		return false
	}

	if block != nil {
		task.LogInfo(" - Block found: " + block.ID())
		return true
	}

	content, err := entity.GetString("content", "")

	if err != nil {
		task.LogError("Error retrieving content: " + err.Error())
		return false
	}

	name, err := entity.GetString("name", "")

	if err != nil {
		task.LogError("Error retrieving title: " + err.Error())
		return false
	}

	status, err := entity.GetString("status", "")

	if err != nil {
		task.LogError("Error retrieving status: " + err.Error())
		return false
	}

	block = cmsstore.NewBlock().
		SetID(entity.ID()).
		SetHandle(entity.Handle()).
		SetName(name).
		SetContent(content).
		SetEditor("codemirror").
		SetStatus(status).
		SetSiteID(site.ID()).
		SetPageID("").
		SetTemplateID("").
		SetType("raw_html").
		SetParentID("").
		SetSequenceInt(-1).
		SetCreatedAt(carbon.CreateFromStdTime(entity.CreatedAt()).ToDateTimeString(carbon.UTC))

	err = config.CmsStore.BlockCreate(context.Background(), block)

	if err != nil {
		task.LogError("Error creating block: " + err.Error())
		return false
	}

	return true
}

func (task *cmsTransferTask) checkAndEnqueueTask() bool {
	// 1. Is the task already enqueued?
	if task.HasQueuedTask() {
		return false
	}

	// 2. Is the task asked to be enqueued?
	if task.GetParam("enqueue") != "yes" {
		return false
	}

	// 3. Enqueue the task
	_, err := task.Enqueue(task.index)

	if err != nil {
		task.LogError("Error enqueuing task: " + err.Error())
	} else {
		task.LogSuccess("Task enqueued.")
	}

	return true
}
