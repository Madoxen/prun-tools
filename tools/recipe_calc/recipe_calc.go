package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fio_connector"
)

type CalculatedRecipe struct {
	RecipeName   string
	BuildingName string
	revenue      float64
	cost         float64
	profit       float64
}

const EXCHANGE_CODE = "NC1"

// Calculates cost and possible profit from production
func main() {
	recipes := fio_connector.Get_all_recipes()

	for _, recipe := range recipes {
		fmt.Println(recipe.RecipeName)
		fmt.Println(recipe.Inputs)
		fmt.Println(recipe.Outputs)
		fmt.Println(recipe.BuildingTicker)
		fmt.Println(recipe.TimeMs)
	}

	cxData := fio_connector.Get_cx_data()

	//Group cxData by material
	priceLookup := make(map[string]map[string]fio_connector.CXData)

	for _, cxEntry := range cxData {
		key := cxEntry.MaterialTicker
		priceLookupEntry, ok := priceLookup[key]
		if !ok {
			priceLookup[key] = make(map[string]fio_connector.CXData)
			priceLookupEntry = priceLookup[key]
		}

		priceLookupEntry[cxEntry.ExchangeCode] = cxEntry
	}

	//Price calculations
	output_data := make([]CalculatedRecipe, 0)
	for _, recipe := range recipes {

		cost := 0.0
		revenue := 0.0

		//get all substrate costs
		for _, substrate := range recipe.Inputs {
			price := priceLookup[substrate.Ticker][EXCHANGE_CODE].Ask
			cost += price * float64(substrate.Amount)
		}

		//get all outputs profits
		for _, output := range recipe.Outputs {
			price := priceLookup[output.Ticker][EXCHANGE_CODE].Bid
			revenue += price * float64(output.Amount)
		}

		profit := revenue - cost

		output_data = append(output_data, CalculatedRecipe{
			RecipeName:   recipe.RecipeName,
			BuildingName: recipe.BuildingTicker,
			revenue:      revenue,
			cost:         cost,
			profit:       profit,
		})
	}

	//Save data to CSV
	file, err := os.Create("recipe_calc.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Ticker", "BuildingName", "Cost", "Revenue", "Profit"}
	err = writer.Write(headers)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for _, data := range output_data {
		row := []string{
			data.RecipeName,
			data.BuildingName,
			fmt.Sprintf("%v", data.cost),
			fmt.Sprintf("%v", data.revenue),
			fmt.Sprintf("%v", data.profit),
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row:", row, "| Error:", err)
			return
		}
	}

}
