package pool

// BuildCache builds the cache for the application
// can cache all entities that are frequently used
// to avoid retrieving them each time
// Can be called:
// - on application start
// - on schedule to keep the cache warm
func BuildCache() {
	// cfmt.Infoln("Building cache for templates...")
	// buildCacheTemplates()
	// cfmt.Infoln("Building cache for pages...")
	// buildCachePages()
	// cfmt.Infoln("Building cache for blocks...")
	// buildCacheBlocks()
	// cfmt.Infoln("Building cache for translations...")
	// buildCacheTranslations()
}

// // buildCachePages stores pages by ID and ALIAS
// func buildCachePages() {
// 	pages, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_PAGE,
// 		Limit:      2000,
// 	})

// 	if err != nil {
// 		cfmt.Errorln(err.Error())
// 		return
// 	}

// 	lo.ForEach(pages, func(page entitystore.Entity, index int) {
// 		attrAlias, err := config.Cms.EntityStore.AttributeFind(page.ID(), "alias")
// 		if err != nil {
// 			return
// 		}
// 		if attrAlias == nil {
// 			return
// 		}
// 		attrContent, err := config.Cms.EntityStore.AttributeFind(page.ID(), "content")
// 		if err != nil {
// 			return
// 		}
// 		if attrContent == nil {
// 			return
// 		}

// 		id := page.ID()
// 		alias := attrAlias.AttributeValue()
// 		content := attrContent.AttributeValue()
// 		key1 := NewPool().KeyPage(alias)
// 		NewPool().SetBlock(key1, content)
// 		key2 := NewPool().KeyPage(id)
// 		NewPool().SetBlock(key2, content)
// 	})
// }

// // buildCacheTemplates stores templates by ID
// func buildCacheTemplates() {
// 	templates, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_TEMPLATE,
// 		Limit:      2000,
// 	})

// 	if err != nil {
// 		cfmt.Errorln(err.Error())
// 		return
// 	}

// 	lo.ForEach(templates, func(template entitystore.Entity, index int) {
// 		attrContent, err := config.Cms.EntityStore.AttributeFind(template.ID(), "content")
// 		if err != nil {
// 			return
// 		}
// 		if attrContent == nil {
// 			return
// 		}

// 		id := template.ID()
// 		content := attrContent.AttributeValue()
// 		key := NewPool().KeyTemplate(id)
// 		NewPool().SetTemplate(key, content)
// 	})
// }

// // buildCacheBlocks stores blocks by ID
// func buildCacheBlocks() {
// 	blocks, err := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_BLOCK,
// 		Limit:      2000,
// 	})

// 	if err != nil {
// 		cfmt.Errorln(err.Error())
// 		return
// 	}

// 	lo.ForEach(blocks, func(block entitystore.Entity, index int) {
// 		attrContent, err := config.Cms.EntityStore.AttributeFind(block.ID(), "content")
// 		if err != nil {
// 			return
// 		}
// 		if attrContent == nil {
// 			return
// 		}

// 		id := block.ID()
// 		content := attrContent.AttributeValue()
// 		key := NewPool().KeyBlock(id)
// 		NewPool().SetBlock(key, content)
// 	})
// }

// // buildCacheTranslations stores translation by ID and language, and KEY and language
// func buildCacheTranslations() {
// 	translations, errTranslations := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_TRANSLATION,
// 		Limit:      2000,
// 	})

// 	if errTranslations != nil {
// 		cfmt.Errorln(errTranslations.Error())
// 		return
// 	}

// 	lo.ForEach(translations, func(translation entitystore.Entity, index int) {
// 		attributes, errAttributes := translation.GetAttributes()
// 		if errAttributes != nil {
// 			cfmt.Errorln(errTranslations.Error())
// 			return
// 		}

// 		id := translation.ID()
// 		handle := translation.Handle()
// 		for _, attr := range attributes {
// 			language := attr.AttributeKey()
// 			translationText := attr.AttributeValue()
// 			key1 := NewPool().KeyTranslation(id, language)
// 			key2 := NewPool().KeyTranslation(handle, language)
// 			NewPool().SetTranslation(key1, translationText)
// 			NewPool().SetTranslation(key2, translationText)
// 		}
// 	})
// }
