package exception

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

type ExceptionPerson struct {
	Name     string `json:"name"`
	National string `json:"national"`
}

type ExceptionStore struct {
	ExceptionMap map[string]string
}

var expRespons ExceptionPerson
var exceptionName []ExceptionPerson

func New() *ExceptionStore {
	exceptionMap := make(map[string]string, len(exceptionName))

	file, err := os.Open("././data/exception.csv")
	if err != nil {
		log.Println("File exception.csv is not found:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Could not read:", err)
		}
		expRespons.Name = record[0]
		expRespons.National = record[1]
		exceptionMap[expRespons.Name] = expRespons.National
	}

	return &ExceptionStore{
		ExceptionMap: exceptionMap,
	}
}

func (es *ExceptionStore) AddExcStore(exception ExceptionPerson) error {
	file, err := os.Open("././data/exception.csv")
	if err != nil {
		log.Println("File exception.csv is not found:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	exceptionData := []string{exception.Name, exception.National}
	writer.Write(exceptionData)

	if val, ok := es.ExceptionMap[exception.Name]; !ok {
		es.ExceptionMap[exception.Name] = exception.National

		return nil
	} else if val == exception.National && ok {
		err := errors.New("Исключение уже существует")
		return err
	} else if val != exception.National && ok {
		es.ExceptionMap[exception.Name] = exception.National

		return nil
	}
	err = errors.New("Добавление исключения невозможно! Обратитесь к разработчику")
	return err
}

func (es *ExceptionStore) ExpetionCheck(name string) ExceptionPerson {
	val, ok := es.ExceptionMap[name]
	if !ok {
		return ExceptionPerson{}
	}
	expRespons.Name = name
	expRespons.National = val

	return expRespons
}
