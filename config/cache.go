package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var cache *Cache

type Cache struct {
	Prefix string
	Driver string
}

func init() {
	cache = NewCache()
}

func NewCache() *Cache {
	cache = new(Cache)
	cache.Prefix = fmt.Sprintf("%s_cache_", viper.Get("APP_NAME"))
	if viper.GetBool("CACHE_DRIVER") == true {
		cache.Driver = fmt.Sprintf("%s", viper.Get("CACHE_DRIVER"))
	} else {
		cache.Driver = "file"
	}
	return cache
}
