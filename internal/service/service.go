package service

import (
	"getNationalClient/internal/exception"
	"getNationalClient/internal/model"
	"getNationalClient/internal/nationalpredict"
	"getNationalClient/pkg/cache"
)

type National interface {
	GetNational(string) (string, error)
}

type Service struct {
	National          National
	NationalPredicter *nationalpredict.NationalPredicter
	Exception         *exception.ExceptionStore
}

func New(np *nationalpredict.NationalPredicter, exc *exception.ExceptionStore) *Service {
	return &Service{
		NationalPredicter: np,
		Exception:         exc,
	}
}

func (sv *Service) NationalName(name string) (string, error) {
	var user model.User
	user.Name = name
	const ttlMs = 5
	cacheUsers := cache.NewCash(uint32(ttlMs))
	go cacheUsers.Clean()

	for {
		cacheUser, err := cacheUsers.GetCaheVal(user.Name)
		if err == nil {
			user = cacheUser

			return user.National, nil
		}
		exception := sv.Exception.ExpetionCheck(name)
		if exception.Name != "" {
			user.Name = exception.Name
			user.National = exception.National
			cacheUsers.AddWord(user)

			return user.National, nil
		}
		user.National, err = sv.NationalPredicter.GetNational(user)
		if err == nil {
			cacheUsers.AddWord(user)

			return user.National, nil
		}

	}
}
