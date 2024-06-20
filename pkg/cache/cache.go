package cache

import (
	"fmt"
	"getNationalClient/internal/model"
	"log"
	"sync"
	"time"
)

type Cache struct {
	Records map[string]model.User
	TTL     uint32
	mu      sync.Mutex
}

func NewCash(ttlMs uint32, count int) *Cache {
	Record := make(map[string]model.User, count)

	return &Cache{
		Records: Record,
		TTL:     ttlMs,
	}

}

func (c *Cache) GetCaheVal(name string) (model.User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, ok := c.Records[name]; ok {
		value.Lastusedgetime = time.Now()
		c.Records[name] = value

		return value, nil
	} else {
		return model.User{}, fmt.Errorf("record is not found")
	}

}

func (c *Cache) UpdRecord(value model.User) {
	c.mu.Lock()
	if _, ok := c.Records[value.Name]; ok {
		log.Printf("Кэш был обновлен!")
		c.Records[value.Name] = value
		c.mu.Unlock()
	}

}

func (c *Cache) AddRecodr(user model.User) {
	c.mu.Lock()
	if _, ok := c.Records[user.Name]; !ok {
		c.Records[user.Name] = user
		log.Printf("Добавлена новая запись в кеш: %v", user)
		c.mu.Unlock()
	}
}

func (c *Cache) Clean() {
	for {
		now := time.Now()
		for _, val := range c.Records {
			timeUse := val.Lastusedgetime
			timecheck := timeUse.Add(time.Duration(c.TTL) * time.Minute)
			if int(now.Sub(timecheck)) > 0 {
				delete(c.Records, val.Name)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
