package cached

import (
	"fmt"
	"log"
	"time"
)

type CachedItem struct {
	Value     interface{}
	Timestamp time.Time
}

type CachedMemory struct {
	Cache map[string]CachedItem
}

var cachedMem *CachedMemory

func init() {
	log.Println(fmt.Sprintf("%s", "test"))
	cachedMem = &CachedMemory{
		Cache: map[string]CachedItem{},
	}
}

func (c *CachedMemory) Put(key string, v interface{}) error {
	c.Cache[key] = CachedItem{
		Value:     v,
		Timestamp: time.Now(),
	}
	return nil
}

func (c *CachedMemory) Get(key string) interface{} {
	if v, e := c.Cache[key]; e {
		return v
	} else {
		return nil
	}
}

func Expired() error {

	return nil
}

func Put(key string, v interface{}) error {
	return cachedMem.Put(key, v)
}

func Get(key string) interface{} {

	return cachedMem.Get(key)
}
