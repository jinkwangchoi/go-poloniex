package poloniex

import "github.com/shopspring/decimal"

type Currency struct {
	Name               string          `json:"name"`
	MaxDailyWithdrawal string          `json:"maxDailyWithdrawal"`
	TxFee              decimal.Decimal `json:"txFee,string"`
	MinConf            int             `json:"minConf"`
	Disabled           int             `json:"disabled"`
	Frozen             int             `json:"frozen"`
	Delisted           int             `json:"delisted"`
}

type Currencies struct {
	Pair map[string]Currency
}
