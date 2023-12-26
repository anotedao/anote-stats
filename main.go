package main

import (
	"log"

	"github.com/anonutopia/gowaves"
	macaron "gopkg.in/macaron.v1"
)

var mac *macaron.Macaron

var cch *Cache

var anc *gowaves.WavesNodeClient

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	anc = initAnote()

	cch = initCache()

	mac = initMacaron()

	mac.Run("127.0.0.1", Port)
}
