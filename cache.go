package main

import (
	"log"
	"time"
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
	c.StatsCache.Price = price
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
	Price        float64
	AmountTlg    float64
	AmountMobile float64
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