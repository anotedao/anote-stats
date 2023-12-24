package main

import macaron "gopkg.in/macaron.v1"

func statsView(ctx *macaron.Context) {
	ctx.Data["Price"] = 39.48
	ctx.HTML(200, "stats")
}
