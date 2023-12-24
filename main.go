package main

import (
	"log"

	macaron "gopkg.in/macaron.v1"
)

var mac *macaron.Macaron

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mac = initMacaron()

	mac.Run("127.0.0.1", Port)
}
