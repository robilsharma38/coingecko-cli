package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robilsharma38/coingecko-cli/config"
)

type CoinList struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func GetCoinList() ([]CoinList, error) {
	dotenv := config.GoDotEnvVariable("token")
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.CoinListURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cg-demo-api-key", dotenv)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var coins []CoinList
	err = json.NewDecoder(resp.Body).Decode(&coins)
	if err != nil {
		fmt.Println(err)
	}
	return coins[:15], nil
}
