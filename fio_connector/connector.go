package fio_connector

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var priceLookup map[string]map[string]CXData = nil
var buildingLookup map[string]Building = nil
var workforceUpkeepLookup map[string]WorkforceNeeds = nil
var cxDataCache []CXData = nil
var cxStationDataCache []CXStation = nil
var recipeDataCache []Recipe = nil
var buildingDataCache []Building = nil
var workforceNeedsDataCache []WorkforceNeeds = nil

const API_URL = "https://rest.fnar.net/"

func get_fio(endpoint string) (*http.Response, error) {
	fmt.Println("Requesting: " + API_URL + endpoint)
	return http.Get(API_URL + endpoint)
}

// Makes GET request to FIO endpoint and then tries to decode JSON into selected
// model
func get_fio_object(endpoint string, target interface{}) {
	resp, err := get_fio(endpoint)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Get_cx_data() []CXData {
	if cxDataCache == nil {
		get_fio_object("exchange/full", &cxDataCache)
	}
	return cxDataCache
}

func Get_all_recipes() []Recipe {
	if recipeDataCache == nil {
		get_fio_object("recipes/allrecipes", &recipeDataCache)
	}
	return recipeDataCache
}

func Get_cx_stations() []CXStation {
	if cxStationDataCache == nil {
		get_fio_object("exchange/station", &cxStationDataCache)
	}
	return cxStationDataCache
}

func Get_buildings() []Building {
	if buildingDataCache == nil {
		get_fio_object("building/allbuildings", &buildingDataCache)
	}
	return buildingDataCache
}

func Get_building(ticker string) Building {
	if buildingLookup == nil {
		buildingLookup = make(map[string]Building)
		for _, b := range Get_buildings() {
			key := b.Ticker
			buildingLookup[key] = b
		}
	}
	return buildingLookup[ticker]
}

func Get_workforce_upkeep(workforce string) WorkforceNeeds {
	if workforceNeedsDataCache == nil {
		get_fio_object("global/workforceneeds", &workforceNeedsDataCache)
	}

	if workforceUpkeepLookup == nil {
		workforceUpkeepLookup = make(map[string]WorkforceNeeds, 0)
		for _, b := range workforceNeedsDataCache {
			workforceUpkeepLookup[b.WorkforceType] = b
		}
	}
	return workforceUpkeepLookup[workforce]
}

func Get_all_workforce_upkeep() map[string]WorkforceNeeds {
	if workforceNeedsDataCache == nil {
		get_fio_object("global/workforceneeds", &workforceNeedsDataCache)
	}

	if workforceUpkeepLookup == nil {
		workforceUpkeepLookup = make(map[string]WorkforceNeeds, 0)
		for _, b := range workforceNeedsDataCache {
			workforceUpkeepLookup[b.WorkforceType] = b
		}
	}
	return workforceUpkeepLookup
}

func Get_price(ticker string, exchange string) CXData {
	//Fill price lookup if it does not exist yet
	if priceLookup == nil {
		cxData := Get_cx_data()
		//Group cxData by material
		priceLookup = make(map[string]map[string]CXData)

		for _, cxEntry := range cxData {
			key := cxEntry.MaterialTicker
			priceLookupEntry, ok := priceLookup[key]
			if !ok {
				priceLookup[key] = make(map[string]CXData)
				priceLookupEntry = priceLookup[key]
			}

			priceLookupEntry[cxEntry.ExchangeCode] = cxEntry
		}
	}

	return priceLookup[ticker][exchange]
}
