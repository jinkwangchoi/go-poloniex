package poloniex

import "github.com/shopspring/decimal"

type CandleStick struct {
	Date            PoloniexDate    `json:"date"`
	High            decimal.Decimal `json:"high"`
	Low             decimal.Decimal `json:"low"`
	Open            decimal.Decimal `json:"open"`
	Close           decimal.Decimal `json:"close"`
	Volume          decimal.Decimal `json:"volume"`
	QuoteVolume     decimal.Decimal `json:"quoteVolume"`
	WeightedAverage decimal.Decimal `json:"weightedAverage"`
}
