package service

import (
	"fmt"
	"getNationalClient/internal/model"
	"getNationalClient/internal/nationalpredict"
	"getNationalClient/pkg/cache"
	"log"
	"os"
	"time"
)

type National interface {
	GetNational(string) (string, error)
}

type Service struct {
	National          National
	NationalPredicter *nationalpredict.NationalPredicter
}

func New(np *nationalpredict.NationalPredicter) *Service {
	return &Service{

		NationalPredicter: np,
	}
}

func (sv *Service) Start() (string, error) {
	var user model.User
	// var err error
	const ttlMs = 5
	cacheUsers := cache.NewCash(uint32(ttlMs))
	go cacheUsers.Clean()

	for {
		fmt.Fscan(os.Stdin, &user.Name)
		if user.Name == "Владимир" {
			user.National = "Slavic"
			user.ID = uint32(time.Now().Unix())
			fmt.Printf("ID: %v\nName: %s\nNational: Slavic\n", user.ID, user.Name)
			// return "Slavic", nil
		} else {
			cu, err := cacheUsers.GetUserByName(user.Name)
			if err != nil {
				log.Println(err)
				user.National, err = sv.NationalPredicter.GetNational(user)
				if err != nil {
					log.Println("Error find national:", err)
					// return "", err
				}
				user.ID = uint32(time.Now().Unix())
				cacheUsers.AddWord(user)
				fmt.Printf("ID: %v\nName: %s\nNational: %s\n", user.ID, user.Name, user.National)

				// return user.National, nil
			} else {
				fmt.Printf("ID: %v\nName: %s\nNational: %s\n", cu.ID, cu.Name, cu.National)

				// return cu.National, nil
			}

		}
	}
}
