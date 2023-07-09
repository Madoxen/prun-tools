package fio_connector

// TODO: change floats to decimal type for handling financial data
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

type CXStation struct {
	StationID               string `json:"StationId"`
	NaturalID               string `json:"NaturalId"`
	Name                    string `json:"Name"`
	SystemID                string `json:"SystemId"`
	SystemNaturalID         string `json:"SystemNaturalId"`
	SystemName              string `json:"SystemName"`
	CommisionTimeEpochMs    int64  `json:"CommisionTimeEpochMs"`
	ComexID                 string `json:"ComexId"`
	ComexName               string `json:"ComexName"`
	ComexCode               string `json:"ComexCode"`
	WarehouseID             string `json:"WarehouseId"`
	CountryID               string `json:"CountryId"`
	CountryCode             string `json:"CountryCode"`
	CountryName             string `json:"CountryName"`
	CurrencyNumericCode     int    `json:"CurrencyNumericCode"`
	CurrencyCode            string `json:"CurrencyCode"`
	CurrencyName            string `json:"CurrencyName"`
	CurrencyDecimals        int    `json:"CurrencyDecimals"`
	GovernorID              any    `json:"GovernorId"`
	GovernorUserName        any    `json:"GovernorUserName"`
	GovernorCorporationID   string `json:"GovernorCorporationId"`
	GovernorCorporationName string `json:"GovernorCorporationName"`
	GovernorCorporationCode string `json:"GovernorCorporationCode"`
	UserNameSubmitted       string `json:"UserNameSubmitted"`
	Timestamp               string `json:"Timestamp"`
}

type Building struct {
	BuildingCosts []struct {
		CommodityName   string  `json:"CommodityName"`
		CommodityTicker string  `json:"CommodityTicker"`
		Weight          float64 `json:"Weight"`
		Volume          float64 `json:"Volume"`
		Amount          int     `json:"Amount"`
	} `json:"BuildingCosts"`
	Recipes []struct {
		Inputs             []any  `json:"Inputs"`
		Outputs            []any  `json:"Outputs"`
		BuildingRecipeID   string `json:"BuildingRecipeId"`
		DurationMs         int    `json:"DurationMs"`
		RecipeName         string `json:"RecipeName"`
		StandardRecipeName string `json:"StandardRecipeName"`
	} `json:"Recipes"`
	BuildingID        string `json:"BuildingId"`
	Name              string `json:"Name"`
	Ticker            string `json:"Ticker"`
	Expertise         string `json:"Expertise"`
	Pioneers          int    `json:"Pioneers"`
	Settlers          int    `json:"Settlers"`
	Technicians       int    `json:"Technicians"`
	Engineers         int    `json:"Engineers"`
	Scientists        int    `json:"Scientists"`
	AreaCost          int    `json:"AreaCost"`
	UserNameSubmitted string `json:"UserNameSubmitted"`
	Timestamp         string `json:"Timestamp"`
}

type WorkforceNeeds struct {
	Needs []struct {
		MaterialID       string  `json:"MaterialId"`
		MaterialName     string  `json:"MaterialName"`
		MaterialTicker   string  `json:"MaterialTicker"`
		MaterialCategory string  `json:"MaterialCategory"`
		Amount           float64 `json:"Amount"`
	} `json:"Needs"`
	WorkforceType string `json:"WorkforceType"`
}
