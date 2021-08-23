package services

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/galloaleonardo/api-marvel-heroes/helpers"
	"github.com/galloaleonardo/api-marvel-heroes/structs"
)

func importHeroes() []structs.Hero {
	lines, err := helpers.ReadCsv()

	if err != nil {
		panic(err)
	}

	heroes := []structs.Hero{}

	for _, line := range lines[1:] {
		height, err := strconv.ParseFloat(line[6], 64)

		if err != nil {
			height = 0
		}

		width, err := strconv.ParseFloat(line[10], 64)

		if err != nil {
			width = 0
		}

		data := structs.Hero{
			Name:      line[1],
			Gender:    line[2],
			EyeColor:  line[3],
			Race:      line[4],
			HairColor: line[5],
			Height:    height,
			Publisher: line[7],
			SkinColor: line[8],
			Alignment: line[9],
			Width:     width,
		}

		heroes = append(heroes, data)
	}

	return heroes
}

func filterHeroes(params url.Values, heroes []structs.Hero) []structs.Hero {
	var searched []structs.Hero

	for k, v := range params {
		if k == "name" {
			for i := range heroes {
				if strings.Contains(strings.ToLower(heroes[i].Name), strings.ToLower(strings.Join(v, ""))) {
					searched = append(searched, heroes[i])
				}
			}
		}

		if k == "gender" {
			for i := range heroes {
				if strings.Contains(strings.ToLower(heroes[i].Gender), strings.ToLower(strings.Join(v, ""))) {
					searched = append(searched, heroes[i])
				}
			}
		}

		if k == "race" {
			for i := range heroes {
				if strings.Contains(strings.ToLower(heroes[i].Race), strings.ToLower(strings.Join(v, ""))) {
					searched = append(searched, heroes[i])
				}
			}
		}

		if k == "hairColor" {
			for i := range heroes {
				if strings.Contains(strings.ToLower(heroes[i].HairColor), strings.ToLower(strings.Join(v, ""))) {
					searched = append(searched, heroes[i])
				}
			}
		}

		if k == "publisher" {
			for i := range heroes {
				if strings.Contains(strings.ToLower(heroes[i].Publisher), strings.ToLower(strings.Join(v, ""))) {
					searched = append(searched, heroes[i])
				}
			}
		}
	}

	return searched
}

func Heroes(params url.Values) []structs.Hero {
	heroes := importHeroes()

	if len(params) != 0 {
		heroes = filterHeroes(params, heroes)
	}

	return heroes
}

func MarvelHeroes() []structs.Hero {
	heroes := importHeroes()

	params := url.Values{}

	params.Add("publisher", "Marvel Comics")
	heroes = filterHeroes(params, heroes)

	return heroes
}

func DCHeroes() []structs.Hero {
	heroes := importHeroes()

	params := url.Values{}

	params.Add("publisher", "DC Comics")
	heroes = filterHeroes(params, heroes)

	return heroes
}
