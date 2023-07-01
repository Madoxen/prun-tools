package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/fio_connector"
	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// TODO: change floats to decimal type for handling financial data
type tradeRoute struct {
	materialTicker string
	askCX          string
	ask            float64
	askCount       int
	bidCX          string
	bid            float64
	bidCount       int
	profit         float64
}

type ByProfit []tradeRoute

func (a ByProfit) Len() int           { return len(a) }
func (a ByProfit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByProfit) Less(i, j int) bool { return a[i].profit < a[j].profit }

func main() {
	cxData := fio_connector.Get_cx_data()

	//Group cxData by material
	matData := make(map[string][]fio_connector.CXData)

	for _, cxEntry := range cxData {
		key := cxEntry.MaterialTicker
		matDataEntry, ok := matData[key]
		if !ok {
			matData[key] = []fio_connector.CXData{}
		}
		matData[key] = append(matDataEntry, cxEntry)
	}

	trade_routes := []tradeRoute{}
	//Find best trade routes based on BID/ASK and amount
	for ticker, data := range matData {
		min_ask := math.Inf(1)
		max_bid := math.Inf(-1)

		min_ask_idx := -1
		max_bid_idx := -1

		for idx, entry := range data {
			if entry.Ask < min_ask && entry.AskCount > 0 {
				min_ask = entry.Ask
				min_ask_idx = idx
			}

			if entry.Bid > max_bid && entry.BidCount > 0 {
				max_bid = entry.Bid
				max_bid_idx = idx
			}
		}

		if min_ask_idx == -1 || max_bid_idx == -1 {
			continue
		}

		ask_entry := data[min_ask_idx]
		bid_entry := data[max_bid_idx]

		bidCount := bid_entry.BidCount
		askCount := ask_entry.AskCount
		count := min(bidCount, askCount)
		profit := math.Round(((max_bid*float64(count))-(min_ask*float64(count)))*100) / 100

		trade_routes = append(trade_routes, tradeRoute{
			materialTicker: ticker,
			askCX:          ask_entry.ExchangeCode,
			ask:            min_ask,
			askCount:       ask_entry.AskCount,
			bidCX:          bid_entry.ExchangeCode,
			bid:            max_bid,
			bidCount:       bid_entry.BidCount,
			profit:         profit,
		})
	}

	file, err := os.Create("trade_finder.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Ticker", "AskCX", "Ask", "AskCount", "BidCX", "Bid", "BidCount", "TotalProfit"}
	err = writer.Write(headers)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	sort.Sort(ByProfit(trade_routes))
	for _, tr := range trade_routes {
		row := []string{
			tr.materialTicker,
			tr.askCX,
			fmt.Sprintf("%v", tr.ask),
			fmt.Sprintf("%v", tr.askCount),
			tr.bidCX,
			fmt.Sprintf("%v", tr.bid),
			fmt.Sprintf("%v", tr.bidCount),
			fmt.Sprintf("%v", tr.profit),
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row:", row, "| Error:", err)
			return
		}
	}

}
