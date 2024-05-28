package service

import (
	"fmt"
	"getNationalClient/internal/model"
	"getNationalClient/internal/nationalpredict"
	"log"
	"os"
	"time"
)

// Взаимодействие с пользователем

type National interface {
	GetNational(string) (string, error)
}

// Объект Сервис

type Service struct {
	National          National
	NationalPredicter *nationalpredict.NationalPredicter
}

// New

func New(np *nationalpredict.NationalPredicter) *Service {
	return &Service{

		NationalPredicter: np,
	}
}

func (sv *Service) Start() (string, error) {
	var user model.User
	var err error

	for {
		fmt.Fscan(os.Stdin, &user.Name)
		if user.Name == "Владимир" {
			user.National = "Slavic"
			user.ID = uint32(time.Now().Unix())
			fmt.Printf("ID: %v\nName: %s\nNational: Slavic\n", user.ID, user.Name)
			// return "Slavic", nil
		} else {
			user.National, err = sv.NationalPredicter.GetNational(user)
			if err != nil {
				log.Println("Error find national:", err)
				// return "", err
			}
			user.ID = uint32(time.Now().Unix())
			fmt.Printf("ID: %v\nName: %s\nNational: %s\n", user.ID, user.Name, user.National)

			// return user.National, nil
		}
	}
}

// Start ( for{} получил в консоли имя - отдал национальность)
