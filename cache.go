package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dustin/go-humanize"
)

type Cache struct {
	StatsCache *StatsCache
}

func (c *Cache) loadStatsCache() {
	at := AddressToday
	priceInt, err := getData("%s__priceAnote", &at)
	if err != nil {
		log.Println(err.Error())
	}
	price := float64(priceInt.(int64)) / float64(MULT8)
	c.StatsCache.Price = fmt.Sprintf("%.4f", price)
	c.StatsCache.AmountNode = getAmountNode()

	bh, err := anc.BlocksHeight()
	if err != nil {
		log.Println(err.Error())
	}
	mined := int64(bh.Height + 1000)
	c.StatsCache.Mined = humanize.Comma(mined)
}

func (c *Cache) start() {
	for {
		c.loadStatsCache()

		time.Sleep(time.Second * 10)
	}
}

type StatsCache struct {
	ActiveMiners int
	Holders      int
	Price        string
	AmountTlg    float64
	AmountMobile float64
	AmountNode   float64
	Mined        string
	Community    string
	Circulation  string
	Active       int
	Payout       int
	Inactive     int
	Referred     int
}

func initCache() *Cache {
	c := &Cache{}
	c.StatsCache = &StatsCache{}
	go c.start()

	return c
}
