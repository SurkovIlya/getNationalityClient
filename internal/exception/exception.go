package exception

type ExpentionPerson struct {
	Name     string
	National string
}

var exceptionName []ExpentionPerson

func ExpetionCheck() map[string]string {
	exceptionMap := make(map[string]string, len(exceptionName))
	exceptionName = append(exceptionName, ExpentionPerson{Name: "Владимир", National: "Slavic"}, ExpentionPerson{Name: "Диана", National: "Сенпай"})
	for _, person := range exceptionName {
		exceptionMap[person.Name] = person.National
	}

	return exceptionMap
}
