package fio_connector

// A material and amount
type MaterialAmount struct {
	Ticker string `json:"Ticker"`
	Amount int    `json:"Amount"`
}

type Recipe struct {
	BuildingTicker string           `json:"BuildingTicker"`
	RecipeName     string           `json:"RecipeName"`
	Inputs         []MaterialAmount `json:"Inputs"`
	Outputs        []MaterialAmount `json:"Outputs"`
	TimeMs         int              `json:"TimeMs"`
}

type Order struct {
	OrderID     string  `json:"OrderId"`
	CompanyID   string  `json:"CompanyId"`
	CompanyName string  `json:"CompanyName"`
	CompanyCode string  `json:"CompanyCode"`
	ItemCount   int     `json:"ItemCount"`
	ItemCost    float64 `json:"ItemCost"`
}

type CXData struct {
	BuyingOrders        []Order `json:"BuyingOrders"`
	SellingOrders       []Order `json:"SellingOrders"`
	CXDataModelID       string  `json:"CXDataModelId"`
	MaterialName        string  `json:"MaterialName"`
	MaterialTicker      string  `json:"MaterialTicker"`
	MaterialID          string  `json:"MaterialId"`
	ExchangeName        string  `json:"ExchangeName"`
	ExchangeCode        string  `json:"ExchangeCode"`
	Currency            string  `json:"Currency"`
	Previous            float64 `json:"Previous"`
	Price               float64 `json:"Price"`
	PriceTimeEpochMs    int64   `json:"PriceTimeEpochMs"`
	High                float64 `json:"High"`
	AllTimeHigh         float64 `json:"AllTimeHigh"`
	Low                 float64 `json:"Low"`
	AllTimeLow          float64 `json:"AllTimeLow"`
	Ask                 float64 `json:"Ask"`
	AskCount            int     `json:"AskCount"`
	Bid                 float64 `json:"Bid"`
	BidCount            int     `json:"BidCount"`
	Supply              int     `json:"Supply"`
	Demand              int     `json:"Demand"`
	Traded              int     `json:"Traded"`
	VolumeAmount        float64 `json:"VolumeAmount"`
	PriceAverage        float64 `json:"PriceAverage"`
	NarrowPriceBandLow  float64 `json:"NarrowPriceBandLow"`
	NarrowPriceBandHigh float64 `json:"NarrowPriceBandHigh"`
	WidePriceBandLow    float64 `json:"WidePriceBandLow"`
	WidePriceBandHigh   float64 `json:"WidePriceBandHigh"`
	MMBuy               float64 `json:"MMBuy"`
	MMSell              float64 `json:"MMSell"`
	UserNameSubmitted   string  `json:"UserNameSubmitted"`
	Timestamp           string  `json:"Timestamp"`
}
