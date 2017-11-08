package main

import (
	"fmt"
	"log"

	coinApi "./go-coinmarketcap"
)

func main() {
	// Get global market data

	// Get info about coin

	UsdInfo, err := coinApi.GetUsdData("usd")
	// if err != nil {
	//	log.Println(err)
	// } else {
	fmt.Println("US Dollar To RMB:\t", UsdInfo.Currencies.CNY)
	Cny:=UsdInfo.Currencies.CNY

	// }

	btcInfo, err := coinApi.GetCoinData("bitcoin")

	if err != nil {
		log.Println(err)
	} else {
		p := btcInfo
		p1 := &btcInfo.MarketCapUsd
		fmt.Printf("%s price=$%v price=%.2frmb 24h=%v%% rank=%v cap=%.0f\n ", p.Symbol, p.PriceUsd,p.PriceUsd*Cny, p.PercentChange24h, p.Rank, *p1)
	}
	ethInfo, err := coinApi.GetCoinData("ethereum")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s price=$%v price=%.2frmb 24h=%v%% rank=%v cap=$%.0f\n", ethInfo.Symbol, ethInfo.PriceUsd,ethInfo.PriceUsd*Cny, ethInfo.PercentChange24h, ethInfo.Rank, ethInfo.MarketCapUsd)
	}

	eosInfo, err := coinApi.GetCoinData("eos")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s price=$%v price=%.2frmb 24h=%v%% rank=%v cap=$%.0f\n", eosInfo.Symbol, eosInfo.PriceUsd,eosInfo.PriceUsd*Cny,eosInfo.PercentChange24h, eosInfo.Rank, eosInfo.MarketCapUsd)
	}
	omgInfo, err := coinApi.GetCoinData("omisego")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s price=$%v price=%.2frmb 24h=%v%% rank=%v cap=$%.0f\n", omgInfo.Symbol, omgInfo.PriceUsd, omgInfo.PriceUsd*Cny,omgInfo.PercentChange24h, omgInfo.Rank, omgInfo.MarketCapUsd)
	}
	payInfo, err := coinApi.GetCoinData("tenx")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s price=$%v price=%.2frmb 24h=%v%% rank=%v cap=$%.0f\n", payInfo.Symbol, payInfo.PriceUsd, payInfo.PriceUsd*Cny,payInfo.PercentChange24h, payInfo.Rank, payInfo.MarketCapUsd)

	}
	
}
