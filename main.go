package main

import (
  "fmt"
  "net/http"
  "net/url"
  "encoding/json"
  "io/ioutil"
  "time"
)
const (
  token = "" // telegram robot token
  chatid = "" // group id for sending the message
  Currency = "" // type of currency to check (format : symbol)
  time_to_check = "" // time to check the price (format : 24h)
)
var (
  Coinlore string = fmt.Sprintf("https://api.coinlore.net/api/ticker/?id=%v", Coins[Currency])
  telegramapi string
)

type info struct {
Id string               `json:"id"`
Symbol string           `json:"symbol"`
Name  string            `json:"name"`
Nameid string           `json:"nameid"`
Rank int                `json:"rank"`
Price string            `json:"price_usd"`
Daychanges string       `json:"percent_change_24h"`
Hourchanges string      `json:"percent_change_1h"`
Weekchanges string      `json:"percent_change_7d"`
Marketcap string        `json:"market_cap_usd"`
Dayvolume string        `json:"volume24"`
Volume24_native string  `json:"volume24_native"`
Csupply string          `json:"csupply"`
Price_btc string        `json:"price_btc"`
Tsupply string          `json:"tsupply"`
Msupply string          `json:"msupply"`
}

var Coins = map[string]int{
  "btc":  90,
  "eth":  80,
  "xrp":  58,
  "usdt": 518,
  "ltc":  1,                //
  "bch":  2321,             //Supported currencies
  "doge": 2,                //
  "xmr": 28,
  "dash": 8,
  "zec": 134,
  "etc": 118,
}

func main() {
  for {
    h, m, _ := time.Now().Clock()
    if  fmt.Sprintf("%v:%v",h, m) == time_to_check {
      output := []info{}
      Input := price()
      err := json.Unmarshal([]byte(Input), &output)
      if err != nil {
        panic(err.Error())
      }
      sendMessage(output)
      _, err = http.Get(telegramapi)
      if err != nil {
        panic(err.Error())
      }
    }
    time.Sleep(60 * time.Second)
  }
}

func price() string {
  res, err := http.Get(Coinlore)
  if err != nil {
    panic(err.Error())
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err.Error())
  }
  return string(body)
}

func sendMessage(output []info){
  var message string = fmt.Sprintf("\nCurrency: %v\nRank: %v\nPrice: %v$\nChanges(1h): %v%%\nChanges(24h): %v%%\nChanges(7d): %v%%\nMarket Cap: %v",
  output[0].Name,
  output[0].Rank,
  output[0].Price,
  output[0].Daychanges,
  output[0].Hourchanges,
  output[0].Weekchanges,
  output[0].Marketcap)
  telegramapi = fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=-%v&text=%v",token, chatid, url.QueryEscape(message))
}
