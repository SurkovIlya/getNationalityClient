package nationalsource

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ApiResponse struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

type NationalSource struct {
	Host string
}

func New(host string) *NationalSource {
	return &NationalSource{
		Host: host,
	}
}

func (np *NationalSource) GetNationalByName(name string) (string, error) {
	url := fmt.Sprintf(np.Host+"/?name=%s", name)

	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Bad status code", resp.StatusCode)

		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll err", err)

		return "", err
	}

	var apiResponse ApiResponse

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("Unmushal err", err)

		return "", err
	}

	// // fmt.Printf("Name: %s\n", apiResponse.Name)
	// // fmt.Printf("Probabilities:\n")
	// // for _, country := range apiResponse.Country {
	// // 	fmt.Printf("Country: %s, Probability: %f\n", country.CountryID, country.Probability)
	// // }

	return apiResponse.Country[0].CountryID, nil
}
