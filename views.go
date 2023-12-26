package main

import macaron "gopkg.in/macaron.v1"

func statsView(ctx *macaron.Context) {
	ctx.Data["Cache"] = cch.StatsCache
	ctx.HTML(200, "stats")
}
