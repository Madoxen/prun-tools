package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fio_connector"
)

type CalculatedRecipe struct {
	RecipeName     string
	BuildingName   string
	revenue        float64
	workforce_cost float64
	cost           float64
	profit         float64
	profit_per_day float64
}

const CX = "NC1"
const DAY_MILISECONDS = 24 * 3600 * 1000

// Calculates cost and possible profit from production
func main() {
	recipes := fio_connector.Get_all_recipes()

	//Prefetch workforce needs as those are static and only will be multiplied
	//by workforce numbers
	workforce_upkeep := fio_connector.Get_all_workforce_upkeep()
	//Workforce cost per person per milisecond
	workforce_cost := make(map[string]float64)
	for name, needs := range workforce_upkeep {
		cost := 0.0
		for _, need := range needs.Needs {
			price := fio_connector.Get_price(need.MaterialTicker, CX).Ask
			cost += price * float64(need.Amount)
		}
		workforce_cost[name] = cost / DAY_MILISECONDS / 100
	}

	pioneer_cost := workforce_cost["PIONEER"]
	settler_cost := workforce_cost["SETTLER"]
	technician_cost := workforce_cost["TECHNICIAN"]
	engineer_cost := workforce_cost["ENGINEER"]
	scientist_cost := workforce_cost["SCIENTIST"]

	//Price calculations
	output_data := make([]CalculatedRecipe, 0)
	for _, recipe := range recipes {

		cost := 0.0
		revenue := 0.0

		//get all substrate costs
		for _, substrate := range recipe.Inputs {
			price := fio_connector.Get_price(substrate.Ticker, CX).Ask
			cost += price * float64(substrate.Amount)
		}

		//get workforce upkeep costs per recipe run
		building := fio_connector.Get_building(recipe.BuildingTicker)
		time := float64(recipe.TimeMs)
		workforce_total_cost := float64(building.Pioneers)*pioneer_cost*time +
			float64(building.Settlers)*settler_cost*time +
			float64(building.Technicians)*technician_cost*time +
			float64(building.Engineers)*engineer_cost*time +
			float64(building.Scientists)*scientist_cost*time

		cost += workforce_total_cost

		//get all outputs profits
		for _, output := range recipe.Outputs {
			price := fio_connector.Get_price(output.Ticker, CX).Bid
			revenue += price * float64(output.Amount)
		}

		profit := revenue - cost
		profit_per_day := (profit / time) * 3600000 * 24

		output_data = append(output_data, CalculatedRecipe{
			RecipeName:     recipe.RecipeName,
			BuildingName:   recipe.BuildingTicker,
			revenue:        revenue,
			workforce_cost: workforce_total_cost,
			cost:           cost,
			profit:         profit,
			profit_per_day: profit_per_day,
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

	headers := []string{"Ticker", "BuildingName", "Workforce Cost", "Cost", "Revenue", "Profit", "Profit Per Day"}
	err = writer.Write(headers)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for _, data := range output_data {
		row := []string{
			data.RecipeName,
			data.BuildingName,
			fmt.Sprintf("%v", data.workforce_cost),
			fmt.Sprintf("%v", data.cost),
			fmt.Sprintf("%v", data.revenue),
			fmt.Sprintf("%v", data.profit),
			fmt.Sprintf("%v", data.profit_per_day),
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row:", row, "| Error:", err)
			return
		}
	}
}
