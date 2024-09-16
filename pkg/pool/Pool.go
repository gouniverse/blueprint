package pool

import (
	"sync"
)

// pool stores data in memory, its private by default
// to avoid accidental access, use NewPool() to access
type pool struct {
	O interface{}
}

// instantiated is a singleton pool
var instantiated *pool

// once ensures that the pool is only instantiated once
var once sync.Once

func NewPool() *pool {
	once.Do(func() {
		instantiated = &pool{}
	})
	return instantiated
}

// func (p *pool) KeyBlock(key string) string {
// 	return "BLOCK_" + key
// }

// func (p *pool) KeyPage(key string) string {
// 	return "PAGE_" + key
// }

// func (p *pool) KeyTemplate(key string) string {
// 	return "TEMPLATE_" + key
// }

// func (p *pool) KeyTranslation(key string, language string) string {
// 	return "TRANSLATION_" + key + "_" + language
// }

// func (p *pool) GetBlock(key string) string {
// 	item := config.InMem.Get(key)
// 	if item == nil {
// 		return ""
// 	}
// 	return utils.ToString(item.Value())
// }

// func (p *pool) SetBlock(key string, value string) {
// 	config.InMem.Set(key, value, 24*time.Hour)
// }

// func (p *pool) GetPage(key string) string {
// 	item := config.InMem.Get(key)
// 	if item == nil {
// 		return ""
// 	}
// 	return utils.ToString(item.Value())
// }

// func (p *pool) SetPage(key string, value string) {
// 	config.InMem.Set(key, value, 24*time.Hour)
// }

// func (p *pool) GetTemplate(key string) string {
// 	item := config.InMem.Get(key)
// 	if item == nil {
// 		return ""
// 	}
// 	return utils.ToString(item.Value())
// }

// func (p *pool) SetTemplate(key string, value string) {
// 	config.InMem.Set(key, value, 24*time.Hour)
// }

// func (p *pool) GetTranslation(key string) string {
// 	item := config.InMem.Get(key)
// 	if item == nil {
// 		return ""
// 	}
// 	return utils.ToString(item.Value())
// }

// func (p *pool) SetTranslation(key string, value string) {
// 	config.InMem.Set(key, value, 24*time.Hour)
// }
