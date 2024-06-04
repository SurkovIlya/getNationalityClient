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

	countrylist := make(map[string]string, len(records))

	for _, row := range records {
		countrylist[row[0]] = row[1]
	}

	return countrylist, nil
}
