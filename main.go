package main

import (
	"log"

	macaron "gopkg.in/macaron.v1"
)

var mac *macaron.Macaron

var cch *Cache

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cch = initCache()

	mac = initMacaron()

	mac.Run("127.0.0.1", Port)
}
