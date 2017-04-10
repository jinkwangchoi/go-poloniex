package poloniex

import (
	"github.com/shopspring/decimal"
)

type Volume map[string]decimal.Decimal

type VolumeCollection struct {
	TotalBTC  decimal.Decimal `json:"totalBTC,string"`
	TotalUSDT decimal.Decimal `json:"totalUSDT,string"`
	TotalXMR  decimal.Decimal `json:"totalXMR,string"`
	TotalXUSD decimal.Decimal `json:"totalXUSD,string"`
	Volumes   map[string]Volume
}
