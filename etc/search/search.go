package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sajari/fuzzy"
)

type Location struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var (
	locations []Location
	model     *fuzzy.Model
)

func LoadLocations(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return errors.New("ошибка парсинга JSON")
	}

	var trainData []string
	for _, loc := range locations {
		trainData = append(trainData, strings.ToLower(loc.City))
	}

	fmt.Println("Train Data:", trainData)

	model = fuzzy.NewModel()
	model.SetThreshold(1)
	model.SetDepth(2)
	model.Train(trainData)

	return nil
}

func GetLocations(query string) ([]Location, error) {
	if query == "" {
		return nil, errors.New("запрос не должен быть пустым")
	}

	query = strings.ToLower(query)
	fmt.Println("Query:", query)

	matches := model.Suggestions(query, false)
	fmt.Println("Matches: ", matches)

	results := []Location{}
	for _, match := range matches {
		for _, loc := range locations {
			if len(query) >= 3 && len(query) < 7 {
				if strings.HasPrefix(strings.ToLower(loc.City), query) {
					results = append(results, loc)
					break
				}
			} else if len(query) == 2 {
				if strings.HasPrefix(strings.ToLower(loc.State), query) {
					results = append(results, loc)
					break
				}
			} else {
				if strings.ToLower(loc.City) == match {
					results = append(results, loc)
					break
				}
			}
		}
		if len(results) >= 15 {
			break
		}
	}

	return results, nil
}
