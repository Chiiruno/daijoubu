package main

import (
	"github.com/Chiiruno/daijoubu/server/auth"
	"github.com/Chiiruno/daijoubu/server/cache"
	"github.com/Chiiruno/daijoubu/server/config"
	"github.com/Chiiruno/daijoubu/server/db"
	"github.com/Chiiruno/daijoubu/server/imager/assets"
	"github.com/Chiiruno/daijoubu/server/lang"
	mlog "github.com/Chiiruno/daijoubu/server/log"
	"github.com/Chiiruno/daijoubu/server/util"
	"github.com/Chiiruno/daijoubu/server/web"
)

func main() {
	if err := config.Server.Load(); err != nil {
		panic(err)
	}

	mlog.Init(mlog.Console)
	mlog.ConsoleHandler.SetDisplayColor(config.Server.Debug)

	if err := util.Parallel(db.Init, assets.Init); err != nil {
		panic(err)
	}

	if err := lang.Init(); err != nil {
		panic(err)
	}

	if err := util.Parallel(cache.Init, auth.Init); err != nil {
		panic(err)
	}

	web.Init()
}
