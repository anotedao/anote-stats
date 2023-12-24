package main

import (
	"github.com/go-macaron/cache"
	macaron "gopkg.in/macaron.v1"
)

func initMacaron() *macaron.Macaron {
	mac := macaron.Classic()

	mac.Use(macaron.Renderer())
	mac.Use(cache.Cacher())

	mac.Get("/", statsView)

	return mac
}
