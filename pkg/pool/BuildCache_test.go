package pool

// import (
// 	"project/config"
// 	"project/internal/testutils"
// 	"testing"

// 	"github.com/gouniverse/cms"
// 	"github.com/gouniverse/entitystore"
// )

// func TestBuildCacheBlocks(t *testing.T) {
// 	testutils.Setup()

// 	block, _ := config.Cms.EntityStore.EntityCreateWithTypeAndAttributes(cms.ENTITY_TYPE_BLOCK, map[string]string{
// 		"content": "Lorem ipsum dolor sit amet",
// 	})

// 	_, errBlocks := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_BLOCK,
// 	})

// 	if errBlocks != nil {
// 		t.Fatal(errBlocks.Error())
// 	}

// 	BuildCache()

// 	key := NewPool().KeyBlock(block.ID())
// 	if NewPool().GetBlock(key) == "" {
// 		t.Fatal(`Block `+key+` Must NOT be empty`, NewPool().GetBlock(key))
// 	}

// 	if NewPool().GetBlock(key) != "Lorem ipsum dolor sit amet" {
// 		t.Fatal(`For key `+key+` Expected "Lorem ipsum dolor sit amet", but found:`, NewPool().GetBlock(key))
// 	}
// }

// func TestBuildCachePages(t *testing.T) {
// 	testutils.Setup()

// 	_, _ = config.Cms.EntityStore.EntityCreateWithTypeAndAttributes(cms.ENTITY_TYPE_PAGE, map[string]string{
// 		"alias":   "/page-1",
// 		"content": "Lorem ipsum dolor sit amet",
// 	})

// 	_, errPages := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_PAGE,
// 	})

// 	if errPages != nil {
// 		t.Fatal(errPages.Error())
// 	}

// 	BuildCache()

// 	key := NewPool().KeyPage("/page-1")
// 	if NewPool().GetPage(key) == "" {
// 		t.Fatal(`Block `+key+` Must NOT be empty`, NewPool().GetPage(key))
// 	}

// 	if NewPool().GetPage(key) != "Lorem ipsum dolor sit amet" {
// 		t.Fatal(`For key `+key+` Expected "Lorem ipsum dolor sit amet", but found:`, NewPool().GetPage(key))
// 	}
// }

// func TestBuildCacheTemplates(t *testing.T) {
// 	testutils.Setup()

// 	template, _ := config.Cms.EntityStore.EntityCreateWithTypeAndAttributes(cms.ENTITY_TYPE_TEMPLATE, map[string]string{
// 		"content": "Lorem ipsum dolor sit amet",
// 	})

// 	_, errTemplates := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_TEMPLATE,
// 	})

// 	if errTemplates != nil {
// 		t.Fatal(errTemplates.Error())
// 	}

// 	BuildCache()

// 	key := NewPool().KeyTemplate(template.ID())
// 	if NewPool().GetTemplate(key) == "" {
// 		t.Fatal(`Block `+key+` Must NOT be empty`, NewPool().GetTemplate(key))
// 	}

// 	if NewPool().GetTemplate(key) != "Lorem ipsum dolor sit amet" {
// 		t.Fatal(`For key `+key+` Expected "Lorem ipsum dolor sit amet", but found:`, NewPool().GetTemplate(key))
// 	}
// }

// func TestBuildCacheTranslations(t *testing.T) {
// 	testutils.Setup()

// 	translation, _ := config.Cms.EntityStore.EntityCreateWithTypeAndAttributes(cms.ENTITY_TYPE_TRANSLATION, map[string]string{
// 		"en": "Hello World",
// 		"bg": "Zdravei sviat",
// 	})

// 	if translation == nil {
// 		t.Fatal("Translation MUST NOT be nil")
// 	}

// 	translation.SetHandle("hello_world")

// 	config.Cms.EntityStore.EntityUpdate(*translation)

// 	_, errTranslations := config.Cms.EntityStore.EntityList(entitystore.EntityQueryOptions{
// 		EntityType: cms.ENTITY_TYPE_TRANSLATION,
// 	})

// 	if errTranslations != nil {
// 		t.Fatal(errTranslations.Error())
// 	}

// 	BuildCache()

// 	key1 := NewPool().KeyTranslation(translation.ID(), "en")
// 	if NewPool().GetTranslation(key1) == "" {
// 		t.Fatal("Key Must NOT be nil", key1, NewPool().GetTranslation(key1))
// 	}
// 	if NewPool().GetTranslation(key1) != "Hello World" {
// 		t.Fatal("Expected, but found:", key1, NewPool().GetTranslation(key1))
// 	}

// 	key2 := NewPool().KeyTranslation(translation.Handle(), "en")
// 	if NewPool().GetTranslation(key2) == "" {
// 		t.Fatal("Key Must NOT be nil", key2, NewPool().GetTranslation(key2))
// 	}
// 	if NewPool().GetTranslation(key2) != "Hello World" {
// 		t.Fatal("Expected, but found:", key2, NewPool().GetTranslation(key2))
// 	}

// 	key3 := NewPool().KeyTranslation(translation.ID(), "bg")
// 	if NewPool().GetTranslation(key3) == "" {
// 		t.Fatal(`Key `+key3+` MUST NOT be nil`, NewPool().GetTranslation(key3))
// 	}
// 	if NewPool().GetTranslation(key3) != "Zdravei sviat" {
// 		t.Fatal(`For key `+key3+` expected 'Zdravei sviat', but found:`, NewPool().GetTranslation(key3))
// 	}

// 	key4 := NewPool().KeyTranslation(translation.Handle(), "bg")
// 	if NewPool().GetTranslation(key4) == "" {
// 		t.Fatal(`Key `+key4+` MUST NOT be nil`, NewPool().GetTranslation(key4))
// 	}
// 	if NewPool().GetTranslation(key4) != "Zdravei sviat" {
// 		t.Fatal(`For key `+key4+` expected 'Zdravei sviat', but found:`, NewPool().GetTranslation(key4))
// 	}
// }
