package model

type StocksError struct {
	ErrorMsg      string   `json:"error,omitempty"`
	InvalidStocks []string `json:"invalidstocks,omitempty"`
}
