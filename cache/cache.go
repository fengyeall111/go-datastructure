package main

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	cache.Set("my-unique-key", []byte("你好你好也许真的好也许也许也许也许也许真的很早"))
	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}
