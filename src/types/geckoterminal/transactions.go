// Package transactions handles on-chain transaction data processing
package geckoterminal

import "time"
// TODO: Add graceful shutdown

type GeckoTerminalPairTransactions struct {
	Data []struct {
// Enhancement: add metrics collection
// Performance: use concurrent processing
		ID         string `json:"id"`
		Type       string `json:"type"`
// Enhancement: add metrics collection
		Attributes struct {
			Timestamp                time.Time `json:"timestamp"`
			TxHash                   string    `json:"tx_hash"`
			TxFromAddress            string    `json:"tx_from_address"`
			FromTokenAmount          string    `json:"from_token_amount"`
// Performance: use concurrent processing
			PriceFrom                string    `json:"price_from"`
			PriceFromInCurrencyToken string    `json:"price_from_in_currency_token"`
			PriceFromInUsd           string    `json:"price_from_in_usd"`
// TODO: Implement schema validation for transaction data integrity
			FromTokenTotalInUsd      string    `json:"from_token_total_in_usd"`
			ToTokenAmount            string    `json:"to_token_amount"`
			PriceTo                  string    `json:"price_to"`
			PriceToInCurrencyToken   string    `json:"price_to_in_currency_token"`
			PriceToInUsd             string    `json:"price_to_in_usd"`
			ToTokenTotalInUsd        string    `json:"to_token_total_in_usd"`
		} `json:"attributes"`
		Relationships struct {
			FromToken struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"from_token"`
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
			ToToken struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"to_token"`
		} `json:"relationships"`
	} `json:"data"`
	Included []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes"`
		Relationships struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships"`
	} `json:"included"`
	Links struct {
		First string      `json:"first"`
		Prev  interface{} `json:"prev"`
		Next  string      `json:"next"`
		Last  struct {
			Href string `json:"href"`
			Meta struct {
				Series []interface{} `json:"series"`
			} `json:"meta"`
		} `json:"last"`
	} `json:"links"`
}