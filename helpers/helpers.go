package helpers

import (
	"encoding/csv"
	"os"
)

func ReadCsv() ([][]string, error) {
	f, err := os.Open("./data/heroes_information.csv")

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, err
}
