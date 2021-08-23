package services

import (
	"encoding/csv"
	"log"
	"os"
)

const HEROES_DATA_PATH = "./data/heroes_information.csv"

func ImportHeroesData() [][]string {
	f, err := os.Open(HEROES_DATA_PATH)

	if err != nil {
		log.Fatal("Unable to read input file"+HEROES_DATA_PATH, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for"+HEROES_DATA_PATH, err)
	}

	return records
}
