package nationalpredict

import (
	"getNationalClient/internal/model"
	"log"
)

type NationalSource interface {
	GetNationalByName(string) (string, error)
}

type NationalPredicter struct {
	CountryList    map[string]string
	NationalSource NationalSource
}

func New(cl map[string]string, ns NationalSource) *NationalPredicter {
	return &NationalPredicter{
		CountryList:    cl,
		NationalSource: ns,
	}
}

func (np *NationalPredicter) GetNational(user model.User) (string, error) {
	iso, err := np.NationalSource.GetNationalByName(user.Name)
	if err != nil {
		log.Println("GetNational error", err)

		return "", err
	}
	national := np.CountryList[iso]

	return national, nil
}
