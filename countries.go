package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type country struct {
	Country   string `json:"country"`
	Neighbors []struct {
		Neighbor string `json:"neighbor"`
		Color    string `json:"color"`
	}
}

const fileName = "/data/countries.json"

func GetCountries() []country {
	var countries []country

	jsonFile, _ := os.Open(fileName)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	if err := json.Unmarshal(byteValue, &countries); err != nil {
		return nil
	}

	return countries
}

func AddCountry(newCountry country) bool {
	countries := GetCountries()

	for _, a := range countries {
		if a.Country == newCountry.Country {
			return false
		}
	}

	countries = append(countries, newCountry)

	result, err := json.Marshal(countries)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(fileName, result, 0644); err != nil {
		log.Fatal(err)
	}

	return true
}
