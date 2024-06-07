package exception

type ExcentionPerson struct {
	Name     string `json:"name"`
	National string `json:"national"`
}

type ExceptionStore struct {
	ExceptionMap map[string]string
}

var exceptionName []ExcentionPerson

func New() *ExceptionStore {
	exceptionMap := make(map[string]string, len(exceptionName))
	exceptionName = append(exceptionName, ExcentionPerson{Name: "Владимир", National: "Slavic"}, ExcentionPerson{Name: "Диана", National: "Сенпай"})
	for _, person := range exceptionName {
		exceptionMap[person.Name] = person.National
	}
	return &ExceptionStore{
		ExceptionMap: exceptionMap,
	}
}

// func (es *ExceptionStore) AddExcStore(exception ExcentionPerson) {
// 	exceptionName = append(exceptionName, exception)
// 	for _, person := range exceptionName {
// 		es.ExceptionMap[person.Name] = person.National
// 	}
// }

func (es *ExceptionStore) ExpetionCheck(name string) ExcentionPerson {

	var expRespons ExcentionPerson
	val, ok := es.ExceptionMap[name]
	if !ok {
		return ExcentionPerson{}
	}
	expRespons.Name = name
	expRespons.National = val

	return expRespons
}
