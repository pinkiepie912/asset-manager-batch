package stock

import (
	"fmt"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
)

func GetQuoteBy(symbol string) *finance.Quote {
	stockQuote, err := quote.Get(symbol)
	if err != nil {
		panic(err)
	}

	fmt.Println(stockQuote)
	return stockQuote
}