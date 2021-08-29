package main

import (
	"encoding/json"
	"io/ioutil"
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

func GetCountry(findCountry string) *country {
	countries := GetCountries()

	for _, a := range countries {
		if a.Country == findCountry {
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

	countries = append(countries, newCountry)
	UpdateCountries(countries)

	return true
}

func UpdateCountry(updateCountry string) bool {
	newCountry := GetCountry(updateCountry)

	if newCountry == nil {
		return false
	}

	countries := GetCountries()

	if len(countries) < 3 {
		return false
	}

	newCountries := append(DeleteFromCountries(countries, updateCountry), *newCountry)

	if err := UpdateCountries(newCountries); err != nil {
		return false
	}

	return true
}

func DeleteCountry(deleteCountry string) bool {
	countries := GetCountries()

	if len(countries) < 3 {
		return false
	}

	newCountries := DeleteFromCountries(countries, deleteCountry)

	if err := UpdateCountries(newCountries); err != nil {
		return false
	}

	return true
}

func UpdateCountries(countries []country) error {
	result, err := json.Marshal(countries)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, result, 0644); err != nil {
		return err
	}

	return nil
}

func DeleteFromCountries(countries []country, deleteCountry string) []country {
	for i, a := range countries {
		if a.Country == deleteCountry {
			return append(countries[:i], countries[i+1:]...)
		}
	}

	return countries
}
