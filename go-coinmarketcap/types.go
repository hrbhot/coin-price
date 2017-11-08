package coinmarketcap

//Coin struct
type Coin struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int     `json:"rank,string"`
	PriceUsd         float64 `json:"price_usd,string"`
	PriceBtc         float64 `json:"price_btc,string"`
	Usd24hVolume     float64 `json:"24h_volume_usd,string"`
	MarketCapUsd     float64 `json:"market_cap_usd,string"`
	AvailableSupply  float64 `json:"available_supply,string"`
	TotalSupply      float64 `json:"total_supply,string"`
	PercentChange1h  float64 `json:"percent_change_1h,string"`
	PercentChange24h float64 `json:"percent_change_24h,string"`
	PercentChange7d  float64 `json:"percent_change_7d,string"`
	LastUpdated      string  `json:"last_updated"`
}

//GlobalMarketData struct
type GlobalMarketData struct {
	TotalMarketCapUsd            float64 `json:"total_market_cap_usd"`
	Total24hVolumeUsd            float64 `json:"total_24h_volume_usd"`
	BitcoinPercentageOfMarketCap float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies             int     `json:"active_currencies"`
	ActiveAssets                 int     `json:"active_assets"`
	ActiveMarkets                int     `json:"active_markets"`
}

// CoinGraph struct
type CoinGraph struct {
	MarketCapByAvailableAupply [][]float64 `json:"market_cap_by_available_supply"`
	PriceBtc                   [][]float64 `json:"price_btc"`
	PriceUsd                   [][]float64 `json:"price_usd"`
	VolumeUsd                  [][]float64 `json:"volume_usd"`
}
// 美元汇率
type Rates struct {
	Base       string `json:"base"`
	Date       string `json:"date"`
	Currencies struct {
		CNY float64 `json:"CNY"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		EUR float64 `json:"EUR"`
		NZD float64 `json:"NZD"`
		RUB float64 `json:"RUB"`
		JPY float64 `json:"JPY"`
		USD float64 `json:"USD"`
	} `json:"rates"`
}