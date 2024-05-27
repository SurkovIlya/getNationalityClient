package nationalpredict

import "ezclient/internal/model"

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
	//np.NationalSource.GetNatByName()

	return "", nil
}
