package main

import macaron "gopkg.in/macaron.v1"

func statsView(ctx *macaron.Context) {
	ctx.HTML(200, "stats")
}
