package nationalpredict

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func GetCountryList() (map[string]string, error) {

	file, err := os.Open("././data/exmp.csv")
	if err != nil {
		log.Println("Open error:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Read error:", err)
		return nil, err
	}
	var cl map[string]string

	for i, row := range records {
		cl[row[0]] = row[i]
	}
	return cl, nil
}
