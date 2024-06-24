package exception

import (
	"encoding/csv"
	"fmt"
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
	if val, ok := es.ExceptionMap[exception.Name]; !ok {
		es.ExceptionMap[exception.Name] = exception.National

		file, err := os.OpenFile("././data/exception.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Println("File exception.csv is not found:", err)

			return fmt.Errorf("добавление исключения невозможно! Обратитесь к разработчику")
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		defer writer.Flush()

		exceptionData := [][]string{{exception.Name, exception.National}}
		writer.WriteAll(exceptionData)

		return nil
	} else if val == exception.National && ok {
		return fmt.Errorf("исключение уже существует")
	} else if val != exception.National && ok {
		es.ExceptionMap[exception.Name] = exception.National

		file, err := os.Open("././data/exception.csv")
		if err != nil {
			log.Println("File exception.csv is not found:", err)

			return fmt.Errorf("добавление исключения невозможно! Обратитесь к разработчику")
		}
		defer file.Close()

		read := csv.NewReader(file)

		record, err := read.ReadAll()
		if err != nil {
			return fmt.Errorf("добавление исключения невозможно! Обратитесь к разработчику")
		}

		fileNew, err := os.Create("././data/exception.csv")
		if err != nil {
			return fmt.Errorf("добавление исключения невозможно! Обратитесь к разработчику")
		}
		defer fileNew.Close()

		write := csv.NewWriter(fileNew)

		for _, value := range record {
			if value[0] == exception.Name {
				value[1] = exception.National
			}
		}

		write.WriteAll(record)
		write.Flush()

		return nil
	}

	return fmt.Errorf("добавление исключения невозможно! Обратитесь к разработчику")
}

func (es *ExceptionStore) DelException(name string) error {
	if _, ok := es.ExceptionMap[name]; !ok {
		return fmt.Errorf("удаление невозможно: для имени %v нет исключений", name)
	}
	delete(es.ExceptionMap, name)

	file, err := os.Open("././data/exception.csv")
	if err != nil {
		log.Println("File exception.csv is not found:", err)

		return fmt.Errorf("удаление исключения невозможно! Обратитесь к разработчику")
	}
	defer file.Close()

	read := csv.NewReader(file)

	record, err := read.ReadAll()
	if err != nil {
		log.Println("Could not read file:", err)

		return fmt.Errorf("удаление исключения невозможно! Обратитесь к разработчику")
	}

	fileNew, err := os.Create("././data/exception.csv")
	if err != nil {
		log.Println("Failed to create a file:", err)

		return fmt.Errorf("удаление исключения невозможно! Обратитесь к разработчику")
	}
	defer fileNew.Close()

	write := csv.NewWriter(fileNew)

	for i, value := range record {
		if value[0] == name {
			record = append(record[:i], record[i+1:]...)
		}
	}

	write.WriteAll(record)
	write.Flush()

	return nil
}

func (es *ExceptionStore) ExpetionCheck(name string) ExceptionPerson {
	if _, ok := es.ExceptionMap[name]; !ok {
		return ExceptionPerson{}
	}

	expRespons.Name = name
	expRespons.National = es.ExceptionMap[name]

	return expRespons
}
