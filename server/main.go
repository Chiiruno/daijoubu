package main

import (
	"github.com/Chiiruno/daijoubu/server/cache"
	"github.com/Chiiruno/daijoubu/server/config"
	"github.com/Chiiruno/daijoubu/server/db"
	"github.com/Chiiruno/daijoubu/server/lang"
	"github.com/Chiiruno/daijoubu/server/web"
)

func main() {
	for _, fn := range [5]func() error{
		cache.Init,
		config.Init,
		db.Init,
		lang.Init,
		web.Init,
	} {
		go func(fn func() error) {
			if err := fn(); err != nil {
				panic(err)
			}
		}(fn)
	}
}
