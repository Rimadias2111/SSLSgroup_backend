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
		trainData = append(trainData, strings.ToLower(loc.City+", "+loc.State))
	}

	fmt.Println("Train Data:", trainData)

	model = fuzzy.NewModel()
	model.SetThreshold(3)
	model.SetDepth(3)
	model.Train(trainData)

	return nil
}

func GetLocations(query string) ([]Location, error) {
	if query == "" {
		return nil, errors.New("запрос не должен быть пустым")
	}

	query = strings.ToLower(query)
	fmt.Println("Query:", query)

	if strings.HasPrefix(query, ",") {
		stateQuery := strings.TrimSpace(strings.TrimPrefix(query, ","))
		results := []Location{}
		for _, loc := range locations {
			if strings.ToLower(loc.State) == stateQuery {
				results = append(results, loc)
			}
		}
		return results, nil
	}

	if len(query) < 5 {
		results := []Location{}
		for _, loc := range locations {
			if strings.Contains(strings.ToLower(loc.State), query) || strings.Contains(strings.ToLower(loc.City), query) {
				results = append(results, loc)
			}
		}
		return results, nil
	}

	matches := model.Suggestions(query, false)
	fmt.Println("Matches: ", matches)

	results := []Location{}
	for _, match := range matches {
		for _, loc := range locations {
			if strings.ToLower(loc.City+", "+loc.State) == match {
				results = append(results, loc)
				break
			}
		}
		if len(results) >= 10 {
			break
		}
	}

	return results, nil
}
