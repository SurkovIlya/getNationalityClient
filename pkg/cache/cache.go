package cache

import (
	"fmt"
	"getNationalClient/internal/model"
	"time"
)

// type UserInCashe struct {
// 	User           model.User
// 	Lastusedgetime time.Time
// }

const UserCount = 1000

type Cache struct {
	Records map[string]model.User
	TTL     uint32
}

func NewCash(ttlMs uint32) *Cache {
	Record := make(map[string]model.User, UserCount)

	return &Cache{
		Records: Record,
		TTL:     ttlMs,
	}

}

func (c *Cache) GetUserByName(name string) (model.User, error) {
	if value, ok := c.Records[name]; ok {
		value.Lastusedgetime = time.Now()
		c.Records[name] = value
		return value, nil
	} else {
		return model.User{}, fmt.Errorf("record is not found")
	}
}
func (c *Cache) AddWord(user model.User) {
	if _, ok := c.Records[user.Name]; !ok {
		c.Records[user.Name] = user
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
