package main

import (
	"github.com/Chiiruno/daijoubu/src/server/cache"
	"github.com/Chiiruno/daijoubu/src/server/config"
	"github.com/Chiiruno/daijoubu/src/server/db"
	"github.com/Chiiruno/daijoubu/src/server/lang"
	"github.com/Chiiruno/daijoubu/src/server/web"
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
