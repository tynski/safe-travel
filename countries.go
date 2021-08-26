package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type country struct {
	Country   string `json:"country"`
	Neighbors struct {
		Germany  string `json:"Germany"`
		Slovakia string `json:"Slovakia"`
		Poland   string `json:"Poland"`
	} `json:"Neighbors"`
}

const fileName = "./data/countries.json"

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

func GetCountry(country string) *country {
	countries := GetCountries()

	for _, a := range countries {
		if a.Country == country {
			return &a
		}
	}

	return nil
}

func AddCountry(newCountry country) bool {
	countries := GetCountries()

	for _, a := range countries {
		if a.Country == newCountry.Country {
			return false
		}
	}

	UpdateCountries(countries, newCountry)

	return true
}

func UpdateCountry(updateCountry country) bool {

	DeleteCountry(updateCountry.Country)
	UpdateCountries(countries, updateCountry)

	return false
}

func UpdateCountries(countries []country, newCountry country) {
	countries = append(countries, newCountry)

	result, err := json.Marshal(countries)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(fileName, result, 0644); err != nil {
		log.Fatal(err)
	}
}

func DeleteCountry(deleteCountry string) {
	countries := GetCountries()

	for i, a := range countries {
		if a.Country == deleteCountry {
			countries = append(countries[:i], countries[i+1:]...)
			return
		}
	}
}
