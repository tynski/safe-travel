package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type country struct {
	Country string `json:"country"`
	Others  []struct {
		Code  string `json:"code"`
		Color string `json:"color"`
	} `json:"others"`
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

	if err := UpdateCountries(countries); err != nil {
		return false
	}

	return true
}

func DeleteCountry(deleteCountry string) bool {
	countries := GetCountries()

	for i, a := range countries {
		if a.Country == deleteCountry {
			countries = append(countries[:i], countries[i+1:]...)
		}
	}

	if err := UpdateCountries(countries); err != nil {
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
