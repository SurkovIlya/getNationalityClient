package service

import (
	"getNationalClient/internal/exception"
	"getNationalClient/internal/model"
	"getNationalClient/internal/nationalpredict"
	"getNationalClient/pkg/cache"
	"time"
)

type National interface {
	GetNational(string) (string, error)
}

type Service struct {
	National          National
	NationalPredicter *nationalpredict.NationalPredicter
	Exception         *exception.ExceptionStore
	UserCache         *cache.Cache
}

func New(np *nationalpredict.NationalPredicter, exc *exception.ExceptionStore, cache *cache.Cache) *Service {
	return &Service{
		NationalPredicter: np,
		Exception:         exc,
		UserCache:         cache,
	}
}

func (sv *Service) NationalName(name string) (model.User, error) {
	var user model.User
	user.Name = name

	go sv.UserCache.Clean()

	for {
		cacheUser, err := sv.UserCache.GetCaheVal(user.Name)
		if err == nil {
			sv.UserCache.UpdRecord(cacheUser)

			return cacheUser, nil
		}
		exception := sv.Exception.ExpetionCheck(name)
		if exception.Name != "" {
			user.Name = exception.Name
			user.National = exception.National
			user.Lastusedgetime = time.Now()
			sv.UserCache.AddRecodr(user)

			return user, nil
		}
		user.National, err = sv.NationalPredicter.GetNational(user)
		if err == nil {
			user.Lastusedgetime = time.Now()
			sv.UserCache.AddRecodr(user)

			return user, nil
		}

	}
}
